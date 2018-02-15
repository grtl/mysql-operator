package cluster

import (
	crv1 "github.com/grtl/mysql-operator/pkg/apis/cr/v1"
)

// ClusterEventType represents type of a ClusterEvent.
type ClusterEventType int

// Available ClusterEvent types.
const (
	ClusterAdded ClusterEventType = iota
	ClusterUpdated
	ClusterDeleted
)

// ClusterEvent represents event processed by cluster controller.
type ClusterEvent struct {
	Type    ClusterEventType
	Cluster *crv1.MySQLCluster
}

// EventsHook is a ControllerHook for cluster controller, which publishes all
// processed events to events channel.
type EventsHook interface {
	OnAdd(object interface{})
	OnUpdate(object interface{})
	OnDelete(object interface{})
	GetEventsChan() <-chan ClusterEvent
}

type eventsHook struct {
	events chan ClusterEvent
}

// NewEventsHook returns a new EventsHook with channel with capacity of channelSize.
func NewEventsHook(channelSize int) EventsHook {
	return &eventsHook{
		events: make(chan ClusterEvent, channelSize),
	}
}

func (h *eventsHook) GetEventsChan() <-chan ClusterEvent {
	return h.events
}

func (h *eventsHook) OnAdd(object interface{}) {
	cluster := object.(*crv1.MySQLCluster)
	h.events <- ClusterEvent{
		Type:    ClusterAdded,
		Cluster: cluster,
	}
}

func (h *eventsHook) OnUpdate(object interface{}) {
	cluster := object.(*crv1.MySQLCluster)
	h.events <- ClusterEvent{
		Type:    ClusterUpdated,
		Cluster: cluster,
	}
}

func (h *eventsHook) OnDelete(object interface{}) {
	cluster := object.(*crv1.MySQLCluster)
	h.events <- ClusterEvent{
		Type:    ClusterDeleted,
		Cluster: cluster,
	}
}
