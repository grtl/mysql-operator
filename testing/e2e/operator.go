package e2e

import (
	"fmt"
	"os"
	"time"

	corev1 "k8s.io/api/core/v1"
	apiextensions "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"

	"github.com/grtl/mysql-operator/pkg/client/clientset/versioned"
	"github.com/grtl/mysql-operator/testing/kubectl"
)

const (
	operatorPod   = "mysql-operator"
	operatorImage = "--image=mysql-operator:testing"
)

// Operator represents MySQL operator pod deployed in the cluster
type Operator interface {
	Start() error
	Stop() error
	KubeClientset() kubernetes.Interface
	ExtClientset() apiextensions.Interface
	Clientset() versioned.Interface
}

type operator struct {
	kubeClientset *kubernetes.Clientset
	extClientset  *apiextensions.Clientset
	clientset     *versioned.Clientset
}

// NewOperator creates new Operator based on ${KUBECONFIG} env variable.
func NewOperator() (Operator, error) {
	kubeconfig, ok := os.LookupEnv("KUBECONFIG")
	if !ok {
		return nil, fmt.Errorf("environment variable ${KUBECONFIG} not defined")
	}

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return nil, err
	}

	newOperator := new(operator)
	err = newOperator.initializeClientsets(config)
	if err != nil {
		return nil, err
	}
	return newOperator, nil
}

func (o *operator) KubeClientset() kubernetes.Interface {
	return o.kubeClientset
}

func (o *operator) ExtClientset() apiextensions.Interface {
	return o.extClientset
}

func (o *operator) Clientset() versioned.Interface {
	return o.clientset
}

func (o *operator) Start() error {
	err := kubectl.Run(operatorPod, operatorImage, "--image-pull-policy=Never", "--restart=Never")
	if err != nil {
		return err
	}
	return o.waitForPodStarted()
}

func (o *operator) Stop() error {
	err := kubectl.Delete("pod", operatorPod)
	if err != nil {
		return err
	}
	return o.waitForPodDeleted()
}

func (o *operator) initializeClientsets(config *rest.Config) error {
	var err error

	o.extClientset, err = apiextensions.NewForConfig(config)
	if err != nil {
		return err
	}

	o.clientset, err = versioned.NewForConfig(config)
	if err != nil {
		return err
	}

	o.kubeClientset, err = kubernetes.NewForConfig(config)
	return err
}

func (o *operator) waitForPodStarted() error {
	return wait.Poll(1*time.Second, 30*time.Second, func() (bool, error) {
		pod, err := o.kubeClientset.CoreV1().Pods("default").Get(operatorPod, metav1.GetOptions{})
		if pod.Status.Phase == corev1.PodRunning {
			return true, nil
		} else if pod.Status.Phase == corev1.PodFailed {
			return false, fmt.Errorf("pod failed -> %v", pod.Status.String())
		}
		return false, err
	})
}

func (o *operator) waitForPodDeleted() error {
	return wait.Poll(1*time.Second, 30*time.Second, func() (bool, error) {
		_, err := o.kubeClientset.CoreV1().Pods("default").Get(operatorPod, metav1.GetOptions{})
		if apierrors.IsNotFound(err) {
			return true, nil
		}
		return false, err
	})
}
