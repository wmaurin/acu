package port

import (
	"github.com/spf13/cobra"
)

// PortCmd represents the port parent command
var PortCmd = &cobra.Command{
	Use:   "port",
	Short: "Port inspection and management utilities (TBD)",
	Long: `Port utilities for inspecting, scanning, and managing network ports.

Provides commands to find what process is using a port, kill processes
bound to specific ports, and scan hosts for open ports.`,
}

func init() {
	// Register subcommands
	PortCmd.AddCommand(findCmd)
	PortCmd.AddCommand(killCmd)
	PortCmd.AddCommand(scanCmd)
}

