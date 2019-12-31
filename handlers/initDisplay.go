package handlers

import (
	"encoding/json"
	"github.com/chyupa/fp700"
	"github.com/chyupa/fp700/commands"
	"io/ioutil"
	"log"
	"net/http"
)

func InitDisplay(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	fp700.Port = r.Header.Get("ComPort")
	var idRequest commands.InitDisplayRequest

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(reqBody, &idRequest)

	commands.InitDisplay(idRequest)

	w.Write([]byte("Success"))
}
