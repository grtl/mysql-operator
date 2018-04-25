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

// FakeMySQLBackupSchedules implements MySQLBackupScheduleInterface
type FakeMySQLBackupSchedules struct {
	Fake *FakeCrV1
	ns   string
}

var mysqlbackupschedulesResource = schema.GroupVersionResource{Group: "cr.mysqloperator.grtl.github.com", Version: "v1", Resource: "mysqlbackupschedules"}

var mysqlbackupschedulesKind = schema.GroupVersionKind{Group: "cr.mysqloperator.grtl.github.com", Version: "v1", Kind: "MySQLBackupSchedule"}

// Get takes name of the mySQLBackupSchedule, and returns the corresponding mySQLBackupSchedule object, and an error if there is any.
func (c *FakeMySQLBackupSchedules) Get(name string, options v1.GetOptions) (result *cr_v1.MySQLBackupSchedule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(mysqlbackupschedulesResource, c.ns, name), &cr_v1.MySQLBackupSchedule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*cr_v1.MySQLBackupSchedule), err
}

// List takes label and field selectors, and returns the list of MySQLBackupSchedules that match those selectors.
func (c *FakeMySQLBackupSchedules) List(opts v1.ListOptions) (result *cr_v1.MySQLBackupScheduleList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(mysqlbackupschedulesResource, mysqlbackupschedulesKind, c.ns, opts), &cr_v1.MySQLBackupScheduleList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &cr_v1.MySQLBackupScheduleList{}
	for _, item := range obj.(*cr_v1.MySQLBackupScheduleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested mySQLBackupSchedules.
func (c *FakeMySQLBackupSchedules) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(mysqlbackupschedulesResource, c.ns, opts))

}

// Create takes the representation of a mySQLBackupSchedule and creates it.  Returns the server's representation of the mySQLBackupSchedule, and an error, if there is any.
func (c *FakeMySQLBackupSchedules) Create(mySQLBackupSchedule *cr_v1.MySQLBackupSchedule) (result *cr_v1.MySQLBackupSchedule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(mysqlbackupschedulesResource, c.ns, mySQLBackupSchedule), &cr_v1.MySQLBackupSchedule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*cr_v1.MySQLBackupSchedule), err
}

// Update takes the representation of a mySQLBackupSchedule and updates it. Returns the server's representation of the mySQLBackupSchedule, and an error, if there is any.
func (c *FakeMySQLBackupSchedules) Update(mySQLBackupSchedule *cr_v1.MySQLBackupSchedule) (result *cr_v1.MySQLBackupSchedule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(mysqlbackupschedulesResource, c.ns, mySQLBackupSchedule), &cr_v1.MySQLBackupSchedule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*cr_v1.MySQLBackupSchedule), err
}

// Delete takes name of the mySQLBackupSchedule and deletes it. Returns an error if one occurs.
func (c *FakeMySQLBackupSchedules) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(mysqlbackupschedulesResource, c.ns, name), &cr_v1.MySQLBackupSchedule{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeMySQLBackupSchedules) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(mysqlbackupschedulesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &cr_v1.MySQLBackupScheduleList{})
	return err
}

// Patch applies the patch and returns the patched mySQLBackupSchedule.
func (c *FakeMySQLBackupSchedules) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *cr_v1.MySQLBackupSchedule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(mysqlbackupschedulesResource, c.ns, name, data, subresources...), &cr_v1.MySQLBackupSchedule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*cr_v1.MySQLBackupSchedule), err
}
