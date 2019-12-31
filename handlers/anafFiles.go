package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/chyupa/fp700"
	"github.com/chyupa/fp700/commands"
	"io/ioutil"
	"log"
	"net/http"
)

func AnafFiles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	var anafFilesRequest commands.AnafFilesRequest

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(reqBody, &anafFilesRequest)

	fp700.Port = r.Header.Get("ComPort")
	response, e := commands.AnafFiles(anafFilesRequest)

	fmt.Println(response, e)

	convJson, _  := json.Marshal(response)
	w.Write(convJson)
}
