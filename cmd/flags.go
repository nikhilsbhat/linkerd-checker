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
	cmd.PersistentFlags().BoolVarP(&analyse.All, "all", "", false,
		"enable this to Render output in JSON format")
	cmd.PersistentFlags().StringSliceVarP(&analyse.Category, "category", "", nil,
		"categories from the output to be analysed")
}