package gen

import (
	"github.com/spf13/cobra"
)

// GenCmd represents the gen parent command
var GenCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generator utilities for passwords, UUIDs, and secrets (TBD)",
	Long: `Generator utilities for creating secure random values.

Provides commands to generate passwords, UUIDs, and base64-encoded
secrets suitable for configuration files and API keys.`,
}

func init() {
	// Register subcommands
	GenCmd.AddCommand(passwordCmd)
	GenCmd.AddCommand(uuidCmd)
	GenCmd.AddCommand(secretCmd)
}

