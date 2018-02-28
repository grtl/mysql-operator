package crd

import (
	"time"

	"github.com/sirupsen/logrus"

	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	apiextensions "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/errors"
	"k8s.io/apimachinery/pkg/util/wait"

	"github.com/grtl/mysql-operator/util"
)

// waitForCRDEstablished stops the execution until Custom Resource Definition
// is registered or a timeout occurs.
func waitForCRDEstablished(clientset *apiextensions.Clientset, CRDName string) error {
	return wait.Poll(500*time.Millisecond, 60*time.Second, func() (bool, error) {
		crd, err := clientset.ApiextensionsV1beta1().CustomResourceDefinitions().Get(CRDName, metav1.GetOptions{})
		if err != nil {
			return false, err
		}

		for _, cond := range crd.Status.Conditions {
			switch cond.Type {
			case apiextensionsv1.Established:
				if cond.Status == apiextensionsv1.ConditionTrue {
					return true, err
				}
			case apiextensionsv1.NamesAccepted:
				if cond.Status == apiextensionsv1.ConditionFalse {
					logrus.WithField("reason", cond.Reason).Warn("Name conflict")
				}
			}
		}
		return false, err
	})
}

// createCRD registers given custom resource definition into the kubernetes api.
func createCRD(clientset *apiextensions.Clientset, filename string) error {
	crd := new(apiextensionsv1.CustomResourceDefinition)
	err := util.ObjectFromFile(filename, crd)
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

	err = waitForCRDEstablished(clientset, crd.ObjectMeta.Name)
	if err != nil {
		deleteErr := crdInterface.Delete(crd.ObjectMeta.Name, nil)
		return errors.NewAggregate([]error{err, deleteErr})
	}

	return nil
}
