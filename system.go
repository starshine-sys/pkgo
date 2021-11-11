package pkgo

import (
	"time"

	"github.com/google/uuid"
)

// System holds all the data for a system
type System struct {
	ID   string    `json:"id"`
	UUID uuid.UUID `json:"uuid"`

	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Tag         string `json:"tag,omitempty"`
	AvatarURL   string `json:"avatar_url,omitempty"`
	Banner      string `json:"banner,omitempty"`
	Color       string `json:"color,omitempty"`

	Created  time.Time `json:"created"`
	Timezone string    `json:"timezone,omitempty"`

	Privacy *SystemPrivacy `json:"privacy,omitempty"`
}

// SystemPrivacy ...
type SystemPrivacy struct {
	DescriptionPrivacy  Privacy `json:"description_privacy,omitempty"`
	MemberListPrivacy   Privacy `json:"member_list_privacy,omitempty"`
	GroupListPrivacy    Privacy `json:"group_list_privacy,omitempty"`
	FrontPrivacy        Privacy `json:"front_privacy,omitempty"`
	FrontHistoryPrivacy Privacy `json:"front_history_privacy,omitempty"`
}

// Me gets the current token's system.
// If force is set to true, this will always fetch the system from the API.
func (s *Session) Me(force bool) (sys System, err error) {
	if s.token == "" {
		return sys, ErrNoToken
	}

	if !force && s.system != nil {
		return *s.system, nil
	}

	sys, err = s.System("@me")
	if err != nil {
		return
	}

	s.system = &sys
	return
}

// System gets a system by its 5-character system ID.
// Some fields may be empty if unauthenticated and the system has chosen to make those fields private.
func (s *Session) System(id string) (sys System, err error) {
	err = s.RequestJSON("GET", "/systems/"+id, &sys)
	return
}

// SystemByUUID gets a system by its UUID.
// See s.System for more documentation.
func (s *Session) SystemByUUID(id uuid.UUID) (System, error) {
	return s.System(id.String())
}

// Account gets a system by a Discord snowflake (user ID).
func (s *Session) Account(id Snowflake) (sys System, err error) {
	return s.System(id.String())
}

// EditSystemData ...
type EditSystemData struct {
	Name        NullableString `json:"name,omitempty"`
	Description NullableString `json:"description,omitempty"`
	Tag         NullableString `json:"tag,omitempty"`
	AvatarURL   NullableString `json:"avatar_url,omitempty"`
	Banner      NullableString `json:"banner,omitempty"`
	Color       NullableString `json:"color,omitempty"`
	Timezone    NullableString `json:"tz,omitempty"`
	Privacy     *SystemPrivacy `json:"privacy,omitempty"`
}

// EditSystem edits your system with the provided data.
func (s *Session) EditSystem(psd EditSystemData) (sys System, err error) {
	err = s.RequestJSON("PATCH", "/systems/@me", &sys, WithJSONBody(psd))
	return sys, err
}
