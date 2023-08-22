package widget

import (
	"sync"

	"github.com/guptarohit/asciigraph"
	"github.com/robotjoosen/go-brew/pkg/brew"
)

type Graph struct {
	mux      *sync.RWMutex
	recipe   brew.Recipe
	position int
}

func NewGraph(recipe brew.Recipe) WidgetAware {
	return &Graph{
		mux:    new(sync.RWMutex),
		recipe: recipe,
	}
}

func (g *Graph) SetPosition(pos int) WidgetAware {
	g.mux.RLock()
	defer g.mux.RUnlock()

	g.position = pos

	return g
}

func (g *Graph) Render() (string, error) {
	g.mux.RLock()
	defer g.mux.RUnlock()

	graphData := g.schemaToPoints(g.recipe.Schema)
	graphSize := []int{22, 7}

	return asciigraph.PlotMany(
		g.plotGraphSegment(graphData, g.position+1, graphSize[0]),
		asciigraph.Height(graphSize[1]),
		asciigraph.Precision(0),
		asciigraph.SeriesColors(asciigraph.Default, asciigraph.Red),
	), nil
}

func (g *Graph) schemaToPoints(schema []brew.Pour) (data []float64) {
	var i float64
	var currentValue float64

	// add zero value as a start point
	data = append(data, 0)

	for _, pour := range schema {

		steps := float64(pour.Grams/10) / 10

		for i = 0; i < 10; i++ {
			if i > pour.Duration.Seconds() {
				break
			}

			currentValue += steps * 10
			data = append(data, currentValue)
		}

		remainingTime := pour.Duration.Seconds() - 10
		for i = 0; i < remainingTime; i++ {
			data = append(data, currentValue)
		}
	}

	return data
}

func (g *Graph) plotGraphSegment(data []float64, position int, width int) [][]float64 {
	if position > len(data) {
		position = len(data)
	}

	minC := position - (width / 2)
	if minC < 0 {
		minC = 0
	}
	maxC := position + (width / 2)
	if maxC > len(data) {
		maxC = len(data)
	}
	if maxC < width {
		maxC = width
	}

	current := make([]float64, 1)
	if minC != position {
		current = data[minC:position]
	}

	plot := [][]float64{data[minC:maxC], current}

	return plot
}
