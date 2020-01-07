package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/chyupa/apiServer/utils/logger"
	"github.com/chyupa/fp700/commands"
	"io/ioutil"
	"net/http"
)

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
