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
	ErrInvalidProxyTag  = errors.Sentinel("pkgo: invalid proxy tag")
)

// PKAPIError is an error returned by the PluralKit API
type PKAPIError struct {
	StatusCode int                     `json:"-"`
	Code       int                     `json:"code"`
	RetryAfter *int                    `json:"retry_after,omitempty"`
	Message    string                  `json:"message"`
	Errors     map[string][]ModelError `json:"errors,omitempty"`
}

func (e PKAPIError) Error() string {
	if e.RetryAfter != nil {
		return fmt.Sprintf("rate limited, retry after %dms", *e.RetryAfter)
	}

	return fmt.Sprintf("%d: %s", e.Code, e.Message)
}

// ModelError ...
type ModelError struct {
	Message      string `json:"message"`
	MaxLength    int    `json:"max_length"`
	ActualLength int    `json:"actual_length"`
}

// InvalidError is returned when the data for a PATCH or POST endpoint is invalid.
type InvalidError struct {
	field string
	value string
}

func (e *InvalidError) Error() string {
	return fmt.Sprintf(`invalid value in field "%s": "%s"`, e.field, e.value)
}
