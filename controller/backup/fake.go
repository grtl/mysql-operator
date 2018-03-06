package backup

import (
	"k8s.io/apimachinery/pkg/watch"
	kubeFake "k8s.io/client-go/kubernetes/fake"
	"k8s.io/client-go/testing"

	"github.com/grtl/mysql-operator/controller"
	operator "github.com/grtl/mysql-operator/operator/backup"
	"github.com/grtl/mysql-operator/pkg/client/clientset/versioned/fake"
)

// NewFakeBackupController returns new operator controller among with prepended
// watcher. Created controller uses fake clientSets. Size indicates watcher events
// channel buffer.
func NewFakeBackupController(size int) (*watch.FakeWatcher, controller.Controller) {
	kubeClientset := kubeFake.NewSimpleClientset()
	clientset := fake.NewSimpleClientset()

	watcher := watch.NewFakeWithChanSize(size, false)
	clientset.PrependWatchReactor("mysqlbackups", testing.DefaultWatchReactor(watcher, nil))
	fakeController := NewBackupController(clientset, kubeClientset)
	fakeController.(*backupController).backupOperator = operator.NewFakeOperator()
	return watcher, fakeController
}
