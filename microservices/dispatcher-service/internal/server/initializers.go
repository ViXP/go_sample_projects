package server

import (
	"errors"
	"log"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func ConnectToRabbitMQ() (*amqp.Connection, error) {
	var connection *amqp.Connection
	countdown := 10
	rabbitUrl := os.Getenv("RABBITMQ_URL")

	log.Println("Connecting to RabbitMQ...")

	for countdown > 0 {
		c, err := amqp.Dial(rabbitUrl)

		if err == nil {
			connection = c
			break
		}

		log.Println("RabbitMQ is not yer ready.")
		time.Sleep(2 * time.Second)
		countdown--
	}

	if countdown == 0 {
		return nil, errors.New("could not connect to RabbitMQ")
	}

	return connection, nil
}
