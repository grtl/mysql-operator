package backup

import "github.com/spf13/cobra"

var Cmd = &cobra.Command{
	Use:   "backup",
	Short: "a short description of backup",
	Long: `A longer description of backup with examples:
msp backup create
msp backup delete
msp backup update`,
}
