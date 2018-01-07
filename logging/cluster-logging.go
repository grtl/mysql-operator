package logging

import (
	"github.com/grtl/mysql-operator/controller"
	"github.com/sirupsen/logrus"
)

func logClusterEvent(event controller.ClusterEvent) {
	logrus.WithFields(logrus.Fields{
		"cluster": event.Cluster.Name,
		"event":   event.Type,
	}).Info("Received cluster event")
}
