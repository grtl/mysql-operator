package cluster

import (
	"k8s.io/apimachinery/pkg/watch"
	kubeFake "k8s.io/client-go/kubernetes/fake"
	"k8s.io/client-go/testing"

	"github.com/grtl/mysql-operator/pkg/client/clientset/versioned/fake"
	operator "github.com/grtl/mysql-operator/pkg/operator/cluster"
)

// FakeClusterController can be used as a cluster controller in tests.
// It simulates the cluster controller behaviour with fake clientsets
// and a fake cluster operator.
type FakeClusterController struct {
	*clusterController
	*operator.FakeClusterOperator
}

// NewFakeClusterController returns new operator controller among with prepended
// watcher. Created controller uses fake clientSets and operator. Size indicates
// watcher events channel buffer.
func NewFakeClusterController(size int) (*watch.FakeWatcher, *FakeClusterController) {
	kubeClientset := kubeFake.NewSimpleClientset()
	clientset := fake.NewSimpleClientset()

	watcher := watch.NewFakeWithChanSize(size, false)
	clientset.PrependWatchReactor("mysqlclusters", testing.DefaultWatchReactor(watcher, nil))

	fakeController := NewClusterController(clientset, kubeClientset).(*clusterController)
	fakeOperator := operator.NewFakeOperator()
	fakeController.clusterOperator = fakeOperator
	return watcher, &FakeClusterController{
		clusterController:   fakeController,
		FakeClusterOperator: fakeOperator,
	}
}
