package disk

import (
	"github.com/spf13/cobra"
)

// DiskCmd represents the disk parent command
var DiskCmd = &cobra.Command{
	Use:   "disk",
	Short: "Disk and file management utilities (TBD)",
	Long: `Disk and file utilities for analyzing disk usage and cleaning up files.

Provides commands to find large files/directories and remove common
junk files like node_modules, __pycache__, .DS_Store, etc.`,
}

func init() {
	// Register subcommands
	DiskCmd.AddCommand(largestCmd)
	DiskCmd.AddCommand(cleanCmd)
}

