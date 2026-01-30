package gen

import (
	"fmt"

	"github.com/spf13/cobra"
)

var secretCmd = &cobra.Command{
	Use:   "secret [length]",
	Short: "Generate a base64-encoded secret (TBD)",
	Long: `Generate a cryptographically secure base64-encoded secret.

Suitable for use as API keys, JWT secrets, encryption keys,
and other configuration values. Default length is 32 bytes if not specified.

Examples:
  acu gen secret        # Generate 32-byte secret (44 chars base64)
  acu gen secret 64     # Generate 64-byte secret (88 chars base64)
  acu gen secret 16     # Generate 16-byte secret (24 chars base64)`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("gen secret called")
	},
}

