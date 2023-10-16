package logger

import (
	"log"
	"os"
	"path/filepath"
)

func LogHandlerCall(handlerName string, outputDirectory string, filename string) {

	// If the file and/or path doesn't exist, create, or append to the file
	pathToLogfile := filepath.Join(".", outputDirectory)
	err := os.MkdirAll(pathToLogfile, os.ModePerm)
	if err != nil {
		log.Fatal("Could not create filepath for Handler calls log file. Err =", err)
	}
	file, err := os.OpenFile(outputDirectory+filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal("Could not open log file for Handler calls. Err =", err)
	}

	log.SetOutput(file)

	log.Printf("%s", handlerName)

}
