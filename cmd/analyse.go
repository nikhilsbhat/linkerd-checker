package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/nikhilsbhat/linkerd-checker/pkg/linkerd"
	"github.com/spf13/cobra"
)

func registerAnalyseCommand() *cobra.Command {
	analyseCommand := &cobra.Command{
		Use:     "analyse [flags]",
		Short:   "Command to fetch the version of linkerd-checker installed",
		Long:    `This will help user to find what version of linkerd-checker they are using in their machine.`,
		PreRunE: setCLIClient,
		RunE: func(cmd *cobra.Command, args []string) error {
			cliLogger.Debug("reading linkerd check's output from stdin")
			stdIn := cmd.InOrStdin()

			dec := json.NewDecoder(stdIn)
			var failed bool

			if err := analyse.SetTable(); err != nil {
				return fmt.Errorf("%w", err)
			}

			linkerdCheckErrors := make([]string, 0)

			for {
				config, err := linkerd.GetCheckConfig(dec)
				if errors.Is(err, io.EOF) {
					break
				}
				if err != nil {
					cliLogger.Fatalf("parsing the linkerd check's output errored with: '%s'", err.Error())
				}

				state, err := analyse.Analyse(config)
				if err != nil {
					linkerdCheckErrors = append(linkerdCheckErrors, err.Error())
				}

				if !state {
					failed = true
				}
			}

			analyse.SetStatus(failed)

			if !failed {
				cliLogger.Info("all linkerd checks have succeeded")
			}

			if len(linkerdCheckErrors) != 0 {
				cliLogger.Errorf("linkerd checks failed with: %s", strings.Join(linkerdCheckErrors, "\n"))
			}

			analyse.Render()

			return nil
		},
	}

	analyseCommand.SilenceUsage = true
	analyseCommand.SetUsageTemplate(getUsageTemplate())

	return analyseCommand
}
