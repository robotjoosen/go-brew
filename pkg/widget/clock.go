package widget

import (
	"errors"
	"fmt"
	"math"
	"sync"

	"github.com/robotjoosen/go-brew/pkg/brew"
)

const (
	DotFull        = "●"
	Star           = "*"
	DotEmpty       = "○"
	PourText       = "pour"
	WaitText       = "wait"
	BlankText      = "    "
	PourDuration   = 10
	ClockPositions = 12
)

type Clock struct {
	mux            *sync.RWMutex
	schema         []brew.Pour
	position       int
	additionalTime float64
}

func NewClock(schema []brew.Pour) WidgetAware {
	return &Clock{
		mux:    new(sync.RWMutex),
		schema: schema,
	}
}

func (c *Clock) SetPosition(pos int) WidgetAware {
	c.mux.RLock()
	defer c.mux.RUnlock()

	c.position = pos

	return c
}

func (c *Clock) Render() (string, error) {
	c.mux.RLock()
	defer c.mux.RUnlock()

	pour, err := c.selectPour(c.position)
	if err != nil {
		return "", err
	}

	points := c.renderPoints(pour.Duration.Seconds())
	pourPlaceholder, waitPlaceholder := c.renderPlaceholderText()

	return fmt.Sprintf(
		"    %s  %s\n %s        %s\n%s   %s   %s\n%s   %s   %s\n %s        %s\n    %s  %s", // it's a circle
		points[11], points[0], points[10], points[1], points[9], pourPlaceholder, points[2], points[8], waitPlaceholder, points[3], points[7], points[4], points[6], points[5],
	), nil
}

func (c *Clock) GenerateRaw() (points []string, pourText string, waitText string, err error) {
	c.mux.RLock()
	defer c.mux.RUnlock()

	pour, err := c.selectPour(c.position)
	if err != nil {
		return
	}

	points = c.renderPoints(pour.Duration.Seconds())
	pourText, waitText = c.renderPlaceholderText()

	return
}

func (c *Clock) renderPoints(pourDuration float64) []string {
	pourTime := math.Round((ClockPositions / pourDuration) * PourDuration)
	currentTime := math.Round((ClockPositions / pourDuration) * (float64(c.position) - c.additionalTime))

	points := make([]string, ClockPositions)
	for i := range points {
		if i < int(pourTime) {
			points[i] = Star
		} else {
			points[i] = DotEmpty
		}

		if i < int(currentTime) {
			points[i] = DotFull
		}
	}

	return points
}

func (c *Clock) renderPlaceholderText() (string, string) {
	pourPlaceholder := BlankText
	waitPlaceholder := BlankText
	if c.shouldPour(c.position) {
		pourPlaceholder = PourText
	} else {
		waitPlaceholder = WaitText
	}

	return pourPlaceholder, waitPlaceholder
}

func (c *Clock) shouldPour(pos int) bool {
	if (float64(pos) - c.additionalTime) < PourDuration {
		return true
	}

	return false
}

func (c *Clock) selectPour(pos int) (brew.Pour, error) {
	for _, pour := range c.schema {
		if pos <= int(pour.Duration.Seconds()+c.additionalTime) {
			return pour, nil
		}

		c.additionalTime += pour.Duration.Seconds()
	}

	return brew.Pour{}, errors.New("no pour eligible")
}
