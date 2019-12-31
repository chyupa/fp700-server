package handlers

import (
	"github.com/chyupa/fp700"
	"github.com/chyupa/fp700/commands"
	"net/http"
)

func PrintDiagnostic(w http.ResponseWriter, r *http.Request) {

	fp700.Port = r.Header.Get("ComPort")
	commands.PrintDiagnostic()

	w.Write([]byte("success"))
}
