package widget

type WidgetAware interface {
	SetPosition(int) WidgetAware
	Render() (string, error)
}
