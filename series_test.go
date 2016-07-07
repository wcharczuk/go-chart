package chart

import (
	"testing"
	"time"

	"github.com/blendlabs/go-assert"
)

func TestTimeSeriesGetValue(t *testing.T) {
	assert := assert.New(t)

	ts := TimeSeries{
		Name: "Test",
		XValues: []time.Time{
			time.Now().AddDate(0, 0, -5),
			time.Now().AddDate(0, 0, -4),
			time.Now().AddDate(0, 0, -3),
			time.Now().AddDate(0, 0, -2),
			time.Now().AddDate(0, 0, -1),
		},
		YValues: []float64{
			1.0, 2.0, 3.0, 4.0, 5.0,
		},
	}

	x0, y0 := ts.GetValue(0)
	assert.NotZero(x0)
	assert.Equal(1.0, y0)
}

func TestTimeSeriesRanges(t *testing.T) {
	assert := assert.New(t)

	ts := TimeSeries{
		Name: "Test",
		XValues: []time.Time{
			time.Now().AddDate(0, 0, -5),
			time.Now().AddDate(0, 0, -4),
			time.Now().AddDate(0, 0, -3),
			time.Now().AddDate(0, 0, -2),
			time.Now().AddDate(0, 0, -1),
		},
		YValues: []float64{
			1.0, 2.0, 3.0, 4.0, 5.0,
		},
	}

	xrange := ts.GetXRange(1000)
	x0 := xrange.Translate(time.Now().AddDate(0, 0, -3))
	assert.Equal(500, x0)

	yrange := ts.GetYRange(400)
	y0 := yrange.Translate(3.0)
	assert.Equal(200, y0)
}
