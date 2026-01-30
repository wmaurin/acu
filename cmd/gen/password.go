package gen

import (
	"fmt"

	"github.com/spf13/cobra"
)

var passwordCmd = &cobra.Command{
	Use:   "password [length]",
	Short: "Generate a secure password (TBD)",
	Long: `Generate a cryptographically secure random password.

Uses a mix of uppercase, lowercase, numbers, and special characters.
Default length is 16 characters if not specified.

Examples:
  acu gen password        # Generate 16-character password
  acu gen password 32     # Generate 32-character password
  acu gen password 8      # Generate 8-character password`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("gen password called")
	},
}

