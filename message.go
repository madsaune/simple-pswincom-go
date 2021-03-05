package pswincom

import (
	"fmt"
	"strings"
)

// Message manages the message content
type Message struct {
	Recipient string
	Body      string
	Replace   bool
}

// NewMessage creates a new message
func NewMessage(recipient, message string, replace bool) *Message {
	m := &Message{
		Recipient: recipient,
		Body:      message,
		Replace:   replace,
	}

	m.format()

	return m
}

func (m *Message) format() {
	if strings.Index(m.Recipient, "+") == 0 {
		m.Recipient = string(m.Recipient[1:])
		return
	}

	if strings.Index(m.Recipient, "00") == 0 {
		m.Recipient = string(m.Recipient[2:])
		return
	}

	if len(m.Recipient) == 8 {
		m.Recipient = fmt.Sprintf("47%s", m.Recipient)
		return
	}
}
