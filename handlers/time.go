package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/chyupa/apiServer/utils/logger"
	"github.com/chyupa/fp700/commands"
	"io/ioutil"
	"net/http"
)

func Time(w http.ResponseWriter, r *http.Request) {
	response, err := commands.Time()
	if err != nil {
		fmt.Println(err)
		logger.Error.Println(err)
		http.Error(w, err.Error(), 400)
	}
	json.NewEncoder(w).Encode(response)
}

func SetTime(w http.ResponseWriter, r *http.Request) {
	var timeRequest commands.SetTimeRequest

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		logger.Error.Println(err)
		http.Error(w, err.Error(), 400)
	}

	json.Unmarshal(reqBody, &timeRequest)

	errorCode := commands.SetTime(timeRequest)

	w.Write([]byte(errorCode))
}
