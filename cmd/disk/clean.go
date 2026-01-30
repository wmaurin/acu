package disk

import (
	"fmt"

	"github.com/spf13/cobra"
)

var cleanCmd = &cobra.Command{
	Use:   "clean [path]",
	Short: "Find and remove common junk files (TBD)",
	Long: `Find and remove common junk files and directories.

Removes development artifacts and temporary files including:
  - node_modules/
  - __pycache__/
  - .DS_Store
  - *.pyc
  - .git/ (optional)
  - build artifacts

Defaults to current directory if no path is specified.

Examples:
  acu disk clean             # Clean junk in current directory
  acu disk clean /projects   # Clean junk in /projects
  acu disk clean ./myapp     # Clean junk in ./myapp`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("disk clean called")
	},
}

