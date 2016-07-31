package chart

import (
	"testing"
	"time"

	assert "github.com/blendlabs/go-assert"
)

func TestSequenceFloat64(t *testing.T) {
	assert := assert.New(t)

	asc := Sequence.Float64(1.0, 10.0)
	assert.Len(asc, 10)

	desc := Sequence.Float64(10.0, 1.0)
	assert.Len(desc, 10)
}

func TestSequenceMarketHours(t *testing.T) {
	assert := assert.New(t)

	today := time.Date(2016, 07, 01, 12, 0, 0, 0, Date.Eastern())
	mh := Sequence.MarketHours(today, today, NYSEOpen, NYSEClose, Date.IsNYSEHoliday)
	assert.Len(mh, 7)
}
