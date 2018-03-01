package cluster

import (
	"io/ioutil"
	"testing"

	"github.com/nauyey/factory"
	"github.com/stretchr/testify/suite"

	"github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/test"

	crv1 "github.com/grtl/mysql-operator/pkg/apis/cr/v1"
	testingFactory "github.com/grtl/mysql-operator/testing/factory"
)

type LoggingTestSuite struct {
	suite.Suite
	logrusHook *test.Hook
	cluster    *crv1.MySQLCluster
}

func (suite *LoggingTestSuite) SetupTest() {
	// Turn off logging output
	logrus.SetOutput(ioutil.Discard)

	// Initialize logging hook
	suite.logrusHook = test.NewGlobal()

	// Create fake cluster
	suite.cluster = &crv1.MySQLCluster{}
	err := factory.Build(testingFactory.MySQLClusterFactory).To(suite.cluster)
	suite.Require().Nil(err)
}

func (suite *LoggingTestSuite) TestLogging_OnAdd() {
	logClusterEventBegin(suite.cluster, clusterAdded)
	suite.Require().Equal(1, len(suite.logrusHook.AllEntries()))
	suite.Assert().Equal(logrus.InfoLevel, suite.logrusHook.LastEntry().Level)
	suite.Assert().Equal("Received cluster event", suite.logrusHook.LastEntry().Message)
	suite.Assert().Equal(logrus.Fields{
		"cluster": suite.cluster.Name,
		"event":   clusterAdded,
	}, suite.logrusHook.LastEntry().Data)

	logClusterEventSuccess(suite.cluster, clusterAdded)
	suite.Require().Equal(2, len(suite.logrusHook.AllEntries()))
	suite.Assert().Equal(logrus.InfoLevel, suite.logrusHook.LastEntry().Level)
	suite.Assert().Equal("Successfully processed cluster event", suite.logrusHook.LastEntry().Message)
	suite.Assert().Equal(logrus.Fields{
		"cluster": suite.cluster.Name,
		"event":   clusterAdded,
	}, suite.logrusHook.LastEntry().Data)
}

func (suite *LoggingTestSuite) TestLogging_OnUpdate() {
	logClusterEventBegin(suite.cluster, clusterUpdated)
	suite.Require().Equal(1, len(suite.logrusHook.AllEntries()))
	suite.Assert().Equal(logrus.InfoLevel, suite.logrusHook.LastEntry().Level)
	suite.Assert().Equal("Received cluster event", suite.logrusHook.LastEntry().Message)
	suite.Assert().Equal(logrus.Fields{
		"cluster": suite.cluster.Name,
		"event":   clusterUpdated,
	}, suite.logrusHook.LastEntry().Data)

	logClusterEventSuccess(suite.cluster, clusterUpdated)
	suite.Require().Equal(2, len(suite.logrusHook.AllEntries()))
	suite.Assert().Equal(logrus.InfoLevel, suite.logrusHook.LastEntry().Level)
	suite.Assert().Equal("Successfully processed cluster event", suite.logrusHook.LastEntry().Message)
	suite.Assert().Equal(logrus.Fields{
		"cluster": suite.cluster.Name,
		"event":   clusterUpdated,
	}, suite.logrusHook.LastEntry().Data)
}

func (suite *LoggingTestSuite) TestLogging_OnDelete() {
	logClusterEventBegin(suite.cluster, clusterUpdated)
	suite.Require().Equal(1, len(suite.logrusHook.AllEntries()))
	suite.Assert().Equal(logrus.InfoLevel, suite.logrusHook.LastEntry().Level)
	suite.Assert().Equal("Received cluster event", suite.logrusHook.LastEntry().Message)
	suite.Assert().Equal(logrus.Fields{
		"cluster": suite.cluster.Name,
		"event":   clusterUpdated,
	}, suite.logrusHook.LastEntry().Data)

	logClusterEventSuccess(suite.cluster, clusterUpdated)
	suite.Require().Equal(2, len(suite.logrusHook.AllEntries()))
	suite.Assert().Equal(logrus.InfoLevel, suite.logrusHook.LastEntry().Level)
	suite.Assert().Equal("Successfully processed cluster event", suite.logrusHook.LastEntry().Message)
	suite.Assert().Equal(logrus.Fields{
		"cluster": suite.cluster.Name,
		"event":   clusterUpdated,
	}, suite.logrusHook.LastEntry().Data)
}

func TestLoggingTestSuite(t *testing.T) {
	suite.Run(t, new(LoggingTestSuite))
}
