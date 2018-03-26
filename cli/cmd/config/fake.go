package config

import (
	extclientset "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	extFake "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset/fake"
	"k8s.io/client-go/kubernetes"
	kubeFake "k8s.io/client-go/kubernetes/fake"

	"github.com/grtl/mysql-operator/pkg/client/clientset/versioned"
	"github.com/grtl/mysql-operator/pkg/client/clientset/versioned/fake"
)

type fakeConfig struct {
	clientset    *fake.Clientset
	kubeClienset *kubeFake.Clientset
	extClientset *extFake.Clientset
}

func (c *fakeConfig) Clientset() versioned.Interface {
	return c.clientset
}

func (c *fakeConfig) KubeClientset() kubernetes.Interface {
	return c.kubeClienset
}

func (c *fakeConfig) ExtClientset() extclientset.Interface {
	return c.extClientset
}

// InitFakeConfig initializes global configuration objects with fake clientsets.
func InitFakeConfig() {
	config := new(fakeConfig)
	config.clientset = fake.NewSimpleClientset()
	config.kubeClienset = kubeFake.NewSimpleClientset()
	config.extClientset = extFake.NewSimpleClientset()
	configInstance = config
}
