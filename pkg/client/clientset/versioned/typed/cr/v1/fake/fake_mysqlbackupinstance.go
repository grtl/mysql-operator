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
	cr_v1 "github.com/grtl/mysql-operator/pkg/apis/cr/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeMySQLBackupInstances implements MySQLBackupInstanceInterface
type FakeMySQLBackupInstances struct {
	Fake *FakeCrV1
	ns   string
}

var mysqlbackupinstancesResource = schema.GroupVersionResource{Group: "cr.mysqloperator.grtl.github.com", Version: "v1", Resource: "mysqlbackupinstances"}

var mysqlbackupinstancesKind = schema.GroupVersionKind{Group: "cr.mysqloperator.grtl.github.com", Version: "v1", Kind: "MySQLBackupInstance"}

// Get takes name of the mySQLBackupInstance, and returns the corresponding mySQLBackupInstance object, and an error if there is any.
func (c *FakeMySQLBackupInstances) Get(name string, options v1.GetOptions) (result *cr_v1.MySQLBackupInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(mysqlbackupinstancesResource, c.ns, name), &cr_v1.MySQLBackupInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*cr_v1.MySQLBackupInstance), err
}

// List takes label and field selectors, and returns the list of MySQLBackupInstances that match those selectors.
func (c *FakeMySQLBackupInstances) List(opts v1.ListOptions) (result *cr_v1.MySQLBackupInstanceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(mysqlbackupinstancesResource, mysqlbackupinstancesKind, c.ns, opts), &cr_v1.MySQLBackupInstanceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &cr_v1.MySQLBackupInstanceList{}
	for _, item := range obj.(*cr_v1.MySQLBackupInstanceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested mySQLBackupInstances.
func (c *FakeMySQLBackupInstances) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(mysqlbackupinstancesResource, c.ns, opts))

}

// Create takes the representation of a mySQLBackupInstance and creates it.  Returns the server's representation of the mySQLBackupInstance, and an error, if there is any.
func (c *FakeMySQLBackupInstances) Create(mySQLBackupInstance *cr_v1.MySQLBackupInstance) (result *cr_v1.MySQLBackupInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(mysqlbackupinstancesResource, c.ns, mySQLBackupInstance), &cr_v1.MySQLBackupInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*cr_v1.MySQLBackupInstance), err
}

// Update takes the representation of a mySQLBackupInstance and updates it. Returns the server's representation of the mySQLBackupInstance, and an error, if there is any.
func (c *FakeMySQLBackupInstances) Update(mySQLBackupInstance *cr_v1.MySQLBackupInstance) (result *cr_v1.MySQLBackupInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(mysqlbackupinstancesResource, c.ns, mySQLBackupInstance), &cr_v1.MySQLBackupInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*cr_v1.MySQLBackupInstance), err
}

// Delete takes name of the mySQLBackupInstance and deletes it. Returns an error if one occurs.
func (c *FakeMySQLBackupInstances) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(mysqlbackupinstancesResource, c.ns, name), &cr_v1.MySQLBackupInstance{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeMySQLBackupInstances) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(mysqlbackupinstancesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &cr_v1.MySQLBackupInstanceList{})
	return err
}

// Patch applies the patch and returns the patched mySQLBackupInstance.
func (c *FakeMySQLBackupInstances) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *cr_v1.MySQLBackupInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(mysqlbackupinstancesResource, c.ns, name, data, subresources...), &cr_v1.MySQLBackupInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*cr_v1.MySQLBackupInstance), err
}
