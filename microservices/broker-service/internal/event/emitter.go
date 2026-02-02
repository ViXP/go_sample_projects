package event

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Emitter struct {
	conn         *amqp.Connection
	exchangeName string
}

func (e *Emitter) setup() error {
	channel, err := e.conn.Channel()
	if err != nil {
		return err
	}

	defer channel.Close()
	err = channel.ExchangeDeclare(
		e.exchangeName,
		"topic",
		true,
		false,
		false,
		true,
		nil,
	)

	if err != nil {
		return err
	}

	return nil
}

func (e *Emitter) Emit(event string, severity string) error {
	channel, err := e.conn.Channel()
	if err != nil {
		return err
	}

	defer channel.Close()

	log.Println("Publishing event to RabbitMQ")

	return channel.Publish(e.exchangeName, severity, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(event),
	})
}

func NewEmitter(conn *amqp.Connection, exchangeName string) (*Emitter, error) {
	emitter := Emitter{conn: conn, exchangeName: exchangeName}

	err := emitter.setup()

	if err != nil {
		return &Emitter{}, err
	}

	return &emitter, nil
}
