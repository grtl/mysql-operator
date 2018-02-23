package cluster

import (
	crv1 "github.com/grtl/mysql-operator/pkg/apis/cr/v1"
)

type fakeClusterOperator struct{}

// NewFakeOperator returns new operator that does nothing.
func NewFakeOperator() Operator {
	return new(fakeClusterOperator)
}

func (c *fakeClusterOperator) AddCluster(cluster *crv1.MySQLCluster) error {
	// Just pretend we're adding a cluster. Do nothing.
	return nil
}
