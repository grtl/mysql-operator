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
	time "time"

	cr_v1 "github.com/grtl/mysql-operator/pkg/apis/cr/v1"
	versioned "github.com/grtl/mysql-operator/pkg/client/clientset/versioned"
	internalinterfaces "github.com/grtl/mysql-operator/pkg/client/informers/externalversions/internalinterfaces"
	v1 "github.com/grtl/mysql-operator/pkg/client/listers/cr/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// MySQLBackupInstanceInformer provides access to a shared informer and lister for
// MySQLBackupInstances.
type MySQLBackupInstanceInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.MySQLBackupInstanceLister
}

type mySQLBackupInstanceInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewMySQLBackupInstanceInformer constructs a new informer for MySQLBackupInstance type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewMySQLBackupInstanceInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredMySQLBackupInstanceInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredMySQLBackupInstanceInformer constructs a new informer for MySQLBackupInstance type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredMySQLBackupInstanceInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options meta_v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CrV1().MySQLBackupInstances(namespace).List(options)
			},
			WatchFunc: func(options meta_v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CrV1().MySQLBackupInstances(namespace).Watch(options)
			},
		},
		&cr_v1.MySQLBackupInstance{},
		resyncPeriod,
		indexers,
	)
}

func (f *mySQLBackupInstanceInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredMySQLBackupInstanceInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *mySQLBackupInstanceInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&cr_v1.MySQLBackupInstance{}, f.defaultInformer)
}

func (f *mySQLBackupInstanceInformer) Lister() v1.MySQLBackupInstanceLister {
	return v1.NewMySQLBackupInstanceLister(f.Informer().GetIndexer())
}
