package controller_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/grtl/mysql-operator/controller"

	"github.com/nauyey/factory"

	crv1 "github.com/grtl/mysql-operator/pkg/apis/cr/v1"
	testingFactory "github.com/grtl/mysql-operator/testing/factory"
)

var _ = Describe("Events", func() {
	var (
		eventsHook EventsHook
		cluster    *crv1.MySQLCluster
	)

	BeforeEach(func() {
		eventsHook = NewEventsHook(16)
		cluster = new(crv1.MySQLCluster)
		err := factory.Build(testingFactory.MySQLClusterFactory).To(cluster)
		Expect(err).NotTo(HaveOccurred())
	})

	Describe("Adding cluster object", func() {
		It("should append EventAdded with cluster object to the events channel", func(done Done) {
			go eventsHook.OnAdd(cluster)
			Expect(<-eventsHook.GetEventsChan()).To(Equal(Event{
				Type:   EventAdded,
				Object: cluster,
			}))
			close(done)
		})
	})

	Describe("Updating cluster object", func() {
		It("should append EventUpdated with cluster object to the events channel", func(done Done) {
			go eventsHook.OnUpdate(cluster)
			Expect(<-eventsHook.GetEventsChan()).To(Equal(Event{
				Type:   EventUpdated,
				Object: cluster,
			}))
			close(done)
		})
	})

	Describe("Deleting cluster object", func() {
		It("should append EventDeleted with cluster object to the events channel", func(done Done) {
			go eventsHook.OnDelete(cluster)
			Expect(<-eventsHook.GetEventsChan()).To(Equal(Event{
				Type:   EventDeleted,
				Object: cluster,
			}))
			close(done)
		})
	})

	Describe("Getting events channel", func() {
		It("should not be nil", func() {
			Expect(eventsHook.GetEventsChan()).ToNot(BeNil())
		})
	})
})
