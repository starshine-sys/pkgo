package pkgo

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const baseURL = "https://api.pluralkit.me/v"
const version = "1"

// GetEndpoint makes a request to a GET API endpoint
func (s *Session) GetEndpoint(endpoint string, data interface{}) error {
	client := &http.Client{}
	req, err := http.NewRequest("GET", baseURL+version+endpoint, nil)
	if err != nil {
		panic(err)
	}
	if s.Authorized && s.Token != "" {
		req.Header.Add("Authorization", s.Token)
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
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
