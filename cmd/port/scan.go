package port

import (
	"fmt"

	"github.com/spf13/cobra"
)

var scanCmd = &cobra.Command{
	Use:   "scan <host>",
	Short: "Scan common ports on a host (TBD)",
	Long: `Scan a host for commonly used open ports.

Checks common service ports (HTTP, HTTPS, SSH, FTP, databases, etc.)
and reports which ones are open.

Examples:
  acu port scan localhost        # Scan localhost for open ports
  acu port scan 192.168.1.1      # Scan a specific IP
  acu port scan example.com      # Scan a hostname`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("port scan called")
	},
}

