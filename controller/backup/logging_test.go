package backup_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gstruct"

	. "github.com/grtl/mysql-operator/controller/backup"

	"context"
	"io/ioutil"

	"k8s.io/apimachinery/pkg/watch"

	"github.com/nauyey/factory"
	"github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/test"

	"github.com/grtl/mysql-operator/controller"
	crv1 "github.com/grtl/mysql-operator/pkg/apis/cr/v1"
	testingFactory "github.com/grtl/mysql-operator/testing/factory"
)

var _ = Describe("Logging", func() {
	// Turn off logging output
	logrus.SetOutput(ioutil.Discard)

	var (
		logrusHook *test.Hook

		backup *crv1.MySQLBackup

		watcher          *watch.FakeWatcher
		backupController controller.Controller
		eventsHook       controller.EventsHook
	)

	BeforeEach(func() {
		// Initialize logging hook
		logrusHook = test.NewGlobal()

		// Initialize the controller
		watcher, backupController = NewFakeBackupController(16)
		eventsHook = controller.NewEventsHook(16)

		// Setup fake backup
		backup = new(crv1.MySQLBackup)
		err := factory.Build(testingFactory.MySQLBackupFactory).To(backup)
		Expect(err).NotTo(HaveOccurred())
	})

	JustBeforeEach(func() {
		err := backupController.AddHook(eventsHook)
		Expect(err).NotTo(HaveOccurred())

		watcher.Add(backup)
	})

	When("backup is added", func() {
		It("event should be logged", func(done Done) {
			var event controller.Event

			ctx, cancelFunc := context.WithCancel(context.Background())
			go backupController.Run(ctx)
			defer cancelFunc()

			// Wait for
			Eventually(eventsHook.GetEventsChan()).Should(Receive(&event))
			Expect(logrusHook.AllEntries()).To(HaveLen(2))

			By("outputting on event received")
			firstEntry := logrusHook.AllEntries()[0]
			Expect(*firstEntry).To(gstruct.MatchFields(gstruct.IgnoreExtras, gstruct.Fields{
				"Level":   Equal(logrus.InfoLevel),
				"Message": Equal("Received backup event"),
				"Data": Equal(logrus.Fields{
					"event":  BackupAdded,
					"backup": backup.Name,
				}),
			}))

			By("outputting on event processed")
			secondEntry := logrusHook.AllEntries()[1]
			Expect(*secondEntry).To(gstruct.MatchFields(gstruct.IgnoreExtras, gstruct.Fields{
				"Level":   Equal(logrus.InfoLevel),
				"Message": Equal("Successfully processed backup event"),
				"Data": Equal(logrus.Fields{
					"event":  BackupAdded,
					"backup": backup.Name,
				}),
			}))

			close(done)
		})
	})

	When("backup is updated", func() {
		It("event should be logged", func(done Done) {
			var event controller.Event

			ctx, cancelFunc := context.WithCancel(context.Background())
			go backupController.Run(ctx)
			defer cancelFunc()

			// Ignore added event
			Eventually(eventsHook.GetEventsChan()).Should(Receive(&event))
			logrusHook.Reset()
			watcher.Modify(backup)

			// Wait for
			Eventually(eventsHook.GetEventsChan()).Should(Receive(&event))
			Expect(logrusHook.AllEntries()).To(HaveLen(2))

			By("outputting on event received")
			firstEntry := logrusHook.AllEntries()[0]
			Expect(*firstEntry).To(gstruct.MatchFields(gstruct.IgnoreExtras, gstruct.Fields{
				"Level":   Equal(logrus.InfoLevel),
				"Message": Equal("Received backup event"),
				"Data": Equal(logrus.Fields{
					"event":  BackupUpdated,
					"backup": backup.Name,
				}),
			}))

			By("outputting on event processed")
			secondEntry := logrusHook.AllEntries()[1]
			Expect(*secondEntry).To(gstruct.MatchFields(gstruct.IgnoreExtras, gstruct.Fields{
				"Level":   Equal(logrus.InfoLevel),
				"Message": Equal("Successfully processed backup event"),
				"Data": Equal(logrus.Fields{
					"event":  BackupUpdated,
					"backup": backup.Name,
				}),
			}))

			close(done)
		})
	})

	When("backup is deleted", func() {
		It("event should be logged", func(done Done) {
			var event controller.Event

			ctx, cancelFunc := context.WithCancel(context.Background())
			go backupController.Run(ctx)
			defer cancelFunc()

			// Ignore added event
			Eventually(eventsHook.GetEventsChan()).Should(Receive(&event))
			logrusHook.Reset()
			watcher.Delete(backup)

			// Wait for
			Eventually(eventsHook.GetEventsChan()).Should(Receive(&event))
			Expect(logrusHook.AllEntries()).To(HaveLen(2))

			By("outputting on event received")
			firstEntry := logrusHook.AllEntries()[0]
			Expect(*firstEntry).To(gstruct.MatchFields(gstruct.IgnoreExtras, gstruct.Fields{
				"Level":   Equal(logrus.InfoLevel),
				"Message": Equal("Received backup event"),
				"Data": Equal(logrus.Fields{
					"backup": backup.Name,
					"event":  BackupDeleted,
				}),
			}))

			By("outputting on event processed")
			secondEntry := logrusHook.AllEntries()[1]
			Expect(*secondEntry).To(gstruct.MatchFields(gstruct.IgnoreExtras, gstruct.Fields{
				"Level":   Equal(logrus.InfoLevel),
				"Message": Equal("Successfully processed backup event"),
				"Data": Equal(logrus.Fields{
					"backup": backup.Name,
					"event":  BackupDeleted,
				}),
			}))

			close(done)
		})
	})
})
