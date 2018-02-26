package crd

import (
	apiextensionsclient "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/util/errors"

	"github.com/grtl/mysql-operator/util"
)

func CreateMySQLClusterCRD(clientset *apiextensionsclient.Clientset) error {
	crd, err := util.YAMLToCRD("artifacts/mysqlcrd.yaml")
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

	err = util.WaitForCRDEstablished(clientset, crd.ObjectMeta.Name)
	if err != nil {
		deleteErr := crdInterface.Delete(crd.ObjectMeta.Name, nil)
		if deleteErr != nil {
			return errors.NewAggregate([]error{err, deleteErr})
		}
		return err
	}

	return nil
}
