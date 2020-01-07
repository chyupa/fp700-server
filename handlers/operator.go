package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/chyupa/apiServer/utils/logger"
	"github.com/chyupa/fp700/commands"
	"io/ioutil"
	"log"
	"net/http"
)

func GetOperatorName(w http.ResponseWriter, r *http.Request) {
	operatorName, err := commands.GetOperatorName()
	if err != nil {
		fmt.Println(err)
		logger.Error.Println(err)
		http.Error(w, err.Error(), 400)
		return
	}

	w.Write([]byte(operatorName))
}

func SetOperatorName(w http.ResponseWriter, r *http.Request) {
	var opRequest commands.SetOperatorNameRequest
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		logger.Error.Println(err)
	}

	json.Unmarshal(reqBody, &opRequest)

	response, err := commands.SetOperatorName(opRequest)
	if err != nil {
		fmt.Println(err)
		logger.Error.Println(err)
		http.Error(w, err.Error(), 400)
		return
	}

	w.Write([]byte(response))
}

func GetOperatorPassword(w http.ResponseWriter, r *http.Request) {
	operatorPassword, err := commands.GetOperatorPassword()
	if err != nil {
		fmt.Println(err)
		logger.Error.Println(err)
		http.Error(w, err.Error(), 400)
		return
	}

	w.Write([]byte(operatorPassword))
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
		fmt.Println(err)
		logger.Error.Println(err)
		http.Error(w, err.Error(), 400)
		return
	}

	w.Write([]byte("Success"))
}
