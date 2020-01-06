package handlers

import (
	"github.com/chyupa/fp700/commands"
	"net/http"
)

func PrintDiagnostic(w http.ResponseWriter, r *http.Request) {
	commands.PrintDiagnostic()

	w.Write([]byte("success"))
}
