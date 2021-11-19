package webhook

import (
	"encoding/json"

	"github.com/starshine-sys/pkgo/v2"
)

type UnknownEventData struct {
	json.RawMessage
}

var _ Event = (*UnknownEventData)(nil)

func (u UnknownEventData) Type() DispatchEventType {
	return UnknownEventType
}

func (u UnknownEventData) _pld() {}

type CreateSwitchEventData struct {
	pkgo.Switch
}

var _ Event = (*CreateSwitchEventData)(nil)

func (u CreateSwitchEventData) Type() DispatchEventType {
	return CreateSwitchEventType
}

func (u CreateSwitchEventData) _pld() {}
