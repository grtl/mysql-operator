package backup

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

var backupDeleteCmd = &cobra.Command{
	Use:   "delete [backup name]",
	Short: "A short description of backup delete",
	Long: `A longer description of backup delete with examples:
msp backup delete "my-cluster"`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("Backup name is required")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		//TODO: Delete backup logic
		fmt.Println("backup delete called")
	},
}

func init() {
	Cmd.AddCommand(backupDeleteCmd)
}
