package options

import (
	"github.com/spf13/cobra"
)

// ExtractOptions extracts common flags from the cmd and packs them into a struct.
func ExtractOptions(cmd *cobra.Command) *Options {
	var err error
	options := new(Options)

	options.Namespace, err = cmd.InheritedFlags().GetString("namespace")
	if err != nil {
		panic(err)
	}

	options.Force, err = cmd.InheritedFlags().GetBool("force")
	if err != nil {
		panic(err)
	}

	return options
}
