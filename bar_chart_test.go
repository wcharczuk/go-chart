package chart

import (
	"bytes"
	"math"
	"testing"

	"github.com/wcharczuk/go-chart/v2/testutil"
)

func TestBarChartRender(t *testing.T) {
	// replaced new assertions helper

	bc := BarChart{
		Width: 1024,
		Title: "Test Title",
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
	testutil.AssertNil(t, err)
	testutil.AssertNotZero(t, buf.Len())
}

func TestBarChartRenderZero(t *testing.T) {
	// replaced new assertions helper

	bc := BarChart{
		Width: 1024,
		Title: "Test Title",
		Bars: []Value{
			{Value: 0.0, Label: "One"},
			{Value: 0.0, Label: "Two"},
		},
	}

	buf := bytes.NewBuffer([]byte{})
	err := bc.Render(PNG, buf)
	testutil.AssertNotNil(t, err)
}

func TestBarChartProps(t *testing.T) {
	// replaced new assertions helper

	bc := BarChart{}

	testutil.AssertEqual(t, DefaultDPI, bc.GetDPI())
	bc.DPI = 100
	testutil.AssertEqual(t, 100, bc.GetDPI())

	testutil.AssertNil(t, bc.GetFont())
	f, err := GetDefaultFont()
	testutil.AssertNil(t, err)
	bc.Font = f
	testutil.AssertNotNil(t, bc.GetFont())

	testutil.AssertEqual(t, DefaultChartWidth, bc.GetWidth())
	bc.Width = DefaultChartWidth - 1
	testutil.AssertEqual(t, DefaultChartWidth-1, bc.GetWidth())

	testutil.AssertEqual(t, DefaultChartHeight, bc.GetHeight())
	bc.Height = DefaultChartHeight - 1
	testutil.AssertEqual(t, DefaultChartHeight-1, bc.GetHeight())

	testutil.AssertEqual(t, DefaultBarSpacing, bc.GetBarSpacing())
	bc.BarSpacing = 150
	testutil.AssertEqual(t, 150, bc.GetBarSpacing())

	testutil.AssertEqual(t, DefaultBarWidth, bc.GetBarWidth())
	bc.BarWidth = 75
	testutil.AssertEqual(t, 75, bc.GetBarWidth())
}

func TestBarChartRenderNoBars(t *testing.T) {
	// replaced new assertions helper

	bc := BarChart{}
	err := bc.Render(PNG, bytes.NewBuffer([]byte{}))
	testutil.AssertNotNil(t, err)
}

func TestBarChartGetRanges(t *testing.T) {
	// replaced new assertions helper

	bc := BarChart{}

	yr := bc.getRanges()
	testutil.AssertNotNil(t, yr)
	testutil.AssertFalse(t, yr.IsZero())

	testutil.AssertEqual(t, -math.MaxFloat64, yr.GetMax())
	testutil.AssertEqual(t, math.MaxFloat64, yr.GetMin())
}

func TestBarChartGetRangesBarsMinMax(t *testing.T) {
	// replaced new assertions helper

	bc := BarChart{
		Bars: []Value{
			{Value: 1.0},
			{Value: 10.0},
		},
	}

	yr := bc.getRanges()
	testutil.AssertNotNil(t, yr)
	testutil.AssertFalse(t, yr.IsZero())

	testutil.AssertEqual(t, 10, yr.GetMax())
	testutil.AssertEqual(t, 1, yr.GetMin())
}

func TestBarChartGetRangesMinMax(t *testing.T) {
	// replaced new assertions helper

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
	testutil.AssertNotNil(t, yr)
	testutil.AssertFalse(t, yr.IsZero())

	testutil.AssertEqual(t, 15, yr.GetMax())
	testutil.AssertEqual(t, 5, yr.GetMin())
}

func TestBarChartGetRangesTicksMinMax(t *testing.T) {
	// replaced new assertions helper

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
	testutil.AssertNotNil(t, yr)
	testutil.AssertFalse(t, yr.IsZero())

	testutil.AssertEqual(t, 11, yr.GetMax())
	testutil.AssertEqual(t, 7, yr.GetMin())
}

func TestBarChartHasAxes(t *testing.T) {
	// replaced new assertions helper

	bc := BarChart{}
	testutil.AssertTrue(t, bc.hasAxes())
	bc.YAxis = YAxis{
		Style: Hidden(),
	}
	testutil.AssertFalse(t, bc.hasAxes())
}

func TestBarChartGetDefaultCanvasBox(t *testing.T) {
	// replaced new assertions helper

	bc := BarChart{}
	b := bc.getDefaultCanvasBox()
	testutil.AssertFalse(t, b.IsZero())
}

func TestBarChartSetRangeDomains(t *testing.T) {
	// replaced new assertions helper

	bc := BarChart{}
	cb := bc.box()
	yr := bc.getRanges()
	yr2 := bc.setRangeDomains(cb, yr)
	testutil.AssertNotZero(t, yr2.GetDomain())
}

func TestBarChartGetValueFormatters(t *testing.T) {
	// replaced new assertions helper

	bc := BarChart{}
	vf := bc.getValueFormatters()
	testutil.AssertNotNil(t, vf)
	testutil.AssertEqual(t, "1234.00", vf(1234.0))

	bc.YAxis.ValueFormatter = func(_ interface{}) string { return "test" }
	testutil.AssertEqual(t, "test", bc.getValueFormatters()(1234))
}

func TestBarChartGetAxesTicks(t *testing.T) {
	// replaced new assertions helper

	bc := BarChart{
		Bars: []Value{
			{Value: 1.0},
			{Value: 2.0},
			{Value: 3.0},
		},
	}

	r, err := PNG(128, 128)
	testutil.AssertNil(t, err)
	yr := bc.getRanges()
	yf := bc.getValueFormatters()

	bc.YAxis.Style.Hidden = true
	ticks := bc.getAxesTicks(r, yr, yf)
	testutil.AssertEmpty(t, ticks)

	bc.YAxis.Style.Hidden = false
	ticks = bc.getAxesTicks(r, yr, yf)
	testutil.AssertLen(t, ticks, 2)
}

func TestBarChartCalculateEffectiveBarSpacing(t *testing.T) {
	// replaced new assertions helper

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
	testutil.AssertNotZero(t, spacing)

	bc.BarWidth = 250
	spacing = bc.calculateEffectiveBarSpacing(bc.box())
	testutil.AssertZero(t, spacing)
}

func TestBarChartCalculateEffectiveBarWidth(t *testing.T) {
	// replaced new assertions helper

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
	testutil.AssertNotZero(t, spacing)

	barWidth := bc.calculateEffectiveBarWidth(bc.box(), spacing)
	testutil.AssertEqual(t, 10, barWidth)

	bc.BarWidth = 250
	spacing = bc.calculateEffectiveBarSpacing(bc.box())
	testutil.AssertZero(t, spacing)
	barWidth = bc.calculateEffectiveBarWidth(bc.box(), spacing)
	testutil.AssertEqual(t, 199, barWidth)

	testutil.AssertEqual(t, cb.Width()+1, bc.calculateTotalBarWidth(barWidth, spacing))

	bw, bs, total := bc.calculateScaledTotalWidth(cb)
	testutil.AssertEqual(t, spacing, bs)
	testutil.AssertEqual(t, barWidth, bw)
	testutil.AssertEqual(t, cb.Width()+1, total)
}

func TestBarChatGetTitleFontSize(t *testing.T) {
	// replaced new assertions helper
	size := BarChart{Width: 2049, Height: 2049}.getTitleFontSize()
	testutil.AssertEqual(t, 48, size)
	size = BarChart{Width: 1025, Height: 1025}.getTitleFontSize()
	testutil.AssertEqual(t, 24, size)
	size = BarChart{Width: 513, Height: 513}.getTitleFontSize()
	testutil.AssertEqual(t, 18, size)
	size = BarChart{Width: 257, Height: 257}.getTitleFontSize()
	testutil.AssertEqual(t, 12, size)
	size = BarChart{Width: 128, Height: 128}.getTitleFontSize()
	testutil.AssertEqual(t, 10, size)
}
