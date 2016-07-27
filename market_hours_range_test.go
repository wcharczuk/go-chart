package chart

import (
	"testing"
	"time"

	assert "github.com/blendlabs/go-assert"
	"github.com/wcharczuk/go-chart/date"
)

func TestMarketHoursRangeGetDelta(t *testing.T) {
	assert := assert.New(t)

	r := &MarketHoursRange{
		Min:             time.Date(2016, 07, 19, 9, 30, 0, 0, date.Eastern()),
		Max:             time.Date(2016, 07, 22, 16, 00, 0, 0, date.Eastern()),
		MarketOpen:      date.NYSEOpen,
		MarketClose:     date.NYSEClose,
		HolidayProvider: date.IsNYSEHoliday,
	}

	assert.NotZero(r.GetDelta())
}

func TestMarketHoursRangeTranslate(t *testing.T) {
	assert := assert.New(t)

	r := &MarketHoursRange{
		Min:             time.Date(2016, 07, 18, 9, 30, 0, 0, date.Eastern()),
		Max:             time.Date(2016, 07, 22, 16, 00, 0, 0, date.Eastern()),
		MarketOpen:      date.NYSEOpen,
		MarketClose:     date.NYSEClose,
		HolidayProvider: date.IsNYSEHoliday,
		Domain:          1000,
	}

	weds := time.Date(2016, 07, 20, 9, 30, 0, 0, date.Eastern())

	assert.Equal(0, r.Translate(TimeToFloat64(r.Min)))
	assert.Equal(400, r.Translate(TimeToFloat64(weds)))
	assert.Equal(1000, r.Translate(TimeToFloat64(r.Max)))
}

func TestMarketHoursRangeGetTicks(t *testing.T) {
	assert := assert.New(t)

	r := &MarketHoursRange{
		Min:             time.Date(2016, 07, 18, 9, 30, 0, 0, date.Eastern()),
		Max:             time.Date(2016, 07, 22, 16, 00, 0, 0, date.Eastern()),
		MarketOpen:      date.NYSEOpen,
		MarketClose:     date.NYSEClose,
		HolidayProvider: date.IsNYSEHoliday,
		Domain:          1000,
	}

	ticks := r.GetTicks(TimeValueFormatter)
	assert.NotEmpty(ticks)
	assert.Len(ticks, 6)
	assert.NotEqual(TimeToFloat64(r.Min), ticks[0].Value)
	assert.NotEmpty(ticks[0].Label)
}
