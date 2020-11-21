package pkgo

import (
	"time"
)

// Message is a proxied message
type Message struct {
	Timestamp time.Time `json:"timestamp"`
	ID        string    `json:"id"`
	Original  string    `json:"original"`
	Sender    string    `json:"sender"`
	Channel   string    `json:"channel"`
	System    System    `json:"system"`
	Member    Member    `json:"member"`
}

// GetMessage gets a message by Discord snowflake
func (s *Session) GetMessage(id string) (m *Message, err error) {
	if !discordIDre.MatchString(id) {
		return m, &ErrInvalidSnowflake{id}
	}
	err = s.GetEndpoint("/msg/"+id, &m)
	if err != nil {
		switch err.(type) {
		case *ErrStatusNot200:
			return m, &ErrMsgNotFound{ID: id}
		default:
			return m, err
		}
	}
	return m, err
}
