package cluster

import (
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/client-go/kubernetes"

	"github.com/sirupsen/logrus"

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
	err := c.createServiceForCluster(cluster)
	if err != nil {
		// TODO: revert service creation
		return err
	}

	err = c.createStatefulSetForCluster(cluster)
	if err != nil {
		// TODO: revert service and stateful set creation
		return err
	}

	return nil
}

func (c *clusterOperator) createServiceForCluster(cluster *crv1.MySQLCluster) error {
	servicesInterface := c.clientset.CoreV1().Services(cluster.ObjectMeta.Namespace)

	service, err := serviceForCluster(cluster)
	if err != nil {
		return err
	}

	_, err = servicesInterface.Create(service)
	if err != nil && !apierrors.IsAlreadyExists(err) {
		return err
	} else if apierrors.IsAlreadyExists(err) {
		logrus.WithFields(logrus.Fields{
			"cluster": cluster.Name,
		}).Info("Service for cluster already exists")
	}

	return nil
}

func (c *clusterOperator) createStatefulSetForCluster(cluster *crv1.MySQLCluster) error {
	statefulSetsInterface := c.clientset.AppsV1().StatefulSets(cluster.ObjectMeta.Namespace)

	statefulSet, err := statefulSetForCluster(cluster)
	if err != nil {
		return err
	}

	_, err = statefulSetsInterface.Create(statefulSet)
	if err != nil && !apierrors.IsAlreadyExists(err) {
		return err
	} else if apierrors.IsAlreadyExists(err) {
		logrus.WithFields(logrus.Fields{
			"cluster": cluster.Name,
		}).Info("StatefulSet for cluster already exists")
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
