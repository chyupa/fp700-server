package handlers

import (
	"encoding/json"
	"github.com/chyupa/fp700/commands"
	"io/ioutil"
	"log"
	"net/http"
)

func SetFootersData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var footerRequest commands.SetFootersDataRequest
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(reqBody, &footerRequest)

	commands.SetFootersData(footerRequest)

	w.Write([]byte("Success"))
}

func GetFootersData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	footerResponse := commands.FootersData()

	json.NewEncoder(w).Encode(footerResponse)
}
