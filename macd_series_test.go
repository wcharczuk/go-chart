package chart

import (
	"testing"

	"github.com/blendlabs/go-assert"
)

func TestMACDSeries(t *testing.T) {
	assert := assert.New(t)

	mockSeries := mockValueProvider{
		Seq(1.0, 100.0),
		SeqRand(100.0, 256),
	}
	assert.Equal(100, mockSeries.Len())

	mas := &MACDSeries{
		InnerSeries: mockSeries,
	}

	var yvalues []float64
	for x := 0; x < mas.Len(); x++ {
		_, y := mas.GetValue(x)
		yvalues = append(yvalues, y)
	}

	assert.NotEmpty(yvalues)
}
