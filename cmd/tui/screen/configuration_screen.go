package screen

type ConfigurationScreen struct{}

func NewConfigurationScreen() ScreenAware {
	return &ConfigurationScreen{}
}

func (s *ConfigurationScreen) Render() string {
	return ""
}
