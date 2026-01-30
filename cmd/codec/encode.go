package codec

import (
	"fmt"

	"github.com/spf13/cobra"
)

var encodeCmd = &cobra.Command{
	Use:   "encode <format> <string>",
	Short: "Encode a string in the specified format (TBD)",
	Long: `Encode a string using the specified encoding format.

Supported formats:
  - base64: Base64 encoding
  - url: URL/percent encoding
  - html: HTML entity encoding

Examples:
  acu codec encode base64 "Hello World"       # SGVsbG8gV29ybGQ=
  acu codec encode url "Hello World"          # Hello%20World
  acu codec encode html "<script>alert()</script>"`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("codec encode called")
	},
}

