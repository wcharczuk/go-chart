package chart

import (
	"testing"
	"time"

	assert "github.com/blendlabs/go-assert"
	"github.com/wcharczuk/go-chart/date"
)

func TestNYSEMarketHoursDelta(t *testing.T) {
	assert := assert.New(t)

	r := &NYSEMarketHoursRange{
		Min: time.Date(2016, 07, 19, 9, 30, 0, 0, date.Eastern()),
		Max: time.Date(2016, 07, 22, 16, 00, 0, 0, date.Eastern()),
	}

	assert.NotZero(r.GetDelta())
}
