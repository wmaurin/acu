package net

import (
	"fmt"

	"github.com/spf13/cobra"
)

var speedCmd = &cobra.Command{
	Use:   "speed",
	Short: "Run a network speed test (TBD)",
	Long: `Run a network speed test to measure download and upload speeds.

Tests your internet connection speed and displays results
in human-readable format (Mbps).

Examples:
  acu net speed    # Run a speed test`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("net speed called")
	},
}

