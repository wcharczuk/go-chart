package chart

import (
	"bytes"
	"math"
	"testing"

	assert "github.com/blend/go-sdk/assert"
)

func TestBarChartRender(t *testing.T) {
	assert := assert.New(t)

	bc := BarChart{
		Width:      1024,
		Title:      "Test Title",
		TitleStyle: StyleShow(),
		XAxis:      StyleShow(),
		YAxis: YAxis{
			Style: StyleShow(),
		},
		Bars: []Value{
			{Value: 1.0, Label: "One"},
			{Value: 2.0, Label: "Two"},
			{Value: 3.0, Label: "Three"},
			{Value: 4.0, Label: "Four"},
			{Value: 5.0, Label: "Five"},
		},
	}

	buf := bytes.NewBuffer([]byte{})
	err := bc.Render(PNG, buf)
	assert.Nil(err)
	assert.NotZero(buf.Len())
}

func TestBarChartRenderZero(t *testing.T) {
	assert := assert.New(t)

	bc := BarChart{
		Width:      1024,
		Title:      "Test Title",
		TitleStyle: StyleShow(),
		XAxis:      StyleShow(),
		YAxis: YAxis{
			Style: StyleShow(),
		},
		Bars: []Value{
			{Value: 0.0, Label: "One"},
			{Value: 0.0, Label: "Two"},
		},
	}

	buf := bytes.NewBuffer([]byte{})
	err := bc.Render(PNG, buf)
	assert.NotNil(err)
}

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

func TestBarChartGetDefaultCanvasBox(t *testing.T) {
	assert := assert.New(t)

	bc := BarChart{}
	b := bc.getDefaultCanvasBox()
	assert.False(b.IsZero())
}

func TestBarChartSetRangeDomains(t *testing.T) {
	assert := assert.New(t)

	bc := BarChart{}
	cb := bc.box()
	yr := bc.getRanges()
	yr2 := bc.setRangeDomains(cb, yr)
	assert.NotZero(yr2.GetDomain())
}

func TestBarChartGetValueFormatters(t *testing.T) {
	assert := assert.New(t)

	bc := BarChart{}
	vf := bc.getValueFormatters()
	assert.NotNil(vf)
	assert.Equal("1234.00", vf(1234.0))

	bc.YAxis.ValueFormatter = func(_ interface{}) string { return "test" }
	assert.Equal("test", bc.getValueFormatters()(1234))
}

func TestBarChartGetAxesTicks(t *testing.T) {
	assert := assert.New(t)

	bc := BarChart{
		Bars: []Value{
			{Value: 1.0},
			{Value: 2.0},
			{Value: 3.0},
		},
	}

	r, err := PNG(128, 128)
	assert.Nil(err)
	yr := bc.getRanges()
	yf := bc.getValueFormatters()

	ticks := bc.getAxesTicks(r, yr, yf)
	assert.Empty(ticks)

	bc.YAxis.Style.Show = true
	ticks = bc.getAxesTicks(r, yr, yf)
	assert.Len(2, ticks)
}

func TestBarChartCalculateEffectiveBarSpacing(t *testing.T) {
	assert := assert.New(t)

	bc := BarChart{
		Width:    1024,
		BarWidth: 10,
		Bars: []Value{
			{Value: 1.0, Label: "One"},
			{Value: 2.0, Label: "Two"},
			{Value: 3.0, Label: "Three"},
			{Value: 4.0, Label: "Four"},
			{Value: 5.0, Label: "Five"},
		},
	}

	spacing := bc.calculateEffectiveBarSpacing(bc.box())
	assert.NotZero(spacing)

	bc.BarWidth = 250
	spacing = bc.calculateEffectiveBarSpacing(bc.box())
	assert.Zero(spacing)
}

func TestBarChartCalculateEffectiveBarWidth(t *testing.T) {
	assert := assert.New(t)

	bc := BarChart{
		Width:    1024,
		BarWidth: 10,
		Bars: []Value{
			{Value: 1.0, Label: "One"},
			{Value: 2.0, Label: "Two"},
			{Value: 3.0, Label: "Three"},
			{Value: 4.0, Label: "Four"},
			{Value: 5.0, Label: "Five"},
		},
	}

	cb := bc.box()

	spacing := bc.calculateEffectiveBarSpacing(bc.box())
	assert.NotZero(spacing)

	barWidth := bc.calculateEffectiveBarWidth(bc.box(), spacing)
	assert.Equal(10, barWidth)

	bc.BarWidth = 250
	spacing = bc.calculateEffectiveBarSpacing(bc.box())
	assert.Zero(spacing)
	barWidth = bc.calculateEffectiveBarWidth(bc.box(), spacing)
	assert.Equal(199, barWidth)

	assert.Equal(cb.Width()+1, bc.calculateTotalBarWidth(barWidth, spacing))

	bw, bs, total := bc.calculateScaledTotalWidth(cb)
	assert.Equal(spacing, bs)
	assert.Equal(barWidth, bw)
	assert.Equal(cb.Width()+1, total)
}

func TestBarChatGetTitleFontSize(t *testing.T) {
	assert := assert.New(t)
	size := BarChart{Width: 2049, Height: 2049}.getTitleFontSize()
	assert.Equal(48, size)
	size = BarChart{Width: 1025, Height: 1025}.getTitleFontSize()
	assert.Equal(24, size)
	size = BarChart{Width: 513, Height: 513}.getTitleFontSize()
	assert.Equal(18, size)
	size = BarChart{Width: 257, Height: 257}.getTitleFontSize()
	assert.Equal(12, size)
	size = BarChart{Width: 128, Height: 128}.getTitleFontSize()
	assert.Equal(10, size)
}
