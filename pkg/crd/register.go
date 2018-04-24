package crd

import (
	"github.com/grtl/mysql-operator/pkg/util"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	apiextensions "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
)

// RegisterCRD registers given custom resource definition into the kubernetes api.
func RegisterCRD(clientset apiextensions.Interface, filename string) error {
	crd := new(apiextensionsv1.CustomResourceDefinition)
	err := util.ObjectFromFile(filename, crd)
	if err != nil {
		return err
	}

	crdInterface := clientset.ApiextensionsV1beta1().CustomResourceDefinitions()

	_, err = crdInterface.Create(crd)
	if err != nil && apierrors.IsAlreadyExists(err) {
		return nil
	}

	return err
}

// UnregisterCRD removes custom resource definition from kubernetes api.
func UnregisterCRD(clientset apiextensions.Interface, crdName string) error {
	crdInterface := clientset.ApiextensionsV1beta1().CustomResourceDefinitions()
	return crdInterface.Delete(crdName, nil)
}
