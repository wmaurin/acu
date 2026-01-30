package net

import (
	"github.com/spf13/cobra"
)

// NetCmd represents the net parent command
var NetCmd = &cobra.Command{
	Use:   "net",
	Short: "Network diagnostic and testing utilities (TBD)",
	Long: `Network utilities for testing connectivity and measuring performance.

Provides commands for network speed testing and HTTP request analysis
with detailed timing breakdowns.`,
}

func init() {
	// Register subcommands
	NetCmd.AddCommand(speedCmd)
	NetCmd.AddCommand(httpCmd)
}

