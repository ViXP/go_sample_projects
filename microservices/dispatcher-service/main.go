package main

import (
	"dispatcher-service/internal/event"
	"dispatcher-service/internal/server"
	"log"
)

func main() {
	connection, err := server.ConnectToRabbitMQ()

	if err != nil {
		log.Panic(err)
	}

	consumer, err := event.NewConsumer(connection)

	if err != nil {
		log.Panic(err)
	}

	err = consumer.Listen("log.INFO", "auth.INFO", "mail.INFO")

	if err != nil {
		log.Panic(err)
	}

	defer connection.Close()
}
