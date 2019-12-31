package handlers

import (
	"encoding/json"
	"github.com/chyupa/fp700"
	"github.com/chyupa/fp700/commands"
	"io/ioutil"
	"log"
	"net/http"
)

func CloseDmje(w http.ResponseWriter, r *http.Request) {
	var cdRequest commands.CloseDmjeRequest
	w.Header().Set("content-type", "application/json")

	fp700.Port = r.Header.Get("ComPort")

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(reqBody, &cdRequest)

	// get response and send it
	commands.CloseDmje(cdRequest)

	w.Write([]byte("Success"))
}
