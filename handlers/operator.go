package handlers

import (
	"encoding/json"
	"github.com/chyupa/fp700/commands"
	"io/ioutil"
	"log"
	"net/http"
)

func GetOperatorName(w http.ResponseWriter, r *http.Request) {
	operatorName, err := commands.GetOperatorName()
	if err != nil {
		w.WriteHeader(400)
	} else {
		w.Write([]byte(operatorName))
	}
}

func SetOperatorName(w http.ResponseWriter, r *http.Request) {
	var opRequest commands.SetOperatorNameRequest
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(reqBody, &opRequest)

	response, err := commands.SetOperatorName(opRequest)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
	} else {
		w.Write([]byte(response))
	}
}

func GetOperatorPassword(w http.ResponseWriter, r *http.Request) {
	operatorPassword, err := commands.GetOperatorPassword()
	if err != nil {
		w.WriteHeader(400)
	} else {
		w.Write([]byte(operatorPassword))
	}
}

func SetOperatorPassword(w http.ResponseWriter, r *http.Request) {
	var opRequest commands.SetOperatorPasswordRequest
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(reqBody, &opRequest)

	err = commands.SetOperatorPassword(opRequest)
	if err != nil {
		log.Println(err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
	} else {
		w.Write([]byte("Success"))
	}

}
