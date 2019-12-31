package handlers

import (
	"encoding/json"
	"github.com/chyupa/fp700/commands"
	"io/ioutil"
	"log"
	"net/http"
)

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
