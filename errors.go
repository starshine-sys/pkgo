package pkgo

import (
	"fmt"

	"emperror.dev/errors"
)

// Errors
const (
	ErrNoToken          = errors.Sentinel("pkgo: no token in session, can't hit endpoints requiring authentication")
	ErrInvalidID        = errors.Sentinel("pkgo: not a 5-character ID")
	ErrInvalidSnowflake = errors.Sentinel("pkgo: not a valid Discord snowflake")
	ErrMsgNotFound      = errors.Sentinel("pkgo: message not found")
	ErrPrivacyInvalid   = errors.Sentinel("pkgo: invalid privacy setting")
)

// StatusError is returned when a request returns a non-200 status code
type StatusError struct {
	Code   int
	Status string
}

func (e *StatusError) Error() string {
	return fmt.Sprintf("http status code %v: %v", e.Code, e.Status)
}

// InvalidError is returned when the data for a PATCH or POST endpoint is invalid.
type InvalidError struct {
	field string
	value string
}

func (e *InvalidError) Error() string {
	return fmt.Sprintf(`invalid value in field "%s": "%s"`, e.field, e.value)
}
