package screen

import (
	"context"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/robotjoosen/go-brew/pkg/sprite"
	"time"
)

type SplashScreen struct {
	ctx          context.Context
	frameCounter int
}

func NewSplashScreen() ScreenAware {
	s := &SplashScreen{
		ctx: context.Background(),
	}

	go s.tick()

	return s
}

func (s *SplashScreen) Restart() {}

func (s *SplashScreen) Update(msg tea.Msg) tea.Cmd {
	return nil
}

func (s *SplashScreen) Render() string {
	columnStyle := lipgloss.NewStyle().
		Margin(1, 0, 0, 0).
		Padding(0, 1).
		Height(12).
		Border(lipgloss.RoundedBorder())

	return lipgloss.JoinHorizontal(
		lipgloss.Left,
		columnStyle.Width(46).Render(
			lipgloss.JoinHorizontal(
				lipgloss.Left,
				lipgloss.NewStyle().Width(16).Render(""),
				lipgloss.NewStyle().Width(25).Render("\n"+sprite.New().CoffeeAnimated(s.frameCounter)),
			),
		),
	)
}

func (s *SplashScreen) DisableControls() bool {
	return false
}

func (s *SplashScreen) tick() {
	t := time.NewTicker(150 * time.Millisecond)
	for {
		select {
		case <-t.C:
			s.frameCounter++
		case <-s.ctx.Done():
			return
		}
	}
}
