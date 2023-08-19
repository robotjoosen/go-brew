package screen

type SchemaScreen struct{}

func NewSchemaScreen() ScreenAware {
	return &SchemaScreen{}
}

func (s *SchemaScreen) Render() string {
	return ""
}
