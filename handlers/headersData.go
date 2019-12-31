package handlers

import (
	"encoding/json"
	"github.com/chyupa/fp700"
	"github.com/chyupa/fp700/commands"
	"io/ioutil"
	"log"
	"net/http"
)

func HeadersData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	fp700.Port = r.Header.Get("ComPort")
	commandResponse := commands.HeadersData()

	json.NewEncoder(w).Encode(commandResponse)
}

func SetHeadersData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	fp700.Port = r.Header.Get("ComPort")
	var hdrRequest commands.SetHeadersDataRequest

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(reqBody, &hdrRequest)

	commandResponse := commands.SetHeadersData(hdrRequest)

	json.NewEncoder(w).Encode(commandResponse)
}