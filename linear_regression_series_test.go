package chart

import (
	"testing"

	assert "github.com/blend/go-sdk/assert"
	"github.com/wcharczuk/go-chart/seq"
)

func TestLinearRegressionSeries(t *testing.T) {
	assert := assert.New(t)

	mainSeries := ContinuousSeries{
		Name:    "A test series",
		XValues: seq.Range(1.0, 100.0),
		YValues: seq.Range(1.0, 100.0),
	}

	linRegSeries := &LinearRegressionSeries{
		InnerSeries: mainSeries,
	}

	lrx0, lry0 := linRegSeries.GetValues(0)
	assert.InDelta(1.0, lrx0, 0.0000001)
	assert.InDelta(1.0, lry0, 0.0000001)

	lrxn, lryn := linRegSeries.GetLastValues()
	assert.InDelta(100.0, lrxn, 0.0000001)
	assert.InDelta(100.0, lryn, 0.0000001)
}

func TestLinearRegressionSeriesDesc(t *testing.T) {
	assert := assert.New(t)

	mainSeries := ContinuousSeries{
		Name:    "A test series",
		XValues: seq.Range(100.0, 1.0),
		YValues: seq.Range(100.0, 1.0),
	}

	linRegSeries := &LinearRegressionSeries{
		InnerSeries: mainSeries,
	}

	lrx0, lry0 := linRegSeries.GetValues(0)
	assert.InDelta(100.0, lrx0, 0.0000001)
	assert.InDelta(100.0, lry0, 0.0000001)

	lrxn, lryn := linRegSeries.GetLastValues()
	assert.InDelta(1.0, lrxn, 0.0000001)
	assert.InDelta(1.0, lryn, 0.0000001)
}

func TestLinearRegressionSeriesWindowAndOffset(t *testing.T) {
	assert := assert.New(t)

	mainSeries := ContinuousSeries{
		Name:    "A test series",
		XValues: seq.Range(100.0, 1.0),
		YValues: seq.Range(100.0, 1.0),
	}

	linRegSeries := &LinearRegressionSeries{
		InnerSeries: mainSeries,
		Offset:      10,
		Limit:       10,
	}

	assert.Equal(10, linRegSeries.Len())

	lrx0, lry0 := linRegSeries.GetValues(0)
	assert.InDelta(90.0, lrx0, 0.0000001)
	assert.InDelta(90.0, lry0, 0.0000001)

	lrxn, lryn := linRegSeries.GetLastValues()
	assert.InDelta(80.0, lrxn, 0.0000001)
	assert.InDelta(80.0, lryn, 0.0000001)
}
