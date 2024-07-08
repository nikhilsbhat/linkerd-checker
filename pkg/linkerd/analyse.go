package linkerd

import (
	"strings"

	"github.com/fatih/color"
	"github.com/nikhilsbhat/linkerd-checker/pkg/errors"
	"github.com/olekukonko/tablewriter"
	"github.com/sirupsen/logrus"
	"github.com/thoas/go-funk"
)

type Analyse struct {
	All           bool
	NoColor       bool
	Components    []string
	NotComponents []string
	State         string
	File          string
	table         *tablewriter.Table
	logger        *logrus.Logger
}

const (
	StateSuccess = "success"
	StateFailure = "failed"
	StateWarning = "warning"
	StateError   = "error"
)

type CheckIterator []Check

func (analyse *Analyse) Analyse(cfg *CheckConfig) (bool, error) {
	var failed bool

	if analyse.All {
		for _, category := range cfg.Categories {
			if funk.Contains(analyse.NotComponents, category.Name) {
				analyse.logger.Debugf("ignoring linkerd component '%s' from checks since it is on ignore list", category.Name)

				continue
			}

			failed = NewIterator(category.Checks).Iterator(category.Name, analyse)
		}

		return false, nil
	} else {
		for _, category := range analyse.Components {
			for _, cat := range cfg.Categories {
				if cat.Name != category {
					continue
				}

				failed = NewIterator(cat.Checks).Iterator(cat.Name, analyse)
			}
		}
	}

	if failed {
		analyse.logger.Error("not all linkerd checks have succeeded")

		return failed, &errors.CheckerError{Message: "analysing linkerd checks failed"}
	}

	return false, nil
}

func (checks CheckIterator) Iterator(category string, analyse *Analyse) bool {
	var failed bool

	for _, check := range checks {
		if check.Result == StateError {
			failed = true
		}

		description := trimSpace(check.Description)
		errorMsg := trimSpace(check.Error)
		coloredState := analyse.colourCodeState(check.Result)

		analyse.table.Append([]string{category, description, errorMsg, coloredState})
	}

	return failed
}

func NewIterator(check []Check) CheckIterator {
	return check
}

func (analyse *Analyse) SetStatus(status bool) {
	if status {
		analyse.State = StateFailure
	} else {
		analyse.State = StateSuccess
	}
}

func trimSpace(str string) string {
	return strings.TrimSpace(strings.ReplaceAll(str, "\t", ""))
}

func (analyse *Analyse) colourCodeState(state string) string {
	if analyse.NoColor {
		return state
	}

	switch state {
	case StateSuccess:
		return color.GreenString(state)
	case StateFailure, StateError:
		return color.RedString(state)
	case StateWarning:
		return color.YellowString(state)
	default:
		return state
	}
}
