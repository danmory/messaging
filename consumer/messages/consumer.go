package messages

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

type Consumer interface {
	Subscribe(ch *amqp.Channel)
}
