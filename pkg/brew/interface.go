package brew

type Brewable interface {
	SetCoffeeWeight(grams int) Brewable
	GetRecipe() Recipe
	GenerateSchema() []Pour
}
