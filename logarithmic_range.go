package chart

import (
	"fmt"
	"math"
)

// LogarithmicRange represents a boundary for a set of numbers.
type LogarithmicRange struct {
	Min        float64
	Max        float64
	Domain     int
	Descending bool
}

// IsDescending returns if the range is descending.
func (r LogarithmicRange) IsDescending() bool {
	return r.Descending
}

// IsZero returns if the LogarithmicRange has been set or not.
func (r LogarithmicRange) IsZero() bool {
	return (r.Min == 0 || math.IsNaN(r.Min)) &&
		(r.Max == 0 || math.IsNaN(r.Max)) &&
		r.Domain == 0
}

// GetMin gets the min value for the continuous range.
func (r LogarithmicRange) GetMin() float64 {
	return r.Min
}

// SetMin sets the min value for the continuous range.
func (r *LogarithmicRange) SetMin(min float64) {
	r.Min = min
}

// GetMax returns the max value for the continuous range.
func (r LogarithmicRange) GetMax() float64 {
	return r.Max
}

// SetMax sets the max value for the continuous range.
func (r *LogarithmicRange) SetMax(max float64) {
	r.Max = max
}

// GetDelta returns the difference between the min and max value.
func (r LogarithmicRange) GetDelta() float64 {
	return r.Max - r.Min
}

// GetDomain returns the range domain.
func (r LogarithmicRange) GetDomain() int {
	return r.Domain
}

// SetDomain sets the range domain.
func (r *LogarithmicRange) SetDomain(domain int) {
	r.Domain = domain
}

// String returns a simple string for the LogarithmicRange.
func (r LogarithmicRange) String() string {
	return fmt.Sprintf("LogarithmicRange [%.2f,%.2f] => %d", r.Min, r.Max, r.Domain)
}

// Translate maps a given value into the LogarithmicRange space. Modified version from ContinuousRange.
func (r LogarithmicRange) Translate(value float64) int {
	if value < 1 {
		return 0
	}
	normalized := math.Max(value-r.Min, 1)
	ratio := math.Log10(normalized) / math.Log10(r.GetDelta())

	if r.IsDescending() {
		return r.Domain - int(math.Ceil(ratio*float64(r.Domain)))
	}

	return int(math.Ceil(ratio * float64(r.Domain)))
}

// GetTicks calculates the needed ticks for the axis, in log scale. Only supports Y values > 0.
func (r LogarithmicRange) GetTicks(render Renderer, defaults Style, vf ValueFormatter) []Tick {
	var ticks []Tick
	exponentStart := int64(math.Max(0, math.Floor(math.Log10(r.Min))))   // one below min
	exponentEnd := int64(math.Max(0, math.Ceil(math.Log10(r.Max))))      // one above max
	for exp:=exponentStart; exp<=exponentEnd; exp++ {
		tickVal := math.Pow(10, float64(exp))
		ticks = append(ticks, Tick{Value: tickVal, Label: vf(tickVal)})
	}

	return ticks
}
