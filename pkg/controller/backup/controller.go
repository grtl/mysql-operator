package backup

import (
	"context"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"

	crv1 "github.com/grtl/mysql-operator/pkg/apis/cr/v1"
	"github.com/grtl/mysql-operator/pkg/client/clientset/versioned"
	"github.com/grtl/mysql-operator/pkg/client/informers/externalversions"
	"github.com/grtl/mysql-operator/pkg/controller"
	operator "github.com/grtl/mysql-operator/pkg/operator/backup"
)

// NewBackupController returns new backup controller.
func NewBackupController(clientset versioned.Interface, kubeClientset kubernetes.Interface) controller.Controller {
	return &backupController{
		Base:           controller.NewControllerBase(),
		clientset:      clientset,
		backupOperator: operator.NewBackupOperator(clientset, kubeClientset),
	}
}

type backupController struct {
	controller.Base
	clientset      versioned.Interface
	backupOperator operator.Operator
}

func (c *backupController) Run(ctx context.Context) error {
	factory := externalversions.NewSharedInformerFactory(c.clientset, 0)
	informer := factory.Cr().V1().MySQLBackups().Informer()
	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    c.onAdd,
		UpdateFunc: c.onUpdate,
		DeleteFunc: c.onDelete,
	})
	informer.Run(ctx.Done())
	<-ctx.Done()
	return ctx.Err()
}

func (c *backupController) onAdd(obj interface{}) {
	backup := obj.(*crv1.MySQLBackup)

	logBackupEventBegin(backup, BackupAdded)

	err := c.backupOperator.ScheduleBackup(backup)
	if err != nil {
		return
	}

	logBackupEventSuccess(backup, BackupAdded)

	// Run hooks
	for _, hook := range c.GetHooks() {
		hook.OnAdd(backup)
	}
}

func (c *backupController) onUpdate(oldObj, newObj interface{}) {
	newBackup := newObj.(*crv1.MySQLBackup)

	logBackupEventBegin(newBackup, BackupUpdated)

	logBackupEventSuccess(newBackup, BackupUpdated)

	// Run hooks
	for _, hook := range c.GetHooks() {
		hook.OnUpdate(newBackup)
	}
}

func (c *backupController) onDelete(obj interface{}) {
	backup := obj.(*crv1.MySQLBackup)

	logBackupEventBegin(backup, BackupDeleted)

	logBackupEventSuccess(backup, BackupDeleted)

	// Run hooks
	for _, hook := range c.GetHooks() {
		hook.OnDelete(backup)
	}
}
