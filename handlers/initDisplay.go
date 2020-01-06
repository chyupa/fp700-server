package handlers

import (
	"encoding/json"
	"github.com/chyupa/fp700/commands"
	"io/ioutil"
	"log"
	"net/http"
)

func InitDisplay(w http.ResponseWriter, r *http.Request) {
	var idRequest commands.InitDisplayRequest

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(reqBody, &idRequest)

	commands.InitDisplay(idRequest)

	w.Write([]byte("Success"))
}
