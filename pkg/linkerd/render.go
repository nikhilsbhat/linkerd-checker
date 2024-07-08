package linkerd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/olekukonko/tablewriter"
)

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
