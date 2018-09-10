package seq

import (
	"time"

	"github.com/wcharczuk/go-chart/util"
)

// Assert types implement interfaces.
var (
	_ Provider = (*Times)(nil)
)

// Times are an array of times.
// It wraps the array with methods that implement `seq.Provider`.
type Times []time.Time

// Array returns the times to an array.
func (t Times) Array() []time.Time {
	return []time.Time(t)
}

// Len returns the length of the array.
func (t Times) Len() int {
	return len(t)
}

// GetValue returns a value at an index as a time.
func (t Times) GetValue(index int) float64 {
	return util.Time.ToFloat64(t[index])
}
