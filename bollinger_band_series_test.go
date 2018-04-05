package chart

import (
	"fmt"
	"math"
	"testing"

	"github.com/blend/go-sdk/assert"
	"github.com/wcharczuk/go-chart/seq"
)

func TestBollingerBandSeries(t *testing.T) {
	assert := assert.New(t)

	s1 := mockValuesProvider{
		X: seq.Range(1.0, 100.0),
		Y: seq.RandomValuesWithMax(100, 1024),
	}

	bbs := &BollingerBandsSeries{
		InnerSeries: s1,
	}

	xvalues := make([]float64, 100)
	y1values := make([]float64, 100)
	y2values := make([]float64, 100)

	for x := 0; x < 100; x++ {
		xvalues[x], y1values[x], y2values[x] = bbs.GetBoundedValues(x)
	}

	for x := bbs.GetPeriod(); x < 100; x++ {
		assert.True(y1values[x] > y2values[x], fmt.Sprintf("%v vs. %v", y1values[x], y2values[x]))
	}
}

func TestBollingerBandLastValue(t *testing.T) {
	assert := assert.New(t)

	s1 := mockValuesProvider{
		X: seq.Range(1.0, 100.0),
		Y: seq.Range(1.0, 100.0),
	}

	bbs := &BollingerBandsSeries{
		InnerSeries: s1,
	}

	x, y1, y2 := bbs.GetBoundedLastValues()
	assert.Equal(100.0, x)
	assert.Equal(101, math.Floor(y1))
	assert.Equal(83, math.Floor(y2))
}
