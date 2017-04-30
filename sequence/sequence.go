package sequence

import "math"

// Provider is a provider for values for a sequence.
type Provider interface {
	Len() int
	GetValue(int) float64
}

// Seq is a utility wrapper for sequence providers.
type Seq struct {
	Provider
}

// Array enumerates the sequence into a slice.
func (s Seq) Array() (output []float64) {
	output = make([]float64, s.Len())
	for i := 0; i < s.Len(); i++ {
		output[i] = s.GetValue(i)
	}
	return
}

// Each applies the `mapfn` to all values in the value provider.
func (s Seq) Each(mapfn func(int, float64)) {
	for i := 0; i < s.Len(); i++ {
		mapfn(i, s.GetValue(i))
	}
}

// Map applies the `mapfn` to all values in the value provider,
// returning a new sequence.
func (s Seq) Map(mapfn func(i int, v float64) float64) Seq {
	output := make([]float64, s.Len())
	for i := 0; i < s.Len(); i++ {
		mapfn(i, s.GetValue(i))
	}
	return Seq{Array(output)}
}

// FoldLeft collapses a sequence from left to right.
func (s Seq) FoldLeft(mapfn func(i int, v0, v float64) float64) (v0 float64) {
	if s.Len() == 0 {
		return 0
	}

	if s.Len() == 1 {
		return s.GetValue(0)
	}

	v0 = s.GetValue(0)
	for i := 1; i < s.Len(); i++ {
		v0 = mapfn(i, v0, s.GetValue(i))
	}
	return
}

// FoldRight collapses a sequence from right to left.
func (s Seq) FoldRight(mapfn func(i int, v0, v float64) float64) (v0 float64) {
	if s.Len() == 0 {
		return 0
	}

	if s.Len() == 1 {
		return s.GetValue(0)
	}

	v0 = s.GetValue(s.Len() - 1)
	for i := s.Len() - 2; i >= 0; i-- {
		v0 = mapfn(i, v0, s.GetValue(i))
	}
	return
}

// Sum adds all the elements of a series together.
func (s Seq) Sum() (accum float64) {
	if s.Len() == 0 {
		return 0
	}

	for i := 0; i < s.Len(); i++ {
		accum += s.GetValue(i)
	}
	return
}

// Average returns the float average of the values in the buffer.
func (s Seq) Average() float64 {
	if s.Len() == 0 {
		return 0
	}

	return s.Sum() / float64(s.Len())
}

// Variance computes the variance of the buffer.
func (s Seq) Variance() float64 {
	if s.Len() == 0 {
		return 0
	}

	m := s.Average()
	var variance, v float64
	for i := 0; i < s.Len(); i++ {
		v = s.GetValue(i)
		variance += (v - m) * (v - m)
	}

	return variance / float64(s.Len())
}

// StdDev returns the standard deviation.
func (s Seq) StdDev() float64 {
	if s.Len() == 0 {
		return 0
	}

	return math.Pow(s.Variance(), 0.5)
}
