package backupschedule

import (
	crv1 "github.com/grtl/mysql-operator/pkg/apis/cr/v1"
	"github.com/grtl/mysql-operator/pkg/logging"
)

// Event represents an event processed by the Backup Schedule controller.
type Event string

// Available event types.
const (
	BackupScheduleAdded   Event = "Added"
	BackupScheduleUpdated Event = "Updated"
	BackupScheduleDeleted Event = "Deleted"
)

func logBackupScheduleEventBegin(schedule *crv1.MySQLBackupSchedule, event Event) {
	logging.LogBackupSchedule(schedule).WithField(
		"event", event).Info("Received BackupSchedule event")
}

func logBackupScheduleEventSuccess(schedule *crv1.MySQLBackupSchedule, event Event) {
	logging.LogBackupSchedule(schedule).WithField(
		"event", event).Info("Successfully processed BackupSchedule event")
}
