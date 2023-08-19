package recipe

import (
	"github.com/robotjoosen/go-brew/pkg/recipe/tetsu"
)

type Factory struct{}

func NewRecipeFactory() Factory {
	return Factory{}
}

func (f Factory) FourSixMethod() *tetsu.FourSixMethod {
	return tetsu.New()
}
