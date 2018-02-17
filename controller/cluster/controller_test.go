package cluster

import (
	"context"
	"testing"
	"time"

	"github.com/nauyey/factory"
	"github.com/stretchr/testify/suite"

	"k8s.io/apimachinery/pkg/watch"
	kubeFake "k8s.io/client-go/kubernetes/fake"
	kubeTesting "k8s.io/client-go/testing"

	"github.com/grtl/mysql-operator/controller"
	crv1 "github.com/grtl/mysql-operator/pkg/apis/cr/v1"
	"github.com/grtl/mysql-operator/pkg/client/clientset/versioned/fake"
	testFactory "github.com/grtl/mysql-operator/testing/factory"
)

type ClusterControllerTestSuite struct {
	suite.Suite

	cluster *crv1.MySQLCluster

	controller controller.Controller
	watcher    *watch.FakeWatcher
	eventsHook EventsHook

	cancelFunc context.CancelFunc
}

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
	kubeClientset := kubeFake.NewSimpleClientset()
	clientset := fake.NewSimpleClientset()

	suite.watcher = watch.NewFakeWithChanSize(16, false)
	clientset.PrependWatchReactor("mysqlclusters", kubeTesting.DefaultWatchReactor(suite.watcher, nil))
	suite.controller = NewClusterController(clientset, kubeClientset)
	suite.eventsHook = NewEventsHook(16)
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
func (suite *ClusterControllerTestSuite) TestClusterController_OnAdd() {
	suite.testWithTimeout(func(clusterEvent Event) {
		suite.Require().Equal(ClusterAdded, clusterEvent.Type)
		suite.Equal(suite.cluster, clusterEvent.Cluster)
	})
}

// Test if onUpdate function is being called.
func (suite *ClusterControllerTestSuite) TestClusterController_OnUpdate() {
	// Ignore clusterAdded event
	suite.testWithTimeout(func(clusterEvent Event) {
		suite.Require().Equal(ClusterAdded, clusterEvent.Type)
	})

	// Update cluster
	suite.cluster.Spec.Name += "-updated"
	suite.watcher.Modify(suite.cluster)

	suite.testWithTimeout(func(clusterEvent Event) {
		suite.Require().Equal(ClusterUpdated, clusterEvent.Type)
		suite.Equal(suite.cluster, clusterEvent.Cluster)
	})
}

// Test if onDelete function is being called.
func (suite *ClusterControllerTestSuite) TestClusterController_OnDelete() {
	// Ignore clusterAdded event
	suite.testWithTimeout(func(clusterEvent Event) {
		suite.Require().Equal(ClusterAdded, clusterEvent.Type)
	})

	// Delete cluster
	suite.watcher.Delete(suite.cluster)

	suite.testWithTimeout(func(clusterEvent Event) {
		suite.Require().Equal(ClusterDeleted, clusterEvent.Type)
		suite.Equal(suite.cluster, clusterEvent.Cluster)
	})
}

func TestClusterControllerTestSuite(t *testing.T) {
	suite.Run(t, new(ClusterControllerTestSuite))
}

type ClusterControllerHooksTestSuite struct {
	suite.Suite
	controller *clusterController
}

func (suite *ClusterControllerHooksTestSuite) SetupTest() {
	var ok bool
	suite.controller, ok = NewClusterController(nil, nil).(*clusterController)
	suite.Require().True(ok)
}

func (suite *ClusterControllerHooksTestSuite) TestClusterController_AddHook() {
	suite.Require().Equal(0, len(suite.controller.hooks))

	// Add hook
	hook := NewEventsHook(1) // Any hook will do
	err := suite.controller.AddHook(hook)
	suite.Assert().Nil(err)
	suite.Require().Equal(1, len(suite.controller.hooks))

	// Hook already exists
	err = suite.controller.AddHook(hook)
	suite.Assert().NotNil(err)
	suite.Require().Equal(1, len(suite.controller.hooks))

	// Add another hook
	anotherHook := NewEventsHook(1) // Any hook will do
	err = suite.controller.AddHook(anotherHook)
	suite.Assert().Nil(err)
	suite.Require().Equal(2, len(suite.controller.hooks))
}

func (suite *ClusterControllerHooksTestSuite) TestClusterController_RemoveHook() {
	hook := NewEventsHook(1)
	anotherHook := NewEventsHook(1)
	err := suite.controller.AddHook(hook)
	suite.Assert().Nil(err)
	err = suite.controller.AddHook(anotherHook)
	suite.Assert().Nil(err)
	suite.Require().Equal(2, len(suite.controller.hooks))

	// Remove hook
	err = suite.controller.RemoveHook(hook)
	suite.Assert().Nil(err)
	suite.Require().Equal(1, len(suite.controller.hooks))

	// Try to remove the same hook again
	err = suite.controller.RemoveHook(hook)
	suite.Assert().NotNil(err)
	suite.Require().Equal(1, len(suite.controller.hooks))

	// Remove another hook
	err = suite.controller.RemoveHook(anotherHook)
	suite.Assert().Nil(err)
	suite.Require().Equal(0, len(suite.controller.hooks))
}

func TestClusterControllerHooksTestSuite(t *testing.T) {
	suite.Run(t, new(ClusterControllerHooksTestSuite))
}
