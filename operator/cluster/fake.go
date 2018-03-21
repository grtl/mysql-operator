package cluster

import (
	crv1 "github.com/grtl/mysql-operator/pkg/apis/cr/v1"
)

// FakeClusterOperator may be used in tests as a cluster operator.
// It implements all operator functions as simply returning an error
// set via SetError.
type FakeClusterOperator struct {
	err error
}

// NewFakeOperator returns new operator that does nothing.
func NewFakeOperator() *FakeClusterOperator {
	return new(FakeClusterOperator)
}

// SetError can be used to simulate operator failures.
func (c *FakeClusterOperator) SetError(err error) {
	c.err = err
}

// AddCluster simulates adding a cluster. Returns error set via SetError.
func (c *FakeClusterOperator) AddCluster(cluster *crv1.MySQLCluster) error {
	// Just pretend we're adding a cluster. Return err.
	return c.err
}

// UpdateCluster simulates updating a cluster. Returns error set via SetError.
func (c *FakeClusterOperator) UpdateCluster(newCluster *crv1.MySQLCluster) error {
	// Just pretend we're updating a cluster. Do nothing.
	return c.err
}
