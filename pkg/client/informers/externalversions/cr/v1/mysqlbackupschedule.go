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

// MySQLBackupScheduleInformer provides access to a shared informer and lister for
// MySQLBackupSchedules.
type MySQLBackupScheduleInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.MySQLBackupScheduleLister
}

type mySQLBackupScheduleInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewMySQLBackupScheduleInformer constructs a new informer for MySQLBackupSchedule type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewMySQLBackupScheduleInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredMySQLBackupScheduleInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredMySQLBackupScheduleInformer constructs a new informer for MySQLBackupSchedule type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredMySQLBackupScheduleInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options meta_v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CrV1().MySQLBackupSchedules(namespace).List(options)
			},
			WatchFunc: func(options meta_v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CrV1().MySQLBackupSchedules(namespace).Watch(options)
			},
		},
		&cr_v1.MySQLBackupSchedule{},
		resyncPeriod,
		indexers,
	)
}

func (f *mySQLBackupScheduleInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredMySQLBackupScheduleInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *mySQLBackupScheduleInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&cr_v1.MySQLBackupSchedule{}, f.defaultInformer)
}

func (f *mySQLBackupScheduleInformer) Lister() v1.MySQLBackupScheduleLister {
	return v1.NewMySQLBackupScheduleLister(f.Informer().GetIndexer())
}
