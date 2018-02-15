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

type clusterLoggingHook struct{}

func NewClusterLoggingHook() controller.ControllerHook {
	return &clusterLoggingHook{}
}

func (h *clusterLoggingHook) OnAdd(object interface{}) {
	cluster := object.(*crv1.MySQLCluster)
	logrus.WithFields(logrus.Fields{
		"cluster": cluster.Name,
		"event":   clusterAdded,
	}).Info("Received cluster event")
}

func (h *clusterLoggingHook) OnUpdate(object interface{}) {
	cluster := object.(*crv1.MySQLCluster)
	logrus.WithFields(logrus.Fields{
		"cluster": cluster.Name,
		"event":   clusterUpdated,
	}).Info("Received cluster event")
}

func (h *clusterLoggingHook) OnDelete(object interface{}) {
	cluster := object.(*crv1.MySQLCluster)
	logrus.WithFields(logrus.Fields{
		"cluster": cluster.Name,
		"event":   clusterDeleted,
	}).Info("Received cluster event")
}
