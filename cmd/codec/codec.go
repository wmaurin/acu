package codec

import (
	"github.com/spf13/cobra"
)

// CodecCmd represents the codec parent command
var CodecCmd = &cobra.Command{
	Use:   "codec",
	Short: "Encoding, decoding, and hashing utilities (TBD)",
	Long: `Encoding and decoding utilities for various formats.

Provides commands to encode/decode strings in base64, URL, and HTML formats,
as well as generate file checksums (MD5, SHA1, SHA256).`,
}

func init() {
	// Register subcommands
	CodecCmd.AddCommand(encodeCmd)
	CodecCmd.AddCommand(decodeCmd)
	CodecCmd.AddCommand(hashCmd)
}

