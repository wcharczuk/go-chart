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

// GenerateContinuousTicksWithStep generates a set of ticks.
func GenerateContinuousTicksWithStep(ra Range, step float64, vf ValueFormatter) []Tick {
	var ticks []Tick
	min, max := ra.GetMin(), ra.GetMax()
	for cursor := min; cursor <= max; cursor += step {
		ticks = append(ticks, Tick{
			Value: cursor,
			Label: vf(cursor),
		})

		// this guard is in place in case step is super, super small.
		if len(ticks) > DefaultTickCountSanityCheck {
			return ticks
		}
	}
	return ticks
}

// CalculateContinuousTickStep calculates the continous range interval between ticks.
func CalculateContinuousTickStep(r Renderer, ra Range, isVertical bool, style Style, vf ValueFormatter) float64 {
	r.SetFont(style.GetFont())
	r.SetFontSize(style.GetFontSize())
	if isVertical {
		label := vf(ra.GetMin())
		tb := r.MeasureText(label)
		count := int(math.Ceil(float64(ra.GetDomain()) / float64(tb.Height()+DefaultMinimumTickVerticalSpacing)))
		return ra.GetDelta() / float64(count)
	}

	// take a cut at determining the 'widest' value.
	l0 := vf(ra.GetMin())
	ln := vf(ra.GetMax())
	ll := l0
	if len(ln) > len(l0) {
		ll = ln
	}
	llb := r.MeasureText(ll)
	textWidth := llb.Width()
	width := textWidth + DefaultMinimumTickHorizontalSpacing
	count := int(math.Ceil(float64(ra.GetDomain()) / float64(width)))
	return ra.GetDelta() / float64(count)
}
