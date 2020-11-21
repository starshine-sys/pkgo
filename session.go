package pkgo

// Session is the PluralKit API session, including a token
type Session struct {
	Authorized bool
	Token      string
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
			return &Session{Authorized: true, Token: c.Token}
		}
	}
	return &Session{Authorized: false}
}
