package sequence

import (
	"math"
	"math/rand"
	"time"
)

// RandomValues returns an array of random values.
func RandomValues(count int) []float64 {
	return Seq{NewRandom().WithLen(count)}.Array()
}

// RandomValuesWithAverage returns an array of random values with a given average.
func RandomValuesWithAverage(average float64, count int) []float64 {
	return Seq{NewRandom().WithAverage(average).WithLen(count)}.Array()
}

// NewRandom creates a new random sequence.
func NewRandom() *Random {
	return &Random{
		rnd: rand.New(rand.NewSource(time.Now().Unix())),
	}
}

// Random is a random number sequence generator.
type Random struct {
	rnd     *rand.Rand
	scale   *float64
	average *float64
	len     *int
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
	if r.average != nil && r.scale != nil {
		return *r.average + *r.scale - (r.rnd.Float64() * (2 * *r.scale))
	} else if r.scale != nil {
		return r.rnd.Float64() * *r.scale
	}
	return r.rnd.Float64()
}

// WithLen sets a maximum len
func (r *Random) WithLen(length int) *Random {
	r.len = &length
	return r
}

// WithScale sets the scale and returns the Random.
func (r *Random) WithScale(scale float64) *Random {
	r.scale = &scale
	return r
}

// WithAverage sets the average and returns the Random.
func (r *Random) WithAverage(average float64) *Random {
	r.average = &average
	return r
}
