package chart

import (
	"fmt"
	"math"
)

// Range represents a continuous range,
type Range struct {
	Min    float64
	Max    float64
	Domain int
}

// IsZero returns if the range has been set or not.
func (r Range) IsZero() bool {
	return r.Min == 0 && r.Max == 0 && r.Domain == 0
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
// An example would be a 600 px image, with a min of 10 and a max of 100.
// Translate(50) would yield (50.0/90.0)*600 ~= 333.33
func (r Range) Translate(value float64) int {
	finalValue := ((r.Max - value) / r.Delta()) * float64(r.Domain)
	return int(math.Floor(finalValue))
}
