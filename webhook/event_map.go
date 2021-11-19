package webhook

// EventCreators are funcs used to create Event structs to unmarshal to.
var EventCreators = map[DispatchEventType]func() Event{
	CreateSwitchEventType: func() Event { return new(CreateSwitchEventData) },
}
