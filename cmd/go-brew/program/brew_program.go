package program

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/robotjoosen/go-brew/cmd/go-brew/screen"
	"github.com/robotjoosen/go-brew/pkg/brew"
	"github.com/robotjoosen/go-brew/pkg/recipe"
	"github.com/robotjoosen/go-brew/pkg/recipe/tetsu"
	"github.com/robotjoosen/go-brew/pkg/sprite"
	"strings"
	"time"
)

type (
	tickMsg time.Time

	BrewProgram struct {
		screenFactory *screen.ScreenFactory
		screen        screen.ScreenAware
		keymap        keymap
		help          help.Model
	}

	Configuration struct {
	}

	keymap struct {
		helper key.Binding
		schema key.Binding
		config key.Binding
		quit   key.Binding
	}
)

func NewBrewProgram() BrewProgram {
	defaultRecipe := recipe.NewRecipeFactory().
		FourSixMethod().
		SetFlavor(tetsu.BalancedFlavor).
		SetConcentration(tetsu.StrongConcentration).
		SetCoffeeWeight(16)

	screenFactory := screen.NewFactory().
		UpdateRecipeFactory(defaultRecipe)

	p := BrewProgram{
		help:          help.New(),
		screenFactory: screenFactory,
		screen:        screenFactory.New(screen.Splash),
		keymap: keymap{
			helper: key.NewBinding(
				key.WithKeys("s"),
				key.WithHelp("s", "schema"),
			),
			schema: key.NewBinding(
				key.WithKeys("c"),
				key.WithHelp("c", "config"),
			),
			config: key.NewBinding(
				key.WithKeys("h"),
				key.WithHelp("h", "helper"),
			),
			quit: key.NewBinding(
				key.WithKeys("ctrl+c", "esc"),
				key.WithHelp("esc", "quit"),
			),
		},
	}

	return p
}

func (m BrewProgram) Init() tea.Cmd {
	return tea.Batch(tick(), textinput.Blink, tea.EnterAltScreen)
}

func (m BrewProgram) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc": // These keys should exit the program.
			return m, tea.Quit
		}

		if !m.screen.DisableControls() {
			switch msg.String() {
			case "s":
				m.screen = m.screenFactory.New(screen.Schema)
				m.screen.Restart()
			case "c":
				m.screen = m.screenFactory.New(screen.Config)
				m.screen.Restart()
			case "h":
				m.screen = m.screenFactory.New(screen.Helper)
				m.screen.Restart()
			}
		}
	case tickMsg:
		return m, tick()
	case brew.Recipe:
		// todo:
		// 		- do a replacement, which should only be a partial update of the recipe.
		//		- fix water vs ratio calculation with configuration

		m.screenFactory.UpdateRecipeFactory(
			recipe.NewRecipeFactory().
				FourSixMethod().
				SetFlavor(tetsu.BalancedFlavor).
				SetConcentration(tetsu.StrongConcentration).
				SetCoffeeWeight(msg.Coffee),
		)

		// and switch to schema
		m.screen = m.screenFactory.New(screen.Schema)
		m.screen.Restart()
	}

	cmd := m.screen.Update(msg)

	return m, cmd
}

func tick() tea.Cmd {
	// 12 fps
	return tea.Tick(84*time.Millisecond, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func (m BrewProgram) View() (s string) {
	doc := strings.Builder{}

	{ // header logo
		doc.WriteString(lipgloss.JoinHorizontal(lipgloss.Top, (sprite.Sprites{}).Logo()))
	}

	{ // screen
		doc.WriteString(m.screen.Render())
	}

	{ // help
		doc.WriteString(lipgloss.JoinHorizontal(
			lipgloss.Left,
			lipgloss.NewStyle().
				Margin(1, 0, 0, 0).
				Padding(0, 1).
				Border(lipgloss.RoundedBorder()).
				Align(lipgloss.Left).
				Width(46).
				Render(m.help.ShortHelpView([]key.Binding{
					m.keymap.schema,
					m.keymap.helper,
					m.keymap.config,
					m.keymap.quit,
				})),
		))
	}

	return lipgloss.
		NewStyle().
		Padding(1, 2, 1, 2).
		Render(doc.String())
}
