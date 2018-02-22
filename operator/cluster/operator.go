package cluster

import (
	"bytes"
	"text/template"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/client-go/kubernetes"

	"github.com/sirupsen/logrus"

	crv1 "github.com/grtl/mysql-operator/pkg/apis/cr/v1"
)

// AddCluster creates the Kubernetes API objects necessary for a MySQL cluster.
func AddCluster(cluster *crv1.MySQLCluster, kubeClientset kubernetes.Interface) error {
	err := createServiceForCluster(cluster, kubeClientset)
	if err != nil {
		return err
	}
	err = createStatefulSetForCluster(cluster, kubeClientset)
	return err
}

func createServiceForCluster(cluster *crv1.MySQLCluster, kubeClientset kubernetes.Interface) error {
	servicesInterface := kubeClientset.CoreV1().Services(cluster.ObjectMeta.Namespace)

	newService, err := serviceForCluster(cluster)
	if err != nil {
		return err
	}

	_, err = servicesInterface.Create(newService)
	if err != nil && !apierrors.IsAlreadyExists(err) {
		return err
	} else if apierrors.IsAlreadyExists(err) {
		logrus.WithFields(logrus.Fields{
			"cluster": cluster.Name,
		}).Info("Service for cluster already exists")
	}
	return nil
}

func createStatefulSetForCluster(cluster *crv1.MySQLCluster, kubeClientset kubernetes.Interface) error {
	statefulSetsInterface := kubeClientset.AppsV1().StatefulSets(cluster.ObjectMeta.Namespace)

	newStatefulSet, err := statefulSetForCluster(cluster)
	if err != nil {
		return err
	}

	_, err = statefulSetsInterface.Create(newStatefulSet)
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
	err := parseObject(cluster, service, "artifacts/mysql-service.yaml")
	return service, err
}

func statefulSetForCluster(cluster *crv1.MySQLCluster) (*appsv1.StatefulSet, error) {
	statefulSet := new(appsv1.StatefulSet)
	err := parseObject(cluster, statefulSet, "artifacts/mysql-statefulset.yaml")
	return statefulSet, err
}

func parseObject(cluster *crv1.MySQLCluster, object interface{}, file string) error {
	tmpl, err := template.ParseFiles(file)
	if err != nil {
		return err
	}

	var stringBuffer string
	buffer := bytes.NewBufferString(stringBuffer)
	err = tmpl.Execute(buffer, cluster)
	if err != nil {
		return err
	}

	return yaml.NewYAMLOrJSONDecoder(buffer, 64).Decode(object)
}
