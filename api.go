package pkgo

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"emperror.dev/errors"
)

const baseURL = "https://api.pluralkit.me/v"
const version = "1"

// ErrRateLimit is returned when the API rate limit is hit
var ErrRateLimit = errors.New("pkgo: hit API rate limits")

// getEndpoint makes a request to a GET API endpoint
func (s *Session) getEndpoint(endpoint string, data interface{}) error {
	client := &http.Client{}
	req, err := http.NewRequest("GET", baseURL+version+endpoint, nil)
	if err != nil {
		return fmt.Errorf("pkgo: error creating request: %w", err)
	}
	if s.authorized {
		req.Header.Add("Authorization", s.token)
	}

	ctx, cancel := context.WithTimeout(context.Background(), s.Timeout)
	defer cancel()

	err = s.rate.Wait(ctx)
	if err != nil {
		return errors.Wrap(err, "s.getEndpoint")
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// we hit the rate limit
	if resp.StatusCode == 429 {
		return ErrRateLimit
	}

	if resp.StatusCode != 200 {
		return &StatusError{Code: resp.StatusCode, Status: resp.Status}
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.Wrap(err, "s.getEndpoint")
	}
	err = json.Unmarshal(b, data)
	return errors.Wrap(err, "s.getEndpoint")
}

// postEndpoint makes a request to a POST API endpoint
func (s *Session) postEndpoint(endpoint string, data []byte, in interface{}) (interface{}, error) {
	client := &http.Client{}

	req, err := http.NewRequest("POST", baseURL+version+endpoint, bytes.NewBuffer(data))
	if err != nil {
		return in, fmt.Errorf("pkgo: error creating request: %w", err)
	}

	if s.authorized {
		req.Header.Add("Authorization", s.token)
	} else {
		return in, ErrNoToken
	}
	req.Header.Add("Content-Type", "application/json")

	ctx, cancel := context.WithTimeout(context.Background(), s.Timeout)
	defer cancel()

	err = s.rate.Wait(ctx)
	if err != nil {
		return in, errors.Wrap(err, "pkgo: s.postEndpoint")
	}

	resp, err := client.Do(req)
	if err != nil {
		return in, errors.Wrap(err, "pkgo: s.postEndpoint")
	}
	defer resp.Body.Close()

	// we hit the rate limit
	if resp.StatusCode == 429 {
		return in, ErrRateLimit
	}

	if resp.StatusCode == 204 {
		return nil, nil
	}

	if resp.StatusCode != 200 {
		return in, &StatusError{Code: resp.StatusCode, Status: resp.Status}
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return in, errors.Wrap(err, "pkgo: s.postEndpoint")
	}
	err = json.Unmarshal(b, in)
	return in, errors.Wrap(err, "pkgo: s.postEndpoint")
}

func (s *Session) patchEndpoint(endpoint string, data []byte, in interface{}) (err error) {
	client := &http.Client{}

	req, err := http.NewRequest("PATCH", baseURL+version+endpoint, bytes.NewBuffer(data))
	if err != nil {
		return fmt.Errorf("pkgo: error creating request: %w", err)
	}

	if s.authorized {
		req.Header.Add("Authorization", s.token)
	} else {
		return ErrNoToken
	}
	req.Header.Add("Content-Type", "application/json")

	ctx, cancel := context.WithTimeout(context.Background(), s.Timeout)
	defer cancel()

	err = s.rate.Wait(ctx)
	if err != nil {
		return errors.Wrap(err, "s.patchEndpoint")
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// we hit the rate limit
	if resp.StatusCode == 429 {
		return ErrRateLimit
	}

	if resp.StatusCode != 200 {
		return &StatusError{Code: resp.StatusCode, Status: resp.Status}
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.Wrap(err, "s.patchEndpoint")
	}
	err = json.Unmarshal(b, in)
	return errors.Wrap(err, "s.patchEndpoint")
}

func (s *Session) deleteEndpoint(endpoint string) (err error) {
	client := &http.Client{}

	req, err := http.NewRequest("DELETE", baseURL+version+endpoint, nil)
	if err != nil {
		return fmt.Errorf("pkgo: error creating request: %w", err)
	}

	if s.authorized {
		req.Header.Add("Authorization", s.token)
	} else {
		return ErrNoToken
	}

	ctx, cancel := context.WithTimeout(context.Background(), s.Timeout)
	defer cancel()

	err = s.rate.Wait(ctx)
	if err != nil {
		return errors.Wrap(err, "s.deleteEndpoint")
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	return resp.Body.Close()
}
