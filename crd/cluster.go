package crd

import (
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	apiextensions "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/util/errors"

	"github.com/grtl/mysql-operator/util"
)

// CreateClusterCRD registers a MySQLCluster custom resource.
func CreateClusterCRD(clientset *apiextensions.Clientset) error {
	crd := new(apiextensionsv1.CustomResourceDefinition)
	err := util.ObjectFromFile("artifacts/mysql-crd.yaml", crd)
	if err != nil {
		return err
	}

	crdInterface := clientset.ApiextensionsV1beta1().CustomResourceDefinitions()

	_, err = crdInterface.Create(crd)
	if err != nil {
		if apierrors.IsAlreadyExists(err) {
			return nil
		}
		return err
	}

	err = WaitForCRDEstablished(clientset, crd.ObjectMeta.Name)
	if err != nil {
		deleteErr := crdInterface.Delete(crd.ObjectMeta.Name, nil)
		return errors.NewAggregate([]error{err, deleteErr})
	}

	return nil
}
