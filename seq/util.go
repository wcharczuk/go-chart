package seq

import (
	"math"
	"time"

	"github.com/wcharczuk/go-chart/util"
)

func round(input float64, places int) (rounded float64) {
	if math.IsNaN(input) {
		return 0.0
	}

	sign := 1.0
	if input < 0 {
		sign = -1
		input *= -1
	}

	precision := math.Pow(10, float64(places))
	digit := input * precision
	_, decimal := math.Modf(digit)

	if decimal >= 0.5 {
		rounded = math.Ceil(digit)
	} else {
		rounded = math.Floor(digit)
	}

	return rounded / precision * sign
}

func f64i(value float64) int {
	r := round(value, 0)
	return int(r)
}

// SetOfTime is a simple hash set for timestamps as float64s.
type SetOfTime map[float64]bool

// Add adds the value to the hash set.
func (sot SetOfTime) Add(tv time.Time) {
	sot[util.Time.ToFloat64(tv)] = true
}

// Has returns if the set contains a given time.
func (sot SetOfTime) Has(tv time.Time) bool {
	_, hasValue := sot[util.Time.ToFloat64(tv)]
	return hasValue
}

// Remove removes the value from the set.
func (sot SetOfTime) Remove(tv time.Time) {
	delete(sot, util.Time.ToFloat64(tv))
}
