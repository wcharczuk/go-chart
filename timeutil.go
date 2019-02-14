package chart

import "time"

// SecondsPerXYZ
const (
	SecondsPerHour = 60 * 60
	SecondsPerDay  = 60 * 60 * 24
)

// TimeMillis returns a duration as a float millis.
func TimeMillis(d time.Duration) float64 {
	return float64(d) / float64(time.Millisecond)
}

// DiffHours returns the difference in hours between two times.
func DiffHours(t1, t2 time.Time) (hours int) {
	t1n := t1.Unix()
	t2n := t2.Unix()
	var diff int64
	if t1n > t2n {
		diff = t1n - t2n
	} else {
		diff = t2n - t1n
	}
	return int(diff / (SecondsPerHour))
}

// TimeMin returns the minimum and maximum times in a given range.
func TimeMin(times ...time.Time) (min time.Time) {
	if len(times) == 0 {
		return
	}
	min = times[0]
	for index := 1; index < len(times); index++ {
		if times[index].Before(min) {
			min = times[index]
		}

	}
	return
}

// TimeMax returns the minimum and maximum times in a given range.
func TimeMax(times ...time.Time) (max time.Time) {
	if len(times) == 0 {
		return
	}
	max = times[0]

	for index := 1; index < len(times); index++ {
		if times[index].After(max) {
			max = times[index]
		}
	}
	return
}

// TimeMinMax returns the minimum and maximum times in a given range.
func TimeMinMax(times ...time.Time) (min, max time.Time) {
	if len(times) == 0 {
		return
	}
	min = times[0]
	max = times[0]

	for index := 1; index < len(times); index++ {
		if times[index].Before(min) {
			min = times[index]
		}
		if times[index].After(max) {
			max = times[index]
		}
	}
	return
}

// TimeToFloat64 returns a float64 representation of a time.
func TimeToFloat64(t time.Time) float64 {
	return float64(t.UnixNano())
}

// TimeDescending sorts a given list of times ascending, or min to max.
type TimeDescending []time.Time

// Len implements sort.Sorter
func (d TimeDescending) Len() int { return len(d) }

// Swap implements sort.Sorter
func (d TimeDescending) Swap(i, j int) { d[i], d[j] = d[j], d[i] }

// Less implements sort.Sorter
func (d TimeDescending) Less(i, j int) bool { return d[i].After(d[j]) }

// TimeAscending sorts a given list of times ascending, or min to max.
type TimeAscending []time.Time

// Len implements sort.Sorter
func (a TimeAscending) Len() int { return len(a) }

// Swap implements sort.Sorter
func (a TimeAscending) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

// Less implements sort.Sorter
func (a TimeAscending) Less(i, j int) bool { return a[i].Before(a[j]) }
