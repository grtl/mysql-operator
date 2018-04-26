package backupinstance

import (
	crv1 "github.com/grtl/mysql-operator/pkg/apis/cr/v1"
)

// FakeBackupInstanceOperator may be used in tests as a backup instance operator.
// It implements all operator functions as simply returning an fail set via SetError.
type FakeBackupInstanceOperator struct {
	err error
}

// NewFakeBackupInstanceOperator returns new operator that does nothing.
func NewFakeBackupInstanceOperator() *FakeBackupInstanceOperator {
	return new(FakeBackupInstanceOperator)
}

// SetError can be used to simulate operator failures.
func (c *FakeBackupInstanceOperator) SetError(err error) {
	c.err = err
}

// CreateBackup simulates creating a backup. Returns error set via SetError.
func (c *FakeBackupInstanceOperator) CreateBackup(backup *crv1.MySQLBackupInstance) error {
	// Just pretend we're adding a cluster. Return err.
	return c.err
}

// DeleteBackup simulates deleting a backup. Returns error set via SetError.
func (c *FakeBackupInstanceOperator) DeleteBackup(backup *crv1.MySQLBackupInstance) error {
	// Just pretend we're deleting a cluster. Do nothing.
	return c.err
}
