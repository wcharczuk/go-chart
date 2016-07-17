package chart

import (
	"fmt"
	"math"
)

// Range represents a boundary for a set of numbers.
type Range struct {
	Min    float64
	Max    float64
	Domain int
}

// IsZero returns if the range has been set or not.
func (r Range) IsZero() bool {
	return (r.Min == 0 || math.IsNaN(r.Min)) &&
		(r.Max == 0 || math.IsNaN(r.Max)) &&
		r.Domain == 0
}

// Delta returns the difference between the min and max value.
func (r Range) Delta() float64 {
	return r.Max - r.Min
}

// String returns a simple string for the range.
func (r Range) String() string {
	return fmt.Sprintf("Range [%.2f,%.2f] => %d", r.Min, r.Max, r.Domain)
}

// Translate maps a given value into the range space.
func (r Range) Translate(value float64) int {
	normalized := value - r.Min
	ratio := normalized / r.Delta()
	return int(math.Ceil(ratio * float64(r.Domain)))
}

// GetRoundedRangeBounds returns some `prettified` range bounds.
func (r Range) GetRoundedRangeBounds() (min, max float64) {
	delta := r.Max - r.Min
	roundTo := GetRoundToForDelta(delta)
	return RoundDown(r.Min, roundTo), RoundUp(r.Max, roundTo)
}
