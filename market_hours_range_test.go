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

	assert.Equal(0, r.Translate(Time.ToFloat64(r.Min)))
	assert.Equal(400, r.Translate(Time.ToFloat64(weds)))
	assert.Equal(1000, r.Translate(Time.ToFloat64(r.Max)))
}

func TestMarketHoursRangeGetTicks(t *testing.T) {
	assert := assert.New(t)

	r, err := PNG(1024, 1024)
	assert.Nil(err)

	f, err := GetDefaultFont()
	assert.Nil(err)

	defaults := Style{
		Font:      f,
		FontSize:  10,
		FontColor: ColorBlack,
	}

	ra := &MarketHoursRange{
		Min:             Date.On(NYSEOpen, Date.Date(2016, 07, 18, Date.Eastern())),
		Max:             Date.On(NYSEClose, Date.Date(2016, 07, 22, Date.Eastern())),
		MarketOpen:      NYSEOpen,
		MarketClose:     NYSEClose,
		HolidayProvider: Date.IsNYSEHoliday,
		Domain:          1024,
	}

	ticks := ra.GetTicks(r, defaults, TimeValueFormatter)
	assert.NotEmpty(ticks)
	assert.Len(ticks, 5)
	assert.NotEqual(Time.ToFloat64(ra.Min), ticks[0].Value)
	assert.NotEmpty(ticks[0].Label)
}
