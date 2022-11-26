package messages

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/danmory/messaging/producer/models"
	"github.com/danmory/messaging/producer/utils"
	amqp "github.com/rabbitmq/amqp091-go"
)

func Produce(ch *amqp.Channel) {
	q, err := declareQueue(ch)
	utils.FailOnError(err, "Failed to declare a queue")
	go func() {
		ticker := time.Tick(1 * time.Second)
		for range ticker {
			msg, err := getMessage()
			if err != nil {
				log.Printf("Failed to get message: %s", err)
				continue
			}
			if err := ch.PublishWithContext(
				context.Background(),
				"",     // exchange
				q.Name, // routing key
				false,  // mandatory
				false,  // immediate
				amqp.Publishing{
					ContentType: "text/plain",
					Body:        msg,
				}); err != nil {
				log.Printf("Failed to publish a message: %s", err)
			}
			log.Printf(" [x] Sent %s\n", msg)
		}
	}()

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

func getMessage() ([]byte, error) {
	msg := models.Message{
		Text: fmt.Sprintf("Hello World at %s", time.Now()),
		Table: uint8(rand.Uint32() % 2),
	}
	b, err := json.Marshal(msg)
	return b, err
}
