package webhook

// EventCreators are funcs used to create Event structs to unmarshal to.
var EventCreators = map[DispatchEventType]func() Event{
	PingEventType:              func() Event { return new(PingEventData) },
	CreateSwitchEventType:      func() Event { return new(CreateSwitchEventData) },
	UpdateSystemEventType:      func() Event { return new(UpdateSystemEventData) },
	UpdateSettingsEventType:    func() Event { return new(UpdateSettingsEventData) },
	CreateMemberEventType:      func() Event { return new(CreateMemberEventData) },
	UpdateMemberEventType:      func() Event { return new(UpdateMemberEventData) },
	DeleteMemberEventType:      func() Event { return new(DeleteMemberEventData) },
	CreateGroupEventType:       func() Event { return new(CreateGroupEventData) },
	UpdateGroupEventType:       func() Event { return new(UpdateGroupEventData) },
	DeleteGroupEventType:       func() Event { return new(DeleteGroupEventData) },
	LinkAccountEventType:       func() Event { return new(LinkAccountEventData) },
	UnlinkAccountEventType:     func() Event { return new(UnlinkAccountEventData) },
	UpdateSystemGuildEventType: func() Event { return new(UpdateSystemGuildEventData) },
	UpdateMemberGuildEventType: func() Event { return new(UpdateMemberGuildEventData) },
	CreateMessageEventType:     func() Event { return new(CreateMessageEventData) },
	UpdateSwitchEventType:      func() Event { return new(UpdateSwitchEventData) },
	DeleteSwitchEventType:      func() Event { return new(DeleteSwitchEventData) },
	DeleteAllSwitchesEventType: func() Event { return new(DeleteAllSwitchesEventData) },
	SuccessfulImportEventType:  func() Event { return new(SuccessfulImportEventData) },
}
