package cluster

import (
	"context"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"

	"github.com/grtl/mysql-operator/controller"
	"github.com/grtl/mysql-operator/logging"
	operator "github.com/grtl/mysql-operator/operator/cluster"
	crv1 "github.com/grtl/mysql-operator/pkg/apis/cr/v1"
	"github.com/grtl/mysql-operator/pkg/client/clientset/versioned"
	"github.com/grtl/mysql-operator/pkg/client/informers/externalversions"
)

// NewClusterController returns new cluster controller.
func NewClusterController(clientset versioned.Interface, kubeClientset kubernetes.Interface) controller.Controller {
	return &clusterController{
		Base:            controller.NewControllerBase(),
		clientset:       clientset,
		clusterOperator: operator.NewClusterOperator(kubeClientset),
	}
}

type clusterController struct {
	controller.Base
	clientset       versioned.Interface
	clusterOperator operator.Operator
}

func (c *clusterController) Run(ctx context.Context) error {
	factory := externalversions.NewSharedInformerFactory(c.clientset, 0)
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

func (c *clusterController) onAdd(obj interface{}) {
	cluster := obj.(*crv1.MySQLCluster)

	logClusterEventBegin(cluster, ClusterAdded)

	err := c.clusterOperator.AddCluster(cluster)
	if err != nil {
		logging.LogCluster(cluster).WithField("event", ClusterAdded).Error(err)
	} else {
		logClusterEventSuccess(cluster, ClusterAdded)
	}

	// Run hooks
	for _, hook := range c.GetHooks() {
		hook.OnAdd(cluster)
	}
}

func (c *clusterController) onUpdate(oldObj, newObj interface{}) {
	oldCluster := oldObj.(*crv1.MySQLCluster)
	newCluster := newObj.(*crv1.MySQLCluster)

	logClusterEventBegin(newCluster, ClusterUpdated)

	err := c.clusterOperator.UpdateCluster(oldCluster, newCluster)
	if err != nil {
		logging.LogCluster(newCluster).WithField("event", ClusterUpdated).Error(err)
	} else {
		logClusterEventSuccess(newCluster, ClusterUpdated)
	}

	// Run hooks
	for _, hook := range c.GetHooks() {
		hook.OnUpdate(newCluster)
	}
}

func (c *clusterController) onDelete(obj interface{}) {
	cluster := obj.(*crv1.MySQLCluster)

	logClusterEventBegin(cluster, ClusterDeleted)

	logClusterEventSuccess(cluster, ClusterDeleted)

	// Run hooks
	for _, hook := range c.GetHooks() {
		hook.OnDelete(cluster)
	}
}
