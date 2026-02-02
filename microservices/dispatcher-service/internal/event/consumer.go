package event

import (
	"encoding/json"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Consumer struct {
	conn         *amqp.Connection
	exchangeName string
	queueName    string
}

func (c *Consumer) Listen(topics ...string) error {
	channel, err := c.conn.Channel()
	if err != nil {
		return err
	}

	defer channel.Close()

	for _, topic := range topics {
		err = channel.QueueBind(c.queueName, topic, c.exchangeName, false, nil)

		if err != nil {
			return err
		}
	}

	deliveries, err := channel.Consume(c.queueName, "", false, false, false, false, nil)

	if err != nil {
		return err
	}

	emptyChannel := make(chan bool)

	go func() {
		for message := range deliveries {
			var payload Payload
			err = json.Unmarshal(message.Body, &payload)

			if err != nil {
				log.Println("Error unmarshalling message body:", err)
				message.Nack(false, true)
				continue
			}

			message.Ack(false)
			payload.RoutingKey = message.RoutingKey
			log.Println("Received message with topic:", message.RoutingKey)
			go payload.Handle()
		}
	}()

	log.Printf("Dispatcher is listening for RabbitMQ messages in a %s queue\n", c.queueName)

	<-emptyChannel

	return nil
}

func (c *Consumer) setup() error {
	channel, err := c.conn.Channel()
	if err != nil {
		return err
	}

	defer channel.Close()

	err = channel.ExchangeDeclare(
		c.exchangeName,
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

	_, err = channel.QueueDeclare(
		c.queueName,
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

func NewConsumer(conn *amqp.Connection) (*Consumer, error) {
	consumer := Consumer{conn: conn, queueName: "microservices_queue", exchangeName: "microservices_topics"}

	err := consumer.setup()
	if err != nil {
		return &Consumer{}, err
	}

	return &consumer, nil
}
