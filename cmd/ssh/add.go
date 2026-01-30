package ssh

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	addKeyFile string
	addPort    int
)

var addCmd = &cobra.Command{
	Use:   "add <alias> <user@host>",
	Short: "Save an SSH connection (TBD)",
	Long: `Save an SSH connection profile for easy reuse.

Stores the connection details under the specified alias for use with
other ssh commands like tunnel.

Examples:
  acu ssh add prod admin@example.com                    # Basic connection
  acu ssh add dev user@dev.local -p 2222               # Custom port
  acu ssh add staging deploy@staging.io -i ~/.ssh/key  # With key file
  acu ssh add web user@host -i ~/.ssh/id_rsa -p 22     # All options`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ssh add called")
	},
}

func init() {
	addCmd.Flags().StringVarP(&addKeyFile, "identity", "i", "", "Path to SSH private key file")
	addCmd.Flags().IntVarP(&addPort, "port", "p", 22, "SSH port number")
}

