package chart

import (
	"testing"

	assert "github.com/blendlabs/go-assert"
	"github.com/wcharczuk/go-chart/sequence"
)

func TestHistogramSeries(t *testing.T) {
	assert := assert.New(t)

	cs := ContinuousSeries{
		Name:    "Test Series",
		XValues: sequence.Values(1.0, 20.0),
		YValues: sequence.Values(10.0, -10.0),
	}

	hs := HistogramSeries{
		InnerSeries: cs,
	}

	for x := 0; x < hs.Len(); x++ {
		csx, csy := cs.GetValues(0)
		hsx, hsy1, hsy2 := hs.GetBoundedValues(0)
		assert.Equal(csx, hsx)
		assert.True(hsy1 > 0)
		assert.True(hsy2 <= 0)
		assert.True(csy < 0 || (csy > 0 && csy == hsy1))
		assert.True(csy > 0 || (csy < 0 && csy == hsy2))
	}
}
