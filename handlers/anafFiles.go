package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/chyupa/apiServer/utils/logger"
	"github.com/chyupa/fp700/commands"
	"io/ioutil"
	"net/http"
)

func AnafFiles(w http.ResponseWriter, r *http.Request) {
	var anafFilesRequest commands.AnafFilesRequest

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		logger.Error.Println(err)
	}

	json.Unmarshal(reqBody, &anafFilesRequest)

	response, err := commands.AnafFiles(anafFilesRequest)
	if err != nil {
		fmt.Println(err)
		logger.Error.Println(err)
		http.Error(w, err.Error(), 400)
		return
	}

	//convJson, _ := json.Marshal(response)
	//w.Write(convJson)
	json.NewEncoder(w).Encode(response)
}
