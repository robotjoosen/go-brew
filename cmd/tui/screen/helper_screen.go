package screen

type HelperScreen struct{}

func NewHelperScreen() ScreenAware {
	return &HelperScreen{}
}

func (s *HelperScreen) Render() string {
	return ""
}
