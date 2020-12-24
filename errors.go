package pkgo

import (
	"errors"
	"fmt"
)

// Errors
var (
	ErrNoToken          = errors.New("pkgo: no token in session, can't hit endpoints requiring authentication")
	ErrInvalidID        = errors.New("pkgo: not a 5-character ID")
	ErrInvalidSnowflake = errors.New("pkgo: not a valid Discord snowflake")
	ErrMsgNotFound      = errors.New("pkgo: message not found")
)

// ErrStatusNot200 ...
type ErrStatusNot200 struct {
	Code   int
	Status string
}

func (e *ErrStatusNot200) Error() string {
	return fmt.Sprintf("http status code %v: %v", e.Code, e.Status)
}
