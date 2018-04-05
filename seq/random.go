package seq

import (
	"math"
	"math/rand"
	"time"
)

// RandomValues returns an array of random values.
func RandomValues(count int) []float64 {
	return Seq{NewRandom().WithLen(count)}.Array()
}

// RandomValuesWithMax returns an array of random values with a given average.
func RandomValuesWithMax(count int, max float64) []float64 {
	return Seq{NewRandom().WithMax(max).WithLen(count)}.Array()
}

// NewRandom creates a new random seq.
func NewRandom() *Random {
	return &Random{
		rnd: rand.New(rand.NewSource(time.Now().Unix())),
	}
}

// Random is a random number seq generator.
type Random struct {
	rnd *rand.Rand
	max *float64
	min *float64
	len *int
}

// Len returns the number of elements that will be generated.
func (r *Random) Len() int {
	if r.len != nil {
		return *r.len
	}
	return math.MaxInt32
}

// GetValue returns the value.
func (r *Random) GetValue(_ int) float64 {
	if r.min != nil && r.max != nil {
		var delta float64

		if *r.max > *r.min {
			delta = *r.max - *r.min
		} else {
			delta = *r.min - *r.max
		}

		return *r.min + (r.rnd.Float64() * delta)
	} else if r.max != nil {
		return r.rnd.Float64() * *r.max
	} else if r.min != nil {
		return *r.min + (r.rnd.Float64())
	}
	return r.rnd.Float64()
}

// WithLen sets a maximum len
func (r *Random) WithLen(length int) *Random {
	r.len = &length
	return r
}

// Min returns the minimum value.
func (r Random) Min() *float64 {
	return r.min
}

// WithMin sets the scale and returns the Random.
func (r *Random) WithMin(min float64) *Random {
	r.min = &min
	return r
}

// Max returns the maximum value.
func (r Random) Max() *float64 {
	return r.max
}

// WithMax sets the average and returns the Random.
func (r *Random) WithMax(max float64) *Random {
	r.max = &max
	return r
}
