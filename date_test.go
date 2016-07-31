package chart

import (
	"testing"
	"time"

	assert "github.com/blendlabs/go-assert"
)

func parse(v string) time.Time {
	ts, _ := time.Parse("2006-01-02", v)
	return ts
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

func TestNextMarketOpen(t *testing.T) {
	assert := assert.New(t)

	beforeOpen := time.Date(2016, 07, 18, 9, 0, 0, 0, Date.Eastern())
	todayOpen := time.Date(2016, 07, 18, 9, 30, 0, 0, Date.Eastern())

	afterOpen := time.Date(2016, 07, 18, 9, 31, 0, 0, Date.Eastern())
	tomorrowOpen := time.Date(2016, 07, 19, 9, 30, 0, 0, Date.Eastern())

	afterFriday := time.Date(2016, 07, 22, 9, 31, 0, 0, Date.Eastern())
	mondayOpen := time.Date(2016, 07, 25, 9, 30, 0, 0, Date.Eastern())

	weekend := time.Date(2016, 07, 23, 9, 31, 0, 0, Date.Eastern())

	assert.True(todayOpen.Equal(Date.NextMarketOpen(beforeOpen, NYSEOpen, Date.IsNYSEHoliday)))
	assert.True(tomorrowOpen.Equal(Date.NextMarketOpen(afterOpen, NYSEOpen, Date.IsNYSEHoliday)))
	assert.True(mondayOpen.Equal(Date.NextMarketOpen(afterFriday, NYSEOpen, Date.IsNYSEHoliday)))
	assert.True(mondayOpen.Equal(Date.NextMarketOpen(weekend, NYSEOpen, Date.IsNYSEHoliday)))

	testRegression := time.Date(2016, 07, 18, 16, 0, 0, 0, Date.Eastern())
	shouldbe := time.Date(2016, 07, 19, 9, 30, 0, 0, Date.Eastern())

	assert.True(shouldbe.Equal(Date.NextMarketOpen(testRegression, NYSEOpen, Date.IsNYSEHoliday)))
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

	assert.True(todayClose.Equal(Date.NextMarketClose(beforeClose, NYSEClose, Date.IsNYSEHoliday)))
	assert.True(tomorrowClose.Equal(Date.NextMarketClose(afterClose, NYSEClose, Date.IsNYSEHoliday)))
	assert.True(mondayClose.Equal(Date.NextMarketClose(afterFriday, NYSEClose, Date.IsNYSEHoliday)))
	assert.True(mondayClose.Equal(Date.NextMarketClose(weekend, NYSEClose, Date.IsNYSEHoliday)))
}

func TestCalculateMarketSecondsBetween(t *testing.T) {
	assert := assert.New(t)

	start := time.Date(2016, 07, 18, 9, 30, 0, 0, Date.Eastern())
	end := time.Date(2016, 07, 22, 16, 00, 0, 0, Date.Eastern())

	shouldbe := 5 * 6.5 * 60 * 60

	assert.Equal(shouldbe, Date.CalculateMarketSecondsBetween(start, end, NYSEOpen, NYSEClose, Date.IsNYSEHoliday))
}

func TestCalculateMarketSecondsBetween1D(t *testing.T) {
	assert := assert.New(t)

	start := time.Date(2016, 07, 22, 9, 45, 0, 0, Date.Eastern())
	end := time.Date(2016, 07, 22, 15, 45, 0, 0, Date.Eastern())

	shouldbe := 6 * 60 * 60

	assert.Equal(shouldbe, Date.CalculateMarketSecondsBetween(start, end, NYSEOpen, NYSEClose, Date.IsNYSEHoliday))
}

func TestCalculateMarketSecondsBetweenLTM(t *testing.T) {
	assert := assert.New(t)

	start := time.Date(2015, 07, 01, 9, 30, 0, 0, Date.Eastern())
	end := time.Date(2016, 07, 01, 9, 30, 0, 0, Date.Eastern())

	shouldbe := 253 * 6.5 * 60 * 60 //253 full market days since this date last year.
	assert.Equal(shouldbe, Date.CalculateMarketSecondsBetween(start, end, NYSEOpen, NYSEClose, Date.IsNYSEHoliday))
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
