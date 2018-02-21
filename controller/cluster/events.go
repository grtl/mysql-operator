package cluster

import (
	"github.com/grtl/mysql-operator/controller"
	crv1 "github.com/grtl/mysql-operator/pkg/apis/cr/v1"
)

// EventType represents type of a Event.
type EventType int

// Available Event types.
const (
	ClusterAdded EventType = iota
	ClusterUpdated
	ClusterDeleted
)

// Event represents event processed by cluster controller.
type Event struct {
	Type    EventType
	Cluster *crv1.MySQLCluster
}

// EventsHook extends `controller.Hook`. All processed events are published to
// events channel.
type EventsHook interface {
	controller.Hook
	GetEventsChan() <-chan Event
}

type eventsHook struct {
	events chan Event
}

// NewEventsHook returns a new EventsHook with channel with capacity of channelSize.
func NewEventsHook(channelSize int) EventsHook {
	return &eventsHook{
		events: make(chan Event, channelSize),
	}
}

func (h *eventsHook) GetEventsChan() <-chan Event {
	return h.events
}

func (h *eventsHook) OnAdd(object interface{}) {
	cluster := object.(*crv1.MySQLCluster)
	h.events <- Event{
		Type:    ClusterAdded,
		Cluster: cluster,
	}
}

func (h *eventsHook) OnUpdate(object interface{}) {
	cluster := object.(*crv1.MySQLCluster)
	h.events <- Event{
		Type:    ClusterUpdated,
		Cluster: cluster,
	}
}

func (h *eventsHook) OnDelete(object interface{}) {
	cluster := object.(*crv1.MySQLCluster)
	h.events <- Event{
		Type:    ClusterDeleted,
		Cluster: cluster,
	}
}
