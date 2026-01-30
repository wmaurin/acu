package disk

import (
	"fmt"

	"github.com/spf13/cobra"
)

var largestCmd = &cobra.Command{
	Use:   "largest [path]",
	Short: "Find largest files and directories (TBD)",
	Long: `Find and display the largest files and directories in the specified path.

Results are sorted by size in descending order with human-readable sizes.
Defaults to current directory if no path is specified.

Examples:
  acu disk largest             # Find largest files in current directory
  acu disk largest /home       # Find largest files in /home
  acu disk largest ./projects  # Find largest files in ./projects`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("disk largest called")
	},
}

