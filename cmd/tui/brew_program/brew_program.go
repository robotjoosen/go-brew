package brew_program

import (
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/robotjoosen/go-brew/pkg/brew"
	"github.com/robotjoosen/go-brew/pkg/recipe"
	"github.com/robotjoosen/go-brew/pkg/recipe/tetsu"
	"github.com/robotjoosen/go-brew/pkg/widget"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type tickMsg time.Time
type fpsTickMsg time.Time

type BrewProgram struct {
	counter      int
	frameCounter int
	keymap       keymap
	help         help.Model
	layout       *Layout
	sprites      Sprites
	recipe       []brew.Pour
	table        widget.WidgetAware
	list         widget.WidgetAware
	graph        widget.WidgetAware
	clock        widget.WidgetAware
	textInput    textinput.Model
}

type keymap struct {
	quit key.Binding
}

func NewBrewProgram() BrewProgram {
	r := recipe.NewRecipeFactory().
		FourSixMethod().
		SetFlavor(tetsu.BalancedFlavor).
		SetConcentration(tetsu.StrongConcentration).
		SetCoffeeWeight(13).
		GenerateSchema()

	ti := textinput.New()
	ti.Placeholder = "donut"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

	return BrewProgram{
		layout: NewLayout().Set(LayoutSplashScreen),
		help:   help.New(),
		keymap: keymap{
			quit: key.NewBinding(
				key.WithKeys("ctrl+c", "q"),
				key.WithHelp("q", "quit"),
			),
		},
		table:     widget.NewTable(r),
		graph:     widget.NewGraph(r),
		clock:     widget.NewClock(r),
		list:      widget.NewList(r),
		textInput: ti,
		sprites:   Sprites{},
	}
}

func (m BrewProgram) Init() tea.Cmd {
	return tea.Batch(tick(), textinput.Blink, fpsTick(), tea.EnterAltScreen)
}

func tick() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func fpsTick() tea.Cmd {
	return tea.Tick(160*time.Millisecond, func(t time.Time) tea.Msg {
		return fpsTickMsg(t)
	})
}

func (m BrewProgram) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q": // These keys should exit the program.
			return m, tea.Quit
		}
	case fpsTickMsg:
		m.frameCounter++

		return m, fpsTick()
	case tickMsg:
		m.counter++

		return m, tick()
	}

	m.textInput, cmd = m.textInput.Update(msg)

	return m, cmd
}

func (m BrewProgram) View() (s string) {
	doc := strings.Builder{}

	//clockString, _ := m.clock.SetPosition(m.counter).Render()
	graphString, _ := m.graph.SetPosition(m.counter).Render()
	listString, _ := m.list.SetPosition(m.counter).Render()
	tableString, _ := m.table.Render()

	{ // header
		doc.WriteString(lipgloss.JoinHorizontal(
			lipgloss.Top,
			m.sprites.Logo(),
		))
	}

	{ // settings
		columnStyle := lipgloss.NewStyle().
			Margin(1, 0, 0, 0).
			Padding(0, 1).
			Height(12).
			Border(lipgloss.RoundedBorder())
		doc.WriteString(lipgloss.JoinHorizontal(
			lipgloss.Left,
			columnStyle.Width(46).Render(
				m.textInput.View(),
			),
		))
	}

	{ // animation test
		columnStyle := lipgloss.NewStyle().
			Margin(1, 0, 0, 0).
			Padding(0, 1).
			Height(12).
			Border(lipgloss.RoundedBorder())
		doc.WriteString(lipgloss.JoinHorizontal(
			lipgloss.Left,
			columnStyle.Width(46).Render(
				lipgloss.JoinHorizontal(
					lipgloss.Left,
					lipgloss.NewStyle().Width(16).Render(""),
					lipgloss.NewStyle().Width(25).Render("\n"+m.sprites.CoffeeAnimated(m.frameCounter)),
				),
			),
		))
	}

	{ // schedule
		columnStyle := lipgloss.NewStyle().
			Margin(1, 0, 0, 0).
			Padding(0).
			Height(12)
		doc.WriteString(lipgloss.JoinHorizontal(
			lipgloss.Left,
			columnStyle.Width(48).Render(tableString),
		))
	}

	{ // double column
		columnStyle := lipgloss.NewStyle().
			Margin(1, 0, 0, 0).
			Padding(0, 1).
			Height(12).
			Border(lipgloss.RoundedBorder())
		doc.WriteString(lipgloss.JoinHorizontal(
			lipgloss.Left,
			columnStyle.
				Width(16).
				Border(lipgloss.Border{Top: "─", Bottom: "─", Left: "│", Right: "│", TopLeft: "╭", TopRight: "┬", BottomLeft: "╰", BottomRight: "┴"}).
				Copy().
				Align(lipgloss.Left).
				Render(listString),
			columnStyle.
				Width(28).
				Padding(0).
				Border(lipgloss.Border{Top: "─", Bottom: "─", Left: " ", Right: "│", TopLeft: "─", TopRight: "╮", BottomLeft: "─", BottomRight: "╯"}).
				Copy().
				Align(lipgloss.Left).
				AlignVertical(lipgloss.Center).
				Render(graphString),
		))
	}

	{ // help
		doc.WriteString(lipgloss.JoinHorizontal(
			lipgloss.Left,
			lipgloss.NewStyle().
				Margin(1, 0, 0, 0).
				Padding(0, 1).
				Border(lipgloss.RoundedBorder()).
				Align(lipgloss.Right).
				Width(46).
				Render(m.help.ShortHelpView([]key.Binding{
					m.keymap.quit,
				})),
		))
	}

	return lipgloss.NewStyle().Padding(1, 2, 1, 2).Render(doc.String())
}
