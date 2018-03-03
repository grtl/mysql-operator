package cluster

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	storage    string
	backupName string
)

var clusterCreateCmd = &cobra.Command{
	Use:   "create [cluster name]",
	Short: "short description of cluster create",
	Long: `A longer description of cluster create with examples:
msp cluster create "my-cluster" --storage 1Gi`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("Cluster name is required")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		if backupName != "" {
			//TODO: create cluster from backup logic
			fmt.Println("create cluster --storage --from-backup called")
		} else {
			//TODO: create cluster (without backup) logic
			fmt.Println("create cluster --storage called")
		}
	},
}

func init() {
	Cmd.AddCommand(clusterCreateCmd)

	clusterCreateCmd.Flags().StringVar(&storage, "storage", "", "storage value")
	clusterCreateCmd.MarkFlagRequired("storage")
	clusterCreateCmd.Flags().StringVar(&backupName, "from-backup", "", "path for backup")
}
