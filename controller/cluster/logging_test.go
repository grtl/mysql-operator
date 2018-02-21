package cluster

import (
	"testing"

	"github.com/nauyey/factory"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/test"

	crv1 "github.com/grtl/mysql-operator/pkg/apis/cr/v1"
	testingFactory "github.com/grtl/mysql-operator/testing/factory"
)

type LoggingHookTestSuite struct {
	suite.Suite
	loggingHook LoggingHook
	logrusHook  *test.Hook
	cluster     *crv1.MySQLCluster
}

func (suite *LoggingHookTestSuite) SetupTest() {
	suite.loggingHook = NewLoggingHook()
	suite.logrusHook = test.NewGlobal()

	suite.cluster = &crv1.MySQLCluster{}
	err := factory.Build(testingFactory.MySQLClusterFactory).To(suite.cluster)
	suite.Require().Nil(err)
}

func (suite *LoggingHookTestSuite) TestLoggingHook_OnAdd() {
	suite.loggingHook.OnAdd(suite.cluster)
	suite.Assert().Equal(1, len(suite.logrusHook.AllEntries()))
	suite.Assert().Equal(logrus.InfoLevel, suite.logrusHook.LastEntry().Level)
	suite.Assert().Equal("Received cluster event", suite.logrusHook.LastEntry().Message)
	suite.Assert().Equal(logrus.Fields{
		"cluster": suite.cluster.Name,
		"event":   clusterAdded,
	}, suite.logrusHook.LastEntry().Data)
}

func (suite *LoggingHookTestSuite) TestLoggingHook_OnUpdate() {
	suite.loggingHook.OnUpdate(suite.cluster)
	suite.Assert().Equal(1, len(suite.logrusHook.AllEntries()))
	suite.Assert().Equal(logrus.InfoLevel, suite.logrusHook.LastEntry().Level)
	suite.Assert().Equal("Received cluster event", suite.logrusHook.LastEntry().Message)
	suite.Assert().Equal(logrus.Fields{
		"cluster": suite.cluster.Name,
		"event":   clusterUpdated,
	}, suite.logrusHook.LastEntry().Data)
}

func (suite *LoggingHookTestSuite) TestLoggingHook_OnDelete() {
	suite.loggingHook.OnDelete(suite.cluster)
	suite.Assert().Equal(1, len(suite.logrusHook.AllEntries()))
	suite.Assert().Equal(logrus.InfoLevel, suite.logrusHook.LastEntry().Level)
	suite.Assert().Equal("Received cluster event", suite.logrusHook.LastEntry().Message)
	suite.Assert().Equal(logrus.Fields{
		"cluster": suite.cluster.Name,
		"event":   clusterDeleted,
	}, suite.logrusHook.LastEntry().Data)
}

func TestLoggingHookTestSuite(t *testing.T) {
	suite.Run(t, new(LoggingHookTestSuite))
}

func TestLoggingHook_Registers(t *testing.T) {
	hook := NewEventsHook(16)
	clusterController := NewClusterController(nil, nil)
	err := clusterController.AddHook(hook)
	require.Nil(t, err)
}
