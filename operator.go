package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"

	"k8s.io/client-go/tools/clientcmd"

	"github.com/sirupsen/logrus"

	"github.com/grtl/mysql-operator/controller"
	"github.com/grtl/mysql-operator/logging"
	"github.com/grtl/mysql-operator/pkg/client/clientset/versioned"
	"k8s.io/client-go/kubernetes"
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

	clientset, err := versioned.NewForConfig(config)
	if err != nil {
		logrus.Panic(err)
	}

	k_clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()

	clusterController := controller.NewClusterController(clientset, k_clientset)
	go clusterController.Run(ctx)

	go logging.LogEvents(
		ctx,
		clusterController,
	)

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
