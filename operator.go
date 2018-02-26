package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"

	apiextensionsclient "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"

	"github.com/sirupsen/logrus"

	"github.com/grtl/mysql-operator/controller/cluster"
	"github.com/grtl/mysql-operator/crd"
	"github.com/grtl/mysql-operator/pkg/client/clientset/versioned"
)

var (
	kubeconfig = flag.String("kubeconfig", "", "Path to kubeconfig. Only required if out-of-cluster.")
	master     = flag.String("master", "", "The address of the Kubernetes API server. Overrides any value in kubeconfig. Only required if out-of-cluster.")
)

func main() {
	logrus.Info("Starting operator")

	flag.Parse()
	config, err := clientcmd.BuildConfigFromFlags(*master, *kubeconfig)
	if err != nil {
		logrus.Panic(err)
	}

	extClientset, err := apiextensionsclient.NewForConfig(config)
	if err != nil {
		logrus.Panic(err)
	}

	err = crd.CreateMySQLClusterCRD(extClientset)
	if err != nil {
		logrus.Panic(err)
	}

	clientset, err := versioned.NewForConfig(config)
	if err != nil {
		logrus.Panic(err)
	}

	kubeClientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		logrus.Panic(err)
	}

	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()

	clusterController := cluster.NewClusterController(clientset, kubeClientset)
	err = clusterController.AddHook(cluster.NewLoggingHook())
	if err != nil {
		logrus.Warn("Could not initialize logging for cluster controller")
	}
	go clusterController.Run(ctx)

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	for {
		select {
		case s := <-signals:
			logrus.Infof("Received signal %#v", s)
			os.Exit(0)
		}
	}
}
