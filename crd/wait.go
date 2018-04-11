package crd

import (
	"time"

	"github.com/sirupsen/logrus"

	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	apiextensions "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/wait"
)

// WaitForCRDEstablished stops the execution until Custom Resource Definition
// is registered or a timeout occurs.
func WaitForCRDEstablished(clientset apiextensions.Interface, crdName string) error {
	return wait.Poll(250*time.Millisecond, 30*time.Second, func() (bool, error) {
		crd, err := clientset.ApiextensionsV1beta1().CustomResourceDefinitions().Get(crdName, metav1.GetOptions{})
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
