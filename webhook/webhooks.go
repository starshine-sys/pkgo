// Package webhook provides support for PluralKit's outgoing webhooks.
package webhook

import (
	"encoding/json"

	"github.com/google/uuid"
	"github.com/starshine-sys/pkgo/v2"
)

type WebhookEvent struct {
	Type  DispatchEventType `json:"type"`
	Token string            `json:"signing_token"`

	SystemID uuid.UUID      `json:"system_id"`
	EntityID uuid.UUID      `json:"entity_id,omitempty"`
	GuildID  pkgo.Snowflake `json:"guild_id,omitempty"`

	// Raw is the raw JSON data.
	Raw json.RawMessage `json:"data"`

	// Data is the unmarshalled dispatch event payload.
	// Type assertions should be made on this value.
	// Unknown events are unmarshalled as UnknownEventPayload
	Data Event `json:"-"`
}

type DispatchEventType string

const (
	UnknownEventType      DispatchEventType = ""
	PingEventType         DispatchEventType = "PING"
	UpdateSystemEventType DispatchEventType = "UPDATE_SYSTEM"

	CreateMemberEventType DispatchEventType = "CREATE_MEMBER"
	UpdateMemberEventType DispatchEventType = "UPDATE_MEMBER"
	DeleteMemberEventType DispatchEventType = "DELETE_MEMBER"

	CreateGroupEventType        DispatchEventType = "CREATE_GROUP"
	UpdateGroupEventType        DispatchEventType = "UPDATE_GROUP"
	UpdateGroupMembersEventType DispatchEventType = "UPDATE_GROUP_MEMBERS"
	DeleteGroupEventType        DispatchEventType = "DELETE_GROUP"

	LinkAccountEventType   DispatchEventType = "LINK_ACCOUNT"
	UnlinkAccountEventType DispatchEventType = "UNLINK_ACCOUNT"

	UpdateSystemGuildEventType DispatchEventType = "UPDATE_SYSTEM_GUILD"
	UpdateMemberGuildEventType DispatchEventType = "UPDATE_MEMBER_GUILD"

	CreateMessageEventType DispatchEventType = "CREATE_MESSAGE"

	CreateSwitchEventType        DispatchEventType = "CREATE_SWITCH"
	UpdateSwitchEventType        DispatchEventType = "UPDATE_SWITCH"
	UpdateSwitchMembersEventType DispatchEventType = "UPDATE_SWITCH_MEMBERS"
	DeleteSwitchEventType        DispatchEventType = "DELETE_SWITCH"
	DeleteAllSwitchesEventType   DispatchEventType = "DELETE_ALL_SWITCHES"
)

type Event interface {
	Type() DispatchEventType
	_pld()
}
