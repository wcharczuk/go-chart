package chart

import (
	"testing"

	"github.com/blend/go-sdk/assert"
)

func TestPercentageDifferenceSeries(t *testing.T) {
	assert := assert.New(t)

	cs := ContinuousSeries{
		XValues: LinearRange(1.0, 10.0),
		YValues: LinearRange(1.0, 10.0),
	}

	pcs := PercentChangeSeries{
		Name:        "Test Series",
		InnerSeries: cs,
	}

	assert.Equal("Test Series", pcs.GetName())
	assert.Equal(10, pcs.Len())
	x0, y0 := pcs.GetValues(0)
	assert.Equal(1.0, x0)
	assert.Equal(0, y0)

	xn, yn := pcs.GetValues(9)
	assert.Equal(10.0, xn)
	assert.Equal(9.0, yn)

	xn, yn = pcs.GetLastValues()
	assert.Equal(10.0, xn)
	assert.Equal(9.0, yn)
}
