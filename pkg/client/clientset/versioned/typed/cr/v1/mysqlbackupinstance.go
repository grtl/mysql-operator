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
	scheme "github.com/grtl/mysql-operator/pkg/client/clientset/versioned/scheme"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// MySQLBackupInstancesGetter has a method to return a MySQLBackupInstanceInterface.
// A group's client should implement this interface.
type MySQLBackupInstancesGetter interface {
	MySQLBackupInstances(namespace string) MySQLBackupInstanceInterface
}

// MySQLBackupInstanceInterface has methods to work with MySQLBackupInstance resources.
type MySQLBackupInstanceInterface interface {
	Create(*v1.MySQLBackupInstance) (*v1.MySQLBackupInstance, error)
	Update(*v1.MySQLBackupInstance) (*v1.MySQLBackupInstance, error)
	Delete(name string, options *meta_v1.DeleteOptions) error
	DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error
	Get(name string, options meta_v1.GetOptions) (*v1.MySQLBackupInstance, error)
	List(opts meta_v1.ListOptions) (*v1.MySQLBackupInstanceList, error)
	Watch(opts meta_v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.MySQLBackupInstance, err error)
	MySQLBackupInstanceExpansion
}

// mySQLBackupInstances implements MySQLBackupInstanceInterface
type mySQLBackupInstances struct {
	client rest.Interface
	ns     string
}

// newMySQLBackupInstances returns a MySQLBackupInstances
func newMySQLBackupInstances(c *CrV1Client, namespace string) *mySQLBackupInstances {
	return &mySQLBackupInstances{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the mySQLBackupInstance, and returns the corresponding mySQLBackupInstance object, and an error if there is any.
func (c *mySQLBackupInstances) Get(name string, options meta_v1.GetOptions) (result *v1.MySQLBackupInstance, err error) {
	result = &v1.MySQLBackupInstance{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("mysqlbackupinstances").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of MySQLBackupInstances that match those selectors.
func (c *mySQLBackupInstances) List(opts meta_v1.ListOptions) (result *v1.MySQLBackupInstanceList, err error) {
	result = &v1.MySQLBackupInstanceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("mysqlbackupinstances").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested mySQLBackupInstances.
func (c *mySQLBackupInstances) Watch(opts meta_v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("mysqlbackupinstances").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a mySQLBackupInstance and creates it.  Returns the server's representation of the mySQLBackupInstance, and an error, if there is any.
func (c *mySQLBackupInstances) Create(mySQLBackupInstance *v1.MySQLBackupInstance) (result *v1.MySQLBackupInstance, err error) {
	result = &v1.MySQLBackupInstance{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("mysqlbackupinstances").
		Body(mySQLBackupInstance).
		Do().
		Into(result)
	return
}

// Update takes the representation of a mySQLBackupInstance and updates it. Returns the server's representation of the mySQLBackupInstance, and an error, if there is any.
func (c *mySQLBackupInstances) Update(mySQLBackupInstance *v1.MySQLBackupInstance) (result *v1.MySQLBackupInstance, err error) {
	result = &v1.MySQLBackupInstance{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("mysqlbackupinstances").
		Name(mySQLBackupInstance.Name).
		Body(mySQLBackupInstance).
		Do().
		Into(result)
	return
}

// Delete takes name of the mySQLBackupInstance and deletes it. Returns an error if one occurs.
func (c *mySQLBackupInstances) Delete(name string, options *meta_v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("mysqlbackupinstances").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *mySQLBackupInstances) DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("mysqlbackupinstances").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched mySQLBackupInstance.
func (c *mySQLBackupInstances) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.MySQLBackupInstance, err error) {
	result = &v1.MySQLBackupInstance{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("mysqlbackupinstances").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
