package backup

import (
	"context"

	"k8s.io/client-go/tools/cache"

	"github.com/grtl/mysql-operator/controller"
	operator "github.com/grtl/mysql-operator/operator/backup"
	crv1 "github.com/grtl/mysql-operator/pkg/apis/cr/v1"
	"github.com/grtl/mysql-operator/pkg/client/clientset/versioned"
	"github.com/grtl/mysql-operator/pkg/client/informers/externalversions"
	"k8s.io/client-go/kubernetes"
)

// NewBackupController returns new backup controller.
func NewBackupController(clientset versioned.Interface, kubeClientset kubernetes.Interface) controller.Controller {
	return &backupController{
		Base:      controller.NewControllerBase(),
		clientset: clientset,
		backupOperator: operator.NewBackupOperator(kubeClientset),
	}
}

type backupController struct {
	controller.Base
	clientset versioned.Interface
	backupOperator operator.Operator
}

func (b *backupController) Run(ctx context.Context) error {
	factory := externalversions.NewSharedInformerFactory(b.clientset, 0)
	informer := factory.Cr().V1().MySQLBackups().Informer()
	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    b.onAdd,
		UpdateFunc: b.onUpdate,
		DeleteFunc: b.onDelete,
	})
	informer.Run(ctx.Done())
	<-ctx.Done()
	return ctx.Err()
}

func (b *backupController) onAdd(obj interface{}) {
	cluster := obj.(*crv1.MySQLBackup)

	err := b.backupOperator.ScheduleBackup(cluster)
	if err != nil {
		return
	}

	// Run hooks
	for _, hook := range b.GetHooks() {
		hook.OnAdd(cluster)
	}
}

func (b *backupController) onUpdate(oldObj, newObj interface{}) {
	newCluster := newObj.(*crv1.MySQLBackup)

	// Run hooks
	for _, hook := range b.GetHooks() {
		hook.OnUpdate(newCluster)
	}
}

func (b *backupController) onDelete(obj interface{}) {
	cluster := obj.(*crv1.MySQLBackup)

	// Run hooks
	for _, hook := range b.GetHooks() {
		hook.OnDelete(cluster)
	}
}
