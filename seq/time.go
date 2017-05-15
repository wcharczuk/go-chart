package seq

import (
	"time"

	"github.com/wcharczuk/go-chart/util"
)

// TimeUtil is a utility singleton with helper functions for time seq generation.
var TimeUtil timeUtil

type timeUtil struct{}

func (tu timeUtil) MarketHours(from, to time.Time, marketOpen, marketClose time.Time, isHoliday util.HolidayProvider) []time.Time {
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

func (tu timeUtil) MarketHourQuarters(from, to time.Time, marketOpen, marketClose time.Time, isHoliday util.HolidayProvider) []time.Time {
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

func (tu timeUtil) MarketDayCloses(from, to time.Time, marketOpen, marketClose time.Time, isHoliday util.HolidayProvider) []time.Time {
	var times []time.Time
	cursor := util.Date.On(marketOpen, from)
	toClose := util.Date.On(marketClose, to)
	for cursor.Before(toClose) || cursor.Equal(toClose) {
		isValidTradingDay := !isHoliday(cursor) && util.Date.IsWeekDay(cursor.Weekday())
		if isValidTradingDay {
			newValue := util.Date.NoonOn(cursor)
			times = append(times, newValue)
		}

		cursor = util.Date.NextDay(cursor)
	}
	return times
}

func (tu timeUtil) MarketDayAlternateCloses(from, to time.Time, marketOpen, marketClose time.Time, isHoliday util.HolidayProvider) []time.Time {
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

func (tu timeUtil) MarketDayMondayCloses(from, to time.Time, marketOpen, marketClose time.Time, isHoliday util.HolidayProvider) []time.Time {
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

func (tu timeUtil) Hours(start time.Time, totalHours int) []time.Time {
	times := make([]time.Time, totalHours)

	last := start
	for i := 0; i < totalHours; i++ {
		times[i] = last
		last = last.Add(time.Hour)
	}

	return times
}

// HoursFilled adds zero values for the data bounded by the start and end of the xdata array.
func (tu timeUtil) HoursFilled(xdata []time.Time, ydata []float64) ([]time.Time, []float64) {
	start, end := Times(xdata...).MinAndMax()

	totalHours := util.Math.AbsInt(util.Date.DiffHours(start, end))

	finalTimes := tu.Hours(start, totalHours+1)
	finalValues := make([]float64, totalHours+1)

	var hoursFromStart int
	for i, xd := range xdata {
		hoursFromStart = util.Date.DiffHours(start, xd)
		finalValues[hoursFromStart] = ydata[i]
	}

	return finalTimes, finalValues
}
