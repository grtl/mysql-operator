package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"

	apiextensions "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"

	"github.com/sirupsen/logrus"

	"github.com/grtl/mysql-operator/controller/cluster"
	"github.com/grtl/mysql-operator/crd"
	operator "github.com/grtl/mysql-operator/operator/cluster"
	"github.com/grtl/mysql-operator/pkg/client/clientset/versioned"
)

var (
	kubeconfig = flag.String("kubeconfig", "", "Path to kubeconfig. Only required if out-of-cluster.")
	master     = flag.String("master", "", "The address of the Kubernetes API server. Overrides any value in kubeconfig. Only required if out-of-cluster.")
	debug      = flag.Bool("debug", false, "Show debug logs")
)

var (
	clientset     *versioned.Clientset
	kubeClientset *kubernetes.Clientset
	extClientset  *apiextensions.Clientset
)

func main() {
	flag.Parse()

	if *debug {
		logrus.SetLevel(logrus.DebugLevel)
	}

	config, err := clientcmd.BuildConfigFromFlags(*master, *kubeconfig)
	if err != nil {
		logrus.Panic(err)
	}

	logrus.Debug("Initializing clientsets")
	err = initializeClientSets(config)
	if err != nil {
		logrus.Panic(err)
	}

	logrus.Debug("Initializing objects")
	err = initializeObjects()
	if err != nil {
		logrus.Panic(err)
	}

	logrus.Debug("Starting the controller")

	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()

	clusterController := cluster.NewClusterController(clientset, kubeClientset)
	go clusterController.Run(ctx)

	logrus.Info("Listening for events")

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	for {
		select {
		case s := <-signals:
			logrus.WithField("signal", s).Info("Received signal")
			os.Exit(0)
		}
	}
}

func initializeClientSets(config *rest.Config) error {
	var err error

	extClientset, err = apiextensions.NewForConfig(config)
	if err != nil {
		return err
	}

	clientset, err = versioned.NewForConfig(config)
	if err != nil {
		return err
	}

	kubeClientset, err = kubernetes.NewForConfig(config)
	return err
}

func initializeObjects() error {
	err := crd.CreateClusterCRD(extClientset)
	if err != nil {
		return err
	}

	err = crd.CreateBackupCRD(extClientset)
	if err != nil {
		return err
	}

	return operator.CreateConfigMap(kubeClientset)
}
