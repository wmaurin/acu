package docker

import (
	"github.com/spf13/cobra"
)

// DockerCmd represents the docker parent command
var DockerCmd = &cobra.Command{
	Use:   "docker",
	Short: "Docker container and image management utilities (TBD)",
	Long: `Docker utilities for managing containers, images, networks, and system cleanup.

Provides commands to clean up Docker resources including stopped containers,
dangling images, unused networks, and build cache.`,
}

func init() {
	// Register subcommands
	DockerCmd.AddCommand(cleanCmd)
}

