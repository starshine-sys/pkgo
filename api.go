package pkgo

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"emperror.dev/errors"
)

var (
	// BaseURL is the API base url
	BaseURL = "https://api.pluralkit.me/v"
	// Version is the API version
	Version = "1"
)

// Errors returned by Request
const (
	ErrRateLimit = errors.Sentinel("pkgo: hit API rate limits")

	ErrBadRequest    = errors.Sentinel("pkgo: 400 bad request")
	ErrUnauthorized  = errors.Sentinel("pkgo: 401 unauthorized")
	ErrNotFound      = errors.Sentinel("pkgo: 404 not found")
	ErrAlreadyExists = errors.Sentinel("pkgo: 409 resource already exists")
	ErrUnprocessable = errors.Sentinel("pkgo: 422 unprocessable entity")
	ErrUnavailable   = errors.Sentinel("pkgo: 503 service unavailable")
)

type apiError int

func (e apiError) Error() string {
	return fmt.Sprintf("%v %v", int(e), http.StatusText(int(e)))
}

// Request makes a request returning a JSON body.
func (s *Session) Request(method, endpoint string, opts ...RequestOption) (response []byte, err error) {
	req, err := http.NewRequest(method, s.BaseURL+endpoint, nil)
	if err != nil {
		return
	}

	err = s.applyOpts(req, opts)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), s.Timeout)
	defer cancel()

	err = s.rate.Wait(ctx)
	if err != nil {
		return
	}

	resp, err := s.Client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	response, err = io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	switch resp.StatusCode {
	case http.StatusOK, http.StatusNoContent, http.StatusCreated:
	case http.StatusBadRequest:
		return
	case http.StatusUnauthorized:
		return nil, ErrUnauthorized
	case http.StatusNotFound:
		return nil, ErrNotFound
	case http.StatusConflict:
		return nil, ErrAlreadyExists
	case http.StatusUnprocessableEntity:
		return nil, ErrUnprocessable
	case http.StatusServiceUnavailable:
		return nil, ErrUnavailable
	default:
		return nil, apiError(resp.StatusCode)
	}

	return response, err
}

// RequestJSON makes a request returning a JSON body.
func (s *Session) RequestJSON(method, endpoint string, v interface{}, opts ...RequestOption) error {
	resp, err := s.Request(method, endpoint, opts...)
	if err != nil {
		return err
	}

	if v == nil {
		return nil
	}

	return json.Unmarshal(resp, v)
}

// applyOpts applies all options to the given request and returns the last error returned by an option.
func (s *Session) applyOpts(r *http.Request, opts []RequestOption) (err error) {
	// apply global options
	for _, opt := range s.RequestOptions {
		err = opt(r)
	}

	// apply local options
	for _, opt := range opts {
		err = opt(r)
	}
	return
}
