package screen

import (
	"context"
	"strconv"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/robotjoosen/go-brew/pkg/brew"
	"github.com/robotjoosen/go-brew/pkg/widget"
)

type HelperScreen struct {
	ctx     context.Context
	counter time.Duration
	recipe  brew.Recipe
}

func NewHelperScreen(cnf brew.Brewable) ScreenAware {
	h := &HelperScreen{
		ctx:     context.Background(),
		counter: 0,
		recipe:  cnf.GetRecipe(),
	}

	go h.tick()

	return h
}

func (s *HelperScreen) Restart() {
	s.counter = time.Duration(0)
}

func (s *HelperScreen) Update(_ tea.Msg) tea.Cmd {
	return nil
}

func (s *HelperScreen) Render() string {
	columnStyle := lipgloss.NewStyle().
		Margin(1, 0, 0, 0).
		Padding(0, 1).
		Height(12).
		Border(lipgloss.RoundedBorder())

	graphString, _ := widget.
		NewGraph(s.recipe).
		SetPosition(int(s.counter.Seconds())).
		Render()

	_, ctDuration, ctGrams, _ := s.selectPour(s.counter.Seconds(), 0)
	_, ntDuration, ntGrams, _ := s.selectPour(s.counter.Seconds(), 1)

	return lipgloss.JoinHorizontal(
		lipgloss.Left,
		columnStyle.
			Width(16).
			Border(lipgloss.Border{Top: "─", Bottom: "─", Left: "│", Right: "│", TopLeft: "╭", TopRight: "┬", BottomLeft: "╰", BottomRight: "┴"}).
			Copy().
			Align(lipgloss.Left).
			Render(lipgloss.JoinVertical(
				lipgloss.Top,
				"\nnow\n--",
				"time: "+(time.Duration(ctDuration)*time.Second).String(),
				"weight: "+strconv.Itoa(ctGrams)+"g",
				"\nupcoming\n--",
				"time: "+(time.Duration(ntDuration)*time.Second).String(),
				"weight: "+strconv.Itoa(ntGrams)+"g",
			)),
		columnStyle.
			Width(28).
			Padding(0, 0, 1, 0).
			Border(lipgloss.Border{Top: "─", Bottom: "─", Left: " ", Right: "│", TopLeft: "─", TopRight: "╮", BottomLeft: "─", BottomRight: "╯"}).
			Copy().
			Align(lipgloss.Left).
			AlignVertical(lipgloss.Bottom).
			Render(graphString, "\n\n timer:", s.niceTimeString(s.counter)),
	)
}

func (s *HelperScreen) tick() {
	interval := 10 * time.Millisecond
	t := time.NewTicker(interval)
	for {
		select {
		case <-t.C:
			s.counter = s.counter + interval
		case <-s.ctx.Done():
			return
		}
	}
}

func (s *HelperScreen) niceTimeString(t time.Duration) string {
	ts := t.String()

	// todo: make the time look nicer ...

	return ts
}

func (s *HelperScreen) selectPour(pos float64, offset int) (brew.Pour, float64, int, bool) {
	var addTime float64
	var grams int
	var index int
	var selectedPour brew.Pour

	for i, pour := range s.recipe.Schema {
		index = i
		grams += pour.Grams

		if pos <= pour.Duration.Seconds()+addTime {
			selectedPour = pour
			break
		}

		addTime += pour.Duration.Seconds()
	}

	// todo: last one is still returned, should be another one.
	if offset > 0 && offset+index < len(s.recipe.Schema) {
		selectedPour = s.recipe.Schema[offset+index]
		addTime += selectedPour.Duration.Seconds()
		grams += selectedPour.Grams
	}

	return selectedPour, addTime, grams, false
}

func (s *HelperScreen) DisableControls() bool {
	return false
}
