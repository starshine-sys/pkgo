package pkgo

// Session is the PluralKit API session, including a token
type Session struct {
	Authorized bool
	Token      string
	system     string
}

// NewSession returns an empty, unauthorized session
func NewSession() *Session {
	return &Session{Authorized: false}
}

// NewSessionWithToken returns a session with a token
func NewSessionWithToken(token string) *Session {
	return &Session{Authorized: true, Token: token}
}
