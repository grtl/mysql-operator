package backupinstance

import (
	"k8s.io/apimachinery/pkg/watch"
	kubeFake "k8s.io/client-go/kubernetes/fake"
	"k8s.io/client-go/testing"

	"github.com/grtl/mysql-operator/pkg/client/clientset/versioned/fake"
	"github.com/grtl/mysql-operator/pkg/controller"
)

// NewFakeBackupInstanceController returns new operator controller among with
// prepended watcher. Created controller uses fake clientSets. Size indicates
// watcher events channel buffer.
func NewFakeBackupInstanceController(size int) (*watch.FakeWatcher, controller.Controller) {
	kubeClientset := kubeFake.NewSimpleClientset()
	clientset := fake.NewSimpleClientset()

	watcher := watch.NewFakeWithChanSize(size, false)
	clientset.PrependWatchReactor("mysqlbackupinstances", testing.DefaultWatchReactor(watcher, nil))
	fakeController := NewBackupInstanceController(clientset, kubeClientset)
	return watcher, fakeController
}
