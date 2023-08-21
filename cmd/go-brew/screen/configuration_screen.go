package screen

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/robotjoosen/go-brew/pkg/brew"
	"strconv"
	"strings"
)

var (
	focusedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	blurredStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	cursorStyle  = focusedStyle.Copy()
	noStyle      = lipgloss.NewStyle()
)

const (
	gramsPlaceholder = "grams"
	waterPlaceholder = "water"
)

type ConfigurationScreen struct {
	cnf        brew.Brewable
	inputs     []textinput.Model
	focusIndex int
}

func NewConfigurationScreen(cnf brew.Brewable) ScreenAware {
	c := &ConfigurationScreen{
		cnf:        cnf,
		inputs:     make([]textinput.Model, 2),
		focusIndex: 2,
	}

	recipe := cnf.GetRecipe()

	var t textinput.Model
	for i := range c.inputs {
		t = textinput.New()
		t.Cursor.Style = cursorStyle
		t.Prompt = ": "
		t.CharLimit = 32

		switch i {
		case 0:
			t.Placeholder = gramsPlaceholder
			t.SetValue(strconv.Itoa(recipe.Coffee))
		case 1:
			t.Placeholder = waterPlaceholder
			t.SetValue(strconv.Itoa(recipe.Water))
		}

		c.inputs[i] = t
	}

	return c
}

func (s *ConfigurationScreen) Restart() {}

func (s *ConfigurationScreen) Update(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(s.inputs))

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "tab", "enter", "up", "down":
			tempCmds, handled := s.saveInputValues(msg, cmds)
			if !handled {
				tempCmds = s.navigateInputs(msg, cmds)
			}

			cmds = tempCmds
		}
	}

	// Only text inputs with Focus() set will respond, so it's safe to simply
	// update all of them here without any further logic.
	for i := range s.inputs {
		s.inputs[i], cmds[i] = s.inputs[i].Update(msg)
	}

	return tea.Batch(cmds...)
}

func (s *ConfigurationScreen) Render() string {
	var b strings.Builder

	for i := range s.inputs {
		b.WriteString(s.inputs[i].Placeholder + " ")
		b.WriteString(s.inputs[i].View())
		if i < len(s.inputs)-1 {
			b.WriteRune('\n')
		}
	}

	var submitButton string
	if s.focusIndex == len(s.inputs) {
		submitButton = focusedStyle.Copy().Render("[ Save ]")
	} else {
		submitButton = blurredStyle.Copy().Render("[ Save ]")
	}

	return lipgloss.JoinHorizontal(
		lipgloss.Left,
		lipgloss.NewStyle().
			Margin(1, 0, 0, 0).
			Padding(0, 1).
			Height(12).
			Border(lipgloss.RoundedBorder()).
			Width(46).
			Render(
				b.String(),
				"\n\n",
				submitButton,
			),
	)
}

func (s *ConfigurationScreen) DisableControls() bool {
	return len(s.inputs) > s.focusIndex
}

func (s *ConfigurationScreen) saveInputValues(msg tea.KeyMsg, cmds []tea.Cmd) ([]tea.Cmd, bool) {
	if s.focusIndex == len(s.inputs) && msg.String() == "enter" {
		recipeMsg := brew.Recipe{}
		for _, input := range s.inputs {
			switch input.Placeholder {
			case gramsPlaceholder:
				v, _ := strconv.Atoi(input.Value())
				recipeMsg.Coffee = v
			case waterPlaceholder:
				v, _ := strconv.Atoi(input.Value())
				recipeMsg.Water = v
			}
		}

		return append(cmds, func() tea.Msg { return recipeMsg }), true
	}

	return cmds, false
}

func (s *ConfigurationScreen) navigateInputs(msg tea.KeyMsg, cmds []tea.Cmd) []tea.Cmd {
	if msg.String() == "up" {
		s.focusIndex--
	} else {
		s.focusIndex++
	}

	if s.focusIndex > len(s.inputs) {
		s.focusIndex = 0
	} else if s.focusIndex < 0 {
		s.focusIndex = len(s.inputs)
	}

	for i := 0; i <= len(s.inputs)-1; i++ {
		if i == s.focusIndex {
			cmds = append(cmds, s.inputs[i].Focus())
			s.inputs[i].PromptStyle = focusedStyle
			s.inputs[i].TextStyle = focusedStyle

			continue
		}

		s.inputs[i].Blur()
		s.inputs[i].PromptStyle = noStyle
		s.inputs[i].TextStyle = noStyle
	}

	return cmds
}
