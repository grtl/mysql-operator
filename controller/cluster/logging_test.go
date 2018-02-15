package cluster

import (
	"testing"

	"github.com/nauyey/factory"
	"github.com/stretchr/testify/suite"

	"github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/test"

	"github.com/grtl/mysql-operator/controller"
	crv1 "github.com/grtl/mysql-operator/pkg/apis/cr/v1"
	testingFactory "github.com/grtl/mysql-operator/testing/factory"
)

type ClusterLoggingHookTestSuite struct {
	suite.Suite
	loggingHook controller.ControllerHook
	logrusHook  *test.Hook
	cluster     *crv1.MySQLCluster
}

func (suite *ClusterLoggingHookTestSuite) SetupTest() {
	suite.loggingHook = NewClusterLoggingHook()
	suite.logrusHook = test.NewGlobal()

	suite.cluster = &crv1.MySQLCluster{}
	err := factory.Build(testingFactory.MySQLClusterFactory).To(suite.cluster)
	suite.Require().Nil(err)
}

func (suite *ClusterLoggingHookTestSuite) TestLogClusterAdd() {
	suite.loggingHook.OnAdd(suite.cluster)
	suite.Assert().Equal(1, len(suite.logrusHook.AllEntries()))
	suite.Assert().Equal(logrus.InfoLevel, suite.logrusHook.LastEntry().Level)
	suite.Assert().Equal("Received cluster event", suite.logrusHook.LastEntry().Message)
	suite.Assert().Equal(logrus.Fields{
		"cluster": suite.cluster.Name,
		"event":   clusterAdded,
	}, suite.logrusHook.LastEntry().Data)
}

func (suite *ClusterLoggingHookTestSuite) TestLogClusterUpdate() {
	suite.loggingHook.OnUpdate(suite.cluster)
	suite.Assert().Equal(1, len(suite.logrusHook.AllEntries()))
	suite.Assert().Equal(logrus.InfoLevel, suite.logrusHook.LastEntry().Level)
	suite.Assert().Equal("Received cluster event", suite.logrusHook.LastEntry().Message)
	suite.Assert().Equal(logrus.Fields{
		"cluster": suite.cluster.Name,
		"event":   clusterUpdated,
	}, suite.logrusHook.LastEntry().Data)
}

func (suite *ClusterLoggingHookTestSuite) TestLogClusterDelete() {
	suite.loggingHook.OnDelete(suite.cluster)
	suite.Assert().Equal(1, len(suite.logrusHook.AllEntries()))
	suite.Assert().Equal(logrus.InfoLevel, suite.logrusHook.LastEntry().Level)
	suite.Assert().Equal("Received cluster event", suite.logrusHook.LastEntry().Message)
	suite.Assert().Equal(logrus.Fields{
		"cluster": suite.cluster.Name,
		"event":   clusterDeleted,
	}, suite.logrusHook.LastEntry().Data)
}

func TestClusterLoggingHookTestSuite(t *testing.T) {
	suite.Run(t, new(ClusterLoggingHookTestSuite))
}
