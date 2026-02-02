package server

import amqp "github.com/rabbitmq/amqp091-go"

type App struct {
	RabbitConn *amqp.Connection
}
