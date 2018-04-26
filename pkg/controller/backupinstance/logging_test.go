package backupinstance_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gstruct"

	. "github.com/grtl/mysql-operator/pkg/controller/backupinstance"

	"context"
	"io/ioutil"

	"k8s.io/apimachinery/pkg/watch"

	"github.com/nauyey/factory"
	"github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/test"

	crv1 "github.com/grtl/mysql-operator/pkg/apis/cr/v1"
	"github.com/grtl/mysql-operator/pkg/controller"
	testingFactory "github.com/grtl/mysql-operator/pkg/testing/factory"
)

var _ = Describe("Logging", func() {
	// Turn off logging output
	logrus.SetOutput(ioutil.Discard)

	var (
		logrusHook *test.Hook

		backup *crv1.MySQLBackupInstance

		watcher            *watch.FakeWatcher
		instanceController controller.Controller
		eventsHook         controller.EventsHook
	)

	BeforeEach(func() {
		// Initialize logging hook
		logrusHook = test.NewGlobal()

		// Initialize the controller
		watcher, instanceController = NewFakeBackupInstanceController(16)
		eventsHook = controller.NewEventsHook(16)

		// Setup fake backup instance
		backup = new(crv1.MySQLBackupInstance)
		err := factory.Build(testingFactory.MySQLBackupInstanceFactory).To(backup)
		Expect(err).NotTo(HaveOccurred())
	})

	JustBeforeEach(func() {
		err := instanceController.AddHook(eventsHook)
		Expect(err).NotTo(HaveOccurred())

		watcher.Add(backup)
	})

	When("Backup Instance is added", func() {
		It("event should be logged", func(done Done) {
			var event controller.Event

			ctx, cancelFunc := context.WithCancel(context.Background())
			go instanceController.Run(ctx)
			defer cancelFunc()

			// Wait for
			Eventually(eventsHook.GetEventsChan()).Should(Receive(&event))
			Expect(logrusHook.AllEntries()).To(HaveLen(2))

			By("outputting on event received")
			firstEntry := logrusHook.AllEntries()[0]
			Expect(*firstEntry).To(gstruct.MatchFields(gstruct.IgnoreExtras, gstruct.Fields{
				"Level":   Equal(logrus.InfoLevel),
				"Message": Equal("Received BackupInstance event"),
				"Data": Equal(logrus.Fields{
					"event":          BackupInstanceAdded,
					"backupInstance": backup.Name,
				}),
			}))

			By("outputting on event processed")
			secondEntry := logrusHook.AllEntries()[1]
			Expect(*secondEntry).To(gstruct.MatchFields(gstruct.IgnoreExtras, gstruct.Fields{
				"Level":   Equal(logrus.InfoLevel),
				"Message": Equal("Successfully processed BackupInstance event"),
				"Data": Equal(logrus.Fields{
					"event":          BackupInstanceAdded,
					"backupInstance": backup.Name,
				}),
			}))

			close(done)
		})
	})

	When("Backup Instance is updated", func() {
		It("event should be logged", func(done Done) {
			var event controller.Event

			ctx, cancelFunc := context.WithCancel(context.Background())
			go instanceController.Run(ctx)
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
				"Message": Equal("Received BackupInstance event"),
				"Data": Equal(logrus.Fields{
					"event":          BackupInstanceUpdated,
					"backupInstance": backup.Name,
				}),
			}))

			By("outputting on event processed")
			secondEntry := logrusHook.AllEntries()[1]
			Expect(*secondEntry).To(gstruct.MatchFields(gstruct.IgnoreExtras, gstruct.Fields{
				"Level":   Equal(logrus.InfoLevel),
				"Message": Equal("Successfully processed BackupInstance event"),
				"Data": Equal(logrus.Fields{
					"event":          BackupInstanceUpdated,
					"backupInstance": backup.Name,
				}),
			}))

			close(done)
		})
	})

	When("Backup Instance is deleted", func() {
		It("event should be logged", func(done Done) {
			var event controller.Event

			ctx, cancelFunc := context.WithCancel(context.Background())
			go instanceController.Run(ctx)
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
				"Message": Equal("Received BackupInstance event"),
				"Data": Equal(logrus.Fields{
					"backupInstance": backup.Name,
					"event":          BackupInstanceDeleted,
				}),
			}))

			By("outputting on event processed")
			secondEntry := logrusHook.AllEntries()[1]
			Expect(*secondEntry).To(gstruct.MatchFields(gstruct.IgnoreExtras, gstruct.Fields{
				"Level":   Equal(logrus.InfoLevel),
				"Message": Equal("Successfully processed BackupInstance event"),
				"Data": Equal(logrus.Fields{
					"backupInstance": backup.Name,
					"event":          BackupInstanceDeleted,
				}),
			}))

			close(done)
		})
	})
})
