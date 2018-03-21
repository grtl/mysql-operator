package cluster

import "github.com/spf13/cobra"

// Cmd provides commands for managing cluster resources.
var Cmd = &cobra.Command{
	Use:   "cluster",
	Short: "a short description of cluster",
	Long: `A longer description of cluster with examples:
msp cluster create
msp cluster delete
msp cluster update`,
}
