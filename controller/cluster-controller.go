package controller

import (
	"context"
	"fmt"

	"k8s.io/client-go/tools/cache"

	crv1 "github.com/grtl/mysql-operator/pkg/apis/cr/v1"
	"github.com/grtl/mysql-operator/pkg/client/clientset/versioned"
	"github.com/grtl/mysql-operator/pkg/client/informers/externalversions"
)

// ClusterController observes changes on MySQLCluster custom resource and reacts to them.
type ClusterController struct {
	Clientset *versioned.Clientset
}

// Run starts the controller thread.
func (c *ClusterController) Run(ctx context.Context) error {
	factory := externalversions.NewSharedInformerFactory(c.Clientset, 0)
	informer := factory.Cr().V1().MySQLClusters().Informer()
	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    c.onAdd,
		UpdateFunc: c.onUpdate,
		DeleteFunc: c.onDelete,
	})
	informer.Run(ctx.Done())
	<-ctx.Done()
	return ctx.Err()
}

func (c *ClusterController) onAdd(obj interface{}) {
	cluster := obj.(*crv1.MySQLCluster)
	fmt.Printf("On create %s\n", cluster.Spec.Name)
}

func (c *ClusterController) onUpdate(oldObj, newObj interface{}) {
	cluster := oldObj.(*crv1.MySQLCluster)
	fmt.Printf("On update %s\n", cluster.Spec.Name)
}

func (c *ClusterController) onDelete(obj interface{}) {
	cluster := obj.(*crv1.MySQLCluster)
	fmt.Printf("On delete %s\n", cluster.Spec.Name)
}
