package port

import (
	"fmt"

	"github.com/spf13/cobra"
)

var findCmd = &cobra.Command{
	Use:   "find <port>",
	Short: "Find what process is using a port (TBD)",
	Long: `Find the process that is currently using the specified port.

Displays the process ID (PID), process name, and user running the process.

Examples:
  acu port find 8080     # Find what's using port 8080
  acu port find 3000     # Find what's using port 3000`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("port find called")
	},
}

