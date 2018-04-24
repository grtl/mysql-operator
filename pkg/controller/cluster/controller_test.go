package cluster_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/grtl/mysql-operator/pkg/controller/cluster"

	"context"
	"io/ioutil"

	"k8s.io/apimachinery/pkg/watch"

	"github.com/nauyey/factory"
	"github.com/sirupsen/logrus"

	crv1 "github.com/grtl/mysql-operator/pkg/apis/cr/v1"
	"github.com/grtl/mysql-operator/pkg/controller"
	testingFactory "github.com/grtl/mysql-operator/pkg/testing/factory"
)

var _ = Describe("Cluster Controller", func() {
	// Turn off logging output
	logrus.SetOutput(ioutil.Discard)

	var (
		cluster *crv1.MySQLCluster

		watcher           *watch.FakeWatcher
		clusterController controller.Controller
		eventsHook        controller.EventsHook
	)

	BeforeEach(func() {
		// Initialize the controller
		watcher, clusterController = NewFakeClusterController(16)
		eventsHook = controller.NewEventsHook(16)

		// Setup fake cluster
		cluster = new(crv1.MySQLCluster)
		err := factory.Build(testingFactory.MySQLClusterFactory).To(cluster)
		Expect(err).NotTo(HaveOccurred())
	})

	JustBeforeEach(func() {
		err := clusterController.AddHook(eventsHook)
		Expect(err).NotTo(HaveOccurred())

		watcher.Add(cluster)
	})

	When("Cluster is added", func() {
		It("should get processed by the controller", func(done Done) {
			var event controller.Event

			ctx, cancelFunc := context.WithCancel(context.Background())
			go clusterController.Run(ctx)
			defer cancelFunc()

			Eventually(eventsHook.GetEventsChan()).Should(Receive(&event))
			Expect(event.Type).To(Equal(controller.EventAdded))
			Expect(event.Object).To(Equal(cluster))

			close(done)
		})
	})

	When("Cluster is updated", func() {
		It("should get processed by the controller", func(done Done) {
			var event controller.Event

			ctx, cancelFunc := context.WithCancel(context.Background())
			go clusterController.Run(ctx)
			defer cancelFunc()

			// Ignore added event
			Eventually(eventsHook.GetEventsChan()).Should(Receive(&event))
			Expect(event.Type).To(Equal(controller.EventAdded))

			// Update cluster
			watcher.Modify(cluster)

			Eventually(eventsHook.GetEventsChan()).Should(Receive(&event))
			Expect(event.Type).To(Equal(controller.EventUpdated))
			Expect(event.Object).To(Equal(cluster))

			close(done)
		})
	})

	When("Cluster is deleted", func() {
		It("should get processed by the controller", func(done Done) {
			var event controller.Event

			ctx, cancelFunc := context.WithCancel(context.Background())
			go clusterController.Run(ctx)
			defer cancelFunc()

			// Ignore added event
			Eventually(eventsHook.GetEventsChan()).Should(Receive(&event))
			Expect(event.Type).To(Equal(controller.EventAdded))

			// Update cluster
			watcher.Delete(cluster)

			Eventually(eventsHook.GetEventsChan()).Should(Receive(&event))
			Expect(event.Type).To(Equal(controller.EventDeleted))
			Expect(event.Object).To(Equal(cluster))

			close(done)
		})
	})
})
