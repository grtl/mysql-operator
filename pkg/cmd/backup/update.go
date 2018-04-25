package backup

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

var backupUpdateCmd = &cobra.Command{
	Use:   "update [backupschedule name]",
	Short: "A short description of backup update",
	Long: `A longer description of backup update with usage:
msp backup update "my-backup"`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("Backup name is required")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		//TODO: update backup logic
		fmt.Println("backup update called")
	},
}

func init() {
	Cmd.AddCommand(backupUpdateCmd)
}
