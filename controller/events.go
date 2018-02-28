package controller

// EventType represents type of a Event.
type EventType int

// Available Event types.
const (
	EventAdded EventType = iota
	EventUpdated
	EventDeleted
)

// Event represents event processed by the controller.
type Event struct {
	Type   EventType
	Object interface{}
}

// EventsHook extends `Hook`. All processed events are published to the events channel.
type EventsHook interface {
	Hook
	GetEventsChan() <-chan Event
}

type eventsHook struct {
	events chan Event
}

// NewEventsHook returns a new EventsHook with channel with the capacity of channelSize.
func NewEventsHook(channelSize int) EventsHook {
	return &eventsHook{
		events: make(chan Event, channelSize),
	}
}

func (h *eventsHook) GetEventsChan() <-chan Event {
	return h.events
}

func (h *eventsHook) OnAdd(object interface{}) {
	h.events <- Event{
		Type:   EventAdded,
		Object: object,
	}
}

func (h *eventsHook) OnUpdate(object interface{}) {
	h.events <- Event{
		Type:   EventUpdated,
		Object: object,
	}
}

func (h *eventsHook) OnDelete(object interface{}) {
	h.events <- Event{
		Type:   EventDeleted,
		Object: object,
	}
}
