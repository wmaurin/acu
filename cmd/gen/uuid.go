package gen

import (
	"fmt"

	"github.com/spf13/cobra"
)

var uuidCmd = &cobra.Command{
	Use:   "uuid",
	Short: "Generate a UUID v4 (TBD)",
	Long: `Generate a random UUID version 4.

Produces a universally unique identifier in the standard format:
xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx

Examples:
  acu gen uuid    # Generate a new UUID v4`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("gen uuid called")
	},
}

