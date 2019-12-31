package handlers

import (
	"encoding/json"
	"github.com/chyupa/fp700"
	"github.com/chyupa/fp700/commands"
	"io/ioutil"
	"log"
	"net/http"
)

func ReportMf(w http.ResponseWriter, r *http.Request) {
	var mfRequest commands.MfRequest
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(reqBody, &mfRequest)

	fp700.Port = r.Header.Get("ComPort")
	commandResponse, err := commands.ReadMf(mfRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	w.Write([]byte(commandResponse))
}
