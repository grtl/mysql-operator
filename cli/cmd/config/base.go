package config

import (
	extclientset "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"

	"github.com/grtl/mysql-operator/pkg/client/clientset/versioned"
)

type baseConfig struct {
	clientset    *versioned.Clientset
	kubeClienset *kubernetes.Clientset
	extClientset *extclientset.Clientset
}

func (c *baseConfig) Clientset() versioned.Interface {
	return c.clientset
}

func (c *baseConfig) KubeClientset() kubernetes.Interface {
	return c.kubeClienset
}

func (c *baseConfig) ExtClientset() extclientset.Interface {
	return c.extClientset
}

// InitBaseConfig initializes global configuration objects with standard values.
func InitBaseConfig(kubeconfig string) error {
	config := new(baseConfig)

	restConfig, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return err
	}

	config.clientset, err = versioned.NewForConfig(restConfig)
	if err != nil {
		return err
	}

	config.kubeClienset, err = kubernetes.NewForConfig(restConfig)
	if err != nil {
		return err
	}

	config.extClientset, err = extclientset.NewForConfig(restConfig)
	if err != nil {
		return err
	}

	configInstance = config
	return nil
}
