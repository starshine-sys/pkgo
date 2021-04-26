package pkgo

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
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
	if s.authorized && s.token != "" {
		req.Header.Add("Authorization", s.token)
	}

	// block so we don't hit the rate limit
	s.RateLimit()

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
		return &ErrStatusNot200{Code: resp.StatusCode, Status: resp.Status}
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, data)
}

// postEndpoint makes a request to a GET API endpoint
func (s *Session) postEndpoint(endpoint string, data []byte, in interface{}) (interface{}, error) {
	client := &http.Client{}

	req, err := http.NewRequest("POST", baseURL+version+endpoint, bytes.NewBuffer(data))
	if err != nil {
		return in, fmt.Errorf("pkgo: error creating request: %w", err)
	}

	if s.authorized && s.token != "" {
		req.Header.Add("Authorization", s.token)
	} else {
		return in, ErrNoToken
	}
	req.Header.Add("Content-Type", "application/json")

	// block so we don't hit the rate limit
	s.RateLimit()

	resp, err := client.Do(req)
	if err != nil {
		return in, err
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
		return in, &ErrStatusNot200{Code: resp.StatusCode, Status: resp.Status}
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return in, err
	}
	err = json.Unmarshal(b, in)
	return in, err
}

func (s *Session) patchEndpoint(endpoint string, data []byte, in interface{}) (err error) {
	client := &http.Client{}

	req, err := http.NewRequest("PATCH", baseURL+version+endpoint, bytes.NewBuffer(data))
	if err != nil {
		return fmt.Errorf("pkgo: error creating request: %w", err)
	}

	if s.authorized && s.token != "" {
		req.Header.Add("Authorization", s.token)
	} else {
		return ErrNoToken
	}
	req.Header.Add("Content-Type", "application/json")

	// block so we don't hit the rate limit
	s.RateLimit()

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
		return &ErrStatusNot200{Code: resp.StatusCode, Status: resp.Status}
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, in)
}
