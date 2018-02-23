package cluster

import (
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/errors"
	"k8s.io/client-go/kubernetes"

	"github.com/grtl/mysql-operator/logging"
	crv1 "github.com/grtl/mysql-operator/pkg/apis/cr/v1"
	"github.com/grtl/mysql-operator/util"
)

const (
	serviceTemplate     = "artifacts/mysql-service.yaml"
	statefulSetTemplate = "artifacts/mysql-statefulset.yaml"
)

// Operator represents an object to manipulate MySQLCluster custom resources.
type Operator interface {
	// AddCluster creates the Kubernetes API objects necessary for a MySQL cluster.
	AddCluster(cluster *crv1.MySQLCluster) error
}

type clusterOperator struct {
	clientset kubernetes.Interface
}

// NewClusterOperator returns a new Operator.
func NewClusterOperator(clientset kubernetes.Interface) Operator {
	return &clusterOperator{
		clientset: clientset,
	}
}

func (c *clusterOperator) AddCluster(cluster *crv1.MySQLCluster) error {
	logging.LogCluster(cluster).Debug("Creating service.")

	err := c.createService(cluster)
	if err != nil {
		return err
	}

	logging.LogCluster(cluster).Debug("Creating stateful set.")

	err = c.createStatefulSet(cluster)
	if err != nil {
		// Cleanup - remove already created service
		logging.LogCluster(cluster).WithField(
			"error", err).Warn("Reverting service creation.")
		removeErr := c.removeService(cluster)
		return errors.NewAggregate([]error{err, removeErr})
	}

	return nil
}

func (c *clusterOperator) createService(cluster *crv1.MySQLCluster) error {
	serviceInterface := c.clientset.CoreV1().Services(cluster.Namespace)
	service, err := serviceForCluster(cluster)
	if err != nil {
		return err
	}

	_, err = serviceInterface.Create(service)
	if err != nil && !apierrors.IsAlreadyExists(err) {
		return err
	} else if apierrors.IsAlreadyExists(err) {
		logging.LogCluster(cluster).Warn("Service for cluster already exists")
	}

	return nil
}

func (c *clusterOperator) createStatefulSet(cluster *crv1.MySQLCluster) error {
	statefulSetInterface := c.clientset.AppsV1().StatefulSets(cluster.Namespace)
	statefulSet, err := statefulSetForCluster(cluster)
	if err != nil {
		return err
	}

	_, err = statefulSetInterface.Create(statefulSet)
	if err != nil && !apierrors.IsAlreadyExists(err) {
		return err
	} else if apierrors.IsAlreadyExists(err) {
		logging.LogCluster(cluster).Warn("StatefulSet for cluster already exists")
	}

	return nil
}

func serviceForCluster(cluster *crv1.MySQLCluster) (*corev1.Service, error) {
	service := new(corev1.Service)
	err := util.ObjectFromTemplate(cluster, service, serviceTemplate)
	return service, err
}

func statefulSetForCluster(cluster *crv1.MySQLCluster) (*appsv1.StatefulSet, error) {
	statefulSet := new(appsv1.StatefulSet)
	err := util.ObjectFromTemplate(cluster, statefulSet, statefulSetTemplate)
	return statefulSet, err
}

func (c *clusterOperator) removeService(cluster *crv1.MySQLCluster) error {
	serviceInterface := c.clientset.CoreV1().Services(cluster.Namespace)
	return serviceInterface.Delete(cluster.Spec.Name, new(metav1.DeleteOptions))
}

func (c *clusterOperator) removeStatefulSet(cluster *crv1.MySQLCluster) error {
	statefulSetInterface := c.clientset.AppsV1().StatefulSets(cluster.Namespace)
	return statefulSetInterface.Delete(cluster.Spec.Name, new(metav1.DeleteOptions))
}
