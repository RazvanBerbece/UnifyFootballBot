package logger

import (
	"log"
	"os"
)

func LogHandlerCall(handlerName string) {

	// If the file doesn't exist, create it or append to the file
	file, err := os.OpenFile("logs/handler_calls.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)

	log.Printf("%s", handlerName)

}
