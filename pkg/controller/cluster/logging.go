package cluster

import (
	crv1 "github.com/grtl/mysql-operator/pkg/apis/cr/v1"
	"github.com/grtl/mysql-operator/pkg/logging"
)

// Event represents an event processed by cluster controller.
type Event string

// Available event types.
const (
	ClusterAdded   Event = "Added"
	ClusterUpdated Event = "Updated"
	ClusterDeleted Event = "Deleted"
)

func logClusterEventBegin(cluster *crv1.MySQLCluster, event Event) {
	logging.LogCluster(cluster).WithField(
		"event", event).Info("Received cluster event")
}

func logClusterEventSuccess(cluster *crv1.MySQLCluster, event Event) {
	logging.LogCluster(cluster).WithField(
		"event", event).Info("Successfully processed cluster event")
}
