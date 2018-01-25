package util

import (
	"os"
	"time"

	"github.com/sirupsen/logrus"

	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	apiextensionsclient "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/apimachinery/pkg/util/yaml"
)

func YAMLToCRD(filepath string) (*apiextensions.CustomResourceDefinition, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}

	crd := &apiextensions.CustomResourceDefinition{}
	yaml.NewYAMLOrJSONDecoder(f, 32).Decode(crd)

	return crd, nil
}

func WaitForCRDEstablished(clientset *apiextensionsclient.Clientset, CRDName string) error {
	return wait.Poll(500*time.Millisecond, 60*time.Second, func() (bool, error) {
		crd, err := clientset.ApiextensionsV1beta1().CustomResourceDefinitions().Get(CRDName, v1.GetOptions{})
		if err != nil {
			return false, err
		}

		for _, cond := range crd.Status.Conditions {
			switch cond.Type {
			case apiextensions.Established:
				if cond.Status == apiextensions.ConditionTrue {
					return true, err
				}
			case apiextensions.NamesAccepted:
				if cond.Status == apiextensions.ConditionFalse {
					logrus.Infof("Name conflict: %v", cond.Reason)
				}
			}
		}
		return false, err
	})
}
