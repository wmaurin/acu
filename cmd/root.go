package cmd

import (
	"os"

	"acu/cmd/codec"
	"acu/cmd/disk"
	"acu/cmd/docker"
	"acu/cmd/gen"
	"acu/cmd/net"
	"acu/cmd/port"
	"acu/cmd/ssh"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// Color definitions
var (
	blue   = color.New(color.FgBlue, color.Bold).SprintFunc()
	cyan   = color.New(color.FgCyan).SprintFunc()
	yellow = color.New(color.FgYellow).SprintFunc()
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "acu",
	Short: "A cross-platform CLI toolkit for common terminal tasks",
	Long: `Acu is a cross-platform CLI toolkit for common terminal tasks.

Use "acu [command] --help" for more information about a command.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Register all parent commands
	rootCmd.AddCommand(docker.DockerCmd)
	rootCmd.AddCommand(port.PortCmd)
	rootCmd.AddCommand(disk.DiskCmd)
	rootCmd.AddCommand(net.NetCmd)
	rootCmd.AddCommand(ssh.SSHCmd)
	rootCmd.AddCommand(gen.GenCmd)
	rootCmd.AddCommand(codec.CodecCmd)

	// Set custom usage template with blue command references
	cobra.AddTemplateFunc("blue", blue)
	cobra.AddTemplateFunc("cyan", cyan)
	cobra.AddTemplateFunc("yellow", yellow)

	usageTemplate := `{{yellow "Usage:"}}{{if .Runnable}}
  {{blue .UseLine}}{{end}}{{if .HasAvailableSubCommands}}
  {{blue .CommandPath}} [command]{{end}}{{if gt (len .Aliases) 0}}

{{yellow "Aliases:"}}
  {{.NameAndAliases}}{{end}}{{if .HasExample}}

{{yellow "Examples:"}}
{{.Example}}{{end}}{{if .HasAvailableSubCommands}}{{$cmds := .Commands}}{{if eq (len .Groups) 0}}

{{yellow "Available Commands:"}}{{range $cmds}}{{if (or .IsAvailableCommand (eq .Name "help"))}}
  {{blue (rpad .Name .NamePadding)}} {{.Short}}{{end}}{{end}}{{else}}{{range $group := .Groups}}

{{.Title}}{{range $cmds}}{{if (and (eq .GroupID $group.ID) (or .IsAvailableCommand (eq .Name "help")))}}
  {{blue (rpad .Name .NamePadding)}} {{.Short}}{{end}}{{end}}{{end}}{{if not .AllChildCommandsHaveGroup}}

{{yellow "Additional Commands:"}}{{range $cmds}}{{if (and (eq .GroupID "") (or .IsAvailableCommand (eq .Name "help")))}}
  {{blue (rpad .Name .NamePadding)}} {{.Short}}{{end}}{{end}}{{end}}{{end}}{{end}}{{if .HasAvailableLocalFlags}}

{{yellow "Flags:"}}
{{.LocalFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasAvailableInheritedFlags}}

{{yellow "Global Flags:"}}
{{.InheritedFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasHelpSubCommands}}

{{yellow "Additional help topics:"}}{{range .Commands}}{{if .IsAdditionalHelpTopicCommand}}
  {{blue (rpad .CommandPath .CommandPathPadding)}} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableSubCommands}}

Use "{{cyan .CommandPath}} {{cyan "[command] --help"}}" for more information about a command.{{end}}
`
	rootCmd.SetUsageTemplate(usageTemplate)

	helpTemplate := `{{with (or .Long .Short)}}{{. | trimTrailingWhitespaces}}

{{end}}{{if or .Runnable .HasSubCommands}}{{.UsageString}}{{end}}`
	rootCmd.SetHelpTemplate(helpTemplate)
}

