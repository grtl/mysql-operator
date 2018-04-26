package backupinstance

import (
	crv1 "github.com/grtl/mysql-operator/pkg/apis/cr/v1"
	"github.com/grtl/mysql-operator/pkg/logging"
)

// Event represents an event processed by the Backup Schedule controller.
type Event string

// Available event types.
const (
	BackupInstanceAdded   Event = "Added"
	BackupInstanceUpdated Event = "Updated"
	BackupInstanceDeleted Event = "Deleted"
)

func logBackupInstanceEventBegin(backup *crv1.MySQLBackupInstance, event Event) {
	logging.LogBackupInstance(backup).WithField(
		"event", event).Info("Received BackupInstance event")
}

func logBackupInstanceEventSuccess(backup *crv1.MySQLBackupInstance, event Event) {
	logging.LogBackupInstance(backup).WithField(
		"event", event).Info("Successfully processed BackupInstance event")
}
