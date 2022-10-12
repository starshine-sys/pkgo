package pkgo

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var (
	// BaseURL is the API base url
	BaseURL = "https://api.pluralkit.me/v"
	// Version is the API version
	Version = "2"
)

// HTTPError is an error code returned by the API.
// This is only returned if the API did not return a valid error object.
type HTTPError int

func (h HTTPError) Error() string {
	return fmt.Sprintf("%v %v", int(h), http.StatusText(int(h)))
}

// Request makes a request returning a JSON body.
func (s *Session) Request(method, endpoint string, opts ...RequestOption) (response []byte, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), s.Timeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, method, s.BaseURL+endpoint, nil)
	if err != nil {
		return
	}

	err = s.applyOpts(req, opts)
	if err != nil {
		return nil, err
	}

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

	if resp.StatusCode >= 400 {
		var e PKAPIError
		err = json.Unmarshal(response, &e)
		if err != nil {
			return nil, HTTPError(resp.StatusCode)
		}
		e.StatusCode = resp.StatusCode
		return nil, &e
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
