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

func InitDisplay(w http.ResponseWriter, r *http.Request) {
	var idRequest commands.InitDisplayRequest

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		logger.Error.Println(err)
		http.Error(w, err.Error(), 400)
		return
	}

	json.Unmarshal(reqBody, &idRequest)

	commands.InitDisplay(idRequest)

	w.Write([]byte("Success"))

}

func MaintenanceData(w http.ResponseWriter, r *http.Request) {
	commandResponse, err := commands.MaintenanceData()
	if err != nil {
		fmt.Println(err)
		logger.Error.Println(err)
		http.Error(w, err.Error(), 400)
		return
	}

	json.NewEncoder(w).Encode(commandResponse)
}

func GetFiscalization(w http.ResponseWriter, r *http.Request) {
	response, err := commands.GetFiscalization()
	if err != nil {
		fmt.Println(err)
		logger.Error.Println(err)
		http.Error(w, err.Error(), 400)
		return
	}

	w.Write([]byte(response))
}

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

func PrintDiagnostic(w http.ResponseWriter, r *http.Request) {
	commands.PrintDiagnostic()

	w.Write([]byte("success"))
}

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

func CloseDmje(w http.ResponseWriter, r *http.Request) {
	var cdRequest commands.CloseDmjeRequest

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(reqBody, &cdRequest)

	// get response and send it
	commands.CloseDmje(cdRequest)

	w.Write([]byte("Success"))
}

func FabricationNumber(w http.ResponseWriter, r *http.Request) {
	commands.FabricationNumber()

	w.Write([]byte("Success"))
}

func Fiscalize(w http.ResponseWriter, r *http.Request) {
	var fiscalizeRequest commands.FiscalizeRequest

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		logger.Error.Println(err)
		http.Error(w, err.Error(), 400)
		return
	}

	json.Unmarshal(reqBody, &fiscalizeRequest)

	fiscalizeResponse, err := commands.Fiscalize(fiscalizeRequest)
	if err != nil {
		fmt.Println(err)
		logger.Error.Println(err)
		http.Error(w, err.Error(), 400)
		return
	}

	json.NewEncoder(w).Encode(fiscalizeResponse)
}

func GetVat(w http.ResponseWriter, r *http.Request) {
	vat, err := commands.GetVat()
	if err != nil {
		fmt.Println(err)
		logger.Error.Println(err)
		http.Error(w, err.Error(), 400)
		return
	}

	w.Write([]byte(vat))
}

func SetVat(w http.ResponseWriter, r *http.Request) {
	var setVatRequest commands.SetVatRequest

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		logger.Error.Println(err)
	}

	json.Unmarshal(reqBody, &setVatRequest)

	setVatResponse, err := commands.SetVat(setVatRequest)
	if err != nil {
		fmt.Println(err)
		logger.Error.Println(err)
		http.Error(w, err.Error(), 400)
		return
	}

	json.NewEncoder(w).Encode(setVatResponse)
}

func GetVatChanges(w http.ResponseWriter, r *http.Request) {
	vat, err := commands.GetVatChanges()
	if err != nil {
		fmt.Println(err)
		logger.Error.Println(err)
		http.Error(w, err.Error(), 400)
		return
	}

	json.NewEncoder(w).Encode(vat)
}

func ChangeServicePassword(w http.ResponseWriter, r *http.Request) {
	var cspReq commands.ChangeServicePasswordRequest

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		logger.Error.Println(err)
	}

	json.Unmarshal(reqBody, &cspReq)

	err = commands.ChangeServicePassword(cspReq)
	if err != nil {
		fmt.Println(err)
		logger.Error.Println(err)
		http.Error(w, err.Error(), 400)
		return
	}

	w.Write([]byte("success"))
}

func ActivateServicePassword(w http.ResponseWriter, r *http.Request) {
	var aspReq commands.ActivateServicePasswordRequest

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		logger.Error.Println(err)
	}

	json.Unmarshal(reqBody, &aspReq)

	err = commands.ActivateServicePassword(aspReq)
	if err != nil {
		fmt.Println(err)
		logger.Error.Println(err)
		http.Error(w, err.Error(), 400)
		return
	}

	w.Write([]byte("success"))
}

func SetPrinterMode(w http.ResponseWriter, r *http.Request) {
	type setPrinterModeRequest struct {
		Mode string `json:"mode"`
	}

	var spmReq setPrinterModeRequest

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		logger.Error.Println(err)
		http.Error(w, err.Error(), 400)
		return
	}

	json.Unmarshal(reqBody, spmReq)

	response, err := commands.SetPrinterMode(spmReq.Mode)
	if err != nil {
		fmt.Println(err)
		logger.Error.Println(err)
		http.Error(w, err.Error(), 400)
		return
	}

	json.NewEncoder(w).Encode(response)
}

func GetLastZDate(w http.ResponseWriter, r *http.Request) {

	lzDate, err := commands.GetLastZDate()
	if err != nil {
		fmt.Println(err)
		logger.Error.Println(err)
		http.Error(w, err.Error(), 400)
		return
	}

	json.NewEncoder(w).Encode(lzDate)
}
