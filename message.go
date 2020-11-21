package pkgo

import "time"

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
