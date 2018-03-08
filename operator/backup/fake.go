package backup

import (
	crv1 "github.com/grtl/mysql-operator/pkg/apis/cr/v1"
)

type fakeBackupOperator struct{}

// NewFakeOperator returns new operator that does nothing.
func NewFakeOperator() Operator {
	return new(fakeBackupOperator)
}

func (b *fakeBackupOperator) ScheduleBackup(backup *crv1.MySQLBackup) error {
	// Just pretend we're scheduling backup. Do nothing.
	return nil
}
