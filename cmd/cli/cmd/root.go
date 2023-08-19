package cmd

import (
	"fmt"
	"github.com/robotjoosen/go-brew/pkg/brew"
	"github.com/robotjoosen/go-brew/pkg/recipe"
	"github.com/robotjoosen/go-brew/pkg/recipe/tetsu"
	"os"
	"strconv"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/robotjoosen/go-brew/pkg/domain"
	"github.com/rs/zerolog/log"
	"github.com/schollz/progressbar/v3"
	"github.com/spf13/cobra"
)

var Flavor string
var Concentration string
var Coffee string

var RootCmd = &cobra.Command{
	Short: "GO Brew - a simple CLI brew tool",
	Run: func(cmd *cobra.Command, args []string) {
		coffeeWeight, err := strconv.Atoi(Coffee)
		if err != nil {
			log.Error().Err(err).Send()

			return
		}

		factory := recipe.NewRecipeFactory()
		r := factory.FourSixMethod().
			SetFlavor(stringToFlavor(Flavor)).
			SetConcentration(stringToConcentration(Concentration)).
			SetCoffeeWeight(coffeeWeight).
			GetRecipe()

		renderBanner()

		time.Sleep(1 * time.Second)

		fmt.Println("BREW SETTINGS")
		renderBrewSettingsTable(r)

		fmt.Println("BREW SCHEMA")
		renderBrewSchemaTable(r)

		time.Sleep(2 * time.Second)

		startBrewGuide(r)
	},
}

func renderBanner() {
	fmt.Println(domain.SpriteGoBrew())
}

func renderBrewSettingsTable(recipe brew.Recipe) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendRows([]table.Row{
		{"ratio", recipe.Ratio},
		{"coffee", strconv.Itoa(recipe.Coffee) + " g"},
		{"water", strconv.Itoa(recipe.Water) + " ml"},
	})
	t.Render()
}

func renderBrewSchemaTable(recipe brew.Recipe) {
	var totalWeight int
	var totalDuration time.Duration

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"start time", "grams", "to add", "wait for"})

	for _, pour := range recipe.Schema {
		totalWeight += pour.Grams

		t.AppendRow(table.Row{totalDuration, strconv.Itoa(totalWeight) + " g", strconv.Itoa(pour.Grams) + " g", pour.Duration})
		totalDuration += pour.Duration
	}

	t.Render()
}

func startBrewGuide(recipe brew.Recipe) {
	startBrewCountDown()

	startTime := time.Now().Local()

	var prevTotalGrams int
	var totalGrams int
	for _, pour := range recipe.Schema {
		prevTotalGrams = totalGrams
		totalGrams += pour.Grams
		fmt.Printf("\a Add %d gram within 10 seconds [%d => %d grams] \n", pour.Grams, prevTotalGrams, totalGrams)

		NewBar(pour.Duration).Start()
	}

	fmt.Printf("\n\aBREW DONE! executed in %s\n", time.Since(startTime).Round(time.Second))
}

func startBrewCountDown() {
	fmt.Print("START BREWING IN")
	fmt.Print(" 3.\a")
	time.Sleep(time.Second)
	fmt.Print(" 2.\a")
	time.Sleep(time.Second)
	fmt.Print(" 1.\a")
	time.Sleep(time.Second)
	fmt.Print(" GO!\n\n")
}

func stringToFlavor(f string) int {
	switch f {
	case "sweet":
		return tetsu.SweetFlavor
	case "standard":
		return tetsu.BalancedFlavor
	case "bright":
		return tetsu.BrightFlavor
	}

	return 0
}

func stringToConcentration(s string) int {
	switch s {
	case "light":
		return tetsu.LightConcentration
	case "medium":
		return tetsu.MediumConcentration
	case "strong":
		return tetsu.StrongConcentration
	}

	return 0
}

type Bar struct {
	progressbar *progressbar.ProgressBar
	duration    int64
}

func NewBar(duration time.Duration) *Bar {
	return &Bar{
		duration: int64(duration.Seconds()),
		progressbar: progressbar.NewOptions64(
			int64(duration.Seconds()),
			progressbar.OptionSetPredictTime(false),
			progressbar.OptionSetWriter(os.Stderr),
			progressbar.OptionShowCount(),
			progressbar.OptionOnCompletion(func() {
				if _, err := fmt.Fprint(os.Stderr, "\n"); err != nil {
					return
				}
			}),
			progressbar.OptionSpinnerType(14),
			progressbar.OptionSetWidth(int(duration.Seconds())),
		),
	}
}

func (b *Bar) Start() {
	ticker := time.NewTicker(time.Second)
	for i := 0; int64(i) < b.duration; i++ {
		select {
		case <-ticker.C:
			if err := b.progressbar.Add(1); err != nil {
				return
			}
		}
	}
}
