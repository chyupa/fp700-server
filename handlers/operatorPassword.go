package handlers

import (
	"github.com/chyupa/fp700"
	"github.com/chyupa/fp700/commands"
	"net/http"
)

func GetOperatorPassword(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	fp700.Port = r.Header.Get("ComPort")

	operatorPassword, err := commands.GetOperatorPassword()
	if err != nil {
		w.WriteHeader(400)
	} else {
		w.Write([]byte(operatorPassword))
	}
}
