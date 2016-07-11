package chart

import (
	"fmt"
	"time"
)

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

// Float is an alias for float64 that provides a better .String() method.
type Float float64

// String returns the string representation of a float.
func (f Float) String() string {
	return fmt.Sprintf("%.2f", f)
}
