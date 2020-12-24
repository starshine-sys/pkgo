package pkgo

// Session is the PluralKit API session, including a token
type Session struct {
	authorized bool
	token      string
	system     string
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
