package chart

import (
	"testing"

	"github.com/blend/go-sdk/assert"
	"github.com/wcharczuk/go-chart/seq"
	"github.com/wcharczuk/go-chart/util"
)

type mockValuesProvider struct {
	X []float64
	Y []float64
}

func (m mockValuesProvider) Len() int {
	return util.Math.MinInt(len(m.X), len(m.Y))
}

func (m mockValuesProvider) GetValues(index int) (x, y float64) {
	if index < 0 {
		panic("negative index at GetValue()")
	}
	if index >= util.Math.MinInt(len(m.X), len(m.Y)) {
		panic("index is outside the length of m.X or m.Y")
	}
	x = m.X[index]
	y = m.Y[index]
	return
}

func TestSMASeriesGetValue(t *testing.T) {
	assert := assert.New(t)

	mockSeries := mockValuesProvider{
		seq.Range(1.0, 10.0),
		seq.Range(10, 1.0),
	}
	assert.Equal(10, mockSeries.Len())

	mas := &SMASeries{
		InnerSeries: mockSeries,
		Period:      10,
	}

	var yvalues []float64
	for x := 0; x < mas.Len(); x++ {
		_, y := mas.GetValues(x)
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

func TestSMASeriesGetLastValueWindowOverlap(t *testing.T) {
	assert := assert.New(t)

	mockSeries := mockValuesProvider{
		seq.Range(1.0, 10.0),
		seq.Range(10, 1.0),
	}
	assert.Equal(10, mockSeries.Len())

	mas := &SMASeries{
		InnerSeries: mockSeries,
		Period:      15,
	}

	var yvalues []float64
	for x := 0; x < mas.Len(); x++ {
		_, y := mas.GetValues(x)
		yvalues = append(yvalues, y)
	}

	lx, ly := mas.GetLastValues()
	assert.Equal(10.0, lx)
	assert.Equal(5.5, ly)
	assert.Equal(yvalues[len(yvalues)-1], ly)
}

func TestSMASeriesGetLastValue(t *testing.T) {
	assert := assert.New(t)

	mockSeries := mockValuesProvider{
		seq.Range(1.0, 100.0),
		seq.Range(100, 1.0),
	}
	assert.Equal(100, mockSeries.Len())

	mas := &SMASeries{
		InnerSeries: mockSeries,
		Period:      10,
	}

	var yvalues []float64
	for x := 0; x < mas.Len(); x++ {
		_, y := mas.GetValues(x)
		yvalues = append(yvalues, y)
	}

	lx, ly := mas.GetLastValues()
	assert.Equal(100.0, lx)
	assert.Equal(6, ly)
	assert.Equal(yvalues[len(yvalues)-1], ly)
}
