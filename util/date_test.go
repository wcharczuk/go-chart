package util

import (
	"testing"
	"time"

	assert "github.com/blend/go-sdk/assert"
)

func parse(v string) time.Time {
	ts, _ := time.Parse("2006-01-02", v)
	return ts
}

func TestDateTime(t *testing.T) {
	assert := assert.New(t)

	ts := Date.Time(5, 6, 7, 8, time.UTC)
	assert.Equal(05, ts.Hour())
	assert.Equal(06, ts.Minute())
	assert.Equal(07, ts.Second())
	assert.Equal(8, ts.Nanosecond())
	assert.Equal(time.UTC, ts.Location())
}

func TestDateDate(t *testing.T) {
	assert := assert.New(t)

	ts := Date.Date(2015, 5, 6, time.UTC)
	assert.Equal(2015, ts.Year())
	assert.Equal(5, ts.Month())
	assert.Equal(6, ts.Day())
	assert.Equal(time.UTC, ts.Location())
}

func TestDateOn(t *testing.T) {
	assert := assert.New(t)

	ts := Date.On(Date.Time(5, 4, 3, 2, time.UTC), Date.Date(2016, 6, 7, Date.Eastern()))
	assert.Equal(2016, ts.Year())
	assert.Equal(6, ts.Month())
	assert.Equal(7, ts.Day())
	assert.Equal(5, ts.Hour())
	assert.Equal(4, ts.Minute())
	assert.Equal(3, ts.Second())
	assert.Equal(2, ts.Nanosecond())
	assert.Equal(time.UTC, ts.Location())
}

func TestDateNoonOn(t *testing.T) {
	assert := assert.New(t)
	noon := Date.NoonOn(time.Date(2016, 04, 03, 02, 01, 0, 0, time.UTC))

	assert.Equal(2016, noon.Year())
	assert.Equal(4, noon.Month())
	assert.Equal(3, noon.Day())
	assert.Equal(12, noon.Hour())
	assert.Equal(0, noon.Minute())
	assert.Equal(time.UTC, noon.Location())
}

func TestDateBefore(t *testing.T) {
	assert := assert.New(t)

	assert.True(Date.Before(parse("2015-07-02"), parse("2016-07-01")))
	assert.True(Date.Before(parse("2016-06-01"), parse("2016-07-01")))
	assert.True(Date.Before(parse("2016-07-01"), parse("2016-07-02")))

	assert.False(Date.Before(parse("2016-07-01"), parse("2016-07-01")))
	assert.False(Date.Before(parse("2016-07-03"), parse("2016-07-01")))
	assert.False(Date.Before(parse("2016-08-03"), parse("2016-07-01")))
	assert.False(Date.Before(parse("2017-08-03"), parse("2016-07-01")))
}

func TestDateBeforeHandlesTimezones(t *testing.T) {
	assert := assert.New(t)

	tuesdayUTC := time.Date(2016, 8, 02, 22, 00, 0, 0, time.UTC)
	mondayUTC := time.Date(2016, 8, 01, 1, 00, 0, 0, time.UTC)
	sundayEST := time.Date(2016, 7, 31, 22, 00, 0, 0, Date.Eastern())

	assert.True(Date.Before(sundayEST, tuesdayUTC))
	assert.False(Date.Before(sundayEST, mondayUTC))
}

func TestNextMarketOpen(t *testing.T) {
	assert := assert.New(t)

	beforeOpen := time.Date(2016, 07, 18, 9, 0, 0, 0, Date.Eastern())
	todayOpen := time.Date(2016, 07, 18, 9, 30, 0, 0, Date.Eastern())

	afterOpen := time.Date(2016, 07, 18, 9, 31, 0, 0, Date.Eastern())
	tomorrowOpen := time.Date(2016, 07, 19, 9, 30, 0, 0, Date.Eastern())

	afterFriday := time.Date(2016, 07, 22, 9, 31, 0, 0, Date.Eastern())
	mondayOpen := time.Date(2016, 07, 25, 9, 30, 0, 0, Date.Eastern())

	weekend := time.Date(2016, 07, 23, 9, 31, 0, 0, Date.Eastern())

	assert.True(todayOpen.Equal(Date.NextMarketOpen(beforeOpen, NYSEOpen(), Date.IsNYSEHoliday)))
	assert.True(tomorrowOpen.Equal(Date.NextMarketOpen(afterOpen, NYSEOpen(), Date.IsNYSEHoliday)))
	assert.True(mondayOpen.Equal(Date.NextMarketOpen(afterFriday, NYSEOpen(), Date.IsNYSEHoliday)))
	assert.True(mondayOpen.Equal(Date.NextMarketOpen(weekend, NYSEOpen(), Date.IsNYSEHoliday)))

	assert.Equal(Date.Eastern(), todayOpen.Location())
	assert.Equal(Date.Eastern(), tomorrowOpen.Location())
	assert.Equal(Date.Eastern(), mondayOpen.Location())

	testRegression := time.Date(2016, 07, 18, 16, 0, 0, 0, Date.Eastern())
	shouldbe := time.Date(2016, 07, 19, 9, 30, 0, 0, Date.Eastern())

	assert.True(shouldbe.Equal(Date.NextMarketOpen(testRegression, NYSEOpen(), Date.IsNYSEHoliday)))
}

func TestNextMarketClose(t *testing.T) {
	assert := assert.New(t)

	beforeClose := time.Date(2016, 07, 18, 15, 0, 0, 0, Date.Eastern())
	todayClose := time.Date(2016, 07, 18, 16, 00, 0, 0, Date.Eastern())

	afterClose := time.Date(2016, 07, 18, 16, 1, 0, 0, Date.Eastern())
	tomorrowClose := time.Date(2016, 07, 19, 16, 00, 0, 0, Date.Eastern())

	afterFriday := time.Date(2016, 07, 22, 16, 1, 0, 0, Date.Eastern())
	mondayClose := time.Date(2016, 07, 25, 16, 0, 0, 0, Date.Eastern())

	weekend := time.Date(2016, 07, 23, 9, 31, 0, 0, Date.Eastern())

	assert.True(todayClose.Equal(Date.NextMarketClose(beforeClose, NYSEClose(), Date.IsNYSEHoliday)))
	assert.True(tomorrowClose.Equal(Date.NextMarketClose(afterClose, NYSEClose(), Date.IsNYSEHoliday)))
	assert.True(mondayClose.Equal(Date.NextMarketClose(afterFriday, NYSEClose(), Date.IsNYSEHoliday)))
	assert.True(mondayClose.Equal(Date.NextMarketClose(weekend, NYSEClose(), Date.IsNYSEHoliday)))

	assert.Equal(Date.Eastern(), todayClose.Location())
	assert.Equal(Date.Eastern(), tomorrowClose.Location())
	assert.Equal(Date.Eastern(), mondayClose.Location())
}

func TestCalculateMarketSecondsBetween(t *testing.T) {
	assert := assert.New(t)

	start := time.Date(2016, 07, 18, 9, 30, 0, 0, Date.Eastern())
	end := time.Date(2016, 07, 22, 16, 00, 0, 0, Date.Eastern())

	shouldbe := 5 * 6.5 * 60 * 60

	assert.Equal(shouldbe, Date.CalculateMarketSecondsBetween(start, end, NYSEOpen(), NYSEClose(), Date.IsNYSEHoliday))
}

func TestCalculateMarketSecondsBetween1D(t *testing.T) {
	assert := assert.New(t)

	start := time.Date(2016, 07, 22, 9, 45, 0, 0, Date.Eastern())
	end := time.Date(2016, 07, 22, 15, 45, 0, 0, Date.Eastern())

	shouldbe := 6 * 60 * 60

	assert.Equal(shouldbe, Date.CalculateMarketSecondsBetween(start, end, NYSEOpen(), NYSEClose(), Date.IsNYSEHoliday))
}

func TestCalculateMarketSecondsBetweenLTM(t *testing.T) {
	assert := assert.New(t)

	start := time.Date(2015, 07, 01, 9, 30, 0, 0, Date.Eastern())
	end := time.Date(2016, 07, 01, 9, 30, 0, 0, Date.Eastern())

	shouldbe := 253 * 6.5 * 60 * 60 //253 full market days since this date last year.
	assert.Equal(shouldbe, Date.CalculateMarketSecondsBetween(start, end, NYSEOpen(), NYSEClose(), Date.IsNYSEHoliday))
}

func TestDateNextHour(t *testing.T) {
	assert := assert.New(t)

	start := time.Date(2015, 07, 01, 9, 30, 0, 0, Date.Eastern())
	next := Date.NextHour(start)
	assert.Equal(2015, next.Year())
	assert.Equal(07, next.Month())
	assert.Equal(01, next.Day())
	assert.Equal(10, next.Hour())
	assert.Equal(00, next.Minute())

	next = Date.NextHour(next)
	assert.Equal(11, next.Hour())

	next = Date.NextHour(next)
	assert.Equal(12, next.Hour())

}

func TestDateNextDayOfWeek(t *testing.T) {
	assert := assert.New(t)

	weds := Date.Date(2016, 8, 10, time.UTC)
	fri := Date.Date(2016, 8, 12, time.UTC)
	sun := Date.Date(2016, 8, 14, time.UTC)
	mon := Date.Date(2016, 8, 15, time.UTC)
	weds2 := Date.Date(2016, 8, 17, time.UTC)

	nextFri := Date.NextDayOfWeek(weds, time.Friday)
	nextSunday := Date.NextDayOfWeek(weds, time.Sunday)
	nextMonday := Date.NextDayOfWeek(weds, time.Monday)
	nextWeds := Date.NextDayOfWeek(weds, time.Wednesday)

	assert.Equal(fri.Year(), nextFri.Year())
	assert.Equal(fri.Month(), nextFri.Month())
	assert.Equal(fri.Day(), nextFri.Day())

	assert.Equal(sun.Year(), nextSunday.Year())
	assert.Equal(sun.Month(), nextSunday.Month())
	assert.Equal(sun.Day(), nextSunday.Day())

	assert.Equal(mon.Year(), nextMonday.Year())
	assert.Equal(mon.Month(), nextMonday.Month())
	assert.Equal(mon.Day(), nextMonday.Day())

	assert.Equal(weds2.Year(), nextWeds.Year())
	assert.Equal(weds2.Month(), nextWeds.Month())
	assert.Equal(weds2.Day(), nextWeds.Day())

	assert.Equal(time.UTC, nextFri.Location())
	assert.Equal(time.UTC, nextSunday.Location())
	assert.Equal(time.UTC, nextMonday.Location())
}

func TestDateIsNYSEHoliday(t *testing.T) {
	assert := assert.New(t)

	cursor := time.Date(2013, 01, 01, 0, 0, 0, 0, time.UTC)
	end := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	var holidays int
	for Date.Before(cursor, end) {
		if Date.IsNYSEHoliday(cursor) {
			holidays++
		}
		cursor = cursor.AddDate(0, 0, 1)
	}
	assert.Equal(holidays, 55)
}

func TestDateDiffDays(t *testing.T) {
	assert := assert.New(t)

	t1 := time.Date(2017, 02, 27, 12, 0, 0, 0, time.UTC)
	t2 := time.Date(2017, 01, 10, 3, 0, 0, 0, time.UTC)
	t3 := time.Date(2017, 02, 24, 16, 0, 0, 0, time.UTC)

	assert.Equal(48, Date.DiffDays(t2, t1))
	assert.Equal(2, Date.DiffDays(t3, t1)) // technically we should round down.
}

func TestDateDiffHours(t *testing.T) {
	assert := assert.New(t)

	t1 := time.Date(2017, 02, 27, 12, 0, 0, 0, time.UTC)
	t2 := time.Date(2017, 02, 24, 16, 0, 0, 0, time.UTC)
	t3 := time.Date(2017, 02, 28, 12, 0, 0, 0, time.UTC)

	assert.Equal(68, Date.DiffHours(t2, t1))
	assert.Equal(24, Date.DiffHours(t1, t3))
}
