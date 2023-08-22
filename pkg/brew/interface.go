package brew

type Brewable interface {
	SetCoffeeWeight(grams int) Brewable
	SetWaterWeight(grams int) Brewable
	GetRecipe() Recipe
	GenerateSchema() []Pour
}
