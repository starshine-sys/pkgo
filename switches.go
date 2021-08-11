package pkgo

import (
	"net/url"
	"time"
)

// Switch holds the info for a simple switch, as queried from /s/:id/switches
type Switch struct {
	Timestamp time.Time `json:"timestamp,omitempty"`
	Members   []string  `json:"members"`
}

// Front holds the info for a full switch, as queried from /s/:id/fronters
type Front struct {
	Timestamp time.Time `json:"timestamp,omitempty"`
	Members   []Member  `json:"members"`
}

// Fronters gets the current fronters for a system.
// If the system's fronters are set to private, requires authentication.
func (s *Session) Fronters(id string) (f Front, err error) {
	if id == "" && s.token == "" {
		return f, ErrNoToken
	}
	if id == "" {
		if s.system == nil {
			sys, err := s.Me(false)
			if err != nil {
				return f, err
			}
			id = sys.ID
		} else {
			id = s.system.ID
		}
	}

	err = s.RequestJSON("GET", "/s/"+id+"/fronters", &f)
	return
}

// Switches gets the latest 100 switches for a system.
// For earlier switches, see SwitchesBefore.
// If the system's font history is set to private, requires authentication.
func (s *Session) Switches(id string) (switches []Switch, err error) {
	switches = []Switch{}
	err = s.RequestJSON("GET", "/s/"+id+"/switches", &switches)
	return
}

// SwitchesBefore gets the 100 switches before the given timestamp.
// For the latest switches, see Switches.
// If the system's font history is set to private, requires authentication.
func (s *Session) SwitchesBefore(id string, before time.Time) (switches []Switch, err error) {
	t := before.UTC().Format("2006-01-02T15:04:05")

	switches = []Switch{}
	err = s.RequestJSON("GET", "/s/"+id+"/switches", &switches, WithURLValues(url.Values{
		"before": {t + "Z"},
	}))
	return
}

// RegisterSwitch registers a switch with the given member IDs. Requires authentication.
func (s *Session) RegisterSwitch(ids ...string) (err error) {
	_, err = s.Request("POST", "/s/switches", WithJSONBody(Switch{
		Members: ids,
	}))
	return
}
