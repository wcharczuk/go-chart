package chart

import (
	"bytes"
	"math"
	"testing"

	assert "github.com/blendlabs/go-assert"
)

func TestBarChartProps(t *testing.T) {
	assert := assert.New(t)

	bc := BarChart{}

	assert.Equal(DefaultDPI, bc.GetDPI())
	bc.DPI = 100
	assert.Equal(100, bc.GetDPI())

	assert.Nil(bc.GetFont())
	f, err := GetDefaultFont()
	assert.Nil(err)
	bc.Font = f
	assert.NotNil(bc.GetFont())

	assert.Equal(DefaultChartWidth, bc.GetWidth())
	bc.Width = DefaultChartWidth - 1
	assert.Equal(DefaultChartWidth-1, bc.GetWidth())

	assert.Equal(DefaultChartHeight, bc.GetHeight())
	bc.Height = DefaultChartHeight - 1
	assert.Equal(DefaultChartHeight-1, bc.GetHeight())

	assert.Equal(DefaultBarSpacing, bc.GetBarSpacing())
	bc.BarSpacing = 150
	assert.Equal(150, bc.GetBarSpacing())

	assert.Equal(DefaultBarWidth, bc.GetBarWidth())
	bc.BarWidth = 75
	assert.Equal(75, bc.GetBarWidth())
}

func TestBarChartRenderNoBars(t *testing.T) {
	assert := assert.New(t)

	bc := BarChart{}
	err := bc.Render(PNG, bytes.NewBuffer([]byte{}))
	assert.NotNil(err)
}

func TestBarChartGetRanges(t *testing.T) {
	assert := assert.New(t)

	bc := BarChart{}

	yr := bc.getRanges()
	assert.NotNil(yr)
	assert.False(yr.IsZero())

	assert.Equal(-math.MaxFloat64, yr.GetMax())
	assert.Equal(math.MaxFloat64, yr.GetMin())
}

func TestBarChartGetRangesBarsMinMax(t *testing.T) {
	assert := assert.New(t)

	bc := BarChart{
		Bars: []Value{
			{Value: 1.0},
			{Value: 10.0},
		},
	}

	yr := bc.getRanges()
	assert.NotNil(yr)
	assert.False(yr.IsZero())

	assert.Equal(10, yr.GetMax())
	assert.Equal(1, yr.GetMin())
}

func TestBarChartGetRangesMinMax(t *testing.T) {
	assert := assert.New(t)

	bc := BarChart{
		YAxis: YAxis{
			Range: &ContinuousRange{
				Min: 5.0,
				Max: 15.0,
			},
			Ticks: []Tick{
				{Value: 7.0, Label: "Foo"},
				{Value: 11.0, Label: "Foo2"},
			},
		},
		Bars: []Value{
			{Value: 1.0},
			{Value: 10.0},
		},
	}

	yr := bc.getRanges()
	assert.NotNil(yr)
	assert.False(yr.IsZero())

	assert.Equal(15, yr.GetMax())
	assert.Equal(5, yr.GetMin())
}

func TestBarChartGetRangesTicksMinMax(t *testing.T) {
	assert := assert.New(t)

	bc := BarChart{
		YAxis: YAxis{
			Ticks: []Tick{
				{Value: 7.0, Label: "Foo"},
				{Value: 11.0, Label: "Foo2"},
			},
		},
		Bars: []Value{
			{Value: 1.0},
			{Value: 10.0},
		},
	}

	yr := bc.getRanges()
	assert.NotNil(yr)
	assert.False(yr.IsZero())

	assert.Equal(11, yr.GetMax())
	assert.Equal(7, yr.GetMin())
}

func TestBarChartHasAxes(t *testing.T) {
	assert := assert.New(t)

	bc := BarChart{}
	assert.False(bc.hasAxes())
	bc.YAxis = YAxis{
		Style: StyleShow(),
	}

	assert.True(bc.hasAxes())
}
