package linkerd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/nikhilsbhat/linkerd-checker/pkg/errors"
	"github.com/olekukonko/tablewriter"
	"github.com/sirupsen/logrus"
)

type Analyse struct {
	All      bool
	Category []string
	State    string
	File     string
	table    *tablewriter.Table
	logger   *logrus.Logger
}

func (analyse *Analyse) Analyse(cfg *CheckConfig) (bool, error) {
	if analyse.All {
		if !cfg.Success {
			analyse.logger.Error("not all linkerd checks have succeeded")

			for _, category := range cfg.Categories {
				for _, check := range category.Checks {
					if len(check.Error) != 0 {
						analyse.table.Append([]string{category.Name, trimSpace(check.Description), trimSpace(check.Error), check.Result})
					} else {
						analyse.table.Append([]string{category.Name, trimSpace(check.Description), "", check.Result})
					}
				}
			}

			return true, &errors.CheckerError{Message: "analysing linkerd checks failed"}
		}

		for _, category := range cfg.Categories {
			for _, check := range category.Checks {
				analyse.table.Append([]string{category.Name, trimSpace(check.Description), "", check.Result})
			}
		}

		return false, nil
	}

	var failed bool

	for _, category := range analyse.Category {
		for _, cat := range cfg.Categories {
			if cat.Name != category {
				continue
			}

			for _, check := range cat.Checks {
				if check.Result == "error" {
					failed = true
				}

				if len(check.Error) != 0 {
					analyse.table.Append([]string{cat.Name, trimSpace(check.Description), trimSpace(check.Error), check.Result})
				} else {
					analyse.table.Append([]string{cat.Name, trimSpace(check.Description), "", check.Result})
				}
			}
		}
	}

	if failed {
		return failed, &errors.CheckerError{Message: "analysing linkerd checks failed"}
	}

	return failed, nil
}

func (analyse *Analyse) SetTable() error {
	table := tablewriter.NewWriter(os.Stdout)

	if len(analyse.File) != 0 {
		analyse.logger.Debugf("to-file is enabled, the output would be rendered to file '%s'", analyse.File)

		absPath, err := filepath.Abs(analyse.File)
		if err != nil {
			analyse.logger.Errorf("fetching absolute filepath of '%s' errored with: %s", analyse.File, err.Error())

			return fmt.Errorf("%w", err)
		}

		fileWriter, err := os.Create(absPath)
		if err != nil {
			analyse.logger.Errorf("creating file '%s' errored with: %s", analyse.File, err.Error())

			return fmt.Errorf("%w", err)
		}

		table = tablewriter.NewWriter(fileWriter)
	}

	table.SetHeader([]string{"Category", "Description", "Error Message", "Result"})
	analyse.table = table

	return nil
}

func (analyse *Analyse) SetStatus(status bool) {
	analyse.State = "success"
	if status {
		analyse.State = "failed"
	}
}

func (analyse *Analyse) Render() {
	analyse.table.SetAlignment(tablewriter.ALIGN_CENTER) //nolint:nosnakecase
	analyse.table.SetAutoWrapText(true)
	analyse.table.SetAutoMergeCells(true)
	analyse.table.SetRowLine(true)

	analyse.table.SetFooter([]string{"", "", "State", analyse.State})
	analyse.table.Render()
}

func trimSpace(str string) string {
	str = strings.ReplaceAll(str, "\t", "")
	str = strings.TrimSpace(str)

	return str
}
