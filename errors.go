package pkgo

import "fmt"

// ErrNoToken is returned when an endpoint requiring authentication is hit, but no token is given
type ErrNoToken struct{}

func (e *ErrNoToken) Error() string {
	return "no token in session, can't hit endpoints requiring authentication"
}

// ErrStatusNot200 ...
type ErrStatusNot200 struct {
	Code   int
	Status string
}

func (e *ErrStatusNot200) Error() string {
	return fmt.Sprintf("http status code %v: %v", e.Code, e.Status)
}

// ErrInvalidID is returned when a function requires a 5-letter ID but one isn't given
type ErrInvalidID struct {
	givenID string
}

func (e *ErrInvalidID) Error() string {
	return fmt.Sprintf("5-letter, lowercase ID expected; %v given", e.givenID)
}

// ErrInvalidSnowflake is returned when a function expects a Discord snowflake but one isn't given
type ErrInvalidSnowflake struct {
	givenID string
}

func (e *ErrInvalidSnowflake) Error() string {
	return fmt.Sprintf("discord snowflake expected; %v given", e.givenID)
}

// ErrMsgNotFound is returned when a message isn't found
type ErrMsgNotFound struct {
	ID string
}

func (e *ErrMsgNotFound) Error() string {
	return fmt.Sprintf("message with ID %v not found", e.ID)
}
