package pkgo

import (
	"time"

	"golang.org/x/time/rate"
)

// Session is the PluralKit API session, including a token
type Session struct {
	authorized bool
	token      string
	system     *System

	rate *rate.Limiter

	// Timeout is the maximum time this Session will wait for requests.
	Timeout time.Duration
}

// New returns a session with the given token, or no token if the string is empty.
func New(token string) *Session {
	if token != "" {
		return &Session{
			authorized: true,
			token:      token,
			rate:       rate.NewLimiter(2, 2),
			Timeout:    10 * time.Second,
		}
	}

	return &Session{
		authorized: false,
		rate:       rate.NewLimiter(2, 2),
		Timeout:    10 * time.Second,
	}
}
