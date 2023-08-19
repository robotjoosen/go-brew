package tetsu_test

import (
	"github.com/robotjoosen/go-brew/pkg/brew"
	"github.com/robotjoosen/go-brew/pkg/recipe/tetsu"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFourSixMethod_GenerateSchema(t *testing.T) {
	testCases := []struct {
		withCoffeeWeight  int
		withFlavor        int
		withConcentration int
		expectedSchema    []brew.Pour
	}{
		{
			withCoffeeWeight:  16,
			withFlavor:        tetsu.BalancedFlavor,
			withConcentration: tetsu.StrongConcentration,
			expectedSchema: []brew.Pour{
				{Duration: 45 * time.Second, Grams: 48},
				{Duration: 45 * time.Second, Grams: 48},
				{Duration: 45 * time.Second, Grams: 48},
				{Duration: 45 * time.Second, Grams: 48},
				{Duration: 30 * time.Second, Grams: 48},
			},
		},
		{
			withCoffeeWeight:  16,
			withFlavor:        tetsu.SweetFlavor,
			withConcentration: tetsu.StrongConcentration,
			expectedSchema: []brew.Pour{
				{Duration: 45 * time.Second, Grams: 40},
				{Duration: 45 * time.Second, Grams: 56},
				{Duration: 45 * time.Second, Grams: 48},
				{Duration: 45 * time.Second, Grams: 48},
				{Duration: 30 * time.Second, Grams: 48},
			},
		},
		{
			withCoffeeWeight:  16,
			withFlavor:        tetsu.BrightFlavor,
			withConcentration: tetsu.StrongConcentration,
			expectedSchema: []brew.Pour{
				{Duration: 45 * time.Second, Grams: 56},
				{Duration: 45 * time.Second, Grams: 40},
				{Duration: 45 * time.Second, Grams: 48},
				{Duration: 45 * time.Second, Grams: 48},
				{Duration: 30 * time.Second, Grams: 48},
			},
		},
		{
			withCoffeeWeight:  20,
			withFlavor:        tetsu.BalancedFlavor,
			withConcentration: tetsu.StrongConcentration,
			expectedSchema: []brew.Pour{
				{Duration: 45 * time.Second, Grams: 60},
				{Duration: 45 * time.Second, Grams: 60},
				{Duration: 45 * time.Second, Grams: 60},
				{Duration: 45 * time.Second, Grams: 60},
				{Duration: 30 * time.Second, Grams: 60},
			},
		},
		{
			withCoffeeWeight:  20,
			withFlavor:        tetsu.SweetFlavor,
			withConcentration: tetsu.StrongConcentration,
			expectedSchema: []brew.Pour{
				{Duration: 45 * time.Second, Grams: 50},
				{Duration: 45 * time.Second, Grams: 70},
				{Duration: 45 * time.Second, Grams: 60},
				{Duration: 45 * time.Second, Grams: 60},
				{Duration: 30 * time.Second, Grams: 60},
			},
		},
		{
			withCoffeeWeight:  20,
			withFlavor:        tetsu.BrightFlavor,
			withConcentration: tetsu.StrongConcentration,
			expectedSchema: []brew.Pour{
				{Duration: 45 * time.Second, Grams: 70},
				{Duration: 45 * time.Second, Grams: 50},
				{Duration: 45 * time.Second, Grams: 60},
				{Duration: 45 * time.Second, Grams: 60},
				{Duration: 30 * time.Second, Grams: 60},
			},
		},
		{
			withCoffeeWeight:  20,
			withFlavor:        tetsu.BalancedFlavor,
			withConcentration: tetsu.LightConcentration,
			expectedSchema: []brew.Pour{
				{Duration: 45 * time.Second, Grams: 60},
				{Duration: 45 * time.Second, Grams: 60},
				{Duration: 120 * time.Second, Grams: 180},
			},
		},
		{
			withCoffeeWeight:  20,
			withFlavor:        tetsu.BalancedFlavor,
			withConcentration: tetsu.MediumConcentration,
			expectedSchema: []brew.Pour{
				{Duration: 45 * time.Second, Grams: 60},
				{Duration: 45 * time.Second, Grams: 60},
				{Duration: 60 * time.Second, Grams: 90},
				{Duration: 60 * time.Second, Grams: 90},
			},
		},
	}

	for _, tc := range testCases {
		schema := tetsu.New().
			SetFlavor(tc.withFlavor).
			SetConcentration(tc.withConcentration).
			SetCoffeeWeight(tc.withCoffeeWeight).
			GenerateSchema()

		assert.Equal(t, tc.expectedSchema, schema)
	}
}

func TestFourSixMethod_GetIngredients(t *testing.T) {
	testCases := []struct {
		withCoffeeWeight  int
		withFlavor        int
		withConcentration int
		expectsCoffee     int
		expectsWater      int
	}{
		{
			withCoffeeWeight: 14,
			expectsCoffee:    14,
			expectsWater:     210,
		},
		{
			withCoffeeWeight: 16,
			expectsCoffee:    16,
			expectsWater:     240,
		},
		{
			withCoffeeWeight: 20,
			expectsCoffee:    20,
			expectsWater:     300,
		},
	}

	for _, tc := range testCases {
		recipe := tetsu.New().
			SetFlavor(tc.withFlavor).
			SetConcentration(tc.withConcentration).
			SetCoffeeWeight(tc.withCoffeeWeight).
			GetRecipe()

		assert.Equal(t, recipe.Water, tc.expectsWater)
		assert.Equal(t, recipe.Coffee, tc.expectsCoffee)
	}
}
