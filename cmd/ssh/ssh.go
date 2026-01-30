package ssh

import (
	"github.com/spf13/cobra"
)

// SSHCmd represents the ssh parent command
var SSHCmd = &cobra.Command{
	Use:   "ssh",
	Short: "SSH connection management utilities (TBD)",
	Long: `SSH utilities for managing saved connections and tunnels.

Provides commands to list, add, and remove saved SSH connections,
as well as create SSH tunnels using saved connection profiles.`,
}

func init() {
	// Register subcommands
	SSHCmd.AddCommand(listCmd)
	SSHCmd.AddCommand(addCmd)
	SSHCmd.AddCommand(rmCmd)
	SSHCmd.AddCommand(tunnelCmd)
}

