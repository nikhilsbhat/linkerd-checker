package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/nikhilsbhat/linkerd-checker/pkg/linkerd"
	"github.com/spf13/cobra"
)

func registerAnalyseCommand() *cobra.Command {
	analyseCommand := &cobra.Command{
		Use:     "analyse [flags]",
		Short:   "Command to analyse the linkerd check's outpur",
		Long:    `This will help user to analyse the output of linkerd check [https://linkerd.io/2.14/reference/cli/check/].`,
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

				analyseFailed, err := analyse.Analyse(config)
				if err != nil {
					linkerdCheckErrors = append(linkerdCheckErrors, err.Error())
				}

				if analyseFailed {
					failed = analyseFailed
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

			if failed {
				os.Exit(1)
			}

			return nil
		},
	}

	analyseCommand.SilenceUsage = true
	analyseCommand.SetUsageTemplate(getUsageTemplate())

	return analyseCommand
}
