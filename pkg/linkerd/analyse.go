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

		if failed {
			analyse.logger.Error("not all linkerd checks have succeeded")

			return failed, &errors.CheckerError{Message: "analysing linkerd checks failed"}
		}

		return false, nil
	}

	for _, category := range analyse.Components {
		for _, cat := range cfg.Categories {
			if cat.Name != category {
				continue
			}

			failed = NewIterator(cat.Checks).Iterator(cat.Name, analyse)
		}
	}

	if failed {
		return failed, &errors.CheckerError{Message: "analysing linkerd checks failed"}
	}

	return failed, nil
}

func (checks CheckIterator) Iterator(category string, analyse *Analyse) bool {
	var failed bool

	for _, check := range checks {
		if check.Result == "error" {
			failed = true
		}

		if len(check.Error) != 0 {
			analyse.table.Append([]string{category, trimSpace(check.Description), trimSpace(check.Error), analyse.colourCodeState(check.Result)})

			continue
		}

		analyse.table.Append([]string{category, trimSpace(check.Description), "", analyse.colourCodeState(check.Result)})
	}

	return failed
}

func NewIterator(check []Check) CheckIterator {
	return check
}

func (analyse *Analyse) SetStatus(status bool) {
	analyse.State = StateSuccess
	if status {
		analyse.State = StateFailure
	}
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
