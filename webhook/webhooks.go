// Package webhook provides support for PluralKit's dispatch webhooks.
package webhook

import (
	"encoding/json"
	"strconv"

	"github.com/google/uuid"
	"github.com/starshine-sys/pkgo/v2"
)

type WebhookEvent struct {
	Type  DispatchEventType `json:"type"`
	Token string            `json:"signing_token"`

	SystemID uuid.UUID      `json:"system_id"`
	EntityID string         `json:"id,omitempty"`
	GuildID  pkgo.Snowflake `json:"guild_id,omitempty"`

	// Raw is the raw JSON data.
	Raw json.RawMessage `json:"data"`

	// Data is the unmarshalled dispatch event payload.
	// Type assertions should be made on this value.
	// All types in this value are *pointer* types.
	// Unknown events are unmarshalled as UnknownEventPayload
	Data Event `json:"-"`
}

// UUID tries parsing the event's entity ID into a UUID
func (ev WebhookEvent) UUID() (uuid.UUID, error) {
	return uuid.Parse(ev.EntityID)
}

// Snowflake tries parsing the event's entity ID into an account ID
func (ev WebhookEvent) Snowflake() (pkgo.Snowflake, error) {
	i, err := strconv.ParseUint(ev.EntityID, 10, 64)
	if err != nil {
		return 0, err
	}
	return pkgo.Snowflake(i), nil
}

type DispatchEventType string

const (
	UnknownEventType        DispatchEventType = ""
	PingEventType           DispatchEventType = "PING"
	UpdateSystemEventType   DispatchEventType = "UPDATE_SYSTEM"
	UpdateSettingsEventType DispatchEventType = "UPDATE_SETTINGS"

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

	SuccessfulImportEventType DispatchEventType = "SUCCESSFUL_IMPORT"
)

type Event interface {
	Type() DispatchEventType
}
