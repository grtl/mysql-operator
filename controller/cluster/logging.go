package cluster

import (
	"github.com/grtl/mysql-operator/logging"
	crv1 "github.com/grtl/mysql-operator/pkg/apis/cr/v1"
)

type clusterEvent string

const (
	clusterAdded   clusterEvent = "Added"
	clusterUpdated clusterEvent = "Updated"
	clusterDeleted clusterEvent = "Deleted"
)

func logClusterEventBegin(cluster *crv1.MySQLCluster, event clusterEvent) {
	logging.LogCluster(cluster).WithField(
		"event", event).Info("Received cluster event")
}

func logClusterEventSuccess(cluster *crv1.MySQLCluster, event clusterEvent) {
	logging.LogCluster(cluster).WithField(
		"event", event).Info("Successfully processed cluster event")
}
