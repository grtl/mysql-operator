package cluster

import (
	"context"
	"errors"

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
		clientset:       clientset,
		clusterOperator: operator.NewClusterOperator(kubeClientset),
		hooks:           []controller.Hook{},
	}
}

type clusterController struct {
	clientset       versioned.Interface
	clusterOperator operator.Operator
	hooks           []controller.Hook
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

func (c *clusterController) AddHook(hook controller.Hook) error {
	for _, h := range c.hooks {
		if h == hook {
			return errors.New("Given hook is already installed in the current controller")
		}
	}
	c.hooks = append(c.hooks, hook)
	return nil
}

func (c *clusterController) RemoveHook(hook controller.Hook) error {
	for i, h := range c.hooks {
		if h == hook {
			// Removing hooks is not that common so we can afford it in O(n)
			c.hooks = append(c.hooks[:i], c.hooks[i+1:]...)
			return nil
		}
	}
	return errors.New("Given hook is not installed in the current controller")
}

func (c *clusterController) onAdd(obj interface{}) {
	cluster := obj.(*crv1.MySQLCluster)

	logClusterEventBegin(cluster, clusterAdded)

	err := c.clusterOperator.AddCluster(cluster)
	if err != nil {
		logging.LogCluster(cluster).WithField("event", clusterAdded).Error(err)
		return
	}

	logClusterEventSuccess(cluster, clusterAdded)

	// Run hooks
	for _, hook := range c.hooks {
		hook.OnAdd(cluster)
	}
}

func (c *clusterController) onUpdate(oldObj, newObj interface{}) {
	newCluster := newObj.(*crv1.MySQLCluster)

	logClusterEventBegin(newCluster, clusterUpdated)

	logClusterEventSuccess(newCluster, clusterUpdated)

	// Run hooks
	for _, hook := range c.hooks {
		hook.OnUpdate(newCluster)
	}
}

func (c *clusterController) onDelete(obj interface{}) {
	cluster := obj.(*crv1.MySQLCluster)

	logClusterEventBegin(cluster, clusterDeleted)

	logClusterEventSuccess(cluster, clusterDeleted)

	// Run hooks
	for _, hook := range c.hooks {
		hook.OnDelete(cluster)
	}
}
