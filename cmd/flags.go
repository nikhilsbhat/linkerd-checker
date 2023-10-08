package cmd

import (
	"github.com/nikhilsbhat/linkerd-checker/pkg/linkerd"
	"github.com/spf13/cobra"
)

var (
	analyse linkerd.Analyse
	cliCfg  Config
)

func registerGlobalFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().StringVarP(&cliCfg.LogLevel, "log-level", "l", "info",
		"log level for gocd cli, log levels supported by [https://github.com/sirupsen/logrus] will work")
	cmd.PersistentFlags().BoolVarP(&analyse.NoColor, "no-color", "", false,
		"when enabled does not color encode the output")
	cmd.PersistentFlags().StringVarP(&analyse.File, "to-file", "f", "",
		"CSV file to which the output needs to be written")
	cmd.PersistentFlags().BoolVarP(&analyse.All, "all", "a", false,
		"enable this to Render output in JSON format")
	cmd.PersistentFlags().StringSliceVarP(&analyse.Components, "linkerd-component", "c", nil,
		"list of linkerd components to be considered for analysis")
	cmd.PersistentFlags().StringSliceVarP(&analyse.NotComponents, "not", "n", nil,
		"list of linkerd components to be ignored for the analysis, this has to be used along with flag --all")

	cmd.MarkFlagsMutuallyExclusive("all", "linkerd-component")
}
