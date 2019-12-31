package main

import (
	"github.com/chyupa/apiServer/handlers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
)

type ApiResponse struct {
	ErrorCode int
	Data string
}

func main() {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://exchange.local"},
		AllowedHeaders: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "OPTIONS"},
	})

	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1").Subrouter()

	api.HandleFunc("/time", handlers.Time).Methods(http.MethodGet)

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
	api.HandleFunc("/operator-password", handlers.GetOperatorPassword).Methods(http.MethodGet)

	api.HandleFunc("/print/diagnostic", handlers.PrintDiagnostic).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe(":8082", c.Handler(r)))

}
