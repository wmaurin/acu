package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var askCmd = &cobra.Command{
	Use:   "ask <question>",
	Short: "Ask a question to the terminal agent (TBD)",
	Long: `Ask a question to the AI-powered terminal agent.

The agent can answer questions about commands, help with tasks,
and execute commands on your behalf after approval.

Examples:
  acu ask "How do I find large files?"
  acu ask "What's using port 8080?"
  acu ask "Help me compress this folder"`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ask called")
	},
}

func init() {
	rootCmd.AddCommand(askCmd)
}

