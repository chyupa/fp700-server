package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/chyupa/fp700-server/utils/logger"
	"github.com/chyupa/fp700/commands"
	"io/ioutil"
	"log"
	"net/http"
)

func Transaction(w http.ResponseWriter, r *http.Request) {
	var transactionRequest commands.TransactionRequest

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		logger.Error.Println(err)
	}

	json.Unmarshal(reqBody, &transactionRequest.Commands)

	commandResponse, err := commands.Transaction(transactionRequest)
	if err != nil {
		fmt.Println(err)
		logger.Error.Println(err)
		http.Error(w, err.Error(), 400)
		return
	}

	json.NewEncoder(w).Encode(commandResponse)
}

func ServiceAmount(w http.ResponseWriter, r *http.Request) {
	var saRequest commands.ServiceAmountRequest

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(reqBody, &saRequest)

	// use response of command
	commands.ServiceAmount(saRequest)

	w.Write([]byte("Success"))
}

func GetLastReceipt(w http.ResponseWriter, r *http.Request) {
	type lastReceiptRequest struct {
		Cancel bool `json:"cancel"`
	}

	lrReq := lastReceiptRequest{}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Error.Println(err)
	}

	json.Unmarshal(reqBody, lrReq)

	response, err := commands.LastReceipt(lrReq.Cancel)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	json.NewEncoder(w).Encode(response)
}