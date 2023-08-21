package screen

import "github.com/robotjoosen/go-brew/pkg/brew"

const (
	Splash = "splash"
	Helper = "helper"
	Schema = "schema"
	Config = "config"
)

type ScreenFactory struct {
	brewable brew.Brewable
}

func NewFactory() *ScreenFactory {
	return new(ScreenFactory)
}

func (f *ScreenFactory) UpdateRecipeFactory(brewable brew.Brewable) *ScreenFactory {
	f.brewable = brewable

	return f
}

func (f *ScreenFactory) New(screen string) ScreenAware {
	switch screen {
	case Splash:
		return NewSplashScreen()
	case Helper:
		return NewHelperScreen(f.brewable)
	case Schema:
		return NewSchemaScreen(f.brewable)
	case Config:
		return NewConfigurationScreen(f.brewable)
	}

	return nil
}
