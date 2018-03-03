package cluster

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

var replicasNumber uint32

var clusterUpdateCmd = &cobra.Command{
	Use:   "update [cluster name]",
	Short: "A short description of cluster update",
	Long: `A longer description of cluster update with usage:
msp cluster update "my-cluster" --replicas=4`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("Cluster name is required")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		//TODO: update cluster logic
		fmt.Println("cluster update called")
	},
}

func init() {
	Cmd.AddCommand(clusterUpdateCmd)

	clusterUpdateCmd.Flags().Uint32Var(&replicasNumber, "replicas", 0, "specify number of replicas")
	clusterUpdateCmd.MarkFlagRequired("replicas")
}
