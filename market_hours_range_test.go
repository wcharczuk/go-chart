package chart

import (
	"testing"
	"time"

	assert "github.com/blendlabs/go-assert"
)

func TestMarketHoursRangeGetDelta(t *testing.T) {
	assert := assert.New(t)

	r := &MarketHoursRange{
		Min:             time.Date(2016, 07, 19, 9, 30, 0, 0, Date.Eastern()),
		Max:             time.Date(2016, 07, 22, 16, 00, 0, 0, Date.Eastern()),
		MarketOpen:      NYSEOpen,
		MarketClose:     NYSEClose,
		HolidayProvider: Date.IsNYSEHoliday,
	}

	assert.NotZero(r.GetDelta())
}

func TestMarketHoursRangeTranslate(t *testing.T) {
	assert := assert.New(t)

	r := &MarketHoursRange{
		Min:             time.Date(2016, 07, 18, 9, 30, 0, 0, Date.Eastern()),
		Max:             time.Date(2016, 07, 22, 16, 00, 0, 0, Date.Eastern()),
		MarketOpen:      NYSEOpen,
		MarketClose:     NYSEClose,
		HolidayProvider: Date.IsNYSEHoliday,
		Domain:          1000,
	}

	weds := time.Date(2016, 07, 20, 9, 30, 0, 0, Date.Eastern())

	assert.Equal(0, r.Translate(TimeToFloat64(r.Min)))
	assert.Equal(400, r.Translate(TimeToFloat64(weds)))
	assert.Equal(1000, r.Translate(TimeToFloat64(r.Max)))
}

func TestMarketHoursRangeGetTicks(t *testing.T) {
	assert := assert.New(t)

	r := &MarketHoursRange{
		Min:             time.Date(2016, 07, 18, 9, 30, 0, 0, Date.Eastern()),
		Max:             time.Date(2016, 07, 22, 16, 00, 0, 0, Date.Eastern()),
		MarketOpen:      NYSEOpen,
		MarketClose:     NYSEClose,
		HolidayProvider: Date.IsNYSEHoliday,
		Domain:          1000,
	}

	ticks := r.GetTicks(TimeValueFormatter)
	assert.NotEmpty(ticks)
	assert.Len(ticks, 24)
	assert.NotEqual(TimeToFloat64(r.Min), ticks[0].Value)
	assert.NotEmpty(ticks[0].Label)
}
