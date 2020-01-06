package handlers

import (
	"encoding/json"
	"github.com/chyupa/fp700/commands"
	"io/ioutil"
	"log"
	"net/http"
)

func Transaction(w http.ResponseWriter, r *http.Request) {
	var transactionRequest commands.TransactionRequest

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(reqBody, &transactionRequest.Commands)

	commandResponse := commands.Transaction(transactionRequest)

	json.NewEncoder(w).Encode(commandResponse)
}
