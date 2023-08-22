package tetsu

import (
	"github.com/nethruster/go-fraction"
	"github.com/robotjoosen/go-brew/pkg/brew"
	"strconv"
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

type FourSixMethod struct {
	mux                 *sync.RWMutex
	ratio               []int64
	coffeeFlavor        int
	coffeeConcentration int
	coffeeWeight        int
	waterWeight         int
}

func New() *FourSixMethod {
	return &FourSixMethod{
		mux:          new(sync.RWMutex),
		ratio:        []int64{1, 15},
		coffeeWeight: 20,
		waterWeight:  300,
	}
}

func (fsm *FourSixMethod) SetCoffeeWeight(weight int) brew.Brewable {
	fsm.mux.RLock()
	defer fsm.mux.RUnlock()

	fsm.coffeeWeight = weight
	fsm.ratio = fsm.calcRatio()

	return fsm
}

func (fsm *FourSixMethod) SetWaterWeight(weight int) brew.Brewable {
	fsm.mux.RLock()
	defer fsm.mux.RUnlock()

	fsm.waterWeight = weight
	fsm.ratio = fsm.calcRatio()

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
		Ratio:  strconv.Itoa(int(fsm.ratio[0])) + ":" + strconv.Itoa(int(fsm.ratio[1])),
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

func (fsm *FourSixMethod) calcRatio() []int64 {
	f, err := fraction.New(fsm.coffeeWeight, fsm.waterWeight)
	if err != nil {
		return make([]int64, 2)
	}

	return []int64{f.Numerator(), f.Denominator()}
}
