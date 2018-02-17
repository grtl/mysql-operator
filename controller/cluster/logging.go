package cluster

import (
	"github.com/sirupsen/logrus"

	"github.com/grtl/mysql-operator/controller"
	crv1 "github.com/grtl/mysql-operator/pkg/apis/cr/v1"
)

const (
	clusterAdded   = "Added"
	clusterUpdated = "Updated"
	clusterDeleted = "Deleted"
)

type loggingHook struct{}

// NewLoggingHook returns new Hook for cluster controller.
// LoggingHook logs given objects among with the event type.
func NewLoggingHook() controller.Hook {
	return new(loggingHook)
}

func (h *loggingHook) OnAdd(object interface{}) {
	cluster := object.(*crv1.MySQLCluster)
	logrus.WithFields(logrus.Fields{
		"cluster": cluster.Name,
		"event":   clusterAdded,
	}).Info("Received cluster event")
}

func (h *loggingHook) OnUpdate(object interface{}) {
	cluster := object.(*crv1.MySQLCluster)
	logrus.WithFields(logrus.Fields{
		"cluster": cluster.Name,
		"event":   clusterUpdated,
	}).Info("Received cluster event")
}

func (h *loggingHook) OnDelete(object interface{}) {
	cluster := object.(*crv1.MySQLCluster)
	logrus.WithFields(logrus.Fields{
		"cluster": cluster.Name,
		"event":   clusterDeleted,
	}).Info("Received cluster event")
}
