package logging

import (
	"io/ioutil"
	"testing"

	"github.com/nauyey/factory"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/test"

	crv1 "github.com/grtl/mysql-operator/pkg/apis/cr/v1"
	testingFactory "github.com/grtl/mysql-operator/testing/factory"
)

func TestLogBackup(t *testing.T) {
	// Turn off logging output
	logrus.SetOutput(ioutil.Discard)

	// Initialize logging hook
	logrusHook := test.NewGlobal()
	logrus.SetLevel(logrus.DebugLevel)

	// Create fake backup
	backup := new(crv1.MySQLBackup)
	err := factory.Build(testingFactory.MySQLBackupFactory).To(backup)
	require.Nil(t, err)

	// Debug level
	LogBackup(backup).Debug("Debug")
	require.Equal(t, 1, len(logrusHook.AllEntries()))
	assert.Equal(t, logrus.DebugLevel, logrusHook.LastEntry().Level)
	assert.Equal(t, "Debug", logrusHook.LastEntry().Message)
	assert.Equal(t, logrus.Fields{
		"backup": backup.Name,
	}, logrusHook.LastEntry().Data)

	LogBackup(backup).Info("Info")
	require.Equal(t, 2, len(logrusHook.AllEntries()))
	assert.Equal(t, logrus.InfoLevel, logrusHook.LastEntry().Level)
	assert.Equal(t, "Info", logrusHook.LastEntry().Message)
	assert.Equal(t, logrus.Fields{
		"backup": backup.Name,
	}, logrusHook.LastEntry().Data)

	LogBackup(backup).Warn("Warning")
	require.Equal(t, 3, len(logrusHook.AllEntries()))
	assert.Equal(t, logrus.WarnLevel, logrusHook.LastEntry().Level)
	assert.Equal(t, "Warning", logrusHook.LastEntry().Message)
	assert.Equal(t, logrus.Fields{
		"backup": backup.Name,
	}, logrusHook.LastEntry().Data)

	LogBackup(backup).Error("Error")
	require.Equal(t, 4, len(logrusHook.AllEntries()))
	assert.Equal(t, logrus.ErrorLevel, logrusHook.LastEntry().Level)
	assert.Equal(t, "Error", logrusHook.LastEntry().Message)
	assert.Equal(t, logrus.Fields{
		"backup": backup.Name,
	}, logrusHook.LastEntry().Data)
}
