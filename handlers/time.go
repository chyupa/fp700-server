package handlers

import (
	"encoding/json"
	"github.com/chyupa/fp700"
	"github.com/chyupa/fp700/commands"
	"log"
	"net/http"
)

func Time(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	fp700.Port = r.Header.Get("ComPort")

	response, err := commands.Time()
	if err != nil {
		log.Println(err)
		w.Write([]byte(err.Error()))
	}
	json.NewEncoder(w).Encode(response)
}
