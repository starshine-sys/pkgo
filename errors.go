package pkgo

import "fmt"

// ErrNoToken is returned when an endpoint requiring authentication is hit, but no token is given
type ErrNoToken struct{}

func (e *ErrNoToken) Error() string {
	return "no token in session, can't hit endpoints requiring authentication"
}

// ErrStatusNot200 ...
type ErrStatusNot200 struct {
	Code int
}

func (e *ErrStatusNot200) Error() string {
	return fmt.Sprintf("http status code %v", e.Code)
}
