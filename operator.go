package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"k8s.io/client-go/tools/clientcmd"

	"github.com/grtl/mysql-operator/controller"
	"github.com/grtl/mysql-operator/pkg/client/clientset/versioned"
)

var (
	kubeconfig = flag.String("kubeconfig", "", "Path to kubeconfig. Only required if out-of-cluster.")
	master     = flag.String("master", "", "The address of the Kubernetes API server. Overrides any value in kubeconfig. Only required if out-of-cluster.")
)

var Clientset *versioned.Clientset

func main() {
	flag.Parse()
	config, err := clientcmd.BuildConfigFromFlags(*master, *kubeconfig)
	if err != nil {
		panic(err)
	}

	clientset, err := versioned.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()

	clusterController := controller.ClusterController{Clientset: clientset}
	go clusterController.Run(ctx)

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	for {
		select {
		case s := <-signals:
			fmt.Printf("Received signal %#v\n", s)
			os.Exit(0)
		}
	}
}
