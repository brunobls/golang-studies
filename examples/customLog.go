package main

import (
	"log"
	"os"
	"path"
)

func main() {
	LOGFILE := path.Join(os.TempDir(), "customLog.log")
	file, err := os.OpenFile(LOGFILE, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err == nil {
		defer file.Close()
		iLog := log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
		iLog.Println("Hello World!")
	}
}
