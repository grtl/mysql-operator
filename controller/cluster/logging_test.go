package cluster_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gstruct"

	. "github.com/grtl/mysql-operator/controller/cluster"

	"context"
	"io/ioutil"

	"k8s.io/apimachinery/pkg/watch"

	"github.com/nauyey/factory"
	"github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/test"

	"fmt"
	"github.com/grtl/mysql-operator/controller"
	crv1 "github.com/grtl/mysql-operator/pkg/apis/cr/v1"
	testingFactory "github.com/grtl/mysql-operator/testing/factory"
)

var _ = Describe("Logging", func() {
	// Turn off logging output
	logrus.SetOutput(ioutil.Discard)

	var (
		logrusHook *test.Hook

		cluster *crv1.MySQLCluster

		watcher           *watch.FakeWatcher
		clusterController *FakeClusterController
		eventsHook        controller.EventsHook
	)

	BeforeEach(func() {
		// Initialize logging hook
		logrusHook = test.NewGlobal()

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

	When("cluster is added", func() {
		Describe("and succesfully created", func() {
			It("event should be logged", func(done Done) {
				var event controller.Event

				ctx, cancelFunc := context.WithCancel(context.Background())
				go clusterController.Run(ctx)
				defer cancelFunc()

				// Wait for
				Eventually(eventsHook.GetEventsChan()).Should(Receive(&event))
				Expect(logrusHook.AllEntries()).To(HaveLen(2))

				By("outputting on event received")
				firstEntry := logrusHook.AllEntries()[0]
				Expect(*firstEntry).To(gstruct.MatchFields(gstruct.IgnoreExtras, gstruct.Fields{
					"Level":   Equal(logrus.InfoLevel),
					"Message": Equal("Received cluster event"),
					"Data": Equal(logrus.Fields{
						"event":   ClusterAdded,
						"cluster": cluster.Name,
					}),
				}))

				By("outputting on event processed")
				secondEntry := logrusHook.AllEntries()[1]
				Expect(*secondEntry).To(gstruct.MatchFields(gstruct.IgnoreExtras, gstruct.Fields{
					"Level":   Equal(logrus.InfoLevel),
					"Message": Equal("Successfully processed cluster event"),
					"Data": Equal(logrus.Fields{
						"event":   ClusterAdded,
						"cluster": cluster.Name,
					}),
				}))

				close(done)
			})
		})

		Describe("and error occurs", func() {
			It("event should be logged", func(done Done) {
				var event controller.Event

				clusterController.SetError(fmt.Errorf("Testing error"))

				ctx, cancelFunc := context.WithCancel(context.Background())
				go clusterController.Run(ctx)
				defer cancelFunc()

				// Wait for
				Eventually(eventsHook.GetEventsChan()).Should(Receive(&event))
				Expect(logrusHook.AllEntries()).To(HaveLen(2))

				By("outputting on event received")
				firstEntry := logrusHook.AllEntries()[0]
				Expect(*firstEntry).To(gstruct.MatchFields(gstruct.IgnoreExtras, gstruct.Fields{
					"Level":   Equal(logrus.InfoLevel),
					"Message": Equal("Received cluster event"),
					"Data": Equal(logrus.Fields{
						"event":   ClusterAdded,
						"cluster": cluster.Name,
					}),
				}))

				By("outputting an error after processing failed")
				secondEntry := logrusHook.AllEntries()[1]
				Expect(*secondEntry).To(gstruct.MatchFields(gstruct.IgnoreExtras, gstruct.Fields{
					"Level":   Equal(logrus.ErrorLevel),
					"Message": Equal("Testing error"),
					"Data": Equal(logrus.Fields{
						"event":   ClusterAdded,
						"cluster": cluster.Name,
					}),
				}))

				close(done)
			})
		})
	})

	When("cluster is updated", func() {
		It("event should be logged", func(done Done) {
			var event controller.Event

			ctx, cancelFunc := context.WithCancel(context.Background())
			go clusterController.Run(ctx)
			defer cancelFunc()

			// Ignore added event
			Eventually(eventsHook.GetEventsChan()).Should(Receive(&event))
			logrusHook.Reset()
			watcher.Modify(cluster)

			// Wait for
			Eventually(eventsHook.GetEventsChan()).Should(Receive(&event))
			Expect(logrusHook.AllEntries()).To(HaveLen(2))

			By("outputting on event received")
			firstEntry := logrusHook.AllEntries()[0]
			Expect(*firstEntry).To(gstruct.MatchFields(gstruct.IgnoreExtras, gstruct.Fields{
				"Level":   Equal(logrus.InfoLevel),
				"Message": Equal("Received cluster event"),
				"Data": Equal(logrus.Fields{
					"event":   ClusterUpdated,
					"cluster": cluster.Name,
				}),
			}))

			By("outputting on event processed")
			secondEntry := logrusHook.AllEntries()[1]
			Expect(*secondEntry).To(gstruct.MatchFields(gstruct.IgnoreExtras, gstruct.Fields{
				"Level":   Equal(logrus.InfoLevel),
				"Message": Equal("Successfully processed cluster event"),
				"Data": Equal(logrus.Fields{
					"event":   ClusterUpdated,
					"cluster": cluster.Name,
				}),
			}))

			close(done)
		})
	})

	When("cluster is deleted", func() {
		It("event should be logged", func(done Done) {
			var event controller.Event

			ctx, cancelFunc := context.WithCancel(context.Background())
			go clusterController.Run(ctx)
			defer cancelFunc()

			// Ignore added event
			Eventually(eventsHook.GetEventsChan()).Should(Receive(&event))
			logrusHook.Reset()
			watcher.Delete(cluster)

			// Wait for
			Eventually(eventsHook.GetEventsChan()).Should(Receive(&event))
			Expect(logrusHook.AllEntries()).To(HaveLen(2))

			By("outputting on event received")
			firstEntry := logrusHook.AllEntries()[0]
			Expect(*firstEntry).To(gstruct.MatchFields(gstruct.IgnoreExtras, gstruct.Fields{
				"Level":   Equal(logrus.InfoLevel),
				"Message": Equal("Received cluster event"),
				"Data": Equal(logrus.Fields{
					"cluster": cluster.Name,
					"event":   ClusterDeleted,
				}),
			}))

			By("outputting on event processed")
			secondEntry := logrusHook.AllEntries()[1]
			Expect(*secondEntry).To(gstruct.MatchFields(gstruct.IgnoreExtras, gstruct.Fields{
				"Level":   Equal(logrus.InfoLevel),
				"Message": Equal("Successfully processed cluster event"),
				"Data": Equal(logrus.Fields{
					"cluster": cluster.Name,
					"event":   ClusterDeleted,
				}),
			}))

			close(done)
		})
	})
})
