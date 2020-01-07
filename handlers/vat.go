package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/chyupa/apiServer/utils/logger"
	"github.com/chyupa/fp700/commands"
	"io/ioutil"
	"net/http"
)

func GetVat(w http.ResponseWriter, r *http.Request) {
	vat, err := commands.GetVat()
	if err != nil {
		fmt.Println(err)
		logger.Error.Println(err)
		http.Error(w, err.Error(), 400)
		return
	}

	w.Write([]byte(vat))
}

func SetVat(w http.ResponseWriter, r *http.Request) {
	var setVatRequest commands.SetVatRequest

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		logger.Error.Println(err)
	}

	json.Unmarshal(reqBody, &setVatRequest)

	setVatResponse, err := commands.SetVat(setVatRequest)
	if err != nil {
		fmt.Println(err)
		logger.Error.Println(err)
		http.Error(w, err.Error(), 400)
		return
	}

	json.NewEncoder(w).Encode(setVatResponse)
}

func GetVatChanges(w http.ResponseWriter, r *http.Request) {
	vat, err := commands.GetVatChanges()
	if err != nil {
		fmt.Println(err)
		logger.Error.Println(err)
		http.Error(w, err.Error(), 400)
		return
	}

	json.NewEncoder(w).Encode(vat)
}
