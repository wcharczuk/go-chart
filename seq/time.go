package seq

import (
	"time"

	"github.com/wcharczuk/go-chart/util"
)

// Time is a utility singleton with helper functions for time seq generation.
var Time timeSequence

type timeSequence struct{}

// Days generates a seq of timestamps by day, from -days to today.
func (ts timeSequence) Days(days int) []time.Time {
	var values []time.Time
	for day := days; day >= 0; day-- {
		values = append(values, time.Now().AddDate(0, 0, -day))
	}
	return values
}

func (ts timeSequence) Hours(start time.Time, totalHours int) []time.Time {
	times := make([]time.Time, totalHours)

	last := start
	for i := 0; i < totalHours; i++ {
		times[i] = last
		last = last.Add(time.Hour)
	}

	return times
}

// HoursFilled adds zero values for the data bounded by the start and end of the xdata array.
func (ts timeSequence) HoursFilled(xdata []time.Time, ydata []float64) ([]time.Time, []float64) {
	start, end := util.Time.StartAndEnd(xdata...)
	totalHours := util.Time.DiffHours(start, end)

	finalTimes := ts.Hours(start, totalHours+1)
	finalValues := make([]float64, totalHours+1)

	var hoursFromStart int
	for i, xd := range xdata {
		hoursFromStart = util.Time.DiffHours(start, xd)
		finalValues[hoursFromStart] = ydata[i]
	}

	return finalTimes, finalValues
}
