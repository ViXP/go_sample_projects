package main

import (
	"fmt"
	"log"
	"mailer-service/internal/server"
	"net/http"
)

const serverPort = 80

func main() {
	fmt.Printf("Starting the Mailer service on the port #%v...\n", serverPort)
	err := http.ListenAndServe(fmt.Sprintf(":%v", serverPort), server.Routes())

	if err != nil {
		log.Panic(err)
	}
}
