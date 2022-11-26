package main

import (
	"fmt"
	"os"

	"github.com/danmory/messaging/producer/messages"
	"github.com/danmory/messaging/producer/utils"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Config struct {
	QueueHost string
	QueuePort string
	QueueUser string
	QueuePass string
}

func main() {
	cfg := getConfig()

	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s/", cfg.QueueUser, cfg.QueuePass, cfg.QueueHost, cfg.QueuePort))
	utils.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	utils.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	messages.Produce(ch)

	<-utils.Forever
}

func getConfig() Config {
	return Config{
		QueueHost: os.Getenv("QUEUE_HOST"),
		QueuePort: os.Getenv("QUEUE_PORT"),
		QueueUser: os.Getenv("QUEUE_USER"),
		QueuePass: os.Getenv("QUEUE_PASS"),
	}
}
