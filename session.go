package pkgo

import (
	"sync"
	"time"
)

// Session is the PluralKit API session, including a token
type Session struct {
	authorized bool
	token      string
	system     string

	rate sync.Mutex
}

// Config is the config struct, passed to (Session).NewToken()
type Config struct {
	Token string
}

// NewSession returns a session
func NewSession(c *Config) *Session {
	if c != nil {
		if c.Token != "" {
			return &Session{authorized: true, token: c.Token}
		}
	}
	return &Session{authorized: false}
}

// RateLimit blocks until we can be *sure* we won't hit the rate limit.
// Gets a lock on s.rate, waits 500ms, and unlocks it.
func (s *Session) RateLimit() {
	s.rate.Lock()
	time.Sleep(500 * time.Millisecond)
	s.rate.Unlock()
}
