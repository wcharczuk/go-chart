package util

import (
	"math/rand"
	"time"
)

var (
	// Generate contains some sequence generation utilities.
	// These utilities can be useful for generating test data.
	Generate = &generate{
		rnd: rand.New(rand.NewSource(time.Now().Unix())),
	}
)

type generate struct {
	rnd *rand.Rand
}

// Values produces an array of floats from [start,end] by optional steps.
func (g generate) Values(start, end float64, steps ...float64) Sequence {
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
	return Sequence{Array(values)}
}

// Random generates a fixed length sequence of random values between (0, scale).
func (g generate) RandomValues(samples int, scale float64) Sequence {
	values := make([]float64, samples)

	for x := 0; x < samples; x++ {
		values[x] = g.rnd.Float64() * scale
	}

	return Sequence{Array(values)}
}

// Random generates a fixed length sequence of random values with a given average, above and below that average by (-scale, scale)
func (g generate) RandomValuesWithAverage(samples int, average, scale float64) Sequence {
	values := make([]float64, samples)

	for x := 0; x < samples; x++ {
		jitter := scale - (g.rnd.Float64() * (2 * scale))
		values[x] = average + jitter
	}

	return Sequence{Array(values)}
}
