package fail

import (
	"fmt"
	"os"

	"github.com/grtl/mysql-operator/pkg/cmd/util/options"
)

// ErrorOrForceContinue checks if an error occurred and stops the execution
// unless the force flag was specified.
func ErrorOrForceContinue(err error, opts *options.Options) {
	if err == nil {
		return
	}

	fmt.Fprintln(os.Stderr, err.Error())
	if !opts.Force {
		os.Exit(1)
	}
}

// Error prints out an error to stderr and exits with EXIT_FAILURE.
func Error(err error) {
	fmt.Fprintln(os.Stderr, err.Error())
	os.Exit(1)
}
