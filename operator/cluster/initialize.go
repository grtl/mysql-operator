package cluster

import (
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/client-go/kubernetes"

	"github.com/grtl/mysql-operator/util"
)

// CreateConfigMap registers a "mysql" config map used by MySQL Clusters.
func CreateConfigMap(clientset kubernetes.Interface) error {
	configMap := new(corev1.ConfigMap)
	err := util.ObjectFromFile("artifacts/mysql-configmap.yaml", configMap)
	if err != nil {
		return err
	}

	_, err = clientset.CoreV1().ConfigMaps(corev1.NamespaceDefault).Create(configMap)
	if err != nil && apierrors.IsAlreadyExists(err) {
		return nil
	}
	return err
}
