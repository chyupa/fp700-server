package handlers

import (
	"encoding/json"
	"github.com/chyupa/fp700"
	"github.com/chyupa/fp700/commands"
	"io/ioutil"
	"log"
	"net/http"
)

func Transaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var transactionRequest commands.TransactionRequest

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(reqBody, &transactionRequest.Commands)

	fp700.Port = r.Header.Get("ComPort")
	commandResponse := commands.Transaction(transactionRequest)

	json.NewEncoder(w).Encode(commandResponse)
}
