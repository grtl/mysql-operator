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
	"github.com/grtl/mysql-operator/pkg/client/clientset/versioned"
	"github.com/grtl/mysql-operator/util"
)

const (
	serviceTemplate     = "artifacts/mysql-service.yaml"
	serviceReadTemplate = "artifacts/mysql-service-read.yaml"
	statefulSetTemplate = "artifacts/mysql-statefulset.yaml"
)

// Operator represents an object to manipulate MySQLCluster custom resources.
type Operator interface {
	// AddCluster creates the Kubernetes API objects necessary for a MySQL cluster.
	AddCluster(cluster *crv1.MySQLCluster) error
	UpdateCluster(newCluster *crv1.MySQLCluster) error
}

type clusterOperator struct {
	clientset     versioned.Interface
	kubeClientset kubernetes.Interface
}

// NewClusterOperator returns a new Operator.
func NewClusterOperator(clientset versioned.Interface, kubeClientset kubernetes.Interface) Operator {
	return &clusterOperator{
		clientset:     clientset,
		kubeClientset: kubeClientset,
	}
}

func (c *clusterOperator) AddCluster(cluster *crv1.MySQLCluster) error {
	cluster.WithDefaults()

	logging.LogCluster(cluster).Debug("Creating service.")
	err := c.createService(cluster, serviceTemplate)
	if err != nil {
		return err
	}

	logging.LogCluster(cluster).Debug("Creating read service.")
	err = c.createService(cluster, serviceReadTemplate)
	if err != nil {
		// Cleanup - remove already created service
		logging.LogCluster(cluster).WithField(
			"error", err).Warn("Reverting service creation.")
		removeErr := c.removeService(cluster)
		return errors.NewAggregate([]error{err, removeErr})
	}

	logging.LogCluster(cluster).Debug("Creating stateful set.")
	err = c.createStatefulSet(cluster)
	if err != nil {
		// Cleanup - remove already created services
		logging.LogCluster(cluster).WithField(
			"error", err).Warn("Reverting service creation.")
		removeErr := c.removeService(cluster)
		err = errors.NewAggregate([]error{err, removeErr})

		logging.LogCluster(cluster).WithField(
			"error", err).Warn("Reverting read service creation.")
		removeErr = c.removeReadService(cluster)
		return errors.NewAggregate([]error{err, removeErr})
	}

	return nil
}

func (c *clusterOperator) UpdateCluster(newCluster *crv1.MySQLCluster) error {
	newCluster.WithDefaults()

	logging.LogCluster(newCluster).Debug("Updating stateful set.")
	err := c.updateStatefulSet(newCluster)
	if err != nil {
		logging.LogCluster(newCluster).WithField(
			"error", err).Warn("Setting status")
		setStateErr := c.setClusterState(
			newCluster,
			"Failed update",
			"The provided patch resulted in a StatefulSet update failure",
		)
		return errors.NewAggregate([]error{err, setStateErr})
	}

	return c.setClusterState(newCluster, "Successful update", "")
}

func (c *clusterOperator) setClusterState(cluster *crv1.MySQLCluster, state, message string) error {
	cluster.Status.State = state
	cluster.Status.Message = message
	_, updateErr := c.clientset.CrV1().
		MySQLClusters(cluster.ObjectMeta.Namespace).Update(cluster)

	return updateErr
}

func (c *clusterOperator) createService(cluster *crv1.MySQLCluster, filename string) error {
	serviceInterface := c.kubeClientset.CoreV1().Services(cluster.Namespace)
	service, err := serviceForCluster(cluster, filename)
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
	statefulSetInterface := c.kubeClientset.AppsV1().StatefulSets(cluster.Namespace)
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

func (c *clusterOperator) updateStatefulSet(cluster *crv1.MySQLCluster) error {
	statefulSetInterface := c.kubeClientset.AppsV1().StatefulSets(cluster.Namespace)
	statefulSet, err := statefulSetForCluster(cluster)
	if err != nil {
		return err
	}

	_, err = statefulSetInterface.Update(statefulSet)

	return err
}

func serviceForCluster(cluster *crv1.MySQLCluster, filename string) (*corev1.Service, error) {
	service := new(corev1.Service)
	err := util.ObjectFromTemplate(cluster, service, filename, FuncMap)
	return service, err
}

func statefulSetForCluster(cluster *crv1.MySQLCluster) (*appsv1.StatefulSet, error) {
	statefulSet := new(appsv1.StatefulSet)
	err := util.ObjectFromTemplate(cluster, statefulSet, statefulSetTemplate, FuncMap)
	return statefulSet, err
}

func (c *clusterOperator) removeService(cluster *crv1.MySQLCluster) error {
	serviceInterface := c.kubeClientset.CoreV1().Services(cluster.Namespace)
	return serviceInterface.Delete(ServiceName(cluster.Name), new(metav1.DeleteOptions))
}

func (c *clusterOperator) removeReadService(cluster *crv1.MySQLCluster) error {
	serviceInterface := c.kubeClientset.CoreV1().Services(cluster.Namespace)
	return serviceInterface.Delete(ReadServiceName(cluster.Name), new(metav1.DeleteOptions))
}

func (c *clusterOperator) removeStatefulSet(cluster *crv1.MySQLCluster) error {
	statefulSetInterface := c.kubeClientset.AppsV1().StatefulSets(cluster.Namespace)
	return statefulSetInterface.Delete(StatefulSetName(cluster.Name), new(metav1.DeleteOptions))
}
