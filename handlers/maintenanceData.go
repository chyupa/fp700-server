package handlers

import (
	"encoding/json"
	"github.com/chyupa/fp700/commands"
	"net/http"
)

func MaintenanceData(w http.ResponseWriter, r *http.Request) {
	commandResponse := commands.MaintenanceData()

	json.NewEncoder(w).Encode(commandResponse)
}
