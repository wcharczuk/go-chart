package util

import (
	"testing"
	"time"

	assert "github.com/blendlabs/go-assert"
)

func TestSequenceFloat64(t *testing.T) {
	assert := assert.New(t)

	asc := Generate.Float64(1.0, 10.0)
	assert.Len(asc, 10)

	desc := Generate.Float64(10.0, 1.0)
	assert.Len(desc, 10)
}

func TestSequenceMarketHours(t *testing.T) {
	assert := assert.New(t)

	today := time.Date(2016, 07, 01, 12, 0, 0, 0, Date.Eastern())
	mh := Generate.MarketHours(today, today, NYSEOpen(), NYSEClose(), Date.IsNYSEHoliday)
	assert.Len(mh, 8)
	assert.Equal(Date.Eastern(), mh[0].Location())
}

func TestSequenceMarketQuarters(t *testing.T) {
	assert := assert.New(t)
	today := time.Date(2016, 07, 01, 12, 0, 0, 0, Date.Eastern())
	mh := Generate.MarketHourQuarters(today, today, NYSEOpen(), NYSEClose(), Date.IsNYSEHoliday)
	assert.Len(mh, 4)
	assert.Equal(9, mh[0].Hour())
	assert.Equal(30, mh[0].Minute())
	assert.Equal(Date.Eastern(), mh[0].Location())

	assert.Equal(12, mh[1].Hour())
	assert.Equal(00, mh[1].Minute())
	assert.Equal(Date.Eastern(), mh[1].Location())

	assert.Equal(14, mh[2].Hour())
	assert.Equal(00, mh[2].Minute())
	assert.Equal(Date.Eastern(), mh[2].Location())
}

func TestSequenceHours(t *testing.T) {
	assert := assert.New(t)

	today := time.Date(2016, 07, 01, 12, 0, 0, 0, time.UTC)
	seq := Generate.Hours(today, 24)

	end := Date.End(seq)
	assert.Len(seq, 24)
	assert.Equal(2016, end.Year())
	assert.Equal(07, int(end.Month()))
	assert.Equal(02, end.Day())
	assert.Equal(11, end.Hour())
}

func TestSequenceHoursFill(t *testing.T) {
	assert := assert.New(t)

	xdata := []time.Time{
		time.Date(2016, 07, 01, 12, 0, 0, 0, time.UTC),
		time.Date(2016, 07, 01, 13, 0, 0, 0, time.UTC),
		time.Date(2016, 07, 01, 14, 0, 0, 0, time.UTC),
		time.Date(2016, 07, 02, 4, 0, 0, 0, time.UTC),
		time.Date(2016, 07, 02, 5, 0, 0, 0, time.UTC),
		time.Date(2016, 07, 03, 12, 0, 0, 0, time.UTC),
		time.Date(2016, 07, 03, 14, 0, 0, 0, time.UTC),
	}

	ydata := []float64{
		1.1,
		1.2,
		1.4,
		0.8,
		2.1,
		0.4,
		0.6,
	}

	filledTimes, filledValues := Generate.HoursFilled(xdata, ydata)
	assert.Len(filledTimes, Date.DiffHours(Date.Start(xdata), Date.End(xdata))+1)
	assert.Equal(len(filledValues), len(filledTimes))

	assert.NotZero(filledValues[0])
	assert.NotZero(filledValues[len(filledValues)-1])

	assert.NotZero(filledValues[16])
}
