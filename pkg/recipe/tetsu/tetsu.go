package tetsu

import (
	"github.com/robotjoosen/go-brew/pkg/brew"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	SweetFlavor = iota
	BalancedFlavor
	BrightFlavor
)

const (
	LightConcentration = iota
	MediumConcentration
	StrongConcentration
)

const (
	CoffeeToWaterRatio = "1:15"
)

type FourSixMethod struct {
	mux                 *sync.RWMutex
	ratio               string
	coffeeFlavor        int
	coffeeConcentration int
	coffeeWeight        int
	waterWeight         int
}

func New() *FourSixMethod {
	return &FourSixMethod{
		mux:   new(sync.RWMutex),
		ratio: CoffeeToWaterRatio,
	}
}

func (fsm *FourSixMethod) SetCoffeeWeight(weight int) brew.Brewable {
	fsm.mux.RLock()
	defer fsm.mux.RUnlock()

	fsm.coffeeWeight = weight

	_, waterRatio := convertRatio(CoffeeToWaterRatio)
	fsm.waterWeight = weight * waterRatio

	return fsm
}

func (fsm *FourSixMethod) SetFlavor(flavor int) *FourSixMethod {
	fsm.mux.RLock()
	defer fsm.mux.RUnlock()

	fsm.coffeeFlavor = flavor

	return fsm
}

func (fsm *FourSixMethod) SetConcentration(concentration int) *FourSixMethod {
	fsm.mux.RLock()
	defer fsm.mux.RUnlock()

	fsm.coffeeConcentration = concentration

	return fsm
}

func (fsm *FourSixMethod) GetRecipe() brew.Recipe {
	fsm.mux.RLock()
	defer fsm.mux.RUnlock()

	return brew.Recipe{
		Ratio:  fsm.ratio,
		Coffee: fsm.coffeeWeight,
		Water:  fsm.waterWeight,
		Schema: fsm.GenerateSchema(),
	}
}

func (fsm *FourSixMethod) GenerateSchema() (pours []brew.Pour) {
	fsm.mux.RLock()
	defer fsm.mux.RUnlock()

	pourSize := fsm.waterWeight / 5
	flavorPhase := pourSize * 2

	switch fsm.coffeeFlavor {
	case SweetFlavor:
		pours = append(pours, brew.Pour{Duration: 45 * time.Second, Grams: (flavorPhase / 12) * 5})
		pours = append(pours, brew.Pour{Duration: 45 * time.Second, Grams: (flavorPhase / 12) * 7})
	case BalancedFlavor:
		pours = append(pours, brew.Pour{Duration: 45 * time.Second, Grams: flavorPhase / 2})
		pours = append(pours, brew.Pour{Duration: 45 * time.Second, Grams: flavorPhase / 2})
	case BrightFlavor:
		pours = append(pours, brew.Pour{Duration: 45 * time.Second, Grams: (flavorPhase / 12) * 7})
		pours = append(pours, brew.Pour{Duration: 45 * time.Second, Grams: (flavorPhase / 12) * 5})
	}

	concentrationPhase := pourSize * 3
	switch fsm.coffeeConcentration {
	case LightConcentration:
		pours = append(pours, brew.Pour{Duration: 120 * time.Second, Grams: concentrationPhase})
	case MediumConcentration:
		pours = append(pours, brew.Pour{Duration: 60 * time.Second, Grams: concentrationPhase / 2})
		pours = append(pours, brew.Pour{Duration: 60 * time.Second, Grams: concentrationPhase / 2})
	case StrongConcentration:
		pours = append(pours, brew.Pour{Duration: 45 * time.Second, Grams: concentrationPhase / 3})
		pours = append(pours, brew.Pour{Duration: 45 * time.Second, Grams: concentrationPhase / 3})
		pours = append(pours, brew.Pour{Duration: 30 * time.Second, Grams: concentrationPhase / 3})
	}

	return
}

func convertRatio(ratio string) (coffeeRatio int, waterRatio int) {
	ratios := strings.Split(ratio, ":")
	if len(ratios) != 2 {
		panic("ratio is not what you think")
	}

	coffeeRatio, err := strconv.Atoi(ratios[0])
	if err != nil {
		panic(err.Error())
	}

	waterRatio, err = strconv.Atoi(ratios[1])
	if err != nil {
		panic(err.Error())
	}

	return
}
