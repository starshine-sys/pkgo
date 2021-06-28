package pkgo

import (
	"encoding/json"
	"time"
)

// System holds all the data for a system
type System struct {
	ID      string    `json:"id"`
	Created time.Time `json:"created"`

	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Tag         string `json:"tag,omitempty"`
	AvatarURL   string `json:"avatar_url,omitempty"`
	Timezone    string `json:"tz,omitempty"`

	DescriptionPrivacy  Privacy `json:"description_privacy,omitempty"`
	MemberListPrivacy   Privacy `json:"member_list_privacy,omitempty"`
	FrontPrivacy        Privacy `json:"front_privacy,omitempty"`
	FrontHistoryPrivacy Privacy `json:"front_history_privacy,omitempty"`
}

// Me gets the current token's system.
// If force is set to true, this will always fetch the system from the API.
func (s *Session) Me(force bool) (sys *System, err error) {
	if !s.authorized || s.token == "" {
		return nil, ErrNoToken
	}

	if !force && s.system != nil {
		return s.system, nil
	}

	err = s.getEndpoint("/s", &sys)
	if err != nil {
		return
	}

	s.system = sys
	return
}

// System gets a system by its 5-character system ID.
// Some fields may be empty if unauthenticated and the system has chosen to make those fields private.
func (s *Session) System(id string) (sys *System, err error) {
	if !idRe.MatchString(id) {
		return nil, ErrInvalidID
	}
	err = s.getEndpoint("/s/"+id, &sys)
	return
}

// Account gets a system by a Discord snowflake (user ID).
func (s *Session) Account(id Snowflake) (sys *System, err error) {
	err = s.getEndpoint("/a/"+id.String(), &sys)
	return
}

// EditSystemData ...
type EditSystemData struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Tag         string `json:"tag,omitempty"`
	AvatarURL   string `json:"avatar_url,omitempty"`
	Timezone    string `json:"tz,omitempty"`

	DescriptionPrivacy  Privacy `json:"description_privacy,omitempty"`
	MemberListPrivacy   Privacy `json:"member_list_privacy,omitempty"`
	FrontPrivacy        Privacy `json:"front_privacy,omitempty"`
	FrontHistoryPrivacy Privacy `json:"front_history_privacy,omitempty"`
}

// EditSystem edits your system with the provided data.
func (s *Session) EditSystem(psd EditSystemData) (sys *System, err error) {
	b, err := json.Marshal(psd)
	if err != nil {
		return sys, err
	}

	sys = &System{}
	err = s.patchEndpoint("/s", b, sys)
	if err != nil {
		return nil, err
	}
	return sys, err
}
