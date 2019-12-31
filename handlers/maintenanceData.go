package handlers

import (
	"encoding/json"
	"github.com/chyupa/fp700"
	"github.com/chyupa/fp700/commands"
	"net/http"
)

func MaintenanceData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	fp700.Port = r.Header.Get("ComPort")
	commandResponse := commands.MaintenanceData()

	json.NewEncoder(w).Encode(commandResponse)
}
