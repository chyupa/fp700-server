package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/chyupa/apiServer/utils/logger"
	"github.com/chyupa/fp700/commands"
	"io/ioutil"
	"net/http"
)

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
