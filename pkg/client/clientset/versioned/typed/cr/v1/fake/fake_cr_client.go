/*
Copyright 2017 The MySQL Operator Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package fake

import (
	v1 "github.com/grtl/mysql-operator/pkg/client/clientset/versioned/typed/cr/v1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeCrV1 struct {
	*testing.Fake
}

func (c *FakeCrV1) MySQLBackupInstances(namespace string) v1.MySQLBackupInstanceInterface {
	return &FakeMySQLBackupInstances{c, namespace}
}

func (c *FakeCrV1) MySQLBackupSchedules(namespace string) v1.MySQLBackupScheduleInterface {
	return &FakeMySQLBackupSchedules{c, namespace}
}

func (c *FakeCrV1) MySQLClusters(namespace string) v1.MySQLClusterInterface {
	return &FakeMySQLClusters{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeCrV1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
