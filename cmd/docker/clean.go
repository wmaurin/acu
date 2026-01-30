package docker

import (
	"fmt"

	"github.com/spf13/cobra"
)

var cleanAll bool

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Remove stopped containers, dangling images, and unused networks (TBD)",
	Long: `Clean up Docker resources by removing stopped containers, dangling images,
and unused networks.

By default, removes only stopped/unused resources. Use --all to perform
a complete cleanup including running containers.

Examples:
  acu docker clean           # Clean stopped containers, dangling images, unused networks
  acu docker clean --all     # Stop and remove ALL containers, images, networks, and cache`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("docker clean called")
	},
}

func init() {
	cleanCmd.Flags().BoolVarP(&cleanAll, "all", "a", false, "Stop and remove all containers, images, networks, and cache")
}

