package linkerd

import (
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
