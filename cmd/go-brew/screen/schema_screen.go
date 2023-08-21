package screen

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/robotjoosen/go-brew/pkg/brew"
	"github.com/robotjoosen/go-brew/pkg/widget"
)

type SchemaScreen struct {
	cnf brew.Brewable
}

func NewSchemaScreen(cnf brew.Brewable) ScreenAware {
	return &SchemaScreen{
		cnf: cnf,
	}
}

func (s *SchemaScreen) Restart() {}

func (s *SchemaScreen) Update(msg tea.Msg) tea.Cmd {
	return nil
}

func (s *SchemaScreen) Render() string {
	columnStyle := lipgloss.NewStyle().
		Margin(1, 0, 0, 0).
		Padding(0).
		Height(14)

	configTable, _ := widget.
		NewConfigTable(s.cnf).
		Render()
	schemaTable, _ := widget.
		NewSchemaTable(s.cnf.GenerateSchema()).
		Render()

	return lipgloss.JoinHorizontal(
		lipgloss.Left,
		columnStyle.
			Width(48).
			Render(configTable, schemaTable),
	)
}

func (s *SchemaScreen) DisableControls() bool {
	return false
}
