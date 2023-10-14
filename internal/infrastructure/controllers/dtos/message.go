package dtos

type Message struct {
	Message string `json:"message"`
}

func NewMessage(content string) *Message {
	return &Message{content}
}
