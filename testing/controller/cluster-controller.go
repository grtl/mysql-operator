package controller

import (
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/testing"

	"github.com/grtl/mysql-operator/controller"
	"github.com/grtl/mysql-operator/pkg/client/clientset/versioned/fake"
	kFake "k8s.io/client-go/kubernetes/fake"
)

// NewFakeClusterController returns new fake cluster with prepended watcher on MySQLCluster resource
func NewFakeClusterController() (*watch.FakeWatcher, controller.ClusterController) {
	kClientset := kFake.NewSimpleClientset()

	clientset := fake.NewSimpleClientset()
	watcher := watch.NewFake()
	clientset.PrependWatchReactor("mysqlclusters", testing.DefaultWatchReactor(watcher, nil))
	clusterController := controller.NewClusterController(clientset, kClientset)
	return watcher, clusterController
}
