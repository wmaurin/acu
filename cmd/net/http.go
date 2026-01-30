package net

import (
	"fmt"

	"github.com/spf13/cobra"
)

var httpCmd = &cobra.Command{
	Use:   "http <url>",
	Short: "HTTP request with timing breakdown (TBD)",
	Long: `Make an HTTP request and display detailed timing information.

Shows timing breakdown including:
  - DNS lookup time
  - TCP connection time
  - TLS handshake time
  - Time to first byte (TTFB)
  - Total request time

Examples:
  acu net http https://example.com       # Test HTTPS timing
  acu net http http://localhost:8080     # Test local server timing`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("net http called")
	},
}

