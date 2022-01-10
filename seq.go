package chart

import (
	"math"
	"sort"
)

// Seq is a utility wrapper for seq providers.
type Seq[A Number] struct {
	Sequence[A]
}

// Values enumerates the seq into a slice.
func (s Seq[A]) Values() (output []A) {
	if s.Len() == 0 {
		return
	}

	output = make([]A, s.Len())
	for i := 0; i < s.Len(); i++ {
		output[i] = s.GetValue(i)
	}
	return
}

// Each applies the `mapfn` to all values in the value provider.
func (s Seq[A]) Each(mapfn func(int, A)) {
	for i := 0; i < s.Len(); i++ {
		mapfn(i, s.GetValue(i))
	}
}

// Map applies the `mapfn` to all values in the value provider,
// returning a new seq.
func (s Seq[A]) Map(mapfn func(i int, v A) A) Seq[A] {
	output := make([]A, s.Len())
	for i := 0; i < s.Len(); i++ {
		mapfn(i, s.GetValue(i))
	}
	return Seq[A]{Sequence: Array(output...)}
}

// FoldLeft collapses a seq from left to right.
func (s Seq[A]) FoldLeft(mapfn func(i int, v0, v A) A) (output A) {
	if s.Len() == 0 {
		return
	}

	if s.Len() == 1 {
		return s.GetValue(0)
	}

	output = s.GetValue(0)
	for i := 1; i < s.Len(); i++ {
		output = mapfn(i, output, s.GetValue(i))
	}
	return
}

// FoldRight collapses a seq from right to left.
func (s Seq[A]) FoldRight(mapfn func(i int, v0, v A) A) (output A) {
	if s.Len() == 0 {
		return
	}

	if s.Len() == 1 {
		return s.GetValue(0)
	}

	output = s.GetValue(s.Len() - 1)
	for i := s.Len() - 2; i >= 0; i-- {
		output = mapfn(i, output, s.GetValue(i))
	}
	return
}

// Min returns the minimum value in the seq.
func (s Seq[A]) Min() (min A) {
	if s.Len() == 0 {
		return
	}

	min = s.GetValue(0)
	var value A
	for i := 1; i < s.Len(); i++ {
		value = s.GetValue(i)
		if value < min {
			min = value
		}
	}
	return min
}

// Max returns the maximum value in the seq.
func (s Seq[A]) Max() (max A) {
	if s.Len() == 0 {
		return
	}
	max = s.GetValue(0)
	var value A
	for i := 1; i < s.Len(); i++ {
		value = s.GetValue(i)
		if value > max {
			max = value
		}
	}
	return max
}

// MinMax returns the minimum and the maximum in one pass.
func (s Seq[A]) MinMax() (min, max A) {
	if s.Len() == 0 {
		return
	}
	min = s.GetValue(0)
	max = min
	var value A
	for i := 1; i < s.Len(); i++ {
		value = s.GetValue(i)
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}
	return
}

// Sort returns the seq sorted in ascending order.
// This fully enumerates the seq.
func (s Seq[A]) Sort() Seq[A] {
	if s.Len() == 0 {
		return s
	}
	values := s.Values()
	sort.Slice(values, func(i, j int) bool {
		return values[i] < values[j]
	})
	return Seq[A]{Array(values...)}
}

// Reverse reverses the sequence
func (s Seq[A]) Reverse() Seq[A] {
	if s.Len() == 0 {
		return s
	}

	values := s.Values()
	valuesLen := len(values)
	valuesLen1 := len(values) - 1
	valuesLen2 := valuesLen >> 1
	var i, j A
	for index := 0; index < valuesLen2; index++ {
		i = values[index]
		j = values[valuesLen1-index]
		values[index] = j
		values[valuesLen1-index] = i
	}

	return Seq[A]{Array(values...)}
}

// Sum adds all the elements of a series together.
func (s Seq[A]) Sum() (accum A) {
	if s.Len() == 0 {
		return
	}

	for i := 0; i < s.Len(); i++ {
		accum += s.GetValue(i)
	}
	return
}

// Average returns the float average of the values in the buffer.
func (s Seq[A]) Average() (avg float64) {
	if s.Len() == 0 {
		return
	}

	avg = float64(s.Sum()) / float64((s.Len()))
	return
}

// Variance computes the variance of the buffer.
func (s Seq[A]) Variance() (variance float64) {
	if s.Len() == 0 {
		return 0
	}

	m := s.Average()
	var v float64
	for i := 0; i < s.Len(); i++ {
		v = float64(s.GetValue(i))
		variance += (v - m) * (v - m)
	}

	return variance / float64(s.Len())
}

// StdDev returns the standard deviation.
func (s Seq[A]) StdDev() float64 {
	if s.Len() == 0 {
		return 0
	}

	return math.Pow(float64(s.Variance()), 0.5)
}

//Percentile finds the relative standing in a slice of floats.
// `percent` should be given on the interval [0,1.0).
func (s Seq[A]) Percentile(percent float64) (percentile A) {
	l := s.Len()
	if l == 0 {
		return 0
	}

	if percent < 0 || percent > 1.0 {
		panic("percentile percent out of range [0.0, 1.0)")
	}

	sorted := s.Sort()
	index := percent * float64(l)
	if index == float64(int64(index)) {
		i := f64i(index)
		ci := sorted.GetValue(i - 1)
		c := sorted.GetValue(i)
		percentile = (ci + c) / 2.0
	} else {
		i := f64i(index)
		percentile = sorted.GetValue(i)
	}

	return percentile
}

// Normalize maps every value to the interval [0, 1.0].
func (s Seq[A]) Normalize() Seq[float64] {
	min, max := s.MinMax()

	delta := float64(max - min)
	output := make([]float64, s.Len())
	for i := 0; i < s.Len(); i++ {
		output[i] = (float64(s.GetValue(i)) - float64(min)) / delta
	}

	return Seq[float64]{Array(output...)}
}
