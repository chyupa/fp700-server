package handlers

import (
	"encoding/json"
	"github.com/chyupa/fp700"
	"github.com/chyupa/fp700/commands"
	"io/ioutil"
	"log"
	"net/http"
)

func PrintEj(w http.ResponseWriter, r *http.Request) {
	var ejRequest commands.PrintEjRequest

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(reqBody, &ejRequest)

	fp700.Port = r.Header.Get("ComPort")
	err = commands.PrintEj(ejRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	} else {
		w.Write([]byte("success"))
	}
}
