package backup

import (
	"context"
	"io/ioutil"
	"testing"
	"time"

	"github.com/nauyey/factory"
	"github.com/stretchr/testify/suite"

	"k8s.io/apimachinery/pkg/watch"

	"github.com/sirupsen/logrus"

	"github.com/grtl/mysql-operator/controller"
	crv1 "github.com/grtl/mysql-operator/pkg/apis/cr/v1"
	testFactory "github.com/grtl/mysql-operator/testing/factory"
)

type BackupControllerTestSuite struct {
	suite.Suite

	backup *crv1.MySQLBackup

	controller controller.Controller
	watcher    *watch.FakeWatcher
	eventsHook controller.EventsHook

	cancelFunc context.CancelFunc
}

type eventTest func(controller.Event)

const TIMEOUT = time.Second * 1

func (suite *BackupControllerTestSuite) testWithTimeout(test eventTest) {
	select {
	case event := <-suite.eventsHook.GetEventsChan():
		test(event)
	case <-time.After(TIMEOUT):
		suite.Fail("Timeout while waiting for event")
	}
}

func (suite *BackupControllerTestSuite) SetupTest() {
	// Turn off logging output
	logrus.SetOutput(ioutil.Discard)

	// Initialize the controller
	suite.watcher, suite.controller = NewFakeBackupController(16)
	suite.eventsHook = controller.NewEventsHook(16)
	err := suite.controller.AddHook(suite.eventsHook)
	suite.Require().Nil(err)

	// Test Backup
	suite.backup = new(crv1.MySQLBackup)
	err = factory.Build(testFactory.MySQLBackupFactory).To(suite.backup)
	suite.Require().Nil(err)
	suite.watcher.Add(suite.backup)

	// Start the controller
	ctx, cancelFunc := context.WithCancel(context.Background())
	suite.cancelFunc = cancelFunc

	go suite.controller.Run(ctx)
}

func (suite *BackupControllerTestSuite) TearDownTest() {
	suite.cancelFunc()
}

// Test if onAdd function is being called.
func (suite *BackupControllerTestSuite) TestBackupController_OnAdd() {
	suite.testWithTimeout(func(event controller.Event) {
		suite.Require().Equal(controller.EventAdded, event.Type)
		suite.Equal(suite.backup, event.Object.(*crv1.MySQLBackup))
	})
}

// Test if onUpdate function is being called.
func (suite *BackupControllerTestSuite) TestBackupController_OnUpdate() {
	// Ignore added event
	suite.testWithTimeout(func(event controller.Event) {
		suite.Require().Equal(controller.EventAdded, event.Type)
	})

	// Update backup
	suite.watcher.Modify(suite.backup)

	suite.testWithTimeout(func(event controller.Event) {
		suite.Require().Equal(controller.EventUpdated, event.Type)
		suite.Equal(suite.backup, event.Object.(*crv1.MySQLBackup))
	})
}

// Test if onDelete function is being called.
func (suite *BackupControllerTestSuite) TestBackupController_OnDelete() {
	// Ignore added event
	suite.testWithTimeout(func(event controller.Event) {
		suite.Require().Equal(controller.EventAdded, event.Type)
	})

	// Delete backup
	suite.watcher.Delete(suite.backup)

	suite.testWithTimeout(func(event controller.Event) {
		suite.Require().Equal(controller.EventDeleted, event.Type)
		suite.Equal(suite.backup, event.Object.(*crv1.MySQLBackup))
	})
}

func TestBackupControllerTestSuite(t *testing.T) {
	suite.Run(t, new(BackupControllerTestSuite))
}
