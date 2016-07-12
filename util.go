package chart

import (
	"fmt"
	"math"
	"time"
)

// Float is an alias for float64 that provides a better .String() method.
type Float float64

// String returns the string representation of a float.
func (f Float) String() string {
	return fmt.Sprintf("%.2f", f)
}

// TimeToFloat64 returns a float64 representation of a time.
func TimeToFloat64(t time.Time) float64 {
	return float64(t.Unix())
}

// MinAndMax returns both the min and max in one pass.
func MinAndMax(values ...float64) (min float64, max float64) {
	if len(values) == 0 {
		return
	}
	min = values[0]
	max = values[0]
	for _, v := range values {
		if max < v {
			max = v
		}
		if min > v {
			min = v
		}
	}
	return
}

// MinAndMaxOfTime returns the min and max of a given set of times
// in one pass.
func MinAndMaxOfTime(values ...time.Time) (min time.Time, max time.Time) {
	if len(values) == 0 {
		return
	}

	min = values[0]
	max = values[0]

	for _, v := range values {
		if max.Before(v) {
			max = v
		}
		if min.After(v) {
			min = v
		}
	}
	return
}

// Slices generates N slices that span the total.
// The resulting array will be intermediate indexes until total.
func Slices(count int, total float64) []float64 {
	var values []float64
	sliceWidth := float64(total) / float64(count)
	for cursor := 0.0; cursor < total; cursor += sliceWidth {
		values = append(values, cursor)
	}
	return values
}

// GetRoundToForDelta returns a `roundTo` value for a given delta.
func GetRoundToForDelta(delta float64) float64 {
	startingDeltaBound := math.Pow(10.0, 10.0)
	for cursor := startingDeltaBound; cursor > 0; cursor /= 10.0 {
		if delta > cursor {
			return cursor / 10.0
		}
	}

	return 0.0
}

// RoundUp rounds up to a given roundTo value.
func RoundUp(value, roundTo float64) float64 {
	d1 := math.Ceil(value / roundTo)
	return d1 * roundTo
}

// RoundDown rounds down to a given roundTo value.
func RoundDown(value, roundTo float64) float64 {
	d1 := math.Floor(value / roundTo)
	return d1 * roundTo
}

// MinInt returns the minimum of a set of integers.
func MinInt(values ...int) int {
	min := math.MaxInt32
	for _, v := range values {
		if v < min {
			min = v
		}
	}
	return min
}

// MaxInt returns the maximum of a set of integers.
func MaxInt(values ...int) int {
	max := math.MinInt32
	for _, v := range values {
		if v > max {
			max = v
		}
	}
	return max
}

// AbsInt returns the absolute value of an integer.
func AbsInt(value int) int {
	if value < 0 {
		return -value
	}
	return value
}
