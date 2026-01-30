package ssh

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rmCmd = &cobra.Command{
	Use:   "rm <alias>",
	Short: "Remove a saved SSH connection (TBD)",
	Long: `Remove a saved SSH connection profile by its alias.

Examples:
  acu ssh rm prod      # Remove the 'prod' connection
  acu ssh rm dev       # Remove the 'dev' connection`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ssh rm called")
	},
}

