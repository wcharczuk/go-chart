package util

import (
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

// Date contains utility functions that operate on dates.
var Date date

type date struct{}

func (d date) MustEastern() *time.Location {
	if eastern, err := d.Eastern(); err != nil {
		panic(err)
	} else {
		return eastern
	}
}

// Eastern returns the eastern timezone.
func (d date) Eastern() (*time.Location, error) {
	// Try POSIX
	est, err := time.LoadLocation("America/New_York")
	if err != nil {
		// Try Windows
		est, err = time.LoadLocation("EST")
		if err != nil {
			return nil, err
		}
	}
	return est, nil
}

func (d date) MustPacific() *time.Location {
	if pst, err := d.Pacific(); err != nil {
		panic(err)
	} else {
		return pst
	}
}

// Pacific returns the pacific timezone.
func (d date) Pacific() (*time.Location, error) {
	// Try POSIX
	pst, err := time.LoadLocation("America/Los_Angeles")
	if err != nil {
		// Try Windows
		pst, err = time.LoadLocation("PST")
		if err != nil {
			return nil, err
		}
	}
	return pst, nil
}

// TimeUTC returns a new time.Time for the given clock components in UTC.
// It is meant to be used with the `OnDate` function.
func (d date) TimeUTC(hour, min, sec, nsec int) time.Time {
	return time.Date(0, 0, 0, hour, min, sec, nsec, time.UTC)
}

// Time returns a new time.Time for the given clock components.
// It is meant to be used with the `OnDate` function.
func (d date) Time(hour, min, sec, nsec int, loc *time.Location) time.Time {
	return time.Date(0, 0, 0, hour, min, sec, nsec, loc)
}

// DateUTC returns a new time.Time for the given date comonents at (noon) in UTC.
func (d date) DateUTC(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 12, 0, 0, 0, time.UTC)
}

// DateUTC returns a new time.Time for the given date comonents at (noon) in a given location.
func (d date) Date(year, month, day int, loc *time.Location) time.Time {
	return time.Date(year, time.Month(month), day, 12, 0, 0, 0, loc)
}

// OnDate returns the clock components of clock (hour,minute,second) on the date components of d.
func (d date) OnDate(clock, date time.Time) time.Time {
	tzAdjusted := date.In(clock.Location())
	return time.Date(tzAdjusted.Year(), tzAdjusted.Month(), tzAdjusted.Day(), clock.Hour(), clock.Minute(), clock.Second(), clock.Nanosecond(), clock.Location())
}

// NoonOnDate is a shortcut for On(Time(12,0,0), cd) a.k.a. noon on a given date.
func (d date) NoonOnDate(cd time.Time) time.Time {
	return time.Date(cd.Year(), cd.Month(), cd.Day(), 12, 0, 0, 0, cd.Location())
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

const (
	_secondsPerHour = 60 * 60
	_secondsPerDay  = 60 * 60 * 24
)

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
