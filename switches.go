package pkgo

import (
	"math"
	"net/url"
	"strconv"
	"time"

	"github.com/google/uuid"
)

// Switch holds the info for a simple switch, as queried from /systems/{systemRef}/switches
type Switch struct {
	ID        uuid.UUID `json:"id"`
	Timestamp time.Time `json:"timestamp,omitempty"`
	Members   []string  `json:"members"`
}

// Front holds the info for a full switch, as queried from /systems/{systemRef}/fronters
type Front struct {
	ID        uuid.UUID `json:"id"`
	Timestamp time.Time `json:"timestamp,omitempty"`
	Members   []Member  `json:"members"`
}

// Fronters gets the current fronters for a system.
// If the system's fronters are set to private, requires authentication.
func (s *Session) Fronters(id string) (f Front, err error) {
	err = s.RequestJSON("GET", "/systems/"+id+"/fronters", &f)
	return
}

// Switches gets the latest 100 switches for a system.
// For earlier switches, see SwitchesBefore.
// If the system's font history is set to private, requires authentication.
func (s *Session) Switches(id string) (switches []Switch, err error) {
	switches = []Switch{}
	err = s.RequestJSON("GET", "/systems/"+id+"/switches", &switches)
	return
}

// SwitchesBefore gets `limit` switches before the given timestamp.
// A zero limit will return *all* switches.
// For the latest switches, see Switches.
// If the system's font history is set to private, requires authentication.
func (s *Session) SwitchesBefore(id string, before time.Time, limit uint64) (switches []Switch, err error) {
	if limit > 0 && limit <= 100 {
		return s.switchesBefore(id, before, limit)
	}

	toFetch := int64(limit)
	if limit == 0 {
		toFetch = math.MaxInt64
	}

	for {
		sw, err := s.switchesBefore(id, before, 100)
		if err != nil {
			return switches, err
		}
		switches = append(switches, sw...)
		toFetch -= 100

		// if we didn't get 100 switches, this is the last "page", so return
		if len(sw) < 100 {
			return switches, nil
		}
		// we've fetched all the switches we need
		if toFetch <= 0 {
			return switches, nil
		}

		before = sw[len(sw)-1].Timestamp
	}
}

func (s *Session) switchesBefore(id string, before time.Time, limit uint64) (switches []Switch, err error) {
	t := before.UTC().Format("2006-01-02T15:04:05")

	switches = []Switch{}
	err = s.RequestJSON("GET", "/systems/"+id+"/switches", &switches, WithURLValues(url.Values{
		"before": {t + "Z"},
		"limit":  {strconv.FormatUint(limit, 10)},
	}))
	return
}

// RegisterSwitch registers a switch with the given member IDs. Requires authentication.
func (s *Session) RegisterSwitch(ids ...string) (err error) {
	if len(ids) == 0 {
		ids = []string{}
	}

	dat := struct {
		Members []string
	}{ids}

	_, err = s.Request("POST", "/systems/@me/switches", WithJSONBody(dat))
	return
}

// RegisterSwitchWithTimestamp registers a switch with the given member IDs and the given timestamp. Requires authentication.
func (s *Session) RegisterSwitchWithTimestamp(t time.Time, ids ...string) (err error) {
	if len(ids) == 0 {
		ids = []string{}
	}

	dat := struct {
		Members   []string
		Timestamp time.Time
	}{ids, t.UTC()}

	_, err = s.Request("POST", "/systems/@me/switches", WithJSONBody(dat))
	return
}

// Switch gets the given switch.
func (s *Session) Switch(systemID string, switchID uuid.UUID) (f Front, err error) {
	err = s.RequestJSON("GET", "/systems/"+systemID+"/switches/"+switchID.String(), &f)
	return
}

// UpdateSwitchTimestamp updates the given switch's timestamp.
func (s *Session) UpdateSwitchTimestamp(id uuid.UUID, t time.Time) (f Front, err error) {
	dat := struct {
		Timestamp time.Time `json:"timestamp"`
	}{
		Timestamp: t,
	}

	err = s.RequestJSON("GET", "/systems/@me/switches/"+id.String(), &f, WithJSONBody(dat))
	return
}

// UpdateSwitchMembers updates the given switch's members.
func (s *Session) UpdateSwitchMembers(id uuid.UUID, members []string) (f Front, err error) {
	if members == nil {
		members = []string{}
	}

	err = s.RequestJSON("GET", "/systems/@me/switches/"+id.String()+"/members", &f, WithJSONBody(members))
	return
}

// DeleteSwitch deletes a switch. Requires authentication.
func (s *Session) DeleteSwitch(id uuid.UUID) (err error) {
	_, err = s.Request("DELETE", "/systems/@me/switches/"+id.String())
	return
}
