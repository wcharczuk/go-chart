package chart

import (
	"math/rand"
	"time"
)

var (
	// Sequence contains some sequence utilities.
	// These utilities can be useful for generating test data.
	Sequence = &sequence{
		rnd: rand.New(rand.NewSource(time.Now().Unix())),
	}
)

type sequence struct {
	rnd *rand.Rand
}

// Float64 produces an array of floats from [start,end] by optional steps.
func (s sequence) Float64(start, end float64, steps ...float64) []float64 {
	var values []float64
	step := 1.0
	if len(steps) > 0 {
		step = steps[0]
	}

	if start < end {
		for x := start; x <= end; x += step {
			values = append(values, x)
		}
	} else {
		for x := start; x >= end; x = x - step {
			values = append(values, x)
		}
	}
	return values
}

// Random generates a fixed length sequence of random values between (0, scale).
func (s sequence) Random(samples int, scale float64) []float64 {
	values := make([]float64, samples)

	for x := 0; x < samples; x++ {
		values[x] = s.rnd.Float64() * scale
	}

	return values
}

// Random generates a fixed length sequence of random values with a given average, above and below that average by (-scale, scale)
func (s sequence) RandomWithAverage(samples int, average, scale float64) []float64 {
	values := make([]float64, samples)

	for x := 0; x < samples; x++ {
		jitter := scale - (s.rnd.Float64() * (2 * scale))
		values[x] = average + jitter
	}

	return values
}

// Days generates a sequence of timestamps by day, from -days to today.
func (s sequence) Days(days int) []time.Time {
	var values []time.Time
	for day := days; day >= 0; day-- {
		values = append(values, time.Now().AddDate(0, 0, -day))
	}
	return values
}

func (s sequence) MarketHours(from, to time.Time, marketOpen, marketClose time.Time, isHoliday HolidayProvider) []time.Time {
	var times []time.Time
	cursor := Date.On(marketOpen, from)
	toClose := Date.On(marketClose, to)
	for cursor.Before(toClose) || cursor.Equal(toClose) {
		todayOpen := Date.On(marketOpen, cursor)
		todayClose := Date.On(marketClose, cursor)
		isValidTradingDay := !isHoliday(cursor) && Date.IsWeekDay(cursor.Weekday())

		if (cursor.Equal(todayOpen) || cursor.After(todayOpen)) && (cursor.Equal(todayClose) || cursor.Before(todayClose)) && isValidTradingDay {
			times = append(times, cursor)
		}
		if cursor.After(todayClose) {
			cursor = Date.NextMarketOpen(cursor, marketOpen, isHoliday)
		} else {
			cursor = Date.NextHour(cursor)
		}
	}
	return times
}

func (s sequence) MarketHourQuarters(from, to time.Time, marketOpen, marketClose time.Time, isHoliday HolidayProvider) []time.Time {
	var times []time.Time
	cursor := Date.On(marketOpen, from)
	toClose := Date.On(marketClose, to)
	for cursor.Before(toClose) || cursor.Equal(toClose) {

		isValidTradingDay := !isHoliday(cursor) && Date.IsWeekDay(cursor.Weekday())

		if isValidTradingDay {
			todayOpen := Date.On(marketOpen, cursor)
			todayNoon := Date.NoonOn(cursor)
			today2pm := Date.On(Date.Time(14, 0, 0, 0, cursor.Location()), cursor)
			todayClose := Date.On(marketClose, cursor)
			times = append(times, todayOpen, todayNoon, today2pm, todayClose)
		}

		cursor = Date.NextDay(cursor)
	}
	return times
}

func (s sequence) MarketDayCloses(from, to time.Time, marketOpen, marketClose time.Time, isHoliday HolidayProvider) []time.Time {
	var times []time.Time
	cursor := Date.On(marketOpen, from)
	toClose := Date.On(marketClose, to)
	for cursor.Before(toClose) || cursor.Equal(toClose) {
		isValidTradingDay := !isHoliday(cursor) && Date.IsWeekDay(cursor.Weekday())
		if isValidTradingDay {
			todayClose := Date.On(marketClose, cursor)
			times = append(times, todayClose)
		}

		cursor = Date.NextDay(cursor)
	}
	return times
}

func (s sequence) MarketDayAlternateCloses(from, to time.Time, marketOpen, marketClose time.Time, isHoliday HolidayProvider) []time.Time {
	var times []time.Time
	cursor := Date.On(marketOpen, from)
	toClose := Date.On(marketClose, to)
	for cursor.Before(toClose) || cursor.Equal(toClose) {
		isValidTradingDay := !isHoliday(cursor) && Date.IsWeekDay(cursor.Weekday())
		if isValidTradingDay {
			todayClose := Date.On(marketClose, cursor)
			times = append(times, todayClose)
		}

		cursor = cursor.AddDate(0, 0, 2)
	}
	return times
}

func (s sequence) MarketDayMondayCloses(from, to time.Time, marketOpen, marketClose time.Time, isHoliday HolidayProvider) []time.Time {
	var times []time.Time
	cursor := Date.On(marketClose, from)
	toClose := Date.On(marketClose, to)

	for cursor.Equal(toClose) || cursor.Before(toClose) {
		isValidTradingDay := !isHoliday(cursor) && Date.IsWeekDay(cursor.Weekday())
		if isValidTradingDay {
			times = append(times, cursor)
		}
		cursor = Date.NextDayOfWeek(cursor, time.Monday)
	}
	return times
}

func (s sequence) Hours(start time.Time, totalHours int) []time.Time {
	times := make([]time.Time, totalHours)

	last := start
	for i := 0; i < totalHours; i++ {
		times[i] = last
		last = last.Add(time.Hour)
	}

	return times
}

// HoursFill adds zero values for the data bounded by the start and end of the xdata array.
func (s sequence) HoursFill(xdata []time.Time, ydata []float64) ([]time.Time, []float64) {
	start := Date.Start(xdata)
	end := Date.End(xdata)

	totalHours := Math.AbsInt(Date.DiffHours(start, end))

	finalTimes := s.Hours(start, totalHours+1)
	finalValues := make([]float64, totalHours+1)

	var hoursFromStart int
	for i, xd := range xdata {
		hoursFromStart = Date.DiffHours(start, xd)
		finalValues[hoursFromStart] = ydata[i]
	}

	return finalTimes, finalValues
}
