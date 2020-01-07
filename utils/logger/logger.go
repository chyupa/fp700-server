package logger

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"path/filepath"
)

const (
	LogDir = "/logs"
)

var General *log.Logger
var Error *log.Logger

func init() {
	myself, err := user.Current()
	if err != nil {
		Error.Fatalln("Could not load current user")
	}

	absPath, err := filepath.Abs(myself.HomeDir + LogDir)
	if err != nil {
		fmt.Println(err)
	}

	err = os.MkdirAll(absPath, 0755)
	if err != nil {
		fmt.Println(err)
	}

	generalLog, err := os.OpenFile(absPath+"/log.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	General = log.New(generalLog, "General: ", log.LstdFlags|log.Llongfile)
	Error = log.New(generalLog, "Error: ", log.LstdFlags|log.Llongfile)
}
