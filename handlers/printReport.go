package handlers

import (
	"encoding/json"
	"github.com/chyupa/fp700/commands"
	"io/ioutil"
	"log"
	"net/http"
)

type printReportRequest struct {
	Type string `json:"type"`
}

func PrintReport(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	var prRequest printReportRequest
	err = json.Unmarshal(reqBody, &prRequest)
	log.Println(err)

	response := commands.PrintReport(prRequest.Type)

	json.NewEncoder(w).Encode(response)
}
