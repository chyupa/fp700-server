package handlers

import (
	"encoding/json"
	"github.com/chyupa/fp700/commands"
	"io/ioutil"
	"log"
	"net/http"
)

func GetVat(w http.ResponseWriter, r *http.Request) {
	vat := commands.GetVat()

	w.Write([]byte(vat))
}

func SetVat(w http.ResponseWriter, r *http.Request) {
	var setVatRequest commands.SetVatRequest

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(reqBody, &setVatRequest)

	setVatResponse, err := commands.SetVat(setVatRequest)
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(err.Error())
	} else {
		json.NewEncoder(w).Encode(setVatResponse)
	}
}

func GetVatChanges(w http.ResponseWriter, r *http.Request) {
	vat := commands.GetVatChanges()

	json.NewEncoder(w).Encode(vat)
}
