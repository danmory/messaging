package models

import "fmt"

type Message struct {
	Text  string `json:"text"`
	Table uint8  `json:"table"`
}

func (m *Message) String() string {
	return fmt.Sprintf("Message: %s, Table: %d", m.Text, m.Table)
}
