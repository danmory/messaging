package models

type Message struct {
	Table uint8  `json:"table"`
	Text  string `json:"text"`
}
