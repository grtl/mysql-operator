package controller

import (
	"context"
	"testing"
	"time"

	"github.com/nauyey/factory"
	"github.com/stretchr/testify/suite"

	"k8s.io/apimachinery/pkg/watch"

	"github.com/grtl/mysql-operator/controller"
	"github.com/grtl/mysql-operator/controller/cluster"
	"github.com/grtl/mysql-operator/controller/cluster/fake"
	crv1 "github.com/grtl/mysql-operator/pkg/apis/cr/v1"
	testFactory "github.com/grtl/mysql-operator/testing/factory"
)

type ClusterControllerTestSuite struct {
	suite.Suite

	cluster *crv1.MySQLCluster

	controller controller.Controller
	watcher    *watch.FakeWatcher
	eventsHook cluster.EventsHook

	cancelFunc context.CancelFunc
}

type eventTest func(cluster.ClusterEvent)

const TIMEOUT = time.Second * 1

func (suite *ClusterControllerTestSuite) testWithTimeout(test eventTest) {
	select {
	case event := <-suite.eventsHook.GetEventsChan():
		test(event)
	case <-time.After(TIMEOUT):
		suite.Fail("Timeout while waiting for event")
	}
}

func (suite *ClusterControllerTestSuite) SetupTest() {
	// Initialize the controller
	suite.watcher, suite.controller = fake.NewFakeClusterController(16)
	suite.eventsHook = cluster.NewEventsHook(16)
	err := suite.controller.AddHook(suite.eventsHook)
	suite.Require().Nil(err)

	// Test Cluster
	suite.cluster = &crv1.MySQLCluster{}
	err = factory.Build(testFactory.MySQLClusterFactory).To(suite.cluster)
	suite.Require().Nil(err)
	suite.watcher.Add(suite.cluster)

	// Start the controller
	ctx, cancelFunc := context.WithCancel(context.Background())
	suite.cancelFunc = cancelFunc

	go suite.controller.Run(ctx)
}

func (suite *ClusterControllerTestSuite) TearDownTest() {
	suite.cancelFunc()
}

// Test if onAdd function is being called.
func (suite *ClusterControllerTestSuite) TestClusterControllerAdd() {
	suite.testWithTimeout(func(clusterEvent cluster.ClusterEvent) {
		suite.Require().Equal(cluster.ClusterAdded, clusterEvent.Type)
		suite.Equal(suite.cluster, clusterEvent.Cluster)
	})
}

// Test if onUpdate function is being called.
func (suite *ClusterControllerTestSuite) TestClusterControllerUpdate() {
	// Ignore clusterAdded event
	suite.testWithTimeout(func(clusterEvent cluster.ClusterEvent) {
		suite.Require().Equal(cluster.ClusterAdded, clusterEvent.Type)
	})

	// Update cluster
	suite.cluster.Spec.Name += "-updated"
	suite.watcher.Modify(suite.cluster)

	suite.testWithTimeout(func(clusterEvent cluster.ClusterEvent) {
		suite.Require().Equal(cluster.ClusterUpdated, clusterEvent.Type)
		suite.Equal(suite.cluster, clusterEvent.Cluster)
	})
}

// Test if onDelete function is being called.
func (suite *ClusterControllerTestSuite) TestClusterControllerDelete() {
	// Ignore clusterAdded event
	suite.testWithTimeout(func(clusterEvent cluster.ClusterEvent) {
		suite.Require().Equal(cluster.ClusterAdded, clusterEvent.Type)
	})

	// Delete cluster
	suite.watcher.Delete(suite.cluster)

	suite.testWithTimeout(func(clusterEvent cluster.ClusterEvent) {
		suite.Require().Equal(cluster.ClusterDeleted, clusterEvent.Type)
		suite.Equal(suite.cluster, clusterEvent.Cluster)
	})
}

func TestClusterControllerTestSuite(t *testing.T) {
	suite.Run(t, new(ClusterControllerTestSuite))
}
