package messages

import (
	"encoding/json"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/danmory/messaging/consumer/messages"
	"github.com/danmory/messaging/consumer/models"
	"github.com/danmory/messaging/consumer/utils"
)

type Consumer struct {
	Repository messages.Repository
}

func (c *Consumer) Subscribe(ch *amqp.Channel) {
	q, err := declareQueue(ch)
	utils.FailOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	utils.FailOnError(err, "Failed to register a consumer")

	go func() {
		for d := range msgs {
			c.processMessage(d)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
}

func declareQueue(ch *amqp.Channel) (*amqp.Queue, error) {
	q, err := ch.QueueDeclare(
		"messages", // name
		false,      // durable
		false,      // delete when unused
		false,      // exclusive
		false,      // no-wait
		nil,        // arguments
	)
	return &q, err
}

func (c *Consumer) processMessage(msg amqp.Delivery) {
	message := new(models.Message)
	if err := json.Unmarshal(msg.Body, message); err != nil {
		log.Printf("Failed to unmarshal message: %s", msg.Body)
		return
	}
	log.Printf("Received a message: %s", message)
	
	if err := c.Repository.Save(message); err != nil {
		log.Printf("Failed to save message: %s", err)
		return
	}
}
