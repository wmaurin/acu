package codec

import (
	"fmt"

	"github.com/spf13/cobra"
)

var decodeCmd = &cobra.Command{
	Use:   "decode <format> <string>",
	Short: "Decode a string from the specified format (TBD)",
	Long: `Decode a string from the specified encoding format.

Supported formats:
  - base64: Base64 decoding
  - url: URL/percent decoding
  - html: HTML entity decoding

Examples:
  acu codec decode base64 "SGVsbG8gV29ybGQ="   # Hello World
  acu codec decode url "Hello%20World"         # Hello World
  acu codec decode html "&lt;script&gt;"       # <script>`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("codec decode called")
	},
}

