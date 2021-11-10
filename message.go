package pkgo

import (
	"time"
)

// Message is a proxied message
type Message struct {
	Timestamp time.Time `json:"timestamp"`

	ID       Snowflake `json:"id"`
	Original Snowflake `json:"original"`
	Sender   Snowflake `json:"sender"`
	Channel  Snowflake `json:"channel"`

	System System `json:"system"`

	Member Member `json:"member"`
}

// Message gets a message by Discord snowflake.
func (s *Session) Message(id Snowflake) (m *Message, err error) {
	m = &Message{}
	err = s.RequestJSON("GET", "/messages/"+id.String(), m)
	return
}
