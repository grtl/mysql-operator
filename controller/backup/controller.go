package backup

import (
	"context"

	"k8s.io/client-go/tools/cache"

	"github.com/grtl/mysql-operator/controller"
	crv1 "github.com/grtl/mysql-operator/pkg/apis/cr/v1"
	"github.com/grtl/mysql-operator/pkg/client/clientset/versioned"
	"github.com/grtl/mysql-operator/pkg/client/informers/externalversions"
)

// NewBackupController returns new backup controller.
func NewBackupController(clientset versioned.Interface) controller.Controller {
	return &backupController{
		Base:      controller.NewControllerBase(),
		clientset: clientset,
	}
}

type backupController struct {
	controller.Base
	clientset versioned.Interface
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
	cluster := obj.(*crv1.MySQLBackup)

	// Run hooks
	for _, hook := range c.GetHooks() {
		hook.OnAdd(cluster)
	}
}

func (c *backupController) onUpdate(oldObj, newObj interface{}) {
	newCluster := newObj.(*crv1.MySQLBackup)

	// Run hooks
	for _, hook := range c.GetHooks() {
		hook.OnUpdate(newCluster)
	}
}

func (c *backupController) onDelete(obj interface{}) {
	cluster := obj.(*crv1.MySQLBackup)

	// Run hooks
	for _, hook := range c.GetHooks() {
		hook.OnDelete(cluster)
	}
}
