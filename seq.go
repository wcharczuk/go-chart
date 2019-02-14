package chart

import (
	"math"
	"sort"
	"time"

	"github.com/blend/go-sdk/timeutil"
)

// NewSeq wraps a provider with a seq.
func NewSeq(provider SeqProvider) Seq {
	return Seq{SeqProvider: provider}
}

// Seq is a utility wrapper for seq providers.
type Seq struct {
	SeqProvider
}

// Values enumerates the seq into a slice.
func (s Seq) Values() (output []float64) {
	if s.Len() == 0 {
		return
	}

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
// returning a new seq.
func (s Seq) Map(mapfn func(i int, v float64) float64) Seq {
	output := make([]float64, s.Len())
	for i := 0; i < s.Len(); i++ {
		mapfn(i, s.GetValue(i))
	}
	return Seq{SeqArray(output)}
}

// FoldLeft collapses a seq from left to right.
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

// FoldRight collapses a seq from right to left.
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

// Min returns the minimum value in the seq.
func (s Seq) Min() float64 {
	if s.Len() == 0 {
		return 0
	}
	min := s.GetValue(0)
	var value float64
	for i := 1; i < s.Len(); i++ {
		value = s.GetValue(i)
		if value < min {
			min = value
		}
	}
	return min
}

// Max returns the maximum value in the seq.
func (s Seq) Max() float64 {
	if s.Len() == 0 {
		return 0
	}
	max := s.GetValue(0)
	var value float64
	for i := 1; i < s.Len(); i++ {
		value = s.GetValue(i)
		if value > max {
			max = value
		}
	}
	return max
}

// MinMax returns the minimum and the maximum in one pass.
func (s Seq) MinMax() (min, max float64) {
	if s.Len() == 0 {
		return
	}
	min = s.GetValue(0)
	max = min
	var value float64
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
func (s Seq) Sort() Seq {
	if s.Len() == 0 {
		return s
	}
	values := s.Values()
	sort.Float64s(values)
	return Seq{SeqArray(values)}
}

// Median returns the median or middle value in the sorted seq.
func (s Seq) Median() (median float64) {
	l := s.Len()
	if l == 0 {
		return
	}

	sorted := s.Sort()
	if l%2 == 0 {
		v0 := sorted.GetValue(l/2 - 1)
		v1 := sorted.GetValue(l/2 + 1)
		median = (v0 + v1) / 2
	} else {
		median = float64(sorted.GetValue(l << 1))
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

//Percentile finds the relative standing in a slice of floats.
// `percent` should be given on the interval [0,1.0).
func (s Seq) Percentile(percent float64) (percentile float64) {
	l := s.Len()
	if l == 0 {
		return 0
	}

	if percent < 0 || percent > 1.0 {
		panic("percent out of range [0.0, 1.0)")
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
func (s Seq) Normalize() Seq {
	min, max := s.MinMax()

	delta := max - min
	output := make([]float64, s.Len())
	for i := 0; i < s.Len(); i++ {
		output[i] = (s.GetValue(i) - min) / delta
	}

	return Seq{SeqProvider: SeqArray(output)}
}

// SeqProvider is a provider for values for a seq.
type SeqProvider interface {
	Len() int
	GetValue(int) float64
}

// SeqArray is a wrapper for an array of floats that implements `ValuesProvider`.
type SeqArray []float64

// Len returns the value provider length.
func (a SeqArray) Len() int {
	return len(a)
}

// GetValue returns the value at a given index.
func (a SeqArray) GetValue(index int) float64 {
	return a[index]
}

// SeqDays generates a seq of timestamps by day, from -days to today.
func SeqDays(days int) []time.Time {
	var values []time.Time
	for day := days; day >= 0; day-- {
		values = append(values, time.Now().AddDate(0, 0, -day))
	}
	return values
}

// SeqHours returns a sequence of times by the hour for a given number of hours
// after a given start.
func SeqHours(start time.Time, totalHours int) []time.Time {
	times := make([]time.Time, totalHours)

	last := start
	for i := 0; i < totalHours; i++ {
		times[i] = last
		last = last.Add(time.Hour)
	}

	return times
}

// SeqHoursFilled adds zero values for the data bounded by the start and end of the xdata array.
func SeqHoursFilled(xdata []time.Time, ydata []float64) ([]time.Time, []float64) {
	start, end := TimeMinMax(xdata...)
	totalHours := DiffHours(start, end)

	finalTimes := SeqHours(start, totalHours+1)
	finalValues := make([]float64, totalHours+1)

	var hoursFromStart int
	for i, xd := range xdata {
		hoursFromStart = DiffHours(start, xd)
		finalValues[hoursFromStart] = ydata[i]
	}

	return finalTimes, finalValues
}

// Assert types implement interfaces.
var (
	_ SeqProvider = (*SeqTimes)(nil)
)

// SeqTimes are an array of times.
// It wraps the array with methods that implement `seq.Provider`.
type SeqTimes []time.Time

// Array returns the times to an array.
func (t SeqTimes) Array() []time.Time {
	return []time.Time(t)
}

// Len returns the length of the array.
func (t SeqTimes) Len() int {
	return len(t)
}

// GetValue returns a value at an index as a time.
func (t SeqTimes) GetValue(index int) float64 {
	return timeutil.ToFloat64(t[index])
}

// SeqRange returns the array values of a linear seq with a given start, end and optional step.
func SeqRange(start, end float64) []float64 {
	return Seq{NewSeqLinear().WithStart(start).WithEnd(end).WithStep(1.0)}.Values()
}

// SeqRangeWithStep returns the array values of a linear seq with a given start, end and optional step.
func SeqRangeWithStep(start, end, step float64) []float64 {
	return Seq{NewSeqLinear().WithStart(start).WithEnd(end).WithStep(step)}.Values()
}

// NewSeqLinear returns a new linear generator.
func NewSeqLinear() *SeqLinear {
	return &SeqLinear{step: 1.0}
}

// SeqLinear is a stepwise generator.
type SeqLinear struct {
	start float64
	end   float64
	step  float64
}

// Start returns the start value.
func (lg SeqLinear) Start() float64 {
	return lg.start
}

// End returns the end value.
func (lg SeqLinear) End() float64 {
	return lg.end
}

// Step returns the step value.
func (lg SeqLinear) Step() float64 {
	return lg.step
}

// Len returns the number of elements in the seq.
func (lg SeqLinear) Len() int {
	if lg.start < lg.end {
		return int((lg.end-lg.start)/lg.step) + 1
	}
	return int((lg.start-lg.end)/lg.step) + 1
}

// GetValue returns the value at a given index.
func (lg SeqLinear) GetValue(index int) float64 {
	fi := float64(index)
	if lg.start < lg.end {
		return lg.start + (fi * lg.step)
	}
	return lg.start - (fi * lg.step)
}

// WithStart sets the start and returns the linear generator.
func (lg *SeqLinear) WithStart(start float64) *SeqLinear {
	lg.start = start
	return lg
}

// WithEnd sets the end and returns the linear generator.
func (lg *SeqLinear) WithEnd(end float64) *SeqLinear {
	lg.end = end
	return lg
}

// WithStep sets the step and returns the linear generator.
func (lg *SeqLinear) WithStep(step float64) *SeqLinear {
	lg.step = step
	return lg
}
