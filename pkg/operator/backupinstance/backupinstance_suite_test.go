package backupinstance_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestBackupInstance(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Operator Backup Instance Suite")
}
