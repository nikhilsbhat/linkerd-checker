package linkerd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
	"github.com/nikhilsbhat/linkerd-checker/pkg/errors"
	"github.com/olekukonko/tablewriter"
	"github.com/sirupsen/logrus"
)

type Analyse struct {
	All        bool
	NoColor    bool
	Components []string
	State      string
	File       string
	table      *tablewriter.Table
	logger     *logrus.Logger
}

const (
	StateSuccess = "success"
	StateFailure = "failed"
	StateWarning = "warning"
	StateError   = "error"
)

func (analyse *Analyse) Analyse(cfg *CheckConfig) (bool, error) {
	if analyse.All {
		if !cfg.Success {
			analyse.logger.Error("not all linkerd checks have succeeded")

			for _, category := range cfg.Categories {
				for _, check := range category.Checks {
					if len(check.Error) != 0 {
						analyse.table.Append([]string{category.Name, trimSpace(check.Description), trimSpace(check.Error), analyse.colourCodeState(check.Result)})
					} else {
						analyse.table.Append([]string{category.Name, trimSpace(check.Description), "", analyse.colourCodeState(check.Result)})
					}
				}
			}

			return true, &errors.CheckerError{Message: "analysing linkerd checks failed"}
		}

		for _, category := range cfg.Categories {
			for _, check := range category.Checks {
				analyse.table.Append([]string{category.Name, trimSpace(check.Description), "", analyse.colourCodeState(check.Result)})
			}
		}

		return false, nil
	}

	var failed bool

	for _, category := range analyse.Components {
		for _, cat := range cfg.Categories {
			if cat.Name != category {
				continue
			}

			for _, check := range cat.Checks {
				if check.Result == "error" {
					failed = true
				}

				if len(check.Error) != 0 {
					analyse.table.Append([]string{cat.Name, trimSpace(check.Description), trimSpace(check.Error), analyse.colourCodeState(check.Result)})
				} else {
					analyse.table.Append([]string{cat.Name, trimSpace(check.Description), "", analyse.colourCodeState(check.Result)})
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

	analyse.table = table

	return nil
}

func (analyse *Analyse) SetStatus(status bool) {
	analyse.State = StateSuccess
	if status {
		analyse.State = StateFailure
	}
}

func (analyse *Analyse) Render() {
	analyse.table.SetHeader([]string{"Components", "Description", "Error Message", "Result"})

	if !analyse.NoColor {
		analyse.table.SetHeaderColor(tablewriter.Colors{tablewriter.Bold}, tablewriter.Colors{tablewriter.Bold},
			tablewriter.Colors{tablewriter.Bold}, tablewriter.Colors{tablewriter.Bold})
	}

	analyse.table.SetAlignment(tablewriter.ALIGN_CENTER) //nolint:nosnakecase
	analyse.table.SetAutoWrapText(true)
	analyse.table.SetAutoMergeCells(true)
	analyse.table.SetRowLine(true)

	analyse.table.SetFooter([]string{"", "", "State", analyse.State})

	if !analyse.NoColor {
		switch analyse.State {
		case StateFailure:
			analyse.table.SetFooterColor(tablewriter.Colors{}, tablewriter.Colors{},
				tablewriter.Colors{tablewriter.Bold},
				tablewriter.Colors{tablewriter.FgRedColor})
		default:
			analyse.table.SetFooterColor(tablewriter.Colors{}, tablewriter.Colors{},
				tablewriter.Colors{tablewriter.Bold},
				tablewriter.Colors{tablewriter.FgGreenColor})
		}
	}

	analyse.table.Render()
}

func trimSpace(str string) string {
	str = strings.ReplaceAll(str, "\t", "")
	str = strings.TrimSpace(str)

	return str
}

func (analyse *Analyse) colourCodeState(state string) string {
	if analyse.NoColor {
		return state
	}

	switch state {
	case StateSuccess:
		return color.GreenString(state)
	case StateFailure:
		return color.RedString(state)
	case StateWarning:
		return color.YellowString(state)
	case StateError:
		return color.RedString(state)
	default:
		return ""
	}
}
