package seq

import (
	"sort"
	"time"
)

var (
	// TimeZero is the zero time.
	TimeZero = time.Time{}
)

// Times returns a new time sequence.
func Times(values ...time.Time) TimeSeq {
	return TimeSeq{TimeProvider: ArrayOfTimes(values)}
}

// TimeSeq is a sequence of times.
type TimeSeq struct {
	TimeProvider
}

// Array converts the sequence to times.
func (ts TimeSeq) Array() (output []time.Time) {
	slen := ts.Len()
	if slen == 0 {
		return
	}

	output = make([]time.Time, slen)
	for i := 0; i < slen; i++ {
		output[i] = ts.GetValue(i)
	}
	return
}

// Each applies the `mapfn` to all values in the value provider.
func (ts TimeSeq) Each(mapfn func(int, time.Time)) {
	for i := 0; i < ts.Len(); i++ {
		mapfn(i, ts.GetValue(i))
	}
}

// Map applies the `mapfn` to all values in the value provider,
// returning a new seq.
func (ts TimeSeq) Map(mapfn func(int, time.Time) time.Time) TimeSeq {
	output := make([]time.Time, ts.Len())
	for i := 0; i < ts.Len(); i++ {
		mapfn(i, ts.GetValue(i))
	}
	return TimeSeq{ArrayOfTimes(output)}
}

// FoldLeft collapses a seq from left to right.
func (ts TimeSeq) FoldLeft(mapfn func(i int, v0, v time.Time) time.Time) (v0 time.Time) {
	tslen := ts.Len()
	if tslen == 0 {
		return TimeZero
	}

	if tslen == 1 {
		return ts.GetValue(0)
	}

	v0 = ts.GetValue(0)
	for i := 1; i < tslen; i++ {
		v0 = mapfn(i, v0, ts.GetValue(i))
	}
	return
}

// FoldRight collapses a seq from right to left.
func (ts TimeSeq) FoldRight(mapfn func(i int, v0, v time.Time) time.Time) (v0 time.Time) {
	tslen := ts.Len()
	if tslen == 0 {
		return TimeZero
	}

	if tslen == 1 {
		return ts.GetValue(0)
	}

	v0 = ts.GetValue(tslen - 1)
	for i := tslen - 2; i >= 0; i-- {
		v0 = mapfn(i, v0, ts.GetValue(i))
	}
	return
}

// Sort returns the seq in ascending order.
func (ts TimeSeq) Sort() TimeSeq {
	if ts.Len() == 0 {
		return ts
	}

	values := ts.Array()
	sort.Slice(values, func(i, j int) bool {
		return values[i].Before(values[j])
	})
	return TimeSeq{TimeProvider: ArrayOfTimes(values)}
}

// SortDescending returns the seq in descending order.
func (ts TimeSeq) SortDescending() TimeSeq {
	if ts.Len() == 0 {
		return ts
	}

	values := ts.Array()
	sort.Slice(values, func(i, j int) bool {
		return values[i].After(values[j])
	})
	return TimeSeq{TimeProvider: ArrayOfTimes(values)}
}

// Min returns the minimum (or earliest) time in the sequence.
func (ts TimeSeq) Min() (min time.Time) {
	tslen := ts.Len()
	if tslen == 0 {
		return
	}
	min = ts.GetValue(0)
	var tv time.Time
	for i := 1; i < tslen; i++ {
		tv = ts.GetValue(i)
		if tv.Before(min) {
			min = tv
		}
	}
	return
}

// Start is an alias to `Min`.
func (ts TimeSeq) Start() time.Time {
	return ts.Min()
}

// Max returns the maximum (or latest) time in the sequence.
func (ts TimeSeq) Max() (max time.Time) {
	tslen := ts.Len()
	if tslen == 0 {
		return
	}
	max = ts.GetValue(0)
	var tv time.Time
	for i := 1; i < tslen; i++ {
		tv = ts.GetValue(i)
		if tv.After(max) {
			max = tv
		}
	}
	return
}

// End is an alias to `Max`.
func (ts TimeSeq) End() time.Time {
	return ts.Max()
}

// First returns the first value in the sequence.
func (ts TimeSeq) First() time.Time {
	if ts.Len() == 0 {
		return TimeZero
	}

	return ts.GetValue(0)
}

// Last returns the last value in the sequence.
func (ts TimeSeq) Last() time.Time {
	if ts.Len() == 0 {
		return TimeZero
	}

	return ts.GetValue(ts.Len() - 1)
}

// MinAndMax returns both the earliest and latest value from a sequence in one pass.
func (ts TimeSeq) MinAndMax() (min, max time.Time) {
	tslen := ts.Len()
	if tslen == 0 {
		return
	}
	min = ts.GetValue(0)
	max = ts.GetValue(0)
	var tv time.Time
	for i := 1; i < tslen; i++ {
		tv = ts.GetValue(i)
		if tv.Before(min) {
			min = tv
		}
		if tv.After(max) {
			max = tv
		}
	}
	return
}

// MapDistinct maps values given a map function to their distinct outputs.
func (ts TimeSeq) MapDistinct(mapFn func(time.Time) time.Time) TimeSeq {
	tslen := ts.Len()
	if tslen == 0 {
		return TimeSeq{}
	}

	var output []time.Time
	hourLookup := SetOfTime{}

	// add the initial value
	tv := ts.GetValue(0)
	tvh := mapFn(tv)
	hourLookup.Add(tvh)
	output = append(output, tvh)

	for i := 1; i < tslen; i++ {
		tv = ts.GetValue(i)
		tvh = mapFn(tv)
		if !hourLookup.Has(tvh) {
			hourLookup.Add(tvh)
			output = append(output, tvh)
		}
	}

	return TimeSeq{ArrayOfTimes(output)}
}

// Hours returns times in each distinct hour represented by the sequence.
func (ts TimeSeq) Hours() TimeSeq {
	return ts.MapDistinct(ts.trimToHour)
}

// Days returns times in each distinct day represented by the sequence.
func (ts TimeSeq) Days() TimeSeq {
	return ts.MapDistinct(ts.trimToDay)
}

// Months returns times in each distinct months represented by the sequence.
func (ts TimeSeq) Months() TimeSeq {
	return ts.MapDistinct(ts.trimToMonth)
}

// Years returns times in each distinc year represented by the sequence.
func (ts TimeSeq) Years() TimeSeq {
	return ts.MapDistinct(ts.trimToYear)
}

func (ts TimeSeq) trimToHour(tv time.Time) time.Time {
	return time.Date(tv.Year(), tv.Month(), tv.Day(), tv.Hour(), 0, 0, 0, tv.Location())
}

func (ts TimeSeq) trimToDay(tv time.Time) time.Time {
	return time.Date(tv.Year(), tv.Month(), tv.Day(), 0, 0, 0, 0, tv.Location())
}

func (ts TimeSeq) trimToMonth(tv time.Time) time.Time {
	return time.Date(tv.Year(), tv.Month(), 1, 0, 0, 0, 0, tv.Location())
}

func (ts TimeSeq) trimToYear(tv time.Time) time.Time {
	return time.Date(tv.Year(), 1, 1, 0, 0, 0, 0, tv.Location())
}
