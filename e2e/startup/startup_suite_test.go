package startup_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/grtl/mysql-operator/pkg/testing/e2e"
)

var operator e2e.Operator

func TestStartup(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "(e2e) Startup Suite")
}

var _ = BeforeSuite(func() {
	var err error

	operator, err = e2e.NewOperator()
	Expect(err).NotTo(HaveOccurred())

	err = operator.Start()
	Expect(err).NotTo(HaveOccurred())
})

var _ = AfterSuite(func() {
	if operator == nil {
		// Something went wrong during setup, return to avoid segfault
		return
	}
	err := operator.Stop()
	Expect(err).NotTo(HaveOccurred())
})
