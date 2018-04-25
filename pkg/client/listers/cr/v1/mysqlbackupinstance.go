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

package v1

import (
	v1 "github.com/grtl/mysql-operator/pkg/apis/cr/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// MySQLBackupInstanceLister helps list MySQLBackupInstances.
type MySQLBackupInstanceLister interface {
	// List lists all MySQLBackupInstances in the indexer.
	List(selector labels.Selector) (ret []*v1.MySQLBackupInstance, err error)
	// MySQLBackupInstances returns an object that can list and get MySQLBackupInstances.
	MySQLBackupInstances(namespace string) MySQLBackupInstanceNamespaceLister
	MySQLBackupInstanceListerExpansion
}

// mySQLBackupInstanceLister implements the MySQLBackupInstanceLister interface.
type mySQLBackupInstanceLister struct {
	indexer cache.Indexer
}

// NewMySQLBackupInstanceLister returns a new MySQLBackupInstanceLister.
func NewMySQLBackupInstanceLister(indexer cache.Indexer) MySQLBackupInstanceLister {
	return &mySQLBackupInstanceLister{indexer: indexer}
}

// List lists all MySQLBackupInstances in the indexer.
func (s *mySQLBackupInstanceLister) List(selector labels.Selector) (ret []*v1.MySQLBackupInstance, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.MySQLBackupInstance))
	})
	return ret, err
}

// MySQLBackupInstances returns an object that can list and get MySQLBackupInstances.
func (s *mySQLBackupInstanceLister) MySQLBackupInstances(namespace string) MySQLBackupInstanceNamespaceLister {
	return mySQLBackupInstanceNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// MySQLBackupInstanceNamespaceLister helps list and get MySQLBackupInstances.
type MySQLBackupInstanceNamespaceLister interface {
	// List lists all MySQLBackupInstances in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.MySQLBackupInstance, err error)
	// Get retrieves the MySQLBackupInstance from the indexer for a given namespace and name.
	Get(name string) (*v1.MySQLBackupInstance, error)
	MySQLBackupInstanceNamespaceListerExpansion
}

// mySQLBackupInstanceNamespaceLister implements the MySQLBackupInstanceNamespaceLister
// interface.
type mySQLBackupInstanceNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all MySQLBackupInstances in the indexer for a given namespace.
func (s mySQLBackupInstanceNamespaceLister) List(selector labels.Selector) (ret []*v1.MySQLBackupInstance, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.MySQLBackupInstance))
	})
	return ret, err
}

// Get retrieves the MySQLBackupInstance from the indexer for a given namespace and name.
func (s mySQLBackupInstanceNamespaceLister) Get(name string) (*v1.MySQLBackupInstance, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("mysqlbackupinstance"), name)
	}
	return obj.(*v1.MySQLBackupInstance), nil
}
