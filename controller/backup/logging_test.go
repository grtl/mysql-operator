package backup

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
	backup     *crv1.MySQLBackup
}

func (suite *LoggingTestSuite) SetupTest() {
	// Turn off logging output
	logrus.SetOutput(ioutil.Discard)

	// Initialize logging hook
	suite.logrusHook = test.NewGlobal()

	// Create fake backup
	suite.backup = &crv1.MySQLBackup{}
	err := factory.Build(testingFactory.MySQLBackupFactory).To(suite.backup)
	suite.Require().Nil(err)
}

func (suite *LoggingTestSuite) TestLogging_OnAdd() {
	logBackupEventBegin(suite.backup, backupAdded)
	suite.Require().Equal(1, len(suite.logrusHook.AllEntries()))
	suite.Assert().Equal(logrus.InfoLevel, suite.logrusHook.LastEntry().Level)
	suite.Assert().Equal("Received backup event", suite.logrusHook.LastEntry().Message)
	suite.Assert().Equal(logrus.Fields{
		"backup": suite.backup.Name,
		"event":  backupAdded,
	}, suite.logrusHook.LastEntry().Data)

	logBackupEventSuccess(suite.backup, backupAdded)
	suite.Require().Equal(2, len(suite.logrusHook.AllEntries()))
	suite.Assert().Equal(logrus.InfoLevel, suite.logrusHook.LastEntry().Level)
	suite.Assert().Equal("Successfully processed backup event", suite.logrusHook.LastEntry().Message)
	suite.Assert().Equal(logrus.Fields{
		"backup": suite.backup.Name,
		"event":  backupAdded,
	}, suite.logrusHook.LastEntry().Data)
}

func (suite *LoggingTestSuite) TestLogging_OnUpdate() {
	logBackupEventBegin(suite.backup, backupUpdated)
	suite.Require().Equal(1, len(suite.logrusHook.AllEntries()))
	suite.Assert().Equal(logrus.InfoLevel, suite.logrusHook.LastEntry().Level)
	suite.Assert().Equal("Received backup event", suite.logrusHook.LastEntry().Message)
	suite.Assert().Equal(logrus.Fields{
		"backup": suite.backup.Name,
		"event":  backupUpdated,
	}, suite.logrusHook.LastEntry().Data)

	logBackupEventSuccess(suite.backup, backupUpdated)
	suite.Require().Equal(2, len(suite.logrusHook.AllEntries()))
	suite.Assert().Equal(logrus.InfoLevel, suite.logrusHook.LastEntry().Level)
	suite.Assert().Equal("Successfully processed backup event", suite.logrusHook.LastEntry().Message)
	suite.Assert().Equal(logrus.Fields{
		"backup": suite.backup.Name,
		"event":  backupUpdated,
	}, suite.logrusHook.LastEntry().Data)
}

func (suite *LoggingTestSuite) TestLogging_OnDelete() {
	logBackupEventBegin(suite.backup, backupUpdated)
	suite.Require().Equal(1, len(suite.logrusHook.AllEntries()))
	suite.Assert().Equal(logrus.InfoLevel, suite.logrusHook.LastEntry().Level)
	suite.Assert().Equal("Received backup event", suite.logrusHook.LastEntry().Message)
	suite.Assert().Equal(logrus.Fields{
		"backup": suite.backup.Name,
		"event":  backupUpdated,
	}, suite.logrusHook.LastEntry().Data)

	logBackupEventSuccess(suite.backup, backupUpdated)
	suite.Require().Equal(2, len(suite.logrusHook.AllEntries()))
	suite.Assert().Equal(logrus.InfoLevel, suite.logrusHook.LastEntry().Level)
	suite.Assert().Equal("Successfully processed backup event", suite.logrusHook.LastEntry().Message)
	suite.Assert().Equal(logrus.Fields{
		"backup": suite.backup.Name,
		"event":  backupUpdated,
	}, suite.logrusHook.LastEntry().Data)
}

func TestLoggingTestSuite(t *testing.T) {
	suite.Run(t, new(LoggingTestSuite))
}
