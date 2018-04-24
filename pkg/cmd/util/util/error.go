package util

import (
	"fmt"
	"os"

	"github.com/grtl/mysql-operator/pkg/cmd/util/options"
)

// FailOnErrorOrForceContinue checks if error occurred and stops the execution
// unless the force flag was specified.
func FailOnErrorOrForceContinue(err error, opts *options.Options) {
	if err == nil {
		return
	}

	fmt.Fprintln(os.Stderr, err.Error())
	if !opts.Force {
		os.Exit(1)
	}
}

// Fail with error prints out an error to stderr and exits with EXIT_FAILURE.
func FailWithError(err error) {
	fmt.Fprintln(os.Stderr, err.Error())
	os.Exit(1)
}
