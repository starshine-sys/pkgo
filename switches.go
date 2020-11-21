package pkgo

import (
	"fmt"
	"time"
)

// Switch holds the info for a simple switch, as queried from /s/id/switches
type Switch struct {
	Timestamp time.Time `json:"timestamp"`
	Members   []string  `json:"members"`
}

// Front holds the info for a full switch, as queried from /s/id/fronters
type Front struct {
	Timestamp time.Time `json:"timestamp"`
	Members   []Member  `json:"members"`
}

// GetFronters gets the current fronters for a system
func (s *Session) GetFronters(id string) (f Front, err error) {
	if id == "" && (!s.Authorized || s.Token == "") {
		return f, &ErrNoToken{}
	}
	if id == "" {
		if s.system == "" {
			sys, err := s.GetSystem()
			if err != nil {
				return f, err
			}
			id = sys.ID
		} else {
			id = s.system
		}
	}

	err = s.GetEndpoint("/s/"+id+"/fronters", &f)
	return
}

// RegisterSwitch registers a switch with the given member IDs
func (s *Session) RegisterSwitch(ids ...string) (err error) {
	fmt.Println(ids)
	return
}
