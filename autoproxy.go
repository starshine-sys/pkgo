package pkgo

import (
	"net/url"
	"time"
)

// AutoproxySettings are a system's autoproxy settings.
type AutoproxySettings struct {
	Mode AutoproxyMode `json:"autoproxy_mode"`
	// Member will be empty if Mode is "front"
	Member string `json:"autoproxy_member,omitempty"`
	// LastLatchTimestamp will be zero if Mode is not "latch"
	LastLatchTimestamp time.Time `json:"last_latch_timestamp,omitempty"`
}

type AutoproxyMode string

// Autoproxy modes
const (
	AutoproxyModeOff    AutoproxyMode = "off"
	AutoproxyModeFront  AutoproxyMode = "front"
	AutoproxyModeLatch  AutoproxyMode = "latch"
	AutoproxyModeMember AutoproxyMode = "member"
)

type EditAutoproxyData struct {
	Mode *AutoproxyMode `json:"autoproxy_mode"`
	// Member must be null if Mode is "front"
	Member NullableString `json:"autoproxy_member,omitempty"`
}

// GuildAutoproxy gets the current system's autoproxy settings in the given guild.
func (s *Session) GuildAutoproxy(guildID Snowflake) (as AutoproxySettings, err error) {
	err = s.RequestJSON("GET", "/systems/@me/autoproxy", &as, WithURLValues(url.Values{
		"guild_id": {guildID.String()},
	}))
	return
}

// EditGuildAutoproxy updates the current system's autoproxy settings in the given guild.
func (s *Session) EditGuildAutoproxy(guildID Snowflake, dat EditAutoproxyData) (as AutoproxySettings, err error) {
	err = s.RequestJSON("PATCH", "/systems/@me/autoproxy", &as, WithJSONBody(dat), WithURLValues(url.Values{
		"guild_id": {guildID.String()},
	}))
	return
}
