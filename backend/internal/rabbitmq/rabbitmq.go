package rabbitmq

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
	"log"
	"os"
)

func SetupRabbitMQ() (*amqp.Connection, *amqp.Channel, error) {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	ampqUrl := os.Getenv("AMPQ_URL")
	conn, err := amqp.Dial(ampqUrl)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to connect to RabbitMQ: %w", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to open a channel: %w", err)
	}
	return conn, ch, nil
}

func ConsumeMessages(ch *amqp.Channel, qname string) (<-chan amqp.Delivery, error) {
	msgs, err := ch.Consume(
		qname,
		qname,
		false,
		false,
		false,
		false,
		nil)
	if err != nil {
		return nil, fmt.Errorf("failed to register a consumer: %w", err)
	}
	return msgs, nil
}

func PublishMessages(ch *amqp.Channel, qname string, msg string) error {
	err := ch.Publish(
		"",
		qname,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(msg),
		})
	if err != nil {
		return fmt.Errorf("failed to publish a message: %w", err)
	}
	return nil
}
