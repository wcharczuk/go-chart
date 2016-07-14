package chart

import (
	"testing"

	"github.com/blendlabs/go-assert"
)

type mockValueProvider struct {
	X []float64
	Y []float64
}

func (m mockValueProvider) Len() int {
	return MinInt(len(m.X), len(m.Y))
}

func (m mockValueProvider) GetValue(index int) (x, y float64) {
	x = m.X[index]
	y = m.Y[index]
	return
}

func TestMovingAverageSeriesGetValue(t *testing.T) {
	assert := assert.New(t)

	mockSeries := mockValueProvider{
		Seq(1.0, 10.0),
		Seq(10, 1.0),
	}
	assert.Equal(10, mockSeries.Len())

	mas := &MovingAverageSeries{
		InnerSeries: mockSeries,
		WindowSize:  10,
	}

	var yvalues []float64
	for x := 0; x < mas.Len(); x++ {
		_, y := mas.GetValue(x)
		yvalues = append(yvalues, y)
	}

	assert.Equal(10.0, yvalues[0])
	assert.Equal(9.5, yvalues[1])
	assert.Equal(9.0, yvalues[2])
	assert.Equal(8.5, yvalues[3])
	assert.Equal(8.0, yvalues[4])
	assert.Equal(7.5, yvalues[5])
	assert.Equal(7.0, yvalues[6])
	assert.Equal(6.5, yvalues[7])
	assert.Equal(6.0, yvalues[8])
}

func TestMovingAverageSeriesGetLastValue(t *testing.T) {
	assert := assert.New(t)

	mockSeries := mockValueProvider{
		Seq(1.0, 10.0),
		Seq(10, 1.0),
	}
	assert.Equal(10, mockSeries.Len())

	mas := &MovingAverageSeries{
		InnerSeries: mockSeries,
		WindowSize:  10,
	}

	lx, ly := mas.GetLastValue()
	assert.Equal(10.0, lx)
	assert.Equal(5.5, ly)
}
