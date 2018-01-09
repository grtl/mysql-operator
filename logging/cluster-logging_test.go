package logging

import (
	"testing"

	"github.com/nauyey/factory"
	"github.com/stretchr/testify/suite"

	"github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/test"

	"github.com/grtl/mysql-operator/controller"
	crv1 "github.com/grtl/mysql-operator/pkg/apis/cr/v1"
	testfactory "github.com/grtl/mysql-operator/testing/factory"
)

type LogClusterEventTestSuite struct {
	suite.Suite
	hook    *test.Hook
	cluster *crv1.MySQLCluster
}

func (suite *LogClusterEventTestSuite) SetupTest() {
	suite.hook = test.NewGlobal()

	suite.cluster = &crv1.MySQLCluster{}
	err := factory.Build(testfactory.MySQLClusterFactory).To(suite.cluster)
	suite.Require().Nil(err)
}

func (suite *LogClusterEventTestSuite) TestLogClusterAdd() {
	logClusterEvent(controller.ClusterEvent{Type: controller.ADDED, Cluster: suite.cluster})
	suite.Assert().Equal(1, len(suite.hook.AllEntries()))
	suite.Assert().Equal(logrus.InfoLevel, suite.hook.LastEntry().Level)
	suite.Assert().Equal("Received cluster event", suite.hook.LastEntry().Message)
	suite.Assert().Equal(logrus.Fields{
		"cluster": suite.cluster.Name,
		"event":   controller.ADDED,
	}, suite.hook.LastEntry().Data)
}

func (suite *LogClusterEventTestSuite) TestLogClusterUpdate() {
	logClusterEvent(controller.ClusterEvent{Type: controller.UPDATED, Cluster: suite.cluster})
	suite.Assert().Equal(1, len(suite.hook.AllEntries()))
	suite.Assert().Equal(logrus.InfoLevel, suite.hook.LastEntry().Level)
	suite.Assert().Equal("Received cluster event", suite.hook.LastEntry().Message)
	suite.Assert().Equal(logrus.Fields{
		"cluster": suite.cluster.Name,
		"event":   controller.UPDATED,
	}, suite.hook.LastEntry().Data)
}

func (suite *LogClusterEventTestSuite) TestLogClusterDelete() {
	logClusterEvent(controller.ClusterEvent{Type: controller.DELETED, Cluster: suite.cluster})
	suite.Assert().Equal(1, len(suite.hook.AllEntries()))
	suite.Assert().Equal(logrus.InfoLevel, suite.hook.LastEntry().Level)
	suite.Assert().Equal("Received cluster event", suite.hook.LastEntry().Message)
	suite.Assert().Equal(logrus.Fields{
		"cluster": suite.cluster.Name,
		"event":   controller.DELETED,
	}, suite.hook.LastEntry().Data)
}

func TestLogClusterEventTestSuite(t *testing.T) {
	suite.Run(t, new(LogClusterEventTestSuite))
}
