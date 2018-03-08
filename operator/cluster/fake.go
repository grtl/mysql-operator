package cluster

import (
	crv1 "github.com/grtl/mysql-operator/pkg/apis/cr/v1"
)

type FakeClusterOperator struct {
	err error
}

// NewFakeOperator returns new operator that does nothing.
func NewFakeOperator() *FakeClusterOperator {
	return new(FakeClusterOperator)
}

func (c *FakeClusterOperator) AddCluster(cluster *crv1.MySQLCluster) error {
	// Just pretend we're adding a cluster. Return err.
	return c.err
}

// SetError can be used to simulate operator failures.
func (c *FakeClusterOperator) SetError(err error) {
	c.err = err
}

func (c *FakeClusterOperator) UpdateCluster(oldCluster, newCluster *crv1.MySQLCluster) error {
	// Just pretend we're updating a cluster. Do nothing.
	return c.err
}
