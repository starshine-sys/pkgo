package pkgo

import (
	"encoding/json"
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

// Fronters gets the current fronters for a system
func (s *Session) Fronters(id string) (f Front, err error) {
	if id == "" && (!s.authorized || s.token == "") {
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

	err = s.getEndpoint("/s/"+id+"/fronters", &f)
	return
}

// Switches gets the latest 100 switches for a system.
// For earlier switches, see SwitchesBefore.
func (s *Session) Switches(id string) (switches []Switch, err error) {
	err = s.getEndpoint("/s/"+id+"/switches", &switches)
	return
}

// SwitchesBefore gets the 100 switches before the given timestamp.
// For the latest switches, see Switches.
func (s *Session) SwitchesBefore(id string, before time.Time) (switches []Switch, err error) {
	t := before.UTC().Format("2006-01-02T15:04:05")

	err = s.getEndpoint("/s/"+id+"/switches?before="+t+"Z", &switches)
	return
}

// RegisterSwitch registers a switch with the given member IDs
func (s *Session) RegisterSwitch(ids ...string) (err error) {
	b, err := json.Marshal(Switch{Members: ids})
	if err != nil {
		return err
	}

	_, err = s.postEndpoint("/s/switches", b, nil)
	return
}
