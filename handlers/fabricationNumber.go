package handlers

import (
	"github.com/chyupa/fp700/commands"
	"net/http"
)

func FabricationNumber(w http.ResponseWriter, r *http.Request) {
	commands.FabricationNumber()

	w.Write([]byte("Success"))
}
