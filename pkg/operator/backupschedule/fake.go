package backupschedule

import (
	crv1 "github.com/grtl/mysql-operator/pkg/apis/cr/v1"
)

type fakeBackupScheduleOperator struct{}

// NewFakeOperator returns new operator that does nothing.
func NewFakeOperator() Operator {
	return new(fakeBackupScheduleOperator)
}

func (b *fakeBackupScheduleOperator) AddBackupSchedule(backup *crv1.MySQLBackupSchedule) error {
	// Just pretend we're adding a new Backup Schedule. Do nothing.
	return nil
}
