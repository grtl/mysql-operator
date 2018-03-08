package backup

import (
	"github.com/grtl/mysql-operator/logging"
	crv1 "github.com/grtl/mysql-operator/pkg/apis/cr/v1"
)

type backupEvent string

const (
	backupAdded   backupEvent = "Added"
	backupUpdated backupEvent = "Updated"
	backupDeleted backupEvent = "Deleted"
)

func logBackupEventBegin(backup *crv1.MySQLBackup, event backupEvent) {
	logging.LogBackup(backup).WithField(
		"event", event).Info("Received backup event")
}

func logBackupEventSuccess(backup *crv1.MySQLBackup, event backupEvent) {
	logging.LogBackup(backup).WithField(
		"event", event).Info("Successfully processed backup event")
}
