package backup

import (
	"fmt"

	"github.com/spf13/cobra"
)

var clusterName string

var backupCreateCmd = &cobra.Command{
	Use:   "create [backup name]",
	Short: "short description of backup create",
	Long: `A longer description of backup create with examples:
msp backup create "my-backup" --cluster "my-cluster"`,
	Run: func(cmd *cobra.Command, args []string) {
		backup := "mysql-backup"
		if len(args) >= 1 {
			backup = args[0]
		}

		//TODO: create backup logic
		fmt.Println(backup)
		fmt.Println("backup create--cluster called")
	},
}

func init() {
	Cmd.AddCommand(backupCreateCmd)

	backupCreateCmd.Flags().StringVar(&clusterName, "cluster", "",
		"name of cluster for which the backup is made")
	backupCreateCmd.MarkFlagRequired("cluster")
}
