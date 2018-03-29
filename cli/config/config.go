package config

import (
	"github.com/grtl/mysql-operator/pkg/client/clientset/versioned"
	extclientset "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	"k8s.io/client-go/kubernetes"
)

// Config is used fgr storing global configuration objects.
type Config interface {
	Clientset() versioned.Interface
	KubeClientset() kubernetes.Interface
	ExtClientset() extclientset.Interface
}

var configInstance Config

// GetConfig returns global configuration.
func GetConfig() Config {
	return configInstance
}
