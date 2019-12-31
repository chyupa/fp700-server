package handlers

import (
	"github.com/chyupa/fp700"
	"github.com/chyupa/fp700/commands"
	"net/http"
)

func GetOperatorName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	fp700.Port = r.Header.Get("ComPort")

	operatorName, err := commands.GetOperatorName()
	if err != nil {
		w.WriteHeader(400)
	} else {
		w.Write([]byte(operatorName))
	}
}
