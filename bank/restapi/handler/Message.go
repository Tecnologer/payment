package handler

import (
	"encoding/json"
	"fmt"
)

type Message struct {
	Message string `json:"message"`
}

func NewMessage(message string) *Message {
	return &Message{Message: message}
}

func MessageBytes(message string) []byte {
	data, _ := json.Marshal(NewMessage(message))

	return data
}

func MessageError(msg string, err error) []byte {
	return MessageBytes(fmt.Sprintf("%s: %s", msg, err.Error()))
}

func (m *Message) Bytes() []byte {
	data, _ := json.Marshal(m)

	return data
}
