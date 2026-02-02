package main

import (
	"broker-service/internal/server"
	"log"
)

func main() {
	app := server.App{}

	err := app.ConnectToRabbitMQ()

	if err != nil {
		log.Panic(err)
	}

	defer app.RabbitConn.Close()

	err = app.InitializeServer()

	if err != nil {
		log.Panic(err)
	}
}
