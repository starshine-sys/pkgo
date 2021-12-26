package pkgo

import (
	"time"
)

// Message is a proxied message.
// System and Member may be null if the system/member the message is connected to is deleted.
type Message struct {
	Timestamp time.Time `json:"timestamp"`

	ID       Snowflake `json:"id"`
	Original Snowflake `json:"original"`
	Sender   Snowflake `json:"sender"`
	Channel  Snowflake `json:"channel"`

	System *System `json:"system"`
	Member *Member `json:"member"`
}

// Message gets a message by Discord snowflake.
func (s *Session) Message(id Snowflake) (m Message, err error) {
	err = s.RequestJSON("GET", "/messages/"+id.String(), &m)
	return
}
