package webhook

import (
	"bytes"
	"encoding/json"
	"io"

	"emperror.dev/errors"
	"github.com/google/uuid"
)

// TokenGetter is a function used to get a signing token for the given system ID.
// If it returns an error, unmarshaling is aborted.
type TokenGetter func(systemID uuid.UUID) (string, error)

const ErrInvalidToken = errors.Sentinel("payload token doesn't match")

// Unmarshal unmarshals the given byte slice to a WebhookEvent,
// and also validates the token if tkn is not nil.
// If the token returned by tkn is empty, the event is *not* validated.
func Unmarshal(src []byte, tkn TokenGetter) (ev WebhookEvent, err error) {
	return Decode(bytes.NewReader(src), tkn)
}

// Unmarshal unmarshals the given io.Reader to a WebhookEvent,
// and also validates the token if tkn is not nil.
// If the token returned by tkn is empty, the event is *not* validated.
func Decode(r io.Reader, tkn TokenGetter) (ev WebhookEvent, err error) {
	err = json.NewDecoder(r).Decode(&ev)
	if err != nil {
		return ev, err
	}

	if tkn != nil {
		token, err := tkn(ev.SystemID)
		if err != nil {
			return ev, err
		}

		if token != "" {
			if token != ev.Token {
				return ev, ErrInvalidToken
			}
		}
	}

	if fn, ok := EventCreators[ev.Type]; ok {
		ev.Data = fn()
	} else {
		ev.Data = new(UnknownEventData)
	}

	err = json.Unmarshal(ev.Raw, ev.Data)
	if err != nil {
		return ev, err
	}
	if ev.Data == nil {
		ev.Data = &UnknownEventData{}
	}

	return ev, err
}
