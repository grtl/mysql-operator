package controller

import (
	"context"
	"testing"
	"time"

	"github.com/nauyey/factory"
	"github.com/stretchr/testify/suite"

	"k8s.io/apimachinery/pkg/watch"
	testclient "k8s.io/client-go/testing"

	crv1 "github.com/grtl/mysql-operator/pkg/apis/cr/v1"
	"github.com/grtl/mysql-operator/pkg/client/clientset/versioned/fake"
	testfactory "github.com/grtl/mysql-operator/testing/factory"
	kFake "k8s.io/client-go/kubernetes/fake"
)

type ClusterControllerTestSuite struct {
	suite.Suite
	cluster    *crv1.MySQLCluster
	controller ClusterController
	watcher    *watch.FakeWatcher
	cancelFunc context.CancelFunc
}

type eventTest func(ClusterEvent)

const TIMEOUT = time.Second * 1

func (suite *ClusterControllerTestSuite) testWithTimeout(test eventTest) {
	select {
	case event := <-suite.controller.GetEventsChan():
		test(event)
	case <-time.After(TIMEOUT):
		suite.Fail("Timeout while waiting for event")
	}
}

func (suite *ClusterControllerTestSuite) SetupTest() {
	// Test Cluster
	suite.cluster = &crv1.MySQLCluster{}
	err := factory.Build(testfactory.MySQLClusterFactory).To(suite.cluster)
	suite.Require().Nil(err)

	kClientset := kFake.NewSimpleClientset()

	// Create the clientset with fake watcher
	clientset := fake.NewSimpleClientset(suite.cluster)
	suite.watcher = watch.NewFake()
	clientset.PrependWatchReactor("mysqlclusters", testclient.DefaultWatchReactor(suite.watcher, nil))

	// Create the controller
	suite.controller = NewClusterController(clientset, kClientset)

	ctx, cancelFunc := context.WithCancel(context.Background())
	suite.cancelFunc = cancelFunc

	go suite.controller.Run(ctx)
}

func (suite *ClusterControllerTestSuite) TearDownTest() {
	suite.cancelFunc()
}

// Test if onAdd function is being called.
func (suite *ClusterControllerTestSuite) TestClusterControllerAdd() {
	suite.testWithTimeout(func(clusterEvent ClusterEvent) {
		suite.Require().Equal(ADDED, clusterEvent.Type)
		suite.Equal(suite.cluster, clusterEvent.Cluster)
	})
}

// Test if onUpdate function is being called.
func (suite *ClusterControllerTestSuite) TestClusterControllerUpdate() {
	// Ignore ADDED event
	suite.testWithTimeout(func(clusterEvent ClusterEvent) {
		suite.Require().Equal(ADDED, clusterEvent.Type)
	})

	// Update cluster
	suite.cluster.Spec.Name += "-updated"
	suite.watcher.Modify(suite.cluster)

	suite.testWithTimeout(func(clusterEvent ClusterEvent) {
		suite.Require().Equal(UPDATED, clusterEvent.Type)
		suite.Equal(suite.cluster, clusterEvent.Cluster)
	})
}

// Test if onDelete function is being called.
func (suite *ClusterControllerTestSuite) TestClusterControllerDelete() {
	// Ignore ADDED event
	suite.testWithTimeout(func(clusterEvent ClusterEvent) {
		suite.Require().Equal(ADDED, clusterEvent.Type)
	})

	// Delete cluster
	suite.watcher.Delete(suite.cluster)

	suite.testWithTimeout(func(clusterEvent ClusterEvent) {
		suite.Require().Equal(DELETED, clusterEvent.Type)
		suite.Equal(suite.cluster, clusterEvent.Cluster)
	})
}

func TestClusterControllerTestSuite(t *testing.T) {
	suite.Run(t, new(ClusterControllerTestSuite))
}
