package chart

import (
	"testing"

	"github.com/blendlabs/go-assert"
)

func TestBollingerBandSeries(t *testing.T) {
	assert := assert.New(t)

	s1 := mockValueProvider{
		X: Seq(1.0, 100.0),
		Y: SeqRand(100, 1024),
	}

	bbs := &BollingerBandsSeries{
		InnerSeries: s1,
	}

	xvalues := make([]float64, 100)
	y1values := make([]float64, 100)
	y2values := make([]float64, 100)

	for x := 0; x < 100; x++ {
		xvalues[x], y1values[x], y2values[x] = bbs.GetBoundedValue(x)
	}

	for x := bbs.GetWindowSize(); x < 100; x++ {
		assert.True(y1values[x] > y2values[x])
	}
}
