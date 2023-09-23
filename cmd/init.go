package cmd

import "github.com/spf13/cobra"

func setCLIClient(_ *cobra.Command, _ []string) error {
	SetLogger(cliCfg.LogLevel)
	analyse.SetLogger(cliLogger)

	return nil
}
