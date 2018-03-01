package logging

import (
	crv1 "github.com/grtl/mysql-operator/pkg/apis/cr/v1"
	"github.com/sirupsen/logrus"
)

// LogCluster injects cluster data into logrus fields.
func LogCluster(cluster *crv1.MySQLCluster) *logrus.Entry {
	return logrus.WithFields(logrus.Fields{
		"cluster": cluster.Name,
	})
}
