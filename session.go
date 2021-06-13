package pkgo

import (
	"sync"
	"time"
)

// Session is the PluralKit API session, including a token
type Session struct {
	authorized bool
	token      string
	system     *System

	// NoRate disables client-side rate limiting
	NoRate bool

	rate sync.Mutex
}

// Config is the config struct, passed to (Session).NewConfig()
type Config struct {
	Token  string
	NoRate bool
}

// New returns a session with the given token, or no token if the string is empty.
func New(token string) *Session {
	if token != "" {
		return &Session{authorized: true, token: token}
	}

	return &Session{authorized: false}
}

// NewConfig returns a session
func NewConfig(c *Config) *Session {
	if c != nil {
		if c.Token != "" {
			return &Session{authorized: true, token: c.Token, NoRate: c.NoRate}
		}
		return &Session{authorized: false, NoRate: c.NoRate}
	}
	return &Session{authorized: false, NoRate: false}
}

// RateLimit blocks until we can be *sure* we won't hit the rate limit.
// Gets a lock on s.rate, waits 500ms, and unlocks it.
func (s *Session) RateLimit() {
	if s.NoRate {
		return
	}

	s.rate.Lock()
	time.Sleep(500 * time.Millisecond)
	s.rate.Unlock()
}
