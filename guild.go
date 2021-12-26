package pkgo

// SystemGuild is a system's guild-specific settings
type SystemGuild struct {
	ProxyingEnabled bool          `json:"proxying_enabled"`
	AutoproxyMode   AutoproxyMode `json:"autoproxy_mode"`
	AutoproxyMember string        `json:"autoproxy_member,omitempty"`
	Tag             string        `json:"tag,omitempty"`
	TagEnabled      bool          `json:"tag_enabled"`
}

type AutoproxyMode string

// Autoproxy modes
var (
	AutoproxyModeOff    AutoproxyMode = "off"
	AutoproxyModeFront  AutoproxyMode = "front"
	AutoproxyModeLatch  AutoproxyMode = "latch"
	AutoproxyModeMember AutoproxyMode = "member"
)

// SystemGuild returns the current system's guild settings.
func (s *Session) SystemGuild(guildID Snowflake) (sg SystemGuild, err error) {
	err = s.RequestJSON("GET", "/systems/@me/guilds/"+guildID.String(), &sg)
	return
}

// EditSystemGuildData ...
type EditSystemGuildData struct {
	ProxyingEnabled NullableBool   `json:"proxying_enabled"`
	AutoproxyMode   *AutoproxyMode `json:"autoproxy_mode"`
	AutoproxyMember NullableString `json:"autoproxy_member,omitempty"`
	Tag             NullableString `json:"tag,omitempty"`
	TagEnabled      NullableBool   `json:"tag_enabled"`
}

// EditSystemGuild edits the current system's guild settings.
func (s *Session) EditSystemGuild(guildID Snowflake, dat EditSystemGuildData) (sg SystemGuild, err error) {
	err = s.RequestJSON("PATCH", "/systems/@me/guilds/"+guildID.String(), &sg, WithJSONBody(dat))
	return
}

// MemberGuild is a member's guild-specific settings
type MemberGuild struct {
	DisplayName string `json:"display_name,omitempty"`
	AvatarURL   string `json:"avatar_url,omitempty"`
}

// MemberGuild returns the member's guild settings.
func (s *Session) MemberGuild(memberID string, guildID Snowflake) (mg MemberGuild, err error) {
	err = s.RequestJSON("GET", "/members/"+memberID+"/guilds/"+guildID.String(), &mg)
	return
}

// EditMemberGuildData ...
type EditMemberGuildData struct {
	DisplayName NullableString `json:"display_name,omitempty"`
	AvatarURL   NullableString `json:"avatar_url,omitempty"`
}

// EditMemberGuild edits the member's guild settings.
func (s *Session) EditMemberGuild(memberID string, guildID Snowflake, dat EditMemberGuildData) (mg MemberGuild, err error) {
	err = s.RequestJSON("PATCH", "/members/"+memberID+"/guilds/"+guildID.String(), &mg, WithJSONBody(dat))
	return
}
