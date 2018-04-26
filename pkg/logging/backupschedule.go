package logging

import (
	crv1 "github.com/grtl/mysql-operator/pkg/apis/cr/v1"
	"github.com/sirupsen/logrus"
)

// LogBackupSchedule injects Backup Schedule data into logrus fields.
func LogBackupSchedule(schedule *crv1.MySQLBackupSchedule) *logrus.Entry {
	return logrus.WithFields(logrus.Fields{
		"backupSchedule": schedule.Name,
	})
}
