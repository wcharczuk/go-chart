package chart

import (
	"fmt"
	"image/color"
	"time"
)

// ColorIsZero returns if a color.RGBA is unset or not.
func ColorIsZero(c color.RGBA) bool {
	return c.R == 0 && c.G == 0 && c.B == 0 && c.A == 0
}

// ColorAsString returns if a color.RGBA is unset or not.
func ColorAsString(c color.RGBA) string {
	a := float64(c.A) / float64(255)
	return fmt.Sprintf("rgba(%v,%v,%v,%.1f)", c.R, c.G, c.B, a)
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
