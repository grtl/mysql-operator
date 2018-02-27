package backup

import (
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/testing"

	"github.com/grtl/mysql-operator/controller"
	"github.com/grtl/mysql-operator/pkg/client/clientset/versioned/fake"
)

// NewFakeBackupController returns new operator controller among with prepended
// watcher. Created controller uses fake clientSets. Size indicates watcher events
// channel buffer.
func NewFakeBackupController(size int) (*watch.FakeWatcher, controller.Controller) {
	clientset := fake.NewSimpleClientset()

	watcher := watch.NewFakeWithChanSize(size, false)
	clientset.PrependWatchReactor("mysqlbackups", testing.DefaultWatchReactor(watcher, nil))
	fakeController := NewBackupController(clientset)
	return watcher, fakeController
}
