package handlers

import (
	"encoding/json"
	"github.com/chyupa/fp700/commands"
	"io/ioutil"
	"log"
	"net/http"
)

func Time(w http.ResponseWriter, r *http.Request) {
	response, err := commands.Time()
	if err != nil {
		log.Println(err)
		w.Write([]byte(err.Error()))
	}
	json.NewEncoder(w).Encode(response)
}

func SetTime(w http.ResponseWriter, r *http.Request) {
	var timeRequest commands.SetTimeRequest

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(reqBody, &timeRequest)

	errorCode := commands.SetTime(timeRequest)

	w.Write([]byte(errorCode))
}
