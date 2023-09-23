package linkerd

import (
	"github.com/cheynewallace/tabby"
	"github.com/nikhilsbhat/linkerd-checker/pkg/errors"
	"github.com/sirupsen/logrus"
)

type Analyse struct {
	All      bool
	Category []string
	table    *tabby.Tabby
	logger   *logrus.Logger
}

func (analyse *Analyse) Analyse(cfg *CheckConfig) (bool, error) {
	if analyse.All {
		if !cfg.Success {
			analyse.logger.Error("following linkerd checks have failed")

			for _, category := range cfg.Categories {
				for _, check := range category.Checks {
					if len(check.Error) != 0 {
						analyse.table.AddLine(category.Name, check.Description, check.Error, check.Result)
					} else {
						analyse.table.AddLine(category.Name, check.Description, "", check.Result)
					}
				}
			}

			return false, &errors.CheckerError{Message: "analysing linkerd checks failed"}
		}

		return true, nil
	}

	var failed bool

	for _, category := range analyse.Category {
		for _, cat := range cfg.Categories {
			if cat.Name != category {
				continue
			}

			if !cfg.Success {
				failed = true
			}

			for _, check := range cat.Checks {
				if len(check.Error) != 0 {
					analyse.table.AddLine(cat.Name, check.Description, check.Error, check.Result)
				} else {
					analyse.table.AddLine(cat.Name, check.Description, "", check.Result)
				}
			}
		}
	}

	if failed {
		return failed, &errors.CheckerError{Message: "analysing linkerd checks failed"}
	}

	return failed, nil
}

func (analyse *Analyse) SetTable(table *tabby.Tabby) {
	analyse.table = table
}
