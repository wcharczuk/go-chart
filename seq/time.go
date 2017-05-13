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

func (ts timeSequence) MarketHours(from, to time.Time, marketOpen, marketClose time.Time, isHoliday util.HolidayProvider) []time.Time {
	var times []time.Time
	cursor := util.Date.On(marketOpen, from)
	toClose := util.Date.On(marketClose, to)
	for cursor.Before(toClose) || cursor.Equal(toClose) {
		todayOpen := util.Date.On(marketOpen, cursor)
		todayClose := util.Date.On(marketClose, cursor)
		isValidTradingDay := !isHoliday(cursor) && util.Date.IsWeekDay(cursor.Weekday())

		if (cursor.Equal(todayOpen) || cursor.After(todayOpen)) && (cursor.Equal(todayClose) || cursor.Before(todayClose)) && isValidTradingDay {
			times = append(times, cursor)
		}
		if cursor.After(todayClose) {
			cursor = util.Date.NextMarketOpen(cursor, marketOpen, isHoliday)
		} else {
			cursor = util.Date.NextHour(cursor)
		}
	}
	return times
}

func (ts timeSequence) MarketHourQuarters(from, to time.Time, marketOpen, marketClose time.Time, isHoliday util.HolidayProvider) []time.Time {
	var times []time.Time
	cursor := util.Date.On(marketOpen, from)
	toClose := util.Date.On(marketClose, to)
	for cursor.Before(toClose) || cursor.Equal(toClose) {

		isValidTradingDay := !isHoliday(cursor) && util.Date.IsWeekDay(cursor.Weekday())

		if isValidTradingDay {
			todayOpen := util.Date.On(marketOpen, cursor)
			todayNoon := util.Date.NoonOn(cursor)
			today2pm := util.Date.On(util.Date.Time(14, 0, 0, 0, cursor.Location()), cursor)
			todayClose := util.Date.On(marketClose, cursor)
			times = append(times, todayOpen, todayNoon, today2pm, todayClose)
		}

		cursor = util.Date.NextDay(cursor)
	}
	return times
}

func (ts timeSequence) MarketDayCloses(from, to time.Time, marketOpen, marketClose time.Time, isHoliday util.HolidayProvider) []time.Time {
	var times []time.Time
	cursor := util.Date.On(marketOpen, from)
	toClose := util.Date.On(marketClose, to)
	for cursor.Before(toClose) || cursor.Equal(toClose) {
		isValidTradingDay := !isHoliday(cursor) && util.Date.IsWeekDay(cursor.Weekday())
		if isValidTradingDay {
			todayClose := util.Date.On(marketClose, cursor)
			times = append(times, todayClose)
		}

		cursor = util.Date.NextDay(cursor)
	}
	return times
}

func (ts timeSequence) MarketDayAlternateCloses(from, to time.Time, marketOpen, marketClose time.Time, isHoliday util.HolidayProvider) []time.Time {
	var times []time.Time
	cursor := util.Date.On(marketOpen, from)
	toClose := util.Date.On(marketClose, to)
	for cursor.Before(toClose) || cursor.Equal(toClose) {
		isValidTradingDay := !isHoliday(cursor) && util.Date.IsWeekDay(cursor.Weekday())
		if isValidTradingDay {
			todayClose := util.Date.On(marketClose, cursor)
			times = append(times, todayClose)
		}

		cursor = cursor.AddDate(0, 0, 2)
	}
	return times
}

func (ts timeSequence) MarketDayMondayCloses(from, to time.Time, marketOpen, marketClose time.Time, isHoliday util.HolidayProvider) []time.Time {
	var times []time.Time
	cursor := util.Date.On(marketClose, from)
	toClose := util.Date.On(marketClose, to)

	for cursor.Equal(toClose) || cursor.Before(toClose) {
		isValidTradingDay := !isHoliday(cursor) && util.Date.IsWeekDay(cursor.Weekday())
		if isValidTradingDay {
			times = append(times, cursor)
		}
		cursor = util.Date.NextDayOfWeek(cursor, time.Monday)
	}
	return times
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
	start := Time.Start(xdata)
	end := Time.End(xdata)

	totalHours := util.Math.AbsInt(util.Date.DiffHours(start, end))

	finalTimes := ts.Hours(start, totalHours+1)
	finalValues := make([]float64, totalHours+1)

	var hoursFromStart int
	for i, xd := range xdata {
		hoursFromStart = util.Date.DiffHours(start, xd)
		finalValues[hoursFromStart] = ydata[i]
	}

	return finalTimes, finalValues
}

// Start returns the earliest (min) time in a list of times.
func (ts timeSequence) Start(times []time.Time) time.Time {
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
func (ts timeSequence) End(times []time.Time) time.Time {
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
