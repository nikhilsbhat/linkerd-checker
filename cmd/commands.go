package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/nikhilsbhat/linkerd-checker/version"
	"github.com/spf13/cobra"
)

type cliCommands struct {
	commands []*cobra.Command
}

// Config holds the information of the cli config.
type Config struct {
	NoColor  bool   `yaml:"-"`
	LogLevel string `yaml:"-"`
}

func SetLinkerdCheckerCommands() *cobra.Command {
	return getLinkerdCheckerCommands()
}

// Add an entry in below function to register new command.
func getLinkerdCheckerCommands() *cobra.Command {
	command := new(cliCommands)
	command.commands = append(command.commands, registerVersionCommand())
	command.commands = append(command.commands, registerAnalyseCommand())

	return command.prepareCommands()
}

func (c *cliCommands) prepareCommands() *cobra.Command {
	rootCmd := getRootCommand()
	for _, cmnd := range c.commands {
		rootCmd.AddCommand(cmnd)
	}

	rootCmd.SilenceErrors = true
	registerGlobalFlags(rootCmd)

	return rootCmd
}

func getRootCommand() *cobra.Command {
	rootCommand := &cobra.Command{
		Use:     "linkerd-checker",
		Short:   "Command line interface to analyse the output of linkerd checks",
		Long:    `Command line interface that takes the linkerd check json output and analyse them`,
		PreRunE: setCLIClient,
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Usage() //nolint:wrapcheck
		},
	}
	rootCommand.SetUsageTemplate(getUsageTemplate())

	return rootCommand
}

func registerVersionCommand() *cobra.Command {
	versionCommand := &cobra.Command{
		Use:     "version [flags]",
		Short:   "Command to fetch the version of linkerd-checker installed",
		Long:    `This will help user to find what version of linkerd-checker he/she installed in her machine.`,
		PreRunE: setCLIClient,
		RunE:    AppVersion,
	}
	versionCommand.SetUsageTemplate(getUsageTemplate())

	return versionCommand
}

func AppVersion(_ *cobra.Command, _ []string) error {
	buildInfo, err := json.Marshal(version.GetBuildInfo())
	if err != nil {
		cliLogger.Errorf("fetching version information of linkerd-checker errored with: %s", err.Error())
		os.Exit(1)
	}
	fmt.Printf("linkerd-checker version: %s\n", string(buildInfo))

	return nil
}

func getUsageTemplate() string {
	return `Usage:{{if .Runnable}}
  {{.UseLine}}{{end}}{{if gt (len .Aliases) 0}}{{printf "\n" }}
Aliases:
  {{.NameAndAliases}}{{end}}{{if .HasExample}}{{printf "\n" }}
Examples:
{{.Example}}{{end}}{{if .HasAvailableSubCommands}}{{printf "\n"}}
Available Commands:{{range .Commands}}{{if (or .IsAvailableCommand (eq .Name "help"))}}
  {{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableLocalFlags}}{{printf "\n"}}
Flags:
{{.LocalFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasAvailableInheritedFlags}}{{printf "\n"}}
Global Flags:
{{.InheritedFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasHelpSubCommands}}{{printf "\n"}}
Additional help topics:{{range .Commands}}{{if .IsAdditionalHelpTopicCommand}}
  {{rpad .CommandPath .CommandPathPadding}} {{.Short}}{{end}}{{end}}{{end}}
{{if .HasAvailableSubCommands}}{{printf "\n"}}
Use "{{.CommandPath}} [command] --help" for more information about a command.{{end}}
{{printf "\n"}}`
}
