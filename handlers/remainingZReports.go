package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/chyupa/apiServer/utils/logger"
	"github.com/chyupa/fp700/commands"
	"net/http"
)

func GetRemainingZReports(w http.ResponseWriter, r *http.Request) {
	remainingZReports, err := commands.RemainingZReports()
	if err != nil {
		fmt.Println(err)
		logger.Error.Println(err)
		http.Error(w, err.Error(), 400)
		return
	}

	json.NewEncoder(w).Encode(remainingZReports)
}
