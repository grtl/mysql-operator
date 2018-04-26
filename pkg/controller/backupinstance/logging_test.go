package backupinstance_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gstruct"

	"context"
	"io/ioutil"

	"k8s.io/apimachinery/pkg/watch"

	"github.com/nauyey/factory"
	"github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/test"

	crv1 "github.com/grtl/mysql-operator/pkg/apis/cr/v1"
	"github.com/grtl/mysql-operator/pkg/client/clientset/versioned/fake"
	"github.com/grtl/mysql-operator/pkg/controller"
	. "github.com/grtl/mysql-operator/pkg/controller/backupinstance"
	testingFactory "github.com/grtl/mysql-operator/pkg/testing/factory"
)

var _ = Describe("Logging", func() {
	// Turn off logging output
	logrus.SetOutput(ioutil.Discard)

	var (
		logrusHook *test.Hook

		schedule *crv1.MySQLBackupSchedule
		backup   *crv1.MySQLBackupInstance

		clientset          *fake.Clientset
		watcher            *watch.FakeWatcher
		instanceController controller.Controller
		eventsHook         controller.EventsHook
	)

	BeforeEach(func() {
		// Initialize logging hook
		logrusHook = test.NewGlobal()

		// Initialize the controller
		clientset, watcher, instanceController = NewFakeBackupInstanceController(16)
		eventsHook = controller.NewEventsHook(16)

		// Setup fake backup schedule
		schedule = new(crv1.MySQLBackupSchedule)
		err := factory.Build(testingFactory.MySQLBackupScheduleFactory).To(schedule)
		Expect(err).NotTo(HaveOccurred())

		_, err = clientset.CrV1().MySQLBackupSchedules(schedule.Namespace).Create(schedule)
		Expect(err).NotTo(HaveOccurred())

		// Setup fake backup instance
		backup = new(crv1.MySQLBackupInstance)
		err = factory.Build(testingFactory.MySQLBackupInstanceFactory,
			factory.WithField("Spec.Schedule", schedule)).To(backup)
		Expect(err).NotTo(HaveOccurred())
	})

	JustBeforeEach(func() {
		err := instanceController.AddHook(eventsHook)
		Expect(err).NotTo(HaveOccurred())
	})

	When("Backup Instance is added", func() {
		Describe("with an existing backup schedule", func() {
			JustBeforeEach(func() {
				backup.Spec.Schedule = schedule.Name
				watcher.Add(backup)
			})

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

		Describe("with non-existing backup schedule", func() {
			JustBeforeEach(func() {
				watcher.Add(backup)
			})

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

				By("outputting on event error")
				secondEntry := logrusHook.AllEntries()[1]
				Expect(*secondEntry).To(gstruct.MatchFields(gstruct.IgnoreExtras, gstruct.Fields{
					"Level":   Equal(logrus.ErrorLevel),
					"Message": MatchRegexp("mysqlbackupschedules.cr.mysqloperator.grtl.github.com \"backup-.*\" not found"),
					"Data": Equal(logrus.Fields{
						"event":          BackupInstanceAdded,
						"backupInstance": backup.Name,
					}),
				}))

				close(done)
			})
		})
	})

	When("Backup Instance is updated", func() {
		JustBeforeEach(func() {
			watcher.Add(backup)
		})

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
		JustBeforeEach(func() {
			watcher.Add(backup)
		})

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
