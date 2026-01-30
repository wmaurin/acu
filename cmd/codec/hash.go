package codec

import (
	"fmt"

	"github.com/spf13/cobra"
)

var hashCmd = &cobra.Command{
	Use:   "hash <file>",
	Short: "Generate checksums for a file (TBD)",
	Long: `Generate MD5, SHA1, and SHA256 checksums for a file.

Useful for verifying file integrity and comparing files.

Examples:
  acu codec hash myfile.txt          # Hash a text file
  acu codec hash download.iso        # Hash a downloaded file
  acu codec hash ~/Documents/doc.pdf # Hash with full path`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("codec hash called")
	},
}

