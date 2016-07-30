package chart

import "math"

// TicksProvider is a type that provides ticks.
type TicksProvider interface {
	GetTicks(vf ValueFormatter) []Tick
}

// Tick represents a label on an axis.
type Tick struct {
	Value float64
	Label string
}

// Ticks is an array of ticks.
type Ticks []Tick

// Len returns the length of the ticks set.
func (t Ticks) Len() int {
	return len(t)
}

// Swap swaps two elements.
func (t Ticks) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

// Less returns if i's value is less than j's value.
func (t Ticks) Less(i, j int) bool {
	return t[i].Value < t[j].Value
}

// GenerateContinuousTicks generates a set of ticks.
func GenerateContinuousTicks(r Renderer, ra Range, isVertical bool, style Style, vf ValueFormatter) []Tick {
	var ticks []Tick
	min, max := ra.GetMin(), ra.GetMax()

	ticks = append(ticks, Tick{
		Value: min,
		Label: vf(min),
	})

	minLabel := vf(min)
	style.GetTextOptions().WriteToRenderer(r)
	labelBox := r.MeasureText(minLabel)

	var tickSize int
	if isVertical {
		tickSize = labelBox.Height() + DefaultMinimumTickVerticalSpacing
	} else {
		tickSize = labelBox.Width() + DefaultMinimumTickHorizontalSpacing
	}

	domainRemainder := (ra.GetDomain()) - (tickSize * 2)
	intermediateTickCount := int(math.Floor(float64(domainRemainder) / float64(tickSize)))

	rangeDelta := math.Floor(max - min)
	tickStep := rangeDelta / float64(intermediateTickCount)

	roundTo := Math.GetRoundToForDelta(rangeDelta) / 10

	for x := 1; x < intermediateTickCount; x++ {
		tickValue := min + Math.RoundDown(tickStep*float64(x), roundTo)
		ticks = append(ticks, Tick{
			Value: tickValue,
			Label: vf(tickValue),
		})
	}

	ticks = append(ticks, Tick{
		Value: max,
		Label: vf(max),
	})

	return ticks
}
