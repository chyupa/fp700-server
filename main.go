package main

import (
	"bytes"
	"fmt"
	"github.com/chyupa/apiServer/handlers"
	"github.com/chyupa/fp700"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var GeneralLogger *log.Logger
var ErrorLogger *log.Logger

func init() {
	absPath, err := filepath.Abs("src/github.com/chyupa/apiServer")
	if err != nil {
		fmt.Println(err)
	}

	generalLog, err := os.OpenFile(absPath+"/log.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	GeneralLogger = log.New(generalLog, "General logger: ", log.LstdFlags|log.Llongfile)
	ErrorLogger = log.New(generalLog, "Error logger: ", log.LstdFlags|log.Llongfile)
}

func main() {
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
	api.HandleFunc("/maintenance-data", handlers.MaintenanceData).Methods(http.MethodGet)

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

	api.Use(loggingMiddleware, setPortMiddleware, setJsonContentType)

	ErrorLogger.Fatal(http.ListenAndServe(":8082", c.Handler(r)))
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			ErrorLogger.Println(err)
		}
		GeneralLogger.Println(r.Method, r.RequestURI, string(reqBody))

		r.Body = ioutil.NopCloser(bytes.NewBuffer(reqBody))
		next.ServeHTTP(w, r)
	})
}

func setPortMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fp700.Port = r.Header.Get("ComPort")

		next.ServeHTTP(w, r)
	})
}

func setJsonContentType(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")

		next.ServeHTTP(w, r)
	})
}
