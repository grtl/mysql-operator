package backupinstance_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/grtl/mysql-operator/pkg/controller/backupinstance"

	"context"
	"io/ioutil"

	"k8s.io/apimachinery/pkg/watch"

	"github.com/nauyey/factory"
	"github.com/sirupsen/logrus"

	crv1 "github.com/grtl/mysql-operator/pkg/apis/cr/v1"
	"github.com/grtl/mysql-operator/pkg/client/clientset/versioned/fake"
	"github.com/grtl/mysql-operator/pkg/controller"
	testingFactory "github.com/grtl/mysql-operator/pkg/testing/factory"
)

var _ = Describe("Backup Instance Controller", func() {
	// Turn off logging output
	logrus.SetOutput(ioutil.Discard)

	var (
		schedule *crv1.MySQLBackupSchedule
		backup   *crv1.MySQLBackupInstance

		clientset          *fake.Clientset
		watcher            *watch.FakeWatcher
		scheduleController controller.Controller
		eventsHook         controller.EventsHook
	)

	BeforeEach(func() {
		// Initialize the controller
		clientset, watcher, scheduleController = NewFakeBackupInstanceController(16)
		eventsHook = controller.NewEventsHook(16)

		// Setup fake Backup Schedule
		schedule = new(crv1.MySQLBackupSchedule)
		err := factory.Build(testingFactory.MySQLBackupScheduleFactory).To(schedule)
		Expect(err).NotTo(HaveOccurred())

		_, err = clientset.CrV1().MySQLBackupSchedules(schedule.Namespace).Create(schedule)
		Expect(err).NotTo(HaveOccurred())

		// Setup fake Backup Instance
		backup = new(crv1.MySQLBackupInstance)
		err = factory.Build(testingFactory.MySQLBackupInstanceFactory,
			factory.WithField("Spec.Schedule", schedule.Name)).To(backup)
		Expect(err).NotTo(HaveOccurred())
	})

	JustBeforeEach(func() {
		err := scheduleController.AddHook(eventsHook)
		Expect(err).NotTo(HaveOccurred())

		watcher.Add(backup)
	})

	When("Backup Instance is added", func() {
		It("should get processed by the controller", func(done Done) {
			var event controller.Event

			ctx, cancelFunc := context.WithCancel(context.Background())
			go scheduleController.Run(ctx)
			defer cancelFunc()

			Eventually(eventsHook.GetEventsChan()).Should(Receive(&event))
			Expect(event.Type).To(Equal(controller.EventAdded))
			Expect(event.Object).To(Equal(backup))

			close(done)
		})
	})

	When("Backup Instance is updated", func() {
		It("should get processed by the controller", func(done Done) {
			var event controller.Event

			ctx, cancelFunc := context.WithCancel(context.Background())
			go scheduleController.Run(ctx)
			defer cancelFunc()

			// Ignore added event
			Eventually(eventsHook.GetEventsChan()).Should(Receive(&event))
			Expect(event.Type).To(Equal(controller.EventAdded))

			// Update backup backup
			watcher.Modify(backup)

			Eventually(eventsHook.GetEventsChan()).Should(Receive(&event))
			Expect(event.Type).To(Equal(controller.EventUpdated))
			Expect(event.Object).To(Equal(backup))

			close(done)
		})
	})

	When("Backup Instance is deleted", func() {
		It("should get processed by the controller", func(done Done) {
			var event controller.Event

			ctx, cancelFunc := context.WithCancel(context.Background())
			go scheduleController.Run(ctx)
			defer cancelFunc()

			// Ignore added event
			Eventually(eventsHook.GetEventsChan()).Should(Receive(&event))
			Expect(event.Type).To(Equal(controller.EventAdded))

			// Update Backup Instance
			watcher.Delete(backup)

			Eventually(eventsHook.GetEventsChan()).Should(Receive(&event))
			Expect(event.Type).To(Equal(controller.EventDeleted))
			Expect(event.Object).To(Equal(backup))

			close(done)
		})
	})
})
