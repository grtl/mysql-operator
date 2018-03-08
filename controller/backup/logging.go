package backup

import (
	"github.com/grtl/mysql-operator/logging"
	crv1 "github.com/grtl/mysql-operator/pkg/apis/cr/v1"
)

// Event represents an event processed by backup controller.
type Event string

// Available event types.
const (
	BackupAdded   Event = "Added"
	BackupUpdated Event = "Updated"
	BackupDeleted Event = "Deleted"
)

func logBackupEventBegin(backup *crv1.MySQLBackup, event Event) {
	logging.LogBackup(backup).WithField(
		"event", event).Info("Received backup event")
}

func logBackupEventSuccess(backup *crv1.MySQLBackup, event Event) {
	logging.LogBackup(backup).WithField(
		"event", event).Info("Successfully processed backup event")
}
