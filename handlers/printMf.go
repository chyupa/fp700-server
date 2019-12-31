package handlers

import (
	"encoding/json"
	"github.com/chyupa/fp700"
	"github.com/chyupa/fp700/commands"
	"io/ioutil"
	"log"
	"net/http"
)

func PrintMf(w http.ResponseWriter, r *http.Request) {
	var mfRequest commands.PrintMfRequest
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(reqBody, &mfRequest)

	fp700.Port = r.Header.Get("ComPort")
	err = commands.PrintMf(mfRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	} else {
		w.Write([]byte("success"))
	}

}
