package port

import (
	"fmt"

	"github.com/spf13/cobra"
)

var killCmd = &cobra.Command{
	Use:   "kill <port>",
	Short: "Kill the process using a port (TBD)",
	Long: `Kill the process that is currently using the specified port.

Finds the process bound to the port and terminates it.

Examples:
  acu port kill 8080     # Kill whatever is using port 8080
  acu port kill 3000     # Kill whatever is using port 3000`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("port kill called")
	},
}

