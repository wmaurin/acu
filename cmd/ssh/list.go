package ssh

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List saved SSH connections (TBD)",
	Long: `List all saved SSH connection profiles.

Displays alias, user, host, port, and key file for each saved connection.

Examples:
  acu ssh list    # Show all saved SSH connections`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ssh list called")
	},
}

