package main

import (
	"bytes"
	"fmt"
	"github.com/chyupa/fp700"
	"github.com/chyupa/fp700-server/handlers"
	"github.com/chyupa/fp700-server/utils/logger"
	"github.com/gorilla/mux"
	"github.com/kardianos/service"
	"github.com/rs/cors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type program struct{}

func (p *program) Start(s service.Service) error {
	go p.run()

	return nil
}

func (p *program) run() {
	SetupServer()
}

func (p *program) Stop(s service.Service) error {
	return nil
}

func main() {
	svcConfig := &service.Config{
		Name:        "ExchangeV2",
		DisplayName: "ExchangeV2",
		Description: "ExchangeV2 fiscal printer communication library",
	}

	prg := &program{}

	s, err := service.New(prg, svcConfig)
	if err != nil {
		fmt.Println(err)
		logger.Error.Println(err)
	}

	logg, err := s.Logger(nil)
	if err != nil {
		fmt.Println(err)
		logger.Error.Println(err)
	}

	// read command line args
	var argsWithoutProgName = os.Args[1:]

	for _, param := range argsWithoutProgName {
		if param == "install" {
			err = s.Install()
			if err != nil {
				fmt.Println(err)
				logg.Error(err)
			} else {
				fmt.Println("Service installed successfully.")
			}
		} else if param == "uninstall" {
			err = s.Uninstall()
			if err != nil {
				logg.Error(err)
			} else {
				fmt.Println("Service uninstalled successfully.")
			}
		}
	}

	err = s.Run()
	if err != nil {
		fmt.Println(err)
		logg.Error(err)
	}
}

func SetupServer() {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "OPTIONS"},
	})

	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1").Subrouter()

	api.HandleFunc("/time", handlers.Time).Methods(http.MethodGet)
	api.HandleFunc("/time", handlers.SetTime).Methods(http.MethodPost)

	api.HandleFunc("/report", handlers.PrintReport).Methods(http.MethodPost)
	api.HandleFunc("/report/ej", handlers.ReportEj).Methods(http.MethodPost)
	api.HandleFunc("/print/ej", handlers.PrintEj).Methods(http.MethodPost)
	api.HandleFunc("/report/mf", handlers.ReportMf).Methods(http.MethodPost)
	api.HandleFunc("/print/mf", handlers.PrintMf).Methods(http.MethodPost)

	api.HandleFunc("/init-display", handlers.InitDisplay).Methods(http.MethodPost)

	api.HandleFunc("/transaction", handlers.Transaction).Methods(http.MethodPost)
	api.HandleFunc("/last-receipt", handlers.GetLastReceipt).Methods(http.MethodPost)

	api.HandleFunc("/maintenance-data", handlers.MaintenanceData).Methods(http.MethodGet)
	api.HandleFunc("/check-fiscalized", handlers.GetFiscalization).Methods(http.MethodGet)

	api.HandleFunc("/anaf-files", handlers.AnafFiles).Methods(http.MethodPost)

	api.HandleFunc("/operator-name", handlers.GetOperatorName).Methods(http.MethodGet)
	api.HandleFunc("/operator-name", handlers.SetOperatorName).Methods(http.MethodPost)
	api.HandleFunc("/operator-password", handlers.GetOperatorPassword).Methods(http.MethodGet)
	api.HandleFunc("/operator-password", handlers.SetOperatorPassword).Methods(http.MethodPost)

	api.HandleFunc("/print/diagnostic", handlers.PrintDiagnostic).Methods(http.MethodPost)

	api.HandleFunc("/headers", handlers.HeadersData).Methods(http.MethodGet)
	api.HandleFunc("/headers/changes", handlers.GetHeaderChanges).Methods(http.MethodGet)
	api.HandleFunc("/headers", handlers.SetHeadersData).Methods(http.MethodPost)

	api.HandleFunc("/report/z/remaining", handlers.GetRemainingZReports).Methods(http.MethodGet)

	api.HandleFunc("/footers", handlers.GetFootersData).Methods(http.MethodGet)
	api.HandleFunc("/footers", handlers.SetFootersData).Methods(http.MethodPost)

	api.HandleFunc("/vat", handlers.GetVat).Methods(http.MethodGet)
	api.HandleFunc("/vat/changes", handlers.GetVatChanges).Methods(http.MethodGet)
	api.HandleFunc("/vat", handlers.SetVat).Methods(http.MethodPost)

	api.HandleFunc("/activate-service-password", handlers.ActivateServicePassword).Methods(http.MethodPost)
	api.HandleFunc("/change-service-password", handlers.ChangeServicePassword).Methods(http.MethodPost)

	api.HandleFunc("/set-printer-mode", handlers.SetPrinterMode).Methods(http.MethodPost)

	api.HandleFunc("/fiscalize", handlers.Fiscalize).Methods(http.MethodPost)
	api.HandleFunc("/fabrication-number", handlers.FabricationNumber).Methods(http.MethodPost)

	api.HandleFunc("/close-dmje", handlers.CloseDmje).Methods(http.MethodPost)

	api.HandleFunc("/last-z-date", handlers.GetLastZDate).Methods(http.MethodGet)

	api.Use(loggingMiddleware, setJsonContentType, setPortMiddleware)

	log.Fatal(http.ListenAndServe(":8082", c.Handler(r)))
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println(err)
			logger.Error.Println(err)
		}
		fmt.Println(r.Method, r.RequestURI, string(reqBody))
		logger.General.Println(r.Method, r.RequestURI, string(reqBody))

		r.Body = ioutil.NopCloser(bytes.NewBuffer(reqBody))
		next.ServeHTTP(w, r)
	})
}

func setPortMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		comPort := r.Header.Get("ComPort")
		if len(comPort) < 1 {
			http.Error(w, "110100", 500)
			return
		}

		fp700.Port = comPort
		next.ServeHTTP(w, r)
	})
}

func setJsonContentType(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")

		next.ServeHTTP(w, r)
	})
}
