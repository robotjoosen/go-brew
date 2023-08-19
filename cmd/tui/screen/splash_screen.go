package screen

type SplashScreen struct{}

func NewSplashScreen() ScreenAware {
	return &SplashScreen{}
}

func (s *SplashScreen) Render() string {
	return ""
}
