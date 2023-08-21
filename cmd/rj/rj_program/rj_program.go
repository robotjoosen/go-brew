package rj_program

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/76creates/stickers"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/robotjoosen/go-brew/pkg/domain"
	"golang.org/x/term"
)

type tickMsg time.Time

var (
	statusBarStyle = lipgloss.NewStyle().
			Foreground(lipgloss.AdaptiveColor{Light: "#343433", Dark: "#C1C6B2"}).
			Background(lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#353533"})

	statusStyle = lipgloss.NewStyle().
			Inherit(statusBarStyle).
			Foreground(lipgloss.Color("#FFFDF5")).
			Background(lipgloss.Color("#FF5F87")).
			Padding(0, 1).
			MarginRight(1)

	statusText = lipgloss.NewStyle().Inherit(statusBarStyle)

	descStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#353533")).
			MarginTop(1)
)

type RJProgram struct {
	counter int
	keymap  keymap
	help    help.Model
}

type keymap struct {
	start key.Binding
	stop  key.Binding
	reset key.Binding
	setup key.Binding
	quit  key.Binding
}

func NewRJProgram() RJProgram {
	return RJProgram{
		help: help.New(),
		keymap: keymap{
			start: key.NewBinding(
				key.WithKeys("s"),
				key.WithHelp("s", "start"),
			),
			stop: key.NewBinding(
				key.WithKeys("s"),
				key.WithHelp("s", "stop"),
			),
			reset: key.NewBinding(
				key.WithKeys("r"),
				key.WithHelp("r", "reset"),
			),
			setup: key.NewBinding(
				key.WithKeys("r"),
				key.WithHelp("r", "reset"),
			),
			quit: key.NewBinding(
				key.WithKeys("ctrl+c", "q"),
				key.WithHelp("q", "quit"),
			),
		},
	}
}

func (m RJProgram) Init() tea.Cmd {
	return tea.Batch(tick(), tea.EnterAltScreen)
}

func tick() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func (m RJProgram) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		// These keys should exit the program.
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	case tickMsg:
		m.counter++

		return m, tick()
	}

	// Return the updated program to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

func (m RJProgram) View() (s string) {
	doc := strings.Builder{}
	physicalWidth, physicalHeight, _ := term.GetSize(int(os.Stdout.Fd()))

	// Header
	{
		flexbox := stickers.NewFlexBox(physicalWidth-4, 8)
		flexbox.AddRows([]*stickers.FlexBoxRow{
			flexbox.NewRow().AddCells([]*stickers.FlexBoxCell{
				stickers.NewFlexBoxCell(30, 100).
					SetMinWidth(64).
					SetContent(domain.SpriteRJLogo()),
			}),
		})
		doc.WriteString(flexbox.Render() + "\n")
	}

	// Just the middle
	{
		flexbox := stickers.NewFlexBox(physicalWidth-4, physicalHeight-13)
		flexbox.AddRows([]*stickers.FlexBoxRow{
			flexbox.NewRow().AddCells([]*stickers.FlexBoxCell{
				stickers.NewFlexBoxCell(100, 100).SetStyle(lipgloss.NewStyle()),
				stickers.NewFlexBoxCell(100, 100).SetStyle(lipgloss.NewStyle()),
				stickers.NewFlexBoxCell(100, 100).SetStyle(lipgloss.NewStyle()),
			}).SetStyle(lipgloss.NewStyle().Border(lipgloss.RoundedBorder())),
		})
		doc.WriteString(flexbox.Render() + "\n")
	}

	// Status bar
	{
		hostName, _ := os.Hostname()
		now := time.Now()

		label := statusStyle.Padding(0, 1).Render(hostName)
		bar := lipgloss.JoinHorizontal(lipgloss.Top,
			label,
			statusText.Width(physicalWidth-lipgloss.Width(label)-4).Render("PID "+strconv.Itoa(os.Getpid())+" | TIME "+fmt.Sprintf("%d:%d:%d", now.Hour(), now.Minute(), now.Second())),
		)

		doc.WriteString(bar + "\n\n")
	}

	// Help Menu
	{
		flexbox := stickers.NewFlexBox(physicalWidth-4, 1)
		flexbox.AddRows([]*stickers.FlexBoxRow{
			flexbox.NewRow().AddCells([]*stickers.FlexBoxCell{
				stickers.NewFlexBoxCell(95, 100).SetContent(m.help.ShortHelpView([]key.Binding{
					m.keymap.start,
					m.keymap.stop,
					m.keymap.reset,
					m.keymap.quit,
				})),
				stickers.NewFlexBoxCell(5, 100).SetContent("RJ 🤖").SetStyle(lipgloss.NewStyle().Align(lipgloss.Right)),
			}),
		})
		doc.WriteString(flexbox.Render())
	}

	return lipgloss.NewStyle().Padding(1, 2, 1, 2).Render(doc.String())
}
