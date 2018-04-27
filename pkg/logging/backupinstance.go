package logging

import (
	crv1 "github.com/grtl/mysql-operator/pkg/apis/cr/v1"
	"github.com/sirupsen/logrus"
)

// LogBackupInstance injects Backup Instance data into logrus fields.
func LogBackupInstance(backup *crv1.MySQLBackupInstance) *logrus.Entry {
	return logrus.WithFields(logrus.Fields{
		"backupInstance": backup.Name,
	})
}
