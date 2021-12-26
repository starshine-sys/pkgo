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

	Created time.Time `json:"created"`

	Privacy *SystemPrivacy `json:"privacy,omitempty"`
}

// SystemPrivacy is a system privacy object.
type SystemPrivacy struct {
	DescriptionPrivacy  Privacy `json:"description_privacy,omitempty"`
	MemberListPrivacy   Privacy `json:"member_list_privacy,omitempty"`
	GroupListPrivacy    Privacy `json:"group_list_privacy,omitempty"`
	FrontPrivacy        Privacy `json:"front_privacy,omitempty"`
	FrontHistoryPrivacy Privacy `json:"front_history_privacy,omitempty"`
}

// Me gets the current token's system.
func (s *Session) Me() (sys System, err error) {
	if s.token == "" {
		return sys, ErrNoToken
	}

	return s.System("@me")
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

type SystemSettings struct {
	Timezone     string `json:"timezone"`
	PingsEnabled bool   `json:"pings_enabled"`
	LatchTimeout int    `json:"latch_timeout"`

	MemberDefaultPrivate bool `json:"member_default_private"`
	GroupDefaultPrivate  bool `json:"group_default_private"`
	ShowPrivateInfo      bool `json:"show_private_info"`

	MemberLimit int `json:"member_limit"`
	GroupLimit  int `json:"group_limit"`
}

// SystemSettings returns the current token's system's settings.
func (s *Session) SystemSettings() (settings SystemSettings, err error) {
	err = s.RequestJSON("GET", "/systems/@me/settings", &settings)
	return settings, err
}

type UpdateSystemSettingsData struct {
	Timezone     NullableString `json:"timezone"`
	PingsEnabled NullableBool   `json:"pings_enabled"`
	LatchTimeout NullableInt    `json:"latch_timeout"`

	MemberDefaultPrivate NullableBool `json:"member_default_private"`
	GroupDefaultPrivate  NullableBool `json:"group_default_private"`
	ShowPrivateInfo      NullableBool `json:"show_private_info"`
}

func (s *Session) UpdateSystemSettings(data UpdateSystemSettingsData) (settings SystemSettings, err error) {
	err = s.RequestJSON("PATCH", "/systems/@me/settings", &settings, WithJSONBody(data))
	return settings, err
}
