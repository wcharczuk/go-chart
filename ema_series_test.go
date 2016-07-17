package chart

import (
	"math"
	"testing"

	"github.com/blendlabs/go-assert"
)

func TestEMASeries(t *testing.T) {
	assert := assert.New(t)

	mockSeries := mockValueProvider{
		Seq(1.0, 10.0),
		Seq(10, 1.0),
	}
	assert.Equal(10, mockSeries.Len())

	mas := &EMASeries{
		InnerSeries: mockSeries,
	}

	var yvalues []float64
	for x := 0; x < mas.Len(); x++ {
		_, y := mas.GetValue(x)
		yvalues = append(yvalues, y)
	}

	assert.Equal(10.0, yvalues[0])
	assert.True(math.Abs(yvalues[9]-3.77) < 0.01)

	lvx, lvy := mas.GetLastValue()
	assert.Equal(10.0, lvx)
	assert.True(math.Abs(lvy-3.77) < 0.01)
}
