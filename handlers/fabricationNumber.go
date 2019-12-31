package handlers

import (
	"github.com/chyupa/fp700/commands"
	"net/http"
)

func FabricationNumber(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	commands.FabricationNumber()

	w.Write([]byte("Success"))
}
