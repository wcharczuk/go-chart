package chart

import "math"

// SequenceProvider is a provider for values for a sequence.
type SequenceProvider interface {
	Len() int
	GetValue(int) float64
}

// Sequence is a utility wrapper for sequence providers.
type Sequence struct {
	SequenceProvider
}

// Each applies the `mapfn` to all values in the value provider.
func (s Sequence) Each(mapfn func(int, float64)) {
	for i := 0; i < s.Len(); i++ {
		mapfn(i, s.GetValue(i))
	}
}

// Map applies the `mapfn` to all values in the value provider,
// returning a new sequence.
func (s Sequence) Map(mapfn func(i int, v float64) float64) Sequence {
	output := make([]float64, s.Len())
	for i := 0; i < s.Len(); i++ {
		mapfn(i, s.GetValue(i))
	}
	return Sequence{Array(output)}
}

// Average returns the float average of the values in the buffer.
func (s Sequence) Average() float64 {
	if s.Len() == 0 {
		return 0
	}

	var accum float64
	for i := 0; i < s.Len(); i++ {
		accum += s.GetValue(i)
	}
	return accum / float64(s.Len())
}

// Variance computes the variance of the buffer.
func (s Sequence) Variance() float64 {
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
func (s Sequence) StdDev() float64 {
	if s.Len() == 0 {
		return 0
	}

	return math.Pow(s.Variance(), 0.5)
}

// Array is a wrapper for an array of floats that implements `ValuesProvider`.
type Array []float64

// Len returns the value provider length.
func (a Array) Len() int {
	return len(a)
}

// GetValue returns the value at a given index.
func (a Array) GetValue(index int) float64 {
	return a[index]
}
