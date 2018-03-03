package cluster

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

var clusterDeleteCmd = &cobra.Command{
	Use:   "delete [cluster name]",
	Short: "A short description of cluster delete",
	Long: `A longer description of cluster delete with examples:
msp cluster delete "my-cluster"`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("Cluster name is required")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		//TODO: Delete cluster logic
		fmt.Println("cluster delete called")
	},
}

func init() {
	Cmd.AddCommand(clusterDeleteCmd)
}
