package seq

import "time"

// NewArray creates a new array.
func NewArray(values ...float64) Array {
	return Array(values)
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

// ArrayOfTimes wraps an array of times as a sequence provider.
type ArrayOfTimes []time.Time

// Len returns the length of the array.
func (aot ArrayOfTimes) Len() int {
	return len(aot)
}

// GetValue returns the time at the given index as a time.Time.
func (aot ArrayOfTimes) GetValue(index int) time.Time {
	return aot[index]
}
