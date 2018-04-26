package backupinstance_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestBackupSchedule(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Controller Backup Schedule Suite")
}
