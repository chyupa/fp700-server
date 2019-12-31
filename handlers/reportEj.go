package handlers

import (
	"encoding/json"
	"github.com/chyupa/fp700"
	"github.com/chyupa/fp700/commands"
	"io/ioutil"
	"log"
	"net/http"
)

func ReportEj(w http.ResponseWriter, r *http.Request) {
	var ejRequest commands.EjRequest

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(reqBody, &ejRequest)

	fp700.Port = r.Header.Get("ComPort")
	response, err := commands.ReadEj(ejRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	w.Write([]byte(response))
}
