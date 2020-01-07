package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/chyupa/apiServer/utils/logger"
	"github.com/chyupa/fp700/commands"
	"io/ioutil"
	"net/http"
)

type printReportRequest struct {
	Type string `json:"type"`
}

func PrintReport(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		logger.Error.Println(err)
	}
	var prRequest printReportRequest
	json.Unmarshal(reqBody, &prRequest)

	response, err := commands.PrintReport(prRequest.Type)
	if err != nil {
		fmt.Println(err)
		logger.Error.Println(err)
		http.Error(w, err.Error(), 400)
		return
	}

	json.NewEncoder(w).Encode(response)
}
