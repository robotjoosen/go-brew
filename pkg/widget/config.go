package widget

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/robotjoosen/go-brew/pkg/brew"
	"strconv"
)

type ConfigTable struct {
	cnf brew.Brewable
}

func NewConfigTable(cnf brew.Brewable) WidgetAware {
	return &ConfigTable{
		cnf: cnf,
	}
}

func (t *ConfigTable) SetPosition(int) WidgetAware {
	return t
}

func (t *ConfigTable) Render() (string, error) {
	r := t.cnf.GetRecipe()

	tbl := table.NewWriter()
	tbl.SetStyle(table.StyleRounded)
	tbl.AppendHeader(table.Row{"recipe    ", "        "})
	tbl.AppendRow(table.Row{"ratio", r.Ratio})
	tbl.AppendRow(table.Row{"coffee", strconv.Itoa(r.Coffee) + "g     "})
	tbl.AppendRow(table.Row{"water", strconv.Itoa(r.Water) + "g"})

	render := tbl.Render()

	return render, nil
}
