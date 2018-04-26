package backupschedule

import (
	"context"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"

	crv1 "github.com/grtl/mysql-operator/pkg/apis/cr/v1"
	"github.com/grtl/mysql-operator/pkg/client/clientset/versioned"
	"github.com/grtl/mysql-operator/pkg/client/informers/externalversions"
	"github.com/grtl/mysql-operator/pkg/controller"
	"github.com/grtl/mysql-operator/pkg/logging"
	operator "github.com/grtl/mysql-operator/pkg/operator/backupschedule"
)

// NewBackupScheduleController returns new BackupSchedule controller.
func NewBackupScheduleController(clientset versioned.Interface, kubeClientset kubernetes.Interface) controller.Controller {
	return &backupScheduleController{
		Base:      controller.NewControllerBase(),
		clientset: clientset,
		operator:  operator.NewBackupScheduleOperator(clientset, kubeClientset),
	}
}

type backupScheduleController struct {
	controller.Base
	clientset versioned.Interface
	operator  operator.Operator
}

func (c *backupScheduleController) Run(ctx context.Context) error {
	factory := externalversions.NewSharedInformerFactory(c.clientset, 0)
	informer := factory.Cr().V1().MySQLBackupSchedules().Informer()
	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    c.onAdd,
		UpdateFunc: c.onUpdate,
		DeleteFunc: c.onDelete,
	})
	informer.Run(ctx.Done())
	<-ctx.Done()
	return ctx.Err()
}

func (c *backupScheduleController) onAdd(obj interface{}) {
	schedule := obj.(*crv1.MySQLBackupSchedule)

	logBackupScheduleEventBegin(schedule, BackupScheduleAdded)

	err := c.operator.AddBackupSchedule(schedule)
	if err != nil {
		logging.LogBackupSchedule(schedule).WithField("event", BackupScheduleAdded).Error(err)
	} else {
		logBackupScheduleEventSuccess(schedule, BackupScheduleAdded)
	}

	// Run hooks
	for _, hook := range c.GetHooks() {
		hook.OnAdd(schedule)
	}
}

func (c *backupScheduleController) onUpdate(oldObj, newObj interface{}) {
	newSchedule := newObj.(*crv1.MySQLBackupSchedule)

	logBackupScheduleEventBegin(newSchedule, BackupScheduleUpdated)

	logBackupScheduleEventSuccess(newSchedule, BackupScheduleUpdated)

	// Run hooks
	for _, hook := range c.GetHooks() {
		hook.OnUpdate(newSchedule)
	}
}

func (c *backupScheduleController) onDelete(obj interface{}) {
	schedule := obj.(*crv1.MySQLBackupSchedule)

	logBackupScheduleEventBegin(schedule, BackupScheduleDeleted)

	logBackupScheduleEventSuccess(schedule, BackupScheduleDeleted)

	// Run hooks
	for _, hook := range c.GetHooks() {
		hook.OnDelete(schedule)
	}
}
