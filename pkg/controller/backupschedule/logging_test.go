package backupschedule_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gstruct"

	. "github.com/grtl/mysql-operator/pkg/controller/backupschedule"

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

		schedule *crv1.MySQLBackupSchedule

		watcher            *watch.FakeWatcher
		scheduleController controller.Controller
		eventsHook         controller.EventsHook
	)

	BeforeEach(func() {
		// Initialize logging hook
		logrusHook = test.NewGlobal()

		// Initialize the controller
		watcher, scheduleController = NewFakeBackupScheduleController(16)
		eventsHook = controller.NewEventsHook(16)

		// Setup fake backupschedule
		schedule = new(crv1.MySQLBackupSchedule)
		err := factory.Build(testingFactory.MySQLBackupScheduleFactory).To(schedule)
		Expect(err).NotTo(HaveOccurred())
	})

	JustBeforeEach(func() {
		err := scheduleController.AddHook(eventsHook)
		Expect(err).NotTo(HaveOccurred())

		watcher.Add(schedule)
	})

	When("Backup Schedule is added", func() {
		It("event should be logged", func(done Done) {
			var event controller.Event

			ctx, cancelFunc := context.WithCancel(context.Background())
			go scheduleController.Run(ctx)
			defer cancelFunc()

			// Wait for
			Eventually(eventsHook.GetEventsChan()).Should(Receive(&event))
			Expect(logrusHook.AllEntries()).To(HaveLen(2))

			By("outputting on event received")
			firstEntry := logrusHook.AllEntries()[0]
			Expect(*firstEntry).To(gstruct.MatchFields(gstruct.IgnoreExtras, gstruct.Fields{
				"Level":   Equal(logrus.InfoLevel),
				"Message": Equal("Received BackupSchedule event"),
				"Data": Equal(logrus.Fields{
					"event":          BackupScheduleAdded,
					"backupSchedule": schedule.Name,
				}),
			}))

			By("outputting on event processed")
			secondEntry := logrusHook.AllEntries()[1]
			Expect(*secondEntry).To(gstruct.MatchFields(gstruct.IgnoreExtras, gstruct.Fields{
				"Level":   Equal(logrus.InfoLevel),
				"Message": Equal("Successfully processed BackupSchedule event"),
				"Data": Equal(logrus.Fields{
					"event":          BackupScheduleAdded,
					"backupSchedule": schedule.Name,
				}),
			}))

			close(done)
		})
	})

	When("Backup Schedule is updated", func() {
		It("event should be logged", func(done Done) {
			var event controller.Event

			ctx, cancelFunc := context.WithCancel(context.Background())
			go scheduleController.Run(ctx)
			defer cancelFunc()

			// Ignore added event
			Eventually(eventsHook.GetEventsChan()).Should(Receive(&event))
			logrusHook.Reset()
			watcher.Modify(schedule)

			// Wait for
			Eventually(eventsHook.GetEventsChan()).Should(Receive(&event))
			Expect(logrusHook.AllEntries()).To(HaveLen(2))

			By("outputting on event received")
			firstEntry := logrusHook.AllEntries()[0]
			Expect(*firstEntry).To(gstruct.MatchFields(gstruct.IgnoreExtras, gstruct.Fields{
				"Level":   Equal(logrus.InfoLevel),
				"Message": Equal("Received BackupSchedule event"),
				"Data": Equal(logrus.Fields{
					"event":          BackupScheduleUpdated,
					"backupSchedule": schedule.Name,
				}),
			}))

			By("outputting on event processed")
			secondEntry := logrusHook.AllEntries()[1]
			Expect(*secondEntry).To(gstruct.MatchFields(gstruct.IgnoreExtras, gstruct.Fields{
				"Level":   Equal(logrus.InfoLevel),
				"Message": Equal("Successfully processed BackupSchedule event"),
				"Data": Equal(logrus.Fields{
					"event":          BackupScheduleUpdated,
					"backupSchedule": schedule.Name,
				}),
			}))

			close(done)
		})
	})

	When("Backup Schedule is deleted", func() {
		It("event should be logged", func(done Done) {
			var event controller.Event

			ctx, cancelFunc := context.WithCancel(context.Background())
			go scheduleController.Run(ctx)
			defer cancelFunc()

			// Ignore added event
			Eventually(eventsHook.GetEventsChan()).Should(Receive(&event))
			logrusHook.Reset()
			watcher.Delete(schedule)

			// Wait for
			Eventually(eventsHook.GetEventsChan()).Should(Receive(&event))
			Expect(logrusHook.AllEntries()).To(HaveLen(2))

			By("outputting on event received")
			firstEntry := logrusHook.AllEntries()[0]
			Expect(*firstEntry).To(gstruct.MatchFields(gstruct.IgnoreExtras, gstruct.Fields{
				"Level":   Equal(logrus.InfoLevel),
				"Message": Equal("Received BackupSchedule event"),
				"Data": Equal(logrus.Fields{
					"backupSchedule": schedule.Name,
					"event":          BackupScheduleDeleted,
				}),
			}))

			By("outputting on event processed")
			secondEntry := logrusHook.AllEntries()[1]
			Expect(*secondEntry).To(gstruct.MatchFields(gstruct.IgnoreExtras, gstruct.Fields{
				"Level":   Equal(logrus.InfoLevel),
				"Message": Equal("Successfully processed BackupSchedule event"),
				"Data": Equal(logrus.Fields{
					"backupSchedule": schedule.Name,
					"event":          BackupScheduleDeleted,
				}),
			}))

			close(done)
		})
	})
})
