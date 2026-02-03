package main

import (
	"log"
	"mailer-service/internal/server"
)

func main() {
	go server.StartRPCServer()

	err := server.StartHTTPServer()
	if err != nil {
		log.Panic(err)
	}
}
