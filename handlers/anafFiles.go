package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/chyupa/fp700/commands"
	"io/ioutil"
	"log"
	"net/http"
)

func AnafFiles(w http.ResponseWriter, r *http.Request) {
	var anafFilesRequest commands.AnafFilesRequest

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(reqBody, &anafFilesRequest)

	response, e := commands.AnafFiles(anafFilesRequest)

	fmt.Println(response, e)

	convJson, _ := json.Marshal(response)
	w.Write(convJson)
}
