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

func TestDateEastern(t *testing.T) {
	assert := assert.New(t)
	eastern, err := Date.Eastern()
	assert.Nil(err)
	assert.NotNil(eastern)
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

func TestDateOnDate(t *testing.T) {
	assert := assert.New(t)

	eastern := Date.MustEastern()
	assert.NotNil(eastern)

	ts := Date.OnDate(Date.Time(5, 4, 3, 2, time.UTC), Date.Date(2016, 6, 7, eastern))
	assert.Equal(2016, ts.Year())
	assert.Equal(6, ts.Month())
	assert.Equal(7, ts.Day())
	assert.Equal(5, ts.Hour())
	assert.Equal(4, ts.Minute())
	assert.Equal(3, ts.Second())
	assert.Equal(2, ts.Nanosecond())
	assert.Equal(time.UTC, ts.Location())
}

func TestDateNoonOnDate(t *testing.T) {
	assert := assert.New(t)
	noon := Date.NoonOnDate(time.Date(2016, 04, 03, 02, 01, 0, 0, time.UTC))

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
	sundayEST := time.Date(2016, 7, 31, 22, 00, 0, 0, Date.MustEastern())

	assert.True(Date.Before(sundayEST, tuesdayUTC))
	assert.False(Date.Before(sundayEST, mondayUTC))
}

func TestDateNextHour(t *testing.T) {
	assert := assert.New(t)

	start := time.Date(2015, 07, 01, 9, 30, 0, 0, Date.MustEastern())
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
