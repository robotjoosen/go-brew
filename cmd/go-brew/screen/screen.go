package screen

import tea "github.com/charmbracelet/bubbletea"

type ScreenAware interface {
	Render() string
	Restart()
	Update(msg tea.Msg) tea.Cmd
	DisableControls() bool
}
