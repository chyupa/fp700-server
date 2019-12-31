package handlers

import (
	"encoding/json"
	"github.com/chyupa/fp700"
	"github.com/chyupa/fp700/commands"
	"io/ioutil"
	"log"
	"net/http"
)

func Fiscalize(w http.ResponseWriter, r *http.Request) {
	var fiscalizeRequest commands.FiscalizeRequest
	w.Header().Set("content-type", "application/json")
	fp700.Port = r.Header.Get("ComPort")

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(reqBody, &fiscalizeRequest)

	fiscalizeResponse, _ := commands.Fiscalize(fiscalizeRequest)

	json.NewEncoder(w).Encode(fiscalizeResponse)
}
