package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/chyupa/apiServer/utils/logger"
	"github.com/chyupa/fp700/commands"
	"io/ioutil"
	"net/http"
)

func SetFootersData(w http.ResponseWriter, r *http.Request) {
	var footerRequest commands.SetFootersDataRequest
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		logger.Error.Println(err)
		return
	}

	json.Unmarshal(reqBody, &footerRequest)

	err = commands.SetFootersData(footerRequest)
	if err != nil {
		fmt.Println(err)
		logger.Error.Println(err)
		http.Error(w, err.Error(), 400)
		return
	}

	w.Write([]byte("Success"))
}

func GetFootersData(w http.ResponseWriter, r *http.Request) {
	footerResponse, err := commands.FootersData()
	if err != nil {
		fmt.Println(err)
		logger.Error.Println(err)
		http.Error(w, err.Error(), 400)
		return
	}

	json.NewEncoder(w).Encode(footerResponse)
}
