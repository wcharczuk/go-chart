package chart

import (
	"testing"
	"time"

	assert "github.com/blend/go-sdk/assert"
	"github.com/wcharczuk/go-chart/util"
)

func TestMarketHoursRangeGetDelta(t *testing.T) {
	assert := assert.New(t)

	r := &MarketHoursRange{
		Min:             time.Date(2016, 07, 19, 9, 30, 0, 0, util.Date.Eastern()),
		Max:             time.Date(2016, 07, 22, 16, 00, 0, 0, util.Date.Eastern()),
		MarketOpen:      util.NYSEOpen(),
		MarketClose:     util.NYSEClose(),
		HolidayProvider: util.Date.IsNYSEHoliday,
	}

	assert.NotZero(r.GetDelta())
}

func TestMarketHoursRangeTranslate(t *testing.T) {
	assert := assert.New(t)

	r := &MarketHoursRange{
		Min:             time.Date(2016, 07, 18, 9, 30, 0, 0, util.Date.Eastern()),
		Max:             time.Date(2016, 07, 22, 16, 00, 0, 0, util.Date.Eastern()),
		MarketOpen:      util.NYSEOpen(),
		MarketClose:     util.NYSEClose(),
		HolidayProvider: util.Date.IsNYSEHoliday,
		Domain:          1000,
	}

	weds := time.Date(2016, 07, 20, 9, 30, 0, 0, util.Date.Eastern())

	assert.Equal(0, r.Translate(util.Time.ToFloat64(r.Min)))
	assert.Equal(400, r.Translate(util.Time.ToFloat64(weds)))
	assert.Equal(1000, r.Translate(util.Time.ToFloat64(r.Max)))
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
		Min:             util.Date.On(util.NYSEOpen(), util.Date.Date(2016, 07, 18, util.Date.Eastern())),
		Max:             util.Date.On(util.NYSEClose(), util.Date.Date(2016, 07, 22, util.Date.Eastern())),
		MarketOpen:      util.NYSEOpen(),
		MarketClose:     util.NYSEClose(),
		HolidayProvider: util.Date.IsNYSEHoliday,
		Domain:          1024,
	}

	ticks := ra.GetTicks(r, defaults, TimeValueFormatter)
	assert.NotEmpty(ticks)
	assert.Len(5, ticks)
	assert.NotEqual(util.Time.ToFloat64(ra.Min), ticks[0].Value)
	assert.NotEmpty(ticks[0].Label)
}
