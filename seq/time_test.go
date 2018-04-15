package seq

import (
	"testing"
	"time"

	assert "github.com/blend/go-sdk/assert"
	"github.com/wcharczuk/go-chart/util"
)

func TestTimeMarketHours(t *testing.T) {
	assert := assert.New(t)

	today := time.Date(2016, 07, 01, 12, 0, 0, 0, util.Date.Eastern())
	mh := Time.MarketHours(today, today, util.NYSEOpen(), util.NYSEClose(), util.Date.IsNYSEHoliday)
	assert.Len(8, mh)
	assert.Equal(util.Date.Eastern(), mh[0].Location())
}

func TestTimeMarketHourQuarters(t *testing.T) {
	assert := assert.New(t)
	today := time.Date(2016, 07, 01, 12, 0, 0, 0, util.Date.Eastern())
	mh := Time.MarketHourQuarters(today, today, util.NYSEOpen(), util.NYSEClose(), util.Date.IsNYSEHoliday)
	assert.Len(4, mh)
	assert.Equal(9, mh[0].Hour())
	assert.Equal(30, mh[0].Minute())
	assert.Equal(util.Date.Eastern(), mh[0].Location())

	assert.Equal(12, mh[1].Hour())
	assert.Equal(00, mh[1].Minute())
	assert.Equal(util.Date.Eastern(), mh[1].Location())

	assert.Equal(14, mh[2].Hour())
	assert.Equal(00, mh[2].Minute())
	assert.Equal(util.Date.Eastern(), mh[2].Location())
}

func TestTimeHours(t *testing.T) {
	assert := assert.New(t)

	today := time.Date(2016, 07, 01, 12, 0, 0, 0, time.UTC)
	seq := Time.Hours(today, 24)

	end := Time.End(seq)
	assert.Len(24, seq)
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

	filledTimes, filledValues := Time.HoursFilled(xdata, ydata)
	expected := util.Date.DiffHours(Time.Start(xdata), Time.End(xdata)) + 1
	assert.Len(expected, filledTimes)
	assert.Equal(len(filledValues), len(filledTimes))

	assert.NotZero(filledValues[0])
	assert.NotZero(filledValues[len(filledValues)-1])

	assert.NotZero(filledValues[16])
}

func TestTimeStart(t *testing.T) {
	assert := assert.New(t)

	times := []time.Time{
		time.Now().AddDate(0, 0, -4),
		time.Now().AddDate(0, 0, -2),
		time.Now().AddDate(0, 0, -1),
		time.Now().AddDate(0, 0, -3),
		time.Now().AddDate(0, 0, -5),
	}

	assert.InTimeDelta(Time.Start(times), times[4], time.Millisecond)
}

func TestTimeEnd(t *testing.T) {
	assert := assert.New(t)

	times := []time.Time{
		time.Now().AddDate(0, 0, -4),
		time.Now().AddDate(0, 0, -2),
		time.Now().AddDate(0, 0, -1),
		time.Now().AddDate(0, 0, -3),
		time.Now().AddDate(0, 0, -5),
	}

	assert.InTimeDelta(Time.End(times), times[2], time.Millisecond)
}
