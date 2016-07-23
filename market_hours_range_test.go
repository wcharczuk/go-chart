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
