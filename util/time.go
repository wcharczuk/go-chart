package util

import "time"

var (
	// Time contains time utility functions.
	Time = timeUtil{}
)

type timeUtil struct{}

// TimeToFloat64 returns a float64 representation of a time.
func (tu timeUtil) ToFloat64(t time.Time) float64 {
	return float64(t.UnixNano())
}

// Float64ToTime returns a time from a float64.
func (tu timeUtil) FromFloat64(tf float64) time.Time {
	return time.Unix(0, int64(tf))
}
