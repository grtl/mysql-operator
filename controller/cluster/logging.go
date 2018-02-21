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

// LoggingHook extends `controller.Hook`. All processed events are logged.
type LoggingHook interface {
	controller.Hook
}

type loggingHook struct{}

// NewLoggingHook returns new Hook for cluster controller.
func NewLoggingHook() LoggingHook {
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
