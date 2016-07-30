package date

import (
	"testing"
	"time"

	assert "github.com/blendlabs/go-assert"
)

func TestBeforeDate(t *testing.T) {
	assert := assert.New(t)

	assert.True(BeforeDate(parse("2015-07-02"), parse("2016-07-01")))
	assert.True(BeforeDate(parse("2016-06-01"), parse("2016-07-01")))
	assert.True(BeforeDate(parse("2016-07-01"), parse("2016-07-02")))

	assert.False(BeforeDate(parse("2016-07-01"), parse("2016-07-01")))
	assert.False(BeforeDate(parse("2016-07-03"), parse("2016-07-01")))
	assert.False(BeforeDate(parse("2016-08-03"), parse("2016-07-01")))
	assert.False(BeforeDate(parse("2017-08-03"), parse("2016-07-01")))
}

func TestNextMarketOpen(t *testing.T) {
	assert := assert.New(t)

	beforeOpen := time.Date(2016, 07, 18, 9, 0, 0, 0, Eastern())
	todayOpen := time.Date(2016, 07, 18, 9, 30, 0, 0, Eastern())

	afterOpen := time.Date(2016, 07, 18, 9, 31, 0, 0, Eastern())
	tomorrowOpen := time.Date(2016, 07, 19, 9, 30, 0, 0, Eastern())

	afterFriday := time.Date(2016, 07, 22, 9, 31, 0, 0, Eastern())
	mondayOpen := time.Date(2016, 07, 25, 9, 30, 0, 0, Eastern())

	weekend := time.Date(2016, 07, 23, 9, 31, 0, 0, Eastern())

	assert.True(todayOpen.Equal(NextMarketOpen(beforeOpen, NYSEOpen, IsNYSEHoliday)))
	assert.True(tomorrowOpen.Equal(NextMarketOpen(afterOpen, NYSEOpen, IsNYSEHoliday)))
	assert.True(mondayOpen.Equal(NextMarketOpen(afterFriday, NYSEOpen, IsNYSEHoliday)))
	assert.True(mondayOpen.Equal(NextMarketOpen(weekend, NYSEOpen, IsNYSEHoliday)))

	testRegression := time.Date(2016, 07, 18, 16, 0, 0, 0, Eastern())
	shouldbe := time.Date(2016, 07, 19, 9, 30, 0, 0, Eastern())

	assert.True(shouldbe.Equal(NextMarketOpen(testRegression, NYSEOpen, IsNYSEHoliday)))
}

func TestNextMarketClose(t *testing.T) {
	assert := assert.New(t)

	beforeClose := time.Date(2016, 07, 18, 15, 0, 0, 0, Eastern())
	todayClose := time.Date(2016, 07, 18, 16, 00, 0, 0, Eastern())

	afterClose := time.Date(2016, 07, 18, 16, 1, 0, 0, Eastern())
	tomorrowClose := time.Date(2016, 07, 19, 16, 00, 0, 0, Eastern())

	afterFriday := time.Date(2016, 07, 22, 16, 1, 0, 0, Eastern())
	mondayClose := time.Date(2016, 07, 25, 16, 0, 0, 0, Eastern())

	weekend := time.Date(2016, 07, 23, 9, 31, 0, 0, Eastern())

	assert.True(todayClose.Equal(NextMarketClose(beforeClose, NYSEClose, IsNYSEHoliday)))
	assert.True(tomorrowClose.Equal(NextMarketClose(afterClose, NYSEClose, IsNYSEHoliday)))
	assert.True(mondayClose.Equal(NextMarketClose(afterFriday, NYSEClose, IsNYSEHoliday)))
	assert.True(mondayClose.Equal(NextMarketClose(weekend, NYSEClose, IsNYSEHoliday)))
}

func TestCalculateMarketSecondsBetween(t *testing.T) {
	assert := assert.New(t)

	start := time.Date(2016, 07, 18, 9, 30, 0, 0, Eastern())
	end := time.Date(2016, 07, 22, 16, 00, 0, 0, Eastern())

	shouldbe := 5 * 6.5 * 60 * 60

	assert.Equal(shouldbe, CalculateMarketSecondsBetween(start, end, NYSEOpen, NYSEClose, IsNYSEHoliday))
}

func TestCalculateMarketSecondsBetween1D(t *testing.T) {
	assert := assert.New(t)

	start := time.Date(2016, 07, 22, 9, 45, 0, 0, Eastern())
	end := time.Date(2016, 07, 22, 15, 45, 0, 0, Eastern())

	shouldbe := 6 * 60 * 60

	assert.Equal(shouldbe, CalculateMarketSecondsBetween(start, end, NYSEOpen, NYSEClose, IsNYSEHoliday))
}

func TestCalculateMarketSecondsBetweenLTM(t *testing.T) {
	assert := assert.New(t)

	start := time.Date(2015, 07, 01, 9, 30, 0, 0, Eastern())
	end := time.Date(2016, 07, 01, 9, 30, 0, 0, Eastern())

	shouldbe := 253 * 6.5 * 60 * 60 //253 full market days since this date last year.
	assert.Equal(shouldbe, CalculateMarketSecondsBetween(start, end, NYSEOpen, NYSEClose, IsNYSEHoliday))
}
