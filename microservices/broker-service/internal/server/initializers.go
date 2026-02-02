package server

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

const port = 80

func (app *App) InitializeServer() error {
	log.Printf("Starting the Broker Service on port :%v", port)
	return http.ListenAndServe(fmt.Sprintf(":%v", port), app.Routes())
}

func (app *App) ConnectToRabbitMQ() error {
	countdown := 10
	rabbitUrl := os.Getenv("RABBITMQ_URL")

	log.Println("Connecting to RabbitMQ...")

	for countdown > 0 {
		c, err := amqp.Dial(rabbitUrl)

		if err == nil {
			app.RabbitConn = c
			break
		}

		log.Println("RabbitMQ is not yer ready.")
		time.Sleep(2 * time.Second)
		countdown--
	}

	if countdown == 0 {
		return errors.New("could not connect to RabbitMQ")
	}

	return nil
}
