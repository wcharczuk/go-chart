package chart

import (
	"math/rand"
	"time"
)

var (
	// Sequence contains some sequence utilities.
	// These utilities can be useful for generating test data.
	Sequence = &sequence{}
)

type sequence struct{}

// Float64 produces an array of floats from [start,end] by optional steps.
func (s sequence) Float64(start, end float64, steps ...float64) []float64 {
	var values []float64
	step := 1.0
	if len(steps) > 0 {
		step = steps[0]
	}

	if start < end {
		for x := start; x <= end; x += step {
			values = append(values, x)
		}
	} else {
		for x := start; x >= end; x = x - step {
			values = append(values, x)
		}
	}
	return values
}

// Random generates a fixed length sequence of random values between (0, scale).
func (s sequence) Random(samples int, scale float64) []float64 {
	rnd := rand.New(rand.NewSource(time.Now().Unix()))
	values := make([]float64, samples)

	for x := 0; x < samples; x++ {
		values[x] = rnd.Float64() * scale
	}

	return values
}

// Days generates a sequence of timestamps by day, from -days to today.
func (s sequence) Days(days int) []time.Time {
	var values []time.Time
	for day := days; day >= 0; day-- {
		values = append(values, time.Now().AddDate(0, 0, -day))
	}
	return values
}
