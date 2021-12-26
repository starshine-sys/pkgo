package webhook

import (
	"encoding/json"
	"time"

	"github.com/starshine-sys/pkgo/v2"
)

type UnknownEventData struct {
	json.RawMessage
}

var _ Event = (*UnknownEventData)(nil)

func (*UnknownEventData) Type() DispatchEventType {
	return UnknownEventType
}

type PingEventData struct{}

var _ Event = (*PingEventData)(nil)

func (*PingEventData) Type() DispatchEventType {
	return PingEventType
}

type CreateSwitchEventData struct {
	pkgo.Switch
}

var _ Event = (*CreateSwitchEventData)(nil)

func (*CreateSwitchEventData) Type() DispatchEventType {
	return CreateSwitchEventType
}

type UpdateSystemEventData struct {
	Name        *string             `json:"name,omitempty"`
	Description *string             `json:"description,omitempty"`
	Tag         *string             `json:"tag,omitempty"`
	AvatarURL   *string             `json:"avatar_url,omitempty"`
	Banner      *string             `json:"banner,omitempty"`
	Color       *string             `json:"color,omitempty"`
	Privacy     *pkgo.SystemPrivacy `json:"privacy,omitempty"`
}

var _ Event = (*UpdateSystemEventData)(nil)

func (*UpdateSystemEventData) Type() DispatchEventType {
	return UpdateSystemEventType
}

type UpdateSettingsEventData struct {
	Timezone     *string `json:"timezone,omitempty"`
	PingsEnabled *bool   `json:"pings_enabled,omitempty"`
	LatchTimeout *int    `json:"latch_timeout,omitempty"`

	MemberDefaultPrivate *bool `json:"member_default_private,omitempty"`
	GroupDefaultPrivate  *bool `json:"group_default_private,omitempty"`
	ShowPrivateInfo      *bool `json:"show_private_info,omitempty"`

	MemberLimit *int `json:"member_limit,omitempty"`
	GroupLimit  *int `json:"group_limit,omitempty"`
}

var _ Event = (*UpdateSettingsEventData)(nil)

func (*UpdateSettingsEventData) Type() DispatchEventType {
	return UpdateSettingsEventType
}

type CreateMemberEventData struct {
	Name string `json:"name"`
}

var _ Event = (*CreateMemberEventData)(nil)

func (*CreateMemberEventData) Type() DispatchEventType {
	return CreateMemberEventType
}

type UpdateMemberEventData struct {
	Name        *string        `json:"name,omitempty"`
	DisplayName *string        `json:"display_name,omitempty"`
	Color       *pkgo.Color    `json:"color,omitempty"`
	Birthday    *pkgo.Birthday `json:"birthday,omitempty"`
	Pronouns    *string        `json:"pronouns,omitempty"`
	AvatarURL   *string        `json:"avatar_url,omitempty"`
	Banner      *string        `json:"banner,omitempty"`
	Description *string        `json:"description,omitempty"`

	ProxyTags *[]pkgo.ProxyTag `json:"proxy_tags,omitempty"`
	KeepProxy *bool            `json:"keep_proxy"`

	Privacy *pkgo.MemberPrivacy `json:"privacy,omitempty"`
}

var _ Event = (*UpdateMemberEventData)(nil)

func (*UpdateMemberEventData) Type() DispatchEventType {
	return UpdateMemberEventType
}

type DeleteMemberEventData struct{}

var _ Event = (*DeleteMemberEventData)(nil)

func (*DeleteMemberEventData) Type() DispatchEventType {
	return DeleteMemberEventType
}

type CreateGroupEventData struct {
	Name string `json:"name"`
}

var _ Event = (*CreateGroupEventData)(nil)

func (*CreateGroupEventData) Type() DispatchEventType {
	return CreateGroupEventType
}

type UpdateGroupEventData struct {
	Name        *string     `json:"name,omitempty"`
	DisplayName *string     `json:"display_name,omitempty"`
	Description *string     `json:"description,omitempty"`
	Icon        *string     `json:"icon,omitempty"`
	Banner      *string     `json:"banner,omitempty"`
	Color       *pkgo.Color `json:"color,omitempty"`

	Privacy *pkgo.GroupPrivacy `json:"privacy,omitempty"`
}

var _ Event = (*UpdateGroupEventData)(nil)

func (*UpdateGroupEventData) Type() DispatchEventType {
	return UpdateGroupEventType
}

type DeleteGroupEventData struct{}

var _ Event = (*DeleteGroupEventData)(nil)

func (*DeleteGroupEventData) Type() DispatchEventType {
	return DeleteGroupEventType
}

type LinkAccountEventData struct{}

var _ Event = (*LinkAccountEventData)(nil)

func (*LinkAccountEventData) Type() DispatchEventType {
	return LinkAccountEventType
}

type UnlinkAccountEventData struct{}

var _ Event = (*UnlinkAccountEventData)(nil)

func (*UnlinkAccountEventData) Type() DispatchEventType {
	return UnlinkAccountEventType
}

type UpdateSystemGuildEventData struct {
	GuildID         pkgo.Snowflake      `json:"guild_id"`
	ProxyingEnabled *bool               `json:"proxying_enabled,omitempty"`
	AutoproxyMode   *pkgo.AutoproxyMode `json:"autoproxy_mode,omitempty"`
	AutoproxyMember *string             `json:"autoproxy_member,omitempty"`
	Tag             *string             `json:"tag,omitempty"`
	TagEnabled      *bool               `json:"tag_enabled,omitempty"`
}

var _ Event = (*UpdateSystemGuildEventData)(nil)

func (*UpdateSystemGuildEventData) Type() DispatchEventType {
	return UpdateSystemGuildEventType
}

type UpdateMemberGuildEventData struct {
	GuildID     pkgo.Snowflake `json:"guild_id"`
	DisplayName *string        `json:"display_name,omitempty"`
	AvatarURL   *string        `json:"avatar_url,omitempty"`
}

var _ Event = (*UpdateMemberGuildEventData)(nil)

func (*UpdateMemberGuildEventData) Type() DispatchEventType {
	return UpdateMemberGuildEventType
}

type CreateMessageEventData struct {
	pkgo.Message
}

var _ Event = (*CreateMessageEventData)(nil)

func (*CreateMessageEventData) Type() DispatchEventType {
	return CreateMessageEventType
}

type UpdateSwitchEventData struct {
	Timestamp *time.Time `json:"timestamp,omitempty"`
	Members   *[]string  `json:"members,omitempty"`
}

var _ Event = (*UpdateSwitchEventData)(nil)

func (*UpdateSwitchEventData) Type() DispatchEventType {
	return UpdateSwitchEventType
}

type DeleteSwitchEventData struct{}

var _ Event = (*DeleteSwitchEventData)(nil)

func (*DeleteSwitchEventData) Type() DispatchEventType {
	return DeleteSwitchEventType
}

type DeleteAllSwitchesEventData struct{}

var _ Event = (*DeleteAllSwitchesEventData)(nil)

func (*DeleteAllSwitchesEventData) Type() DispatchEventType {
	return DeleteAllSwitchesEventType
}

type SuccessfulImportEventData struct{}

var _ Event = (*SuccessfulImportEventData)(nil)

func (*SuccessfulImportEventData) Type() DispatchEventType {
	return SuccessfulImportEventType
}
