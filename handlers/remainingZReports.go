package handlers

import (
	"encoding/json"
	"github.com/chyupa/fp700/commands"
	"net/http"
)

func GetRemainingZReports(w http.ResponseWriter, r *http.Request) {
	remainingZReports := commands.RemainingZReports()

	json.NewEncoder(w).Encode(remainingZReports)
}
