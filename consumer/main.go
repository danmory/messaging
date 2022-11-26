package main

import (
	"fmt"
	"os"

	messagesConsumer "github.com/danmory/messaging/consumer/messages/consumer"
	messagesRepository "github.com/danmory/messaging/consumer/messages/repository"

	"github.com/danmory/messaging/consumer/models"
	"github.com/danmory/messaging/consumer/utils"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	cfg := getConfig()

	db := new(messagesRepository.Repository)
	err := db.Init(&cfg)
	utils.FailOnError(err, "Failed to connect to database")
	defer db.Close()

	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s/", cfg.QueueUser, cfg.QueuePass, cfg.QueueHost, cfg.QueuePort))
	utils.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	utils.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	msgConsumer := &messagesConsumer.Consumer{
		Repository: db,
	}
	msgConsumer.Subscribe(ch)

	<-utils.Forever
}

func getConfig() models.Config {
	return models.Config{
		QueueHost: os.Getenv("QUEUE_HOST"),
		QueuePort: os.Getenv("QUEUE_PORT"),
		QueueUser: os.Getenv("QUEUE_USER"),
		QueuePass: os.Getenv("QUEUE_PASS"),
		DBHost:    os.Getenv("DB_HOST"),
		DBPort:    os.Getenv("DB_PORT"),
		DBUser:    os.Getenv("DB_USER"),
		DBPass:    os.Getenv("DB_PASS"),
		DBName:    os.Getenv("DB_NAME"),
	}
}
