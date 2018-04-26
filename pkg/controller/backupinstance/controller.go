package backupinstance

import (
	"context"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"

	crv1 "github.com/grtl/mysql-operator/pkg/apis/cr/v1"
	"github.com/grtl/mysql-operator/pkg/client/clientset/versioned"
	"github.com/grtl/mysql-operator/pkg/client/informers/externalversions"
	"github.com/grtl/mysql-operator/pkg/controller"
	"github.com/grtl/mysql-operator/pkg/logging"
	"github.com/grtl/mysql-operator/pkg/operator/backupinstance"
)

// NewBackupInstanceController returns new backup instance controller.
func NewBackupInstanceController(clientset versioned.Interface, kubeClientset kubernetes.Interface) controller.Controller {
	return &backupInstanceController{
		Base:      controller.NewControllerBase(),
		clientset: clientset,
		operator:  backupinstance.NewBackupInstanceOperator(clientset, kubeClientset),
	}
}

type backupInstanceController struct {
	controller.Base
	clientset versioned.Interface
	operator  backupinstance.Operator
}

func (c *backupInstanceController) Run(ctx context.Context) error {
	factory := externalversions.NewSharedInformerFactory(c.clientset, 0)
	informer := factory.Cr().V1().MySQLBackupInstances().Informer()
	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    c.onAdd,
		UpdateFunc: c.onUpdate,
		DeleteFunc: c.onDelete,
	})
	informer.Run(ctx.Done())
	<-ctx.Done()
	return ctx.Err()
}

func (c *backupInstanceController) onAdd(obj interface{}) {
	backup := obj.(*crv1.MySQLBackupInstance)

	logBackupInstanceEventBegin(backup, BackupInstanceAdded)

	err := c.operator.CreateBackup(backup)
	if err != nil {
		logging.LogBackupInstance(backup).WithField("event", BackupInstanceAdded).Error(err)
	} else {
		logBackupInstanceEventSuccess(backup, BackupInstanceAdded)
	}

	// Run hooks
	for _, hook := range c.GetHooks() {
		hook.OnAdd(backup)
	}
}

func (c *backupInstanceController) onUpdate(oldObj, newObj interface{}) {
	newBackup := newObj.(*crv1.MySQLBackupInstance)

	logBackupInstanceEventBegin(newBackup, BackupInstanceUpdated)

	logBackupInstanceEventSuccess(newBackup, BackupInstanceUpdated)

	// Run hooks
	for _, hook := range c.GetHooks() {
		hook.OnUpdate(newBackup)
	}
}

func (c *backupInstanceController) onDelete(obj interface{}) {
	backup := obj.(*crv1.MySQLBackupInstance)

	logBackupInstanceEventBegin(backup, BackupInstanceDeleted)

	err := c.operator.DeleteBackup(backup)
	if err != nil {
		logging.LogBackupInstance(backup).WithField("event", BackupInstanceDeleted).Error(err)
	} else {
		logBackupInstanceEventSuccess(backup, BackupInstanceDeleted)
	}

	// Run hooks
	for _, hook := range c.GetHooks() {
		hook.OnDelete(backup)
	}
}
