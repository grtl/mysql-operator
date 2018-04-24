package controller_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/grtl/mysql-operator/pkg/controller"
)

var _ = Describe("Base", func() {
	var (
		controllerBase Base
		hook           = NewEventsHook(1) // Any hook will do
		anotherHook    = NewEventsHook(1) // Any hook will do
	)

	BeforeEach(func() {
		controllerBase = NewControllerBase()
	})

	Describe("Adding hook", func() {
		JustBeforeEach(func() {
			err := controllerBase.AddHook(hook)
			Expect(err).NotTo(HaveOccurred())
		})

		Context("which is not in the hooks list", func() {
			Context("to an empty hooks list", func() {
				It("should add the hook to the list", func() {
					Expect(len(controllerBase.GetHooks())).To(Equal(1))
				})
			})
			Context("to an non empty hooks list", func() {
				It("should add the hook to the list", func() {
					err := controllerBase.AddHook(anotherHook)
					Expect(err).NotTo(HaveOccurred())
					Expect(len(controllerBase.GetHooks())).To(Equal(2))
				})
			})
		})

		Context("which has already been added", func() {
			It("should fail", func() {
				err := controllerBase.AddHook(hook)
				Expect(err).To(HaveOccurred())
				Expect(len(controllerBase.GetHooks())).To(Equal(1))
			})
		})
	})

	Describe("Removing hook", func() {
		BeforeEach(func() {
			err := controllerBase.AddHook(hook)
			Expect(err).NotTo(HaveOccurred())
			err = controllerBase.AddHook(anotherHook)
			Expect(err).NotTo(HaveOccurred())
			Expect(len(controllerBase.GetHooks())).To(Equal(2))
		})

		JustBeforeEach(func() {
			err := controllerBase.RemoveHook(hook)
			Expect(err).NotTo(HaveOccurred())
		})

		Context("which is in the hooks list", func() {
			It("should remove the hook from the list", func() {
				Expect(len(controllerBase.GetHooks())).To(Equal(1))
			})
		})

		Context("which has already been removed", func() {
			It("should fail", func() {
				err := controllerBase.RemoveHook(hook)
				Expect(err).To(HaveOccurred())
				Expect(len(controllerBase.GetHooks())).To(Equal(1))
			})
		})
	})
})
