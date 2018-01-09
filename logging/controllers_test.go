package logging

import (
	"context"
	"testing"
	"time"

	"github.com/nauyey/factory"
	"github.com/stretchr/testify/suite"

	"github.com/sirupsen/logrus"
	"k8s.io/apimachinery/pkg/watch"

	"github.com/grtl/mysql-operator/controller"
	crv1 "github.com/grtl/mysql-operator/pkg/apis/cr/v1"
	testcontroller "github.com/grtl/mysql-operator/testing/controller"
	testfactory "github.com/grtl/mysql-operator/testing/factory"
	testlogging "github.com/grtl/mysql-operator/testing/logging"
)

type LogClusterControllerTestSuite struct {
	suite.Suite
	hook    *testlogging.Hook
	cluster *crv1.MySQLCluster

	controller controller.ClusterController
	watcher    *watch.FakeWatcher
	cancelFunc context.CancelFunc
}

const TIMEOUT = time.Second * 1

func (suite *LogClusterControllerTestSuite) testWithTimeout(event controller.ClusterEventType) {
	select {
	case entry := <-suite.hook.GetEntriesChan():
		suite.Assert().Equal(logrus.InfoLevel, entry.Level)
		suite.Assert().Equal("Received cluster event", entry.Message)
		suite.Assert().Equal(logrus.Fields{
			"cluster": suite.cluster.Name,
			"event":   event,
		}, entry.Data)
	case <-time.After(TIMEOUT):
		suite.Fail("Timeout while waiting for event")
	}
}

func (suite *LogClusterControllerTestSuite) SetupTest() {
	suite.hook = testlogging.NewGlobal()

	suite.cluster = &crv1.MySQLCluster{}
	err := factory.Build(testfactory.MySQLClusterFactory).To(suite.cluster)
	suite.Require().Nil(err)

	suite.watcher, suite.controller = testcontroller.NewFakeClusterController()

	ctx, cancelFunc := context.WithCancel(context.Background())
	suite.cancelFunc = cancelFunc

	go suite.controller.Run(ctx)
	go LogEvents(ctx, suite.controller)
}

func (suite *LogClusterControllerTestSuite) TearDownTest() {
	suite.cancelFunc()
}

func (suite *LogClusterControllerTestSuite) TestClusterControllerAdd() {
	suite.watcher.Add(suite.cluster)
	suite.testWithTimeout(controller.ADDED)
}

func (suite *LogClusterControllerTestSuite) TestClusterControllerUpdate() {
	suite.watcher.Add(suite.cluster)
	suite.testWithTimeout(controller.ADDED)
	suite.cluster.Spec.Name += "-updated"
	suite.watcher.Modify(suite.cluster)
	suite.testWithTimeout(controller.UPDATED)
}

func (suite *LogClusterControllerTestSuite) TestLogClusterControllerDelete() {
	suite.watcher.Add(suite.cluster)
	suite.testWithTimeout(controller.ADDED)
	suite.watcher.Delete(suite.cluster)
	suite.testWithTimeout(controller.DELETED)
}

func TestLogClusterControllerTestSuite(t *testing.T) {
	suite.Run(t, new(LogClusterControllerTestSuite))
}
