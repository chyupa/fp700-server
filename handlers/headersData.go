package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/chyupa/apiServer/utils/logger"
	"github.com/chyupa/fp700/commands"
	"io/ioutil"
	"net/http"
)

func HeadersData(w http.ResponseWriter, r *http.Request) {
	commandResponse := commands.HeadersData()

	json.NewEncoder(w).Encode(commandResponse)
}

func SetHeadersData(w http.ResponseWriter, r *http.Request) {
	var hdrRequest commands.SetHeadersDataRequest

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		logger.Error.Println(err)
	}

	json.Unmarshal(reqBody, &hdrRequest)

	commandResponse, err := commands.SetHeadersData(hdrRequest)
	if err != nil {
		fmt.Println(err)
		logger.Error.Println(err)
		http.Error(w, err.Error(), 400)
		return
	}

	json.NewEncoder(w).Encode(commandResponse)
}

func GetHeaderChanges(w http.ResponseWriter, r *http.Request) {
	commandResponse, err := commands.GetHeaderChanges()
	if err != nil {
		fmt.Println(err)
		logger.Error.Println(err)
		http.Error(w, err.Error(), 400)
		return
	}

	json.NewEncoder(w).Encode(commandResponse)
}
