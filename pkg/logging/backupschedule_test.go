package logging_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/grtl/mysql-operator/pkg/logging"

	"io/ioutil"

	"github.com/nauyey/factory"
	"github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/test"

	crv1 "github.com/grtl/mysql-operator/pkg/apis/cr/v1"
	testingFactory "github.com/grtl/mysql-operator/pkg/testing/factory"
)

var _ = Describe("Backup Schedule", func() {
	// Turn off logging output
	logrus.SetOutput(ioutil.Discard)

	var (
		logrusHook *test.Hook
		schedule   *crv1.MySQLBackupSchedule
	)

	BeforeEach(func() {
		// Initialize logging hook
		logrusHook = test.NewGlobal()
		logrus.SetLevel(logrus.DebugLevel)

		// Setup fake backupSchedule
		schedule = new(crv1.MySQLBackupSchedule)
		err := factory.Build(testingFactory.MySQLBackupScheduleFactory).To(schedule)
		Expect(err).NotTo(HaveOccurred())
	})

	Context("Debug", func() {
		It("should log with debug level", func() {
			LogBackupSchedule(schedule).Debug("Debug")
			Expect(logrusHook.Entries).To(HaveLen(1))
			Expect(logrusHook.LastEntry().Level).To(Equal(logrus.DebugLevel))
			Expect(logrusHook.LastEntry().Message).To(Equal("Debug"))
			Expect(logrusHook.LastEntry().Data).To(Equal(logrus.Fields{
				"backupSchedule": schedule.Name,
			}))
		})
	})

	Context("Info", func() {
		It("should log with info level", func() {
			LogBackupSchedule(schedule).Info("Info")
			Expect(logrusHook.Entries).To(HaveLen(1))
			Expect(logrusHook.LastEntry().Level).To(Equal(logrus.InfoLevel))
			Expect(logrusHook.LastEntry().Message).To(Equal("Info"))
			Expect(logrusHook.LastEntry().Data).To(Equal(logrus.Fields{
				"backupSchedule": schedule.Name,
			}))
		})
	})

	Context("Warn", func() {
		It("should log with warn level", func() {
			LogBackupSchedule(schedule).Warn("Warn")
			Expect(logrusHook.Entries).To(HaveLen(1))
			Expect(logrusHook.LastEntry().Level).To(Equal(logrus.WarnLevel))
			Expect(logrusHook.LastEntry().Message).To(Equal("Warn"))
			Expect(logrusHook.LastEntry().Data).To(Equal(logrus.Fields{
				"backupSchedule": schedule.Name,
			}))
		})
	})

	Context("Error", func() {
		It("should log with fail level", func() {
			LogBackupSchedule(schedule).Error("Error")
			Expect(logrusHook.Entries).To(HaveLen(1))
			Expect(logrusHook.LastEntry().Level).To(Equal(logrus.ErrorLevel))
			Expect(logrusHook.LastEntry().Message).To(Equal("Error"))
			Expect(logrusHook.LastEntry().Data).To(Equal(logrus.Fields{
				"backupSchedule": schedule.Name,
			}))
		})
	})
})
