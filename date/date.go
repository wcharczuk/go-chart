package date

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

// Eastern returns the eastern timezone.
func Eastern() *time.Location {
	if _eastern == nil {
		_easternLock.Lock()
		defer _easternLock.Unlock()
		if _eastern == nil {
			_eastern, _ = time.LoadLocation("America/New_York")
		}
	}
	return _eastern
}

// Optional returns a pointer reference to a given time.
func Optional(t time.Time) *time.Time {
	return &t
}

// IsWeekDay returns if the day is a monday->friday.
func IsWeekDay(day time.Weekday) bool {
	return !IsWeekendDay(day)
}

// IsWeekendDay returns if the day is a monday->friday.
func IsWeekendDay(day time.Weekday) bool {
	return day == time.Saturday || day == time.Sunday
}

// BeforeDate returns if a timestamp is strictly before another date (ignoring hours, minutes etc.)
func BeforeDate(before, reference time.Time) bool {
	if before.Year() < reference.Year() {
		return true
	}
	if before.Month() < reference.Month() {
		return true
	}
	return before.Year() == reference.Year() && before.Month() == reference.Month() && before.Day() < reference.Day()
}

// IsNYSEHoliday returns if a date was/is on a nyse holiday day.
func IsNYSEHoliday(t time.Time) bool {
	te := t.In(Eastern())
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

// MarketOpen returns 0930 on a given day.
func MarketOpen(on time.Time) time.Time {
	onEastern := on.In(Eastern())
	return time.Date(onEastern.Year(), onEastern.Month(), onEastern.Day(), 9, 30, 0, 0, Eastern())
}

// MarketClose returns 1600 on a given day.
func MarketClose(on time.Time) time.Time {
	onEastern := on.In(Eastern())
	return time.Date(onEastern.Year(), onEastern.Month(), onEastern.Day(), 16, 0, 0, 0, Eastern())
}

// NextMarketOpen returns the next market open after a given time.
func NextMarketOpen(after time.Time) time.Time {
	afterEastern := after.In(Eastern())
	todaysOpen := MarketOpen(afterEastern)

	if afterEastern.Before(todaysOpen) && IsWeekDay(todaysOpen.Weekday()) && !IsNYSEHoliday(todaysOpen) {
		return todaysOpen
	}

	if afterEastern.Equal(todaysOpen) { //rare but it might happen.
		return todaysOpen
	}

	for cursorDay := 1; cursorDay < 6; cursorDay++ {
		newDay := todaysOpen.AddDate(0, 0, cursorDay)
		if IsWeekDay(newDay.Weekday()) && !IsNYSEHoliday(afterEastern) {
			return time.Date(newDay.Year(), newDay.Month(), newDay.Day(), 9, 30, 0, 0, Eastern())
		}
	}
	return Epoch //we should never reach this.
}

// NextMarketClose returns the next market close after a given time.
func NextMarketClose(after time.Time) time.Time {
	afterEastern := after.In(Eastern())

	todaysClose := MarketClose(afterEastern)
	if afterEastern.Before(todaysClose) && IsWeekDay(todaysClose.Weekday()) && !IsNYSEHoliday(todaysClose) {
		return todaysClose
	}

	if afterEastern.Equal(todaysClose) { //rare but it might happen.
		return todaysClose
	}

	for cursorDay := 1; cursorDay < 6; cursorDay++ {
		newDay := todaysClose.AddDate(0, 0, cursorDay)
		if IsWeekDay(newDay.Weekday()) && !IsNYSEHoliday(newDay) {
			return time.Date(newDay.Year(), newDay.Month(), newDay.Day(), 16, 0, 0, 0, Eastern())
		}
	}
	return Epoch //we should never reach this.
}

// CalculateMarketSecondsBetween calculates the number of seconds the market was open between two dates.
func CalculateMarketSecondsBetween(start, end time.Time) (seconds int64) {
	se := start.In(Eastern())
	ee := end.In(Eastern())

	startMarketOpen := NextMarketOpen(se)
	startMarketClose := NextMarketClose(se)

	if (se.Equal(startMarketOpen) || se.After(startMarketOpen)) && se.Before(startMarketClose) {
		seconds += int64(startMarketClose.Sub(se) / time.Second)
	}

	cursor := NextMarketOpen(startMarketClose)
	for BeforeDate(cursor, ee) {
		if IsWeekDay(cursor.Weekday()) && !IsNYSEHoliday(cursor) {
			close := NextMarketClose(cursor)
			seconds += int64(close.Sub(cursor) / time.Second)
		}
		cursor = cursor.AddDate(0, 0, 1)
	}

	finalMarketOpen := NextMarketOpen(cursor)
	finalMarketClose := NextMarketClose(cursor)
	if end.After(finalMarketOpen) {
		if end.Before(finalMarketClose) {
			seconds += int64(end.Sub(finalMarketOpen) / time.Second)
		} else {
			seconds += int64(finalMarketClose.Sub(finalMarketOpen) / time.Second)
		}
	}

	return
}

// Format returns a string representation of a date.
func format(t time.Time) string {
	return t.Format("2006-01-02")
}

// Parse parses a date from a string.
func parse(str string) time.Time {
	res, _ := time.Parse("2006-01-02", str)
	return res
}
