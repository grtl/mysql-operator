package controller

import (
	"context"

	"github.com/grtl/mysql-operator/operator"

	crv1 "github.com/grtl/mysql-operator/pkg/apis/cr/v1"
	"github.com/grtl/mysql-operator/pkg/client/clientset/versioned"
	"github.com/grtl/mysql-operator/pkg/client/informers/externalversions"
	corev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/tools/cache"
)

// ClusterEventType represents type of a ClusterEvent.
type ClusterEventType int

// Available ClusterEvent types.
const (
	ADDED ClusterEventType = iota
	UPDATED
	DELETED
)

// ClusterEvent is the way to inform about events processed by the controller.
type ClusterEvent struct {
	Type    ClusterEventType
	Cluster *crv1.MySQLCluster
}

// ClusterController processes events on MySQLCluster resources.
type ClusterController interface {
	// Run starts the event listeners.
	Run(ctx context.Context) error
	// GetEventsChan returns the channel consisting of events processed by the controller.
	GetEventsChan() <-chan ClusterEvent
}

// NewClusterController returns new cluster controller.
func NewClusterController(clientset versioned.Interface, corev1_client corev1.CoreV1Interface) ClusterController {
	events := make(chan ClusterEvent, clusterControllerEventsBufferSize)
	return &clusterController{
		clientset:     clientset,
		corev1_client: corev1_client,
		events:        events,
	}
}

type clusterController struct {
	clientset     versioned.Interface
	corev1_client corev1.CoreV1Interface
	events        chan ClusterEvent
}

const clusterControllerEventsBufferSize = 100

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

func (c *clusterController) GetEventsChan() <-chan ClusterEvent {
	return c.events
}

func (c *clusterController) onAdd(obj interface{}) {
	cluster := obj.(*crv1.MySQLCluster)
	c.events <- ClusterEvent{
		Type:    ADDED,
		Cluster: cluster,
	}

	operator.AddCluster(cluster, c.corev1_client)
}

func (c *clusterController) onUpdate(oldObj, newObj interface{}) {
	newCluster := newObj.(*crv1.MySQLCluster)
	c.events <- ClusterEvent{
		Type:    UPDATED,
		Cluster: newCluster,
	}
}

func (c *clusterController) onDelete(obj interface{}) {
	cluster := obj.(*crv1.MySQLCluster)
	c.events <- ClusterEvent{
		Type:    DELETED,
		Cluster: cluster,
	}
}
