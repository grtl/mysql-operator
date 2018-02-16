package cluster

import (
	"testing"
	"time"

	"github.com/nauyey/factory"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	crv1 "github.com/grtl/mysql-operator/pkg/apis/cr/v1"
	testingFactory "github.com/grtl/mysql-operator/testing/factory"
)

type EventsHookTestSuite struct {
	suite.Suite
	eventsHook EventsHook
	cluster    *crv1.MySQLCluster
}

type eventTest func(ClusterEvent)

const TIMEOUT = time.Second * 1

func (suite *EventsHookTestSuite) testWithTimeout(test eventTest) {
	select {
	case event := <-suite.eventsHook.GetEventsChan():
		test(event)
	case <-time.After(TIMEOUT):
		suite.Fail("Timeout while waiting for event")
	}
}

func (suite *EventsHookTestSuite) SetupTest() {
	suite.eventsHook = NewEventsHook(16)

	suite.cluster = &crv1.MySQLCluster{}
	err := factory.Build(testingFactory.MySQLClusterFactory).To(suite.cluster)
	suite.Require().Nil(err)
}

func (suite *EventsHookTestSuite) TestEventsHook_OnAdd() {
	suite.eventsHook.OnAdd(suite.cluster)
	suite.testWithTimeout(func(event ClusterEvent) {
		suite.Assert().Equal(ClusterAdded, event.Type)
		suite.Assert().Equal(suite.cluster, event.Cluster)
	})
}

func (suite *EventsHookTestSuite) TestEventsHook_OnUpdate() {
	suite.eventsHook.OnUpdate(suite.cluster)
	suite.testWithTimeout(func(event ClusterEvent) {
		suite.Assert().Equal(ClusterUpdated, event.Type)
		suite.Assert().Equal(suite.cluster, event.Cluster)
	})
}

func (suite *EventsHookTestSuite) TestEventsHook_OnDelete() {
	suite.eventsHook.OnDelete(suite.cluster)
	suite.testWithTimeout(func(event ClusterEvent) {
		suite.Assert().Equal(ClusterDeleted, event.Type)
		suite.Assert().Equal(suite.cluster, event.Cluster)
	})
}

func (suite *EventsHookTestSuite) TestEventsHook_GetEventsChan() {
	eventsChan := suite.eventsHook.GetEventsChan()
	suite.Require().NotNil(eventsChan)

	hook, ok := suite.eventsHook.(*eventsHook)
	suite.Assert().True(ok)

	event := ClusterEvent{Type: ClusterAdded, Cluster: suite.cluster}
	hook.events <- event

	suite.testWithTimeout(func(clusterEvent ClusterEvent) {
		suite.Require().Equal(event, clusterEvent)
	})
}

func TestEventsHookTestSuite(t *testing.T) {
	suite.Run(t, new(EventsHookTestSuite))
}

func TestEventsHookRegisters(t *testing.T) {
	hook := NewEventsHook(16)
	controller := NewClusterController(nil, nil)
	err := controller.AddHook(hook)
	require.Nil(t, err)
}
