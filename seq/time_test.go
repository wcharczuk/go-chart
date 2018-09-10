package seq

import (
	"testing"
	"time"

	assert "github.com/blend/go-sdk/assert"
	"github.com/wcharczuk/go-chart/util"
)

func TestTimeHours(t *testing.T) {
	assert := assert.New(t)

	today := time.Date(2016, 07, 01, 12, 0, 0, 0, time.UTC)
	seq := Time.Hours(today, 24)

	end := util.Time.End(seq...)
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

	filledTimes, filledValues := Time.HoursFilled(xdata, ydata)
	expected := util.Time.DiffHours(util.Time.Start(xdata...), util.Time.End(xdata...)) + 1
	assert.Len(filledTimes, expected)
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

	assert.InTimeDelta(util.Time.Start(times...), times[4], time.Millisecond)
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

	assert.InTimeDelta(util.Time.End(times...), times[2], time.Millisecond)
}
