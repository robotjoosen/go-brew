package widget

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/robotjoosen/go-brew/pkg/brew"
	"strconv"
	"time"
)

type SchemaTable struct {
	recipe []brew.Pour
}

func NewSchemaTable(recipe []brew.Pour) WidgetAware {
	return &SchemaTable{
		recipe: recipe,
	}
}

func (t *SchemaTable) SetPosition(int) WidgetAware {
	return t
}

func (t *SchemaTable) Render() (string, error) {
	var totalWeight int
	var totalDuration time.Duration

	tbl := table.NewWriter()
	tbl.SetStyle(table.StyleRounded)
	tbl.AppendHeader(table.Row{"start time", "   grams", "  to add", " wait for"})

	for _, pour := range t.recipe {
		totalWeight += pour.Grams

		tbl.AppendRow(table.Row{
			totalDuration,
			strconv.Itoa(totalWeight) + " g",
			strconv.Itoa(pour.Grams) + " g",
			pour.Duration,
		})
		totalDuration += pour.Duration
	}

	render := tbl.Render()

	return render, nil
}
