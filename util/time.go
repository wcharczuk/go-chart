package util

import "time"

var (
	// Time contains time utility functions.
	Time = timeUtil{}
)

type timeUtil struct{}

// Millis returns the duration as milliseconds.
func (tu timeUtil) Millis(d time.Duration) float64 {
	return float64(d) / float64(time.Millisecond)
}

// TimeToFloat64 returns a float64 representation of a time.
func (tu timeUtil) ToFloat64(t time.Time) float64 {
	return float64(t.UnixNano())
}

// Float64ToTime returns a time from a float64.
func (tu timeUtil) FromFloat64(tf float64) time.Time {
	return time.Unix(0, int64(tf))
}

func (tu timeUtil) DiffDays(t1, t2 time.Time) (days int) {
	t1n := t1.Unix()
	t2n := t2.Unix()
	var diff int64
	if t1n > t2n {
		diff = t1n - t2n //yields seconds
	} else {
		diff = t2n - t1n //yields seconds
	}
	return int(diff / (_secondsPerDay))
}

func (tu timeUtil) DiffHours(t1, t2 time.Time) (hours int) {
	t1n := t1.Unix()
	t2n := t2.Unix()
	var diff int64
	if t1n > t2n {
		diff = t1n - t2n
	} else {
		diff = t2n - t1n
	}
	return int(diff / (_secondsPerHour))
}

// Start returns the earliest (min) time in a list of times.
func (tu timeUtil) Start(times ...time.Time) time.Time {
	if len(times) == 0 {
		return time.Time{}
	}

	start := times[0]
	for _, t := range times[1:] {
		if t.Before(start) {
			start = t
		}
	}
	return start
}

// Start returns the earliest (min) time in a list of times.
func (tu timeUtil) End(times ...time.Time) time.Time {
	if len(times) == 0 {
		return time.Time{}
	}

	end := times[0]
	for _, t := range times[1:] {
		if t.After(end) {
			end = t
		}
	}
	return end
}

// StartAndEnd returns the start and end of a given set of time in one pass.
func (tu timeUtil) StartAndEnd(values ...time.Time) (start time.Time, end time.Time) {
	if len(values) == 0 {
		return
	}

	start = values[0]
	end = values[0]

	for _, v := range values[1:] {
		if end.Before(v) {
			end = v
		}
		if start.After(v) {
			start = v
		}
	}
	return
}
