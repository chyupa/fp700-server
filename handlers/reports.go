package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/chyupa/fp700-server/utils/logger"
	"github.com/chyupa/fp700/commands"
	"io/ioutil"
	"log"
	"net/http"
)

func PrintEj(w http.ResponseWriter, r *http.Request) {
	var ejRequest commands.PrintEjRequest

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(reqBody, &ejRequest)

	err = commands.PrintEj(ejRequest)
	if err != nil {
		fmt.Println(err)
		logger.Error.Println(err)
		http.Error(w, err.Error(), 400)
		return
	}

	w.Write([]byte("success"))
}


func PrintMf(w http.ResponseWriter, r *http.Request) {
	var mfRequest commands.PrintMfRequest
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		logger.Error.Println(err)
	}

	json.Unmarshal(reqBody, &mfRequest)

	err = commands.PrintMf(mfRequest)
	if err != nil {
		fmt.Println(err)
		logger.Error.Println(err)
		http.Error(w, err.Error(), 400)
		return
	}

	w.Write([]byte("success"))
}

type printReportRequest struct {
	Type string `json:"type"`
}

func PrintReport(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		logger.Error.Println(err)
	}
	var prRequest printReportRequest
	json.Unmarshal(reqBody, &prRequest)

	response, err := commands.PrintReport(prRequest.Type)
	if err != nil {
		fmt.Println(err)
		logger.Error.Println(err)
		http.Error(w, err.Error(), 400)
		return
	}

	json.NewEncoder(w).Encode(response)
}

func ReportEj(w http.ResponseWriter, r *http.Request) {
	var ejRequest commands.EjRequest

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		logger.Error.Println(err)
	}

	json.Unmarshal(reqBody, &ejRequest)

	response, err := commands.ReadEj(ejRequest)
	if err != nil {
		fmt.Println(err)
		logger.Error.Println(err)
		http.Error(w, err.Error(), 400)
		return
	}

	w.Write([]byte(response))
}

func ReportMf(w http.ResponseWriter, r *http.Request) {
	var mfRequest commands.MfRequest
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		logger.Error.Println(err)
	}

	json.Unmarshal(reqBody, &mfRequest)

	commandResponse, err := commands.ReadMf(mfRequest)
	if err != nil {
		fmt.Println(err)
		logger.Error.Println(err)
		http.Error(w, err.Error(), 400)
		return
	}

	w.Write([]byte(commandResponse))
}


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


func GetRemainingZReports(w http.ResponseWriter, r *http.Request) {
	remainingZReports, err := commands.RemainingZReports()
	if err != nil {
		fmt.Println(err)
		logger.Error.Println(err)
		http.Error(w, err.Error(), 400)
		return
	}

	json.NewEncoder(w).Encode(remainingZReports)
}
