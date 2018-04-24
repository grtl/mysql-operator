package logging

import (
	crv1 "github.com/grtl/mysql-operator/pkg/apis/cr/v1"
	"github.com/sirupsen/logrus"
)

// LogBackup injects backup data into logrus fields.
func LogBackup(backup *crv1.MySQLBackup) *logrus.Entry {
	return logrus.WithFields(logrus.Fields{
		"backup": backup.Name,
	})
}
