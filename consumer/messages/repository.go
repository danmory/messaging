package messages

import "github.com/danmory/messaging/consumer/models"

type Repository interface {
	Init(cfg *models.Config) error
	Save(message *models.Message) error
	Close() error
}
