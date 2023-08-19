package widget

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/robotjoosen/go-brew/pkg/brew"
	"strconv"
	"sync"
)

var (
	special = lipgloss.AdaptiveColor{Light: "#43BF6D", Dark: "#73F59F"}
	list    = lipgloss.NewStyle().
		Border(lipgloss.NormalBorder(), false, true, false, false).
		MarginRight(2).
		Height(8)

	listHeader = lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderBottom(true).
			MarginRight(2).
			Render

	listItem = lipgloss.NewStyle().PaddingLeft(2).Render

	checkMark = lipgloss.NewStyle().SetString("✓").
			Foreground(special).
			PaddingRight(1).
			String()

	listDone = func(s string) string {
		return checkMark + lipgloss.NewStyle().
			Strikethrough(true).
			Foreground(lipgloss.AdaptiveColor{Light: "#969B86", Dark: "#696969"}).
			Render(s)
	}
)

type List struct {
	mux            *sync.RWMutex
	schema         []brew.Pour
	position       int
	additionalTime float64
}

func NewList(schema []brew.Pour) WidgetAware {
	return &List{
		mux:    new(sync.RWMutex),
		schema: schema,
	}
}

func (l *List) SetPosition(pos int) WidgetAware {
	l.mux.RLock()
	defer l.mux.RUnlock()

	l.position = pos

	return l
}

func (l *List) Render() (output string, err error) {
	var ls string
	for i, _ := range l.schema {
		ls += listItem(strconv.Itoa(i)) + "\n"
	}

	output = lipgloss.JoinVertical(lipgloss.Left, listHeader("steps"), ls)

	return
}
