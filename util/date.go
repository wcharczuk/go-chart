package util

import (
	"sync"
	"time"
)

const (
	// AllDaysMask is a bitmask of all the days of the week.
	AllDaysMask = 1<<uint(time.Sunday) | 1<<uint(time.Monday) | 1<<uint(time.Tuesday) | 1<<uint(time.Wednesday) | 1<<uint(time.Thursday) | 1<<uint(time.Friday) | 1<<uint(time.Saturday)
	// WeekDaysMask is a bitmask of all the weekdays of the week.
	WeekDaysMask = 1<<uint(time.Monday) | 1<<uint(time.Tuesday) | 1<<uint(time.Wednesday) | 1<<uint(time.Thursday) | 1<<uint(time.Friday)
	//WeekendDaysMask is a bitmask of the weekend days of the week.
	WeekendDaysMask = 1<<uint(time.Sunday) | 1<<uint(time.Saturday)
)

var (
	// DaysOfWeek are all the time.Weekday in an array for utility purposes.
	DaysOfWeek = []time.Weekday{
		time.Sunday,
		time.Monday,
		time.Tuesday,
		time.Wednesday,
		time.Thursday,
		time.Friday,
		time.Saturday,
	}

	// WeekDays are the business time.Weekday in an array.
	WeekDays = []time.Weekday{
		time.Monday,
		time.Tuesday,
		time.Wednesday,
		time.Thursday,
		time.Friday,
	}

	// WeekendDays are the weekend time.Weekday in an array.
	WeekendDays = []time.Weekday{
		time.Sunday,
		time.Saturday,
	}

	//Epoch is unix epoc saved for utility purposes.
	Epoch = time.Unix(0, 0)
)

var (
	_easternLock sync.Mutex
	_eastern     *time.Location
)

// NYSEOpen is when the NYSE opens.
func NYSEOpen() time.Time { return Date.Time(9, 30, 0, 0, Date.Eastern()) }

// NYSEClose is when the NYSE closes.
func NYSEClose() time.Time { return Date.Time(16, 0, 0, 0, Date.Eastern()) }

// NASDAQOpen is when NASDAQ opens.
func NASDAQOpen() time.Time { return Date.Time(9, 30, 0, 0, Date.Eastern()) }

// NASDAQClose is when NASDAQ closes.
func NASDAQClose() time.Time { return Date.Time(16, 0, 0, 0, Date.Eastern()) }

// NYSEArcaOpen is when NYSEARCA opens.
func NYSEArcaOpen() time.Time { return Date.Time(4, 0, 0, 0, Date.Eastern()) }

// NYSEArcaClose is when NYSEARCA closes.
func NYSEArcaClose() time.Time { return Date.Time(20, 0, 0, 0, Date.Eastern()) }

// HolidayProvider is a function that returns if a given time falls on a holiday.
type HolidayProvider func(time.Time) bool

// defaultHolidayProvider implements `HolidayProvider` and just returns false.
func defaultHolidayProvider(_ time.Time) bool { return false }

var (
	// Date contains utility functions that operate on dates.
	Date = &date{}
)

type date struct{}

// IsNYSEHoliday returns if a date was/is on a nyse holiday day.
func (d date) IsNYSEHoliday(t time.Time) bool {
	te := t.In(d.Eastern())
	if te.Year() == 2013 {
		if te.Month() == 1 {
			return te.Day() == 1 || te.Day() == 21
		} else if te.Month() == 2 {
			return te.Day() == 18
		} else if te.Month() == 3 {
			return te.Day() == 29
		} else if te.Month() == 5 {
			return te.Day() == 27
		} else if te.Month() == 7 {
			return te.Day() == 4
		} else if te.Month() == 9 {
			return te.Day() == 2
		} else if te.Month() == 11 {
			return te.Day() == 28
		} else if te.Month() == 12 {
			return te.Day() == 25
		}
	} else if te.Year() == 2014 {
		if te.Month() == 1 {
			return te.Day() == 1 || te.Day() == 20
		} else if te.Month() == 2 {
			return te.Day() == 17
		} else if te.Month() == 4 {
			return te.Day() == 18
		} else if te.Month() == 5 {
			return te.Day() == 26
		} else if te.Month() == 7 {
			return te.Day() == 4
		} else if te.Month() == 9 {
			return te.Day() == 1
		} else if te.Month() == 11 {
			return te.Day() == 27
		} else if te.Month() == 12 {
			return te.Day() == 25
		}
	} else if te.Year() == 2015 {
		if te.Month() == 1 {
			return te.Day() == 1 || te.Day() == 19
		} else if te.Month() == 2 {
			return te.Day() == 16
		} else if te.Month() == 4 {
			return te.Day() == 3
		} else if te.Month() == 5 {
			return te.Day() == 25
		} else if te.Month() == 7 {
			return te.Day() == 3
		} else if te.Month() == 9 {
			return te.Day() == 7
		} else if te.Month() == 11 {
			return te.Day() == 26
		} else if te.Month() == 12 {
			return te.Day() == 25
		}
	} else if te.Year() == 2016 {
		if te.Month() == 1 {
			return te.Day() == 1 || te.Day() == 18
		} else if te.Month() == 2 {
			return te.Day() == 15
		} else if te.Month() == 3 {
			return te.Day() == 25
		} else if te.Month() == 5 {
			return te.Day() == 30
		} else if te.Month() == 7 {
			return te.Day() == 4
		} else if te.Month() == 9 {
			return te.Day() == 5
		} else if te.Month() == 11 {
			return te.Day() == 24 || te.Day() == 25
		} else if te.Month() == 12 {
			return te.Day() == 26
		}
	} else if te.Year() == 2017 {
		if te.Month() == 1 {
			return te.Day() == 1 || te.Day() == 16
		} else if te.Month() == 2 {
			return te.Day() == 20
		} else if te.Month() == 4 {
			return te.Day() == 15
		} else if te.Month() == 5 {
			return te.Day() == 29
		} else if te.Month() == 7 {
			return te.Day() == 4
		} else if te.Month() == 9 {
			return te.Day() == 4
		} else if te.Month() == 11 {
			return te.Day() == 23
		} else if te.Month() == 12 {
			return te.Day() == 25
		}
	} else if te.Year() == 2018 {
		if te.Month() == 1 {
			return te.Day() == 1 || te.Day() == 15
		} else if te.Month() == 2 {
			return te.Day() == 19
		} else if te.Month() == 3 {
			return te.Day() == 30
		} else if te.Month() == 5 {
			return te.Day() == 28
		} else if te.Month() == 7 {
			return te.Day() == 4
		} else if te.Month() == 9 {
			return te.Day() == 3
		} else if te.Month() == 11 {
			return te.Day() == 22
		} else if te.Month() == 12 {
			return te.Day() == 25
		}
	}
	return false
}

// IsNYSEArcaHoliday returns that returns if a given time falls on a holiday.
func (d date) IsNYSEArcaHoliday(t time.Time) bool {
	return d.IsNYSEHoliday(t)
}

// IsNASDAQHoliday returns if a date was a NASDAQ holiday day.
func (d date) IsNASDAQHoliday(t time.Time) bool {
	return d.IsNYSEHoliday(t)
}

// Time returns a new time.Time for the given clock components.
func (d date) Time(hour, min, sec, nsec int, loc *time.Location) time.Time {
	return time.Date(0, 0, 0, hour, min, sec, nsec, loc)
}

func (d date) Date(year, month, day int, loc *time.Location) time.Time {
	return time.Date(year, time.Month(month), day, 12, 0, 0, 0, loc)
}

// On returns the clock components of clock (hour,minute,second) on the date components of d.
func (d date) On(clock, cd time.Time) time.Time {
	tzAdjusted := cd.In(clock.Location())
	return time.Date(tzAdjusted.Year(), tzAdjusted.Month(), tzAdjusted.Day(), clock.Hour(), clock.Minute(), clock.Second(), clock.Nanosecond(), clock.Location())
}

// NoonOn is a shortcut for On(Time(12,0,0), cd) a.k.a. noon on a given date.
func (d date) NoonOn(cd time.Time) time.Time {
	return time.Date(cd.Year(), cd.Month(), cd.Day(), 12, 0, 0, 0, cd.Location())
}

// Optional returns a pointer reference to a given time.
func (d date) Optional(t time.Time) *time.Time {
	return &t
}

// IsWeekDay returns if the day is a monday->friday.
func (d date) IsWeekDay(day time.Weekday) bool {
	return !d.IsWeekendDay(day)
}

// IsWeekendDay returns if the day is a monday->friday.
func (d date) IsWeekendDay(day time.Weekday) bool {
	return day == time.Saturday || day == time.Sunday
}

// Before returns if a timestamp is strictly before another date (ignoring hours, minutes etc.)
func (d date) Before(before, reference time.Time) bool {
	tzAdjustedBefore := before.In(reference.Location())
	if tzAdjustedBefore.Year() < reference.Year() {
		return true
	}
	if tzAdjustedBefore.Month() < reference.Month() {
		return true
	}
	return tzAdjustedBefore.Year() == reference.Year() && tzAdjustedBefore.Month() == reference.Month() && tzAdjustedBefore.Day() < reference.Day()
}

// NextMarketOpen returns the next market open after a given time.
func (d date) NextMarketOpen(after, openTime time.Time, isHoliday HolidayProvider) time.Time {
	afterLocalized := after.In(openTime.Location())
	todaysOpen := d.On(openTime, afterLocalized)

	if isHoliday == nil {
		isHoliday = defaultHolidayProvider
	}

	todayIsValidTradingDay := d.IsWeekDay(todaysOpen.Weekday()) && !isHoliday(todaysOpen)

	if (afterLocalized.Equal(todaysOpen) || afterLocalized.Before(todaysOpen)) && todayIsValidTradingDay {
		return todaysOpen
	}

	for cursorDay := 1; cursorDay < 7; cursorDay++ {
		newDay := todaysOpen.AddDate(0, 0, cursorDay)
		isValidTradingDay := d.IsWeekDay(newDay.Weekday()) && !isHoliday(newDay)
		if isValidTradingDay {
			return d.On(openTime, newDay)
		}
	}
	panic("Have exhausted day window looking for next market open.")
}

// NextMarketClose returns the next market close after a given time.
func (d date) NextMarketClose(after, closeTime time.Time, isHoliday HolidayProvider) time.Time {
	afterLocalized := after.In(closeTime.Location())

	if isHoliday == nil {
		isHoliday = defaultHolidayProvider
	}

	todaysClose := d.On(closeTime, afterLocalized)
	if afterLocalized.Before(todaysClose) && d.IsWeekDay(todaysClose.Weekday()) && !isHoliday(todaysClose) {
		return todaysClose
	}

	if afterLocalized.Equal(todaysClose) { //rare but it might happen.
		return todaysClose
	}

	for cursorDay := 1; cursorDay < 6; cursorDay++ {
		newDay := todaysClose.AddDate(0, 0, cursorDay)
		if d.IsWeekDay(newDay.Weekday()) && !isHoliday(newDay) {
			return d.On(closeTime, newDay)
		}
	}
	panic("Have exhausted day window looking for next market close.")
}

// CalculateMarketSecondsBetween calculates the number of seconds the market was open between two dates.
func (d date) CalculateMarketSecondsBetween(start, end, marketOpen, marketClose time.Time, isHoliday HolidayProvider) (seconds int64) {
	startEastern := start.In(d.Eastern())
	endEastern := end.In(d.Eastern())

	startMarketOpen := d.On(marketOpen, startEastern)
	startMarketClose := d.On(marketClose, startEastern)

	if !d.IsWeekendDay(startMarketOpen.Weekday()) && !isHoliday(startMarketOpen) {
		if (startEastern.Equal(startMarketOpen) || startEastern.After(startMarketOpen)) && startEastern.Before(startMarketClose) {
			if endEastern.Before(startMarketClose) {
				seconds += int64(endEastern.Sub(startEastern) / time.Second)
			} else {
				seconds += int64(startMarketClose.Sub(startEastern) / time.Second)
			}
		}
	}

	cursor := d.NextMarketOpen(startMarketClose, marketOpen, isHoliday)
	for d.Before(cursor, endEastern) {
		if d.IsWeekDay(cursor.Weekday()) && !isHoliday(cursor) {
			close := d.NextMarketClose(cursor, marketClose, isHoliday)
			seconds += int64(close.Sub(cursor) / time.Second)
		}
		cursor = cursor.AddDate(0, 0, 1)
	}

	finalMarketOpen := d.NextMarketOpen(cursor, marketOpen, isHoliday)
	finalMarketClose := d.NextMarketClose(cursor, marketClose, isHoliday)
	if endEastern.After(finalMarketOpen) {
		if endEastern.Before(finalMarketClose) {
			seconds += int64(endEastern.Sub(finalMarketOpen) / time.Second)
		} else {
			seconds += int64(finalMarketClose.Sub(finalMarketOpen) / time.Second)
		}
	}

	return
}

const (
	_secondsPerHour = 60 * 60
	_secondsPerDay  = 60 * 60 * 24
)

func (d date) DiffDays(t1, t2 time.Time) (days int) {
	t1n := t1.Unix()
	t2n := t2.Unix()
	diff := t2n - t1n //yields seconds
	return int(diff / (_secondsPerDay))
}

func (d date) DiffHours(t1, t2 time.Time) (hours int) {
	t1n := t1.Unix()
	t2n := t2.Unix()
	diff := t2n - t1n //yields seconds
	return int(diff / (_secondsPerHour))
}

// NextDay returns the timestamp advanced a day.
func (d date) NextDay(ts time.Time) time.Time {
	return ts.AddDate(0, 0, 1)
}

// NextHour returns the next timestamp on the hour.
func (d date) NextHour(ts time.Time) time.Time {
	//advance a full hour ...
	advanced := ts.Add(time.Hour)
	minutes := time.Duration(advanced.Minute()) * time.Minute
	final := advanced.Add(-minutes)
	return time.Date(final.Year(), final.Month(), final.Day(), final.Hour(), 0, 0, 0, final.Location())
}

// NextDayOfWeek returns the next instance of a given weekday after a given timestamp.
func (d date) NextDayOfWeek(after time.Time, dayOfWeek time.Weekday) time.Time {
	afterWeekday := after.Weekday()
	if afterWeekday == dayOfWeek {
		return after.AddDate(0, 0, 7)
	}

	// 1 vs 5 ~ add 4 days
	if afterWeekday < dayOfWeek {
		dayDelta := int(dayOfWeek - afterWeekday)
		return after.AddDate(0, 0, dayDelta)
	}

	// 5 vs 1, add 7-(5-1) ~ 3 days
	dayDelta := 7 - int(afterWeekday-dayOfWeek)
	return after.AddDate(0, 0, dayDelta)
}
