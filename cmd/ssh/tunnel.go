package ssh

import (
	"fmt"

	"github.com/spf13/cobra"
)

var tunnelLocalForward string

var tunnelCmd = &cobra.Command{
	Use:   "tunnel <alias>",
	Short: "Create an SSH tunnel with a saved connection (TBD)",
	Long: `Create an SSH tunnel using a saved connection profile.

Sets up local port forwarding to access remote services through the tunnel.

Examples:
  acu ssh tunnel prod -L 8080:localhost:80      # Forward local 8080 to remote 80
  acu ssh tunnel dev -L 5432:localhost:5432     # Forward PostgreSQL
  acu ssh tunnel staging -L 3306:db.local:3306  # Forward to remote host`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ssh tunnel called")
	},
}

func init() {
	tunnelCmd.Flags().StringVarP(&tunnelLocalForward, "local", "L", "", "Local port forwarding (local:remote or local:host:remote)")
	tunnelCmd.MarkFlagRequired("local")
}

