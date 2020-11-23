package chart

import (
	"bytes"
	"image"
	"image/png"
	"math"
	"testing"
	"time"

	"github.com/wcharczuk/go-chart/v2/drawing"
	"github.com/wcharczuk/go-chart/v2/testutil"
)

func TestChartGetDPI(t *testing.T) {
	// replaced new assertions helper

	unset := Chart{}
	testutil.AssertEqual(t, DefaultDPI, unset.GetDPI())
	testutil.AssertEqual(t, 192, unset.GetDPI(192))

	set := Chart{DPI: 128}
	testutil.AssertEqual(t, 128, set.GetDPI())
	testutil.AssertEqual(t, 128, set.GetDPI(192))
}

func TestChartGetFont(t *testing.T) {
	// replaced new assertions helper

	f, err := GetDefaultFont()
	testutil.AssertNil(t, err)

	unset := Chart{}
	testutil.AssertNil(t, unset.GetFont())

	set := Chart{Font: f}
	testutil.AssertNotNil(t, set.GetFont())
}

func TestChartGetWidth(t *testing.T) {
	// replaced new assertions helper

	unset := Chart{}
	testutil.AssertEqual(t, DefaultChartWidth, unset.GetWidth())

	set := Chart{Width: DefaultChartWidth + 10}
	testutil.AssertEqual(t, DefaultChartWidth+10, set.GetWidth())
}

func TestChartGetHeight(t *testing.T) {
	// replaced new assertions helper

	unset := Chart{}
	testutil.AssertEqual(t, DefaultChartHeight, unset.GetHeight())

	set := Chart{Height: DefaultChartHeight + 10}
	testutil.AssertEqual(t, DefaultChartHeight+10, set.GetHeight())
}

func TestChartGetRanges(t *testing.T) {
	// replaced new assertions helper

	c := Chart{
		Series: []Series{
			ContinuousSeries{
				XValues: []float64{-2.0, -1.0, 0, 1.0, 2.0},
				YValues: []float64{1.0, 2.0, 3.0, 4.0, 4.5},
			},
			ContinuousSeries{
				XValues: []float64{1.0, 2.0, 3.0, 4.0, 5.0},
				YValues: []float64{-2.1, -1.0, 0, 1.0, 2.0},
			},
			ContinuousSeries{
				YAxis:   YAxisSecondary,
				XValues: []float64{1.0, 2.0, 3.0, 4.0, 5.0},
				YValues: []float64{10.0, 11.0, 12.0, 13.0, 14.0},
			},
		},
	}

	xrange, yrange, yrangeAlt := c.getRanges()
	testutil.AssertEqual(t, -2.0, xrange.GetMin())
	testutil.AssertEqual(t, 5.0, xrange.GetMax())

	testutil.AssertEqual(t, -2.1, yrange.GetMin())
	testutil.AssertEqual(t, 4.5, yrange.GetMax())

	testutil.AssertEqual(t, 10.0, yrangeAlt.GetMin())
	testutil.AssertEqual(t, 14.0, yrangeAlt.GetMax())

	cSet := Chart{
		XAxis: XAxis{
			Range: &ContinuousRange{Min: 9.8, Max: 19.8},
		},
		YAxis: YAxis{
			Range: &ContinuousRange{Min: 9.9, Max: 19.9},
		},
		YAxisSecondary: YAxis{
			Range: &ContinuousRange{Min: 9.7, Max: 19.7},
		},
		Series: []Series{
			ContinuousSeries{
				XValues: []float64{-2.0, -1.0, 0, 1.0, 2.0},
				YValues: []float64{1.0, 2.0, 3.0, 4.0, 4.5},
			},
			ContinuousSeries{
				XValues: []float64{1.0, 2.0, 3.0, 4.0, 5.0},
				YValues: []float64{-2.1, -1.0, 0, 1.0, 2.0},
			},
			ContinuousSeries{
				YAxis:   YAxisSecondary,
				XValues: []float64{1.0, 2.0, 3.0, 4.0, 5.0},
				YValues: []float64{10.0, 11.0, 12.0, 13.0, 14.0},
			},
		},
	}

	xr2, yr2, yra2 := cSet.getRanges()
	testutil.AssertEqual(t, 9.8, xr2.GetMin())
	testutil.AssertEqual(t, 19.8, xr2.GetMax())

	testutil.AssertEqual(t, 9.9, yr2.GetMin())
	testutil.AssertEqual(t, 19.9, yr2.GetMax())

	testutil.AssertEqual(t, 9.7, yra2.GetMin())
	testutil.AssertEqual(t, 19.7, yra2.GetMax())
}

func TestChartGetRangesUseTicks(t *testing.T) {
	// replaced new assertions helper

	// this test asserts that ticks should supercede manual ranges when generating the overall ranges.

	c := Chart{
		YAxis: YAxis{
			Ticks: []Tick{
				{0.0, "Zero"},
				{1.0, "1.0"},
				{2.0, "2.0"},
				{3.0, "3.0"},
				{4.0, "4.0"},
				{5.0, "Five"},
			},
			Range: &ContinuousRange{
				Min: -5.0,
				Max: 5.0,
			},
		},
		Series: []Series{
			ContinuousSeries{
				XValues: []float64{-2.0, -1.0, 0, 1.0, 2.0},
				YValues: []float64{1.0, 2.0, 3.0, 4.0, 4.5},
			},
		},
	}

	xr, yr, yar := c.getRanges()
	testutil.AssertEqual(t, -2.0, xr.GetMin())
	testutil.AssertEqual(t, 2.0, xr.GetMax())
	testutil.AssertEqual(t, 0.0, yr.GetMin())
	testutil.AssertEqual(t, 5.0, yr.GetMax())
	testutil.AssertTrue(t, yar.IsZero(), yar.String())
}

func TestChartGetRangesUseUserRanges(t *testing.T) {
	// replaced new assertions helper

	c := Chart{
		YAxis: YAxis{
			Range: &ContinuousRange{
				Min: -5.0,
				Max: 5.0,
			},
		},
		Series: []Series{
			ContinuousSeries{
				XValues: []float64{-2.0, -1.0, 0, 1.0, 2.0},
				YValues: []float64{1.0, 2.0, 3.0, 4.0, 4.5},
			},
		},
	}

	xr, yr, yar := c.getRanges()
	testutil.AssertEqual(t, -2.0, xr.GetMin())
	testutil.AssertEqual(t, 2.0, xr.GetMax())
	testutil.AssertEqual(t, -5.0, yr.GetMin())
	testutil.AssertEqual(t, 5.0, yr.GetMax())
	testutil.AssertTrue(t, yar.IsZero(), yar.String())
}

func TestChartGetBackgroundStyle(t *testing.T) {
	// replaced new assertions helper

	c := Chart{
		Background: Style{
			FillColor: drawing.ColorBlack,
		},
	}

	bs := c.getBackgroundStyle()
	testutil.AssertEqual(t, bs.FillColor.String(), drawing.ColorBlack.String())
}

func TestChartGetCanvasStyle(t *testing.T) {
	// replaced new assertions helper

	c := Chart{
		Canvas: Style{
			FillColor: drawing.ColorBlack,
		},
	}

	bs := c.getCanvasStyle()
	testutil.AssertEqual(t, bs.FillColor.String(), drawing.ColorBlack.String())
}

func TestChartGetDefaultCanvasBox(t *testing.T) {
	// replaced new assertions helper

	c := Chart{}
	canvasBoxDefault := c.getDefaultCanvasBox()
	testutil.AssertFalse(t, canvasBoxDefault.IsZero())
	testutil.AssertEqual(t, DefaultBackgroundPadding.Top, canvasBoxDefault.Top)
	testutil.AssertEqual(t, DefaultBackgroundPadding.Left, canvasBoxDefault.Left)
	testutil.AssertEqual(t, c.GetWidth()-DefaultBackgroundPadding.Right, canvasBoxDefault.Right)
	testutil.AssertEqual(t, c.GetHeight()-DefaultBackgroundPadding.Bottom, canvasBoxDefault.Bottom)

	custom := Chart{
		Background: Style{
			Padding: Box{
				Top:    DefaultBackgroundPadding.Top + 1,
				Left:   DefaultBackgroundPadding.Left + 1,
				Right:  DefaultBackgroundPadding.Right + 1,
				Bottom: DefaultBackgroundPadding.Bottom + 1,
			},
		},
	}
	canvasBoxCustom := custom.getDefaultCanvasBox()
	testutil.AssertFalse(t, canvasBoxCustom.IsZero())
	testutil.AssertEqual(t, DefaultBackgroundPadding.Top+1, canvasBoxCustom.Top)
	testutil.AssertEqual(t, DefaultBackgroundPadding.Left+1, canvasBoxCustom.Left)
	testutil.AssertEqual(t, c.GetWidth()-(DefaultBackgroundPadding.Right+1), canvasBoxCustom.Right)
	testutil.AssertEqual(t, c.GetHeight()-(DefaultBackgroundPadding.Bottom+1), canvasBoxCustom.Bottom)
}

func TestChartGetValueFormatters(t *testing.T) {
	// replaced new assertions helper

	c := Chart{
		Series: []Series{
			ContinuousSeries{
				XValues: []float64{-2.0, -1.0, 0, 1.0, 2.0},
				YValues: []float64{1.0, 2.0, 3.0, 4.0, 4.5},
			},
			ContinuousSeries{
				XValues: []float64{1.0, 2.0, 3.0, 4.0, 5.0},
				YValues: []float64{-2.1, -1.0, 0, 1.0, 2.0},
			},
			ContinuousSeries{
				YAxis:   YAxisSecondary,
				XValues: []float64{1.0, 2.0, 3.0, 4.0, 5.0},
				YValues: []float64{10.0, 11.0, 12.0, 13.0, 14.0},
			},
		},
	}

	dxf, dyf, dyaf := c.getValueFormatters()
	testutil.AssertNotNil(t, dxf)
	testutil.AssertNotNil(t, dyf)
	testutil.AssertNotNil(t, dyaf)
}

func TestChartHasAxes(t *testing.T) {
	// replaced new assertions helper

	testutil.AssertTrue(t, Chart{}.hasAxes())
	testutil.AssertFalse(t, Chart{XAxis: XAxis{Style: Hidden()}, YAxis: YAxis{Style: Hidden()}, YAxisSecondary: YAxis{Style: Hidden()}}.hasAxes())

	x := Chart{
		XAxis: XAxis{
			Style: Hidden(),
		},
		YAxis: YAxis{
			Style: Shown(),
		},
		YAxisSecondary: YAxis{
			Style: Hidden(),
		},
	}
	testutil.AssertTrue(t, x.hasAxes())

	y := Chart{
		XAxis: XAxis{
			Style: Shown(),
		},
		YAxis: YAxis{
			Style: Hidden(),
		},
		YAxisSecondary: YAxis{
			Style: Hidden(),
		},
	}
	testutil.AssertTrue(t, y.hasAxes())

	ya := Chart{
		XAxis: XAxis{
			Style: Hidden(),
		},
		YAxis: YAxis{
			Style: Hidden(),
		},
		YAxisSecondary: YAxis{
			Style: Shown(),
		},
	}
	testutil.AssertTrue(t, ya.hasAxes())
}

func TestChartGetAxesTicks(t *testing.T) {
	// replaced new assertions helper

	r, err := PNG(1024, 1024)
	testutil.AssertNil(t, err)

	c := Chart{
		XAxis: XAxis{
			Range: &ContinuousRange{Min: 9.8, Max: 19.8},
		},
		YAxis: YAxis{
			Range: &ContinuousRange{Min: 9.9, Max: 19.9},
		},
		YAxisSecondary: YAxis{
			Range: &ContinuousRange{Min: 9.7, Max: 19.7},
		},
	}
	xr, yr, yar := c.getRanges()

	xt, yt, yat := c.getAxesTicks(r, xr, yr, yar, FloatValueFormatter, FloatValueFormatter, FloatValueFormatter)
	testutil.AssertNotEmpty(t, xt)
	testutil.AssertNotEmpty(t, yt)
	testutil.AssertNotEmpty(t, yat)
}

func TestChartSingleSeries(t *testing.T) {
	// replaced new assertions helper
	now := time.Now()
	c := Chart{
		Title:  "Hello!",
		Width:  1024,
		Height: 400,
		YAxis: YAxis{
			Range: &ContinuousRange{
				Min: 0.0,
				Max: 4.0,
			},
		},
		Series: []Series{
			TimeSeries{
				Name:    "goog",
				XValues: []time.Time{now.AddDate(0, 0, -3), now.AddDate(0, 0, -2), now.AddDate(0, 0, -1)},
				YValues: []float64{1.0, 2.0, 3.0},
			},
		},
	}

	buffer := bytes.NewBuffer([]byte{})
	c.Render(PNG, buffer)
	testutil.AssertNotEmpty(t, buffer.Bytes())
}

func TestChartRegressionBadRanges(t *testing.T) {
	// replaced new assertions helper

	c := Chart{
		Series: []Series{
			ContinuousSeries{
				XValues: []float64{math.Inf(1), math.Inf(1), math.Inf(1), math.Inf(1), math.Inf(1)},
				YValues: []float64{1.0, 2.0, 3.0, 4.0, 4.5},
			},
		},
	}
	buffer := bytes.NewBuffer([]byte{})
	c.Render(PNG, buffer)
	testutil.AssertTrue(t, true, "Render needs to finish.")
}

func TestChartRegressionBadRangesByUser(t *testing.T) {
	// replaced new assertions helper

	c := Chart{
		YAxis: YAxis{
			Range: &ContinuousRange{
				Min: math.Inf(-1),
				Max: math.Inf(1), // this could really happen? eh.
			},
		},
		Series: []Series{
			ContinuousSeries{
				XValues: LinearRange(1.0, 10.0),
				YValues: LinearRange(1.0, 10.0),
			},
		},
	}
	buffer := bytes.NewBuffer([]byte{})
	c.Render(PNG, buffer)
	testutil.AssertTrue(t, true, "Render needs to finish.")
}

func TestChartValidatesSeries(t *testing.T) {
	// replaced new assertions helper

	c := Chart{
		Series: []Series{
			ContinuousSeries{
				XValues: LinearRange(1.0, 10.0),
				YValues: LinearRange(1.0, 10.0),
			},
		},
	}

	testutil.AssertNil(t, c.validateSeries())

	c = Chart{
		Series: []Series{
			ContinuousSeries{
				XValues: LinearRange(1.0, 10.0),
			},
		},
	}

	testutil.AssertNotNil(t, c.validateSeries())
}

func TestChartCheckRanges(t *testing.T) {
	// replaced new assertions helper

	c := Chart{
		Series: []Series{
			ContinuousSeries{
				XValues: []float64{1.0, 2.0},
				YValues: []float64{3.10, 3.14},
			},
		},
	}

	xr, yr, yra := c.getRanges()
	testutil.AssertNil(t, c.checkRanges(xr, yr, yra))
}

func TestChartCheckRangesWithRanges(t *testing.T) {
	// replaced new assertions helper

	c := Chart{
		XAxis: XAxis{
			Range: &ContinuousRange{
				Min: 0,
				Max: 10,
			},
		},
		YAxis: YAxis{
			Range: &ContinuousRange{
				Min: 0,
				Max: 5,
			},
		},
		Series: []Series{
			ContinuousSeries{
				XValues: []float64{1.0, 2.0},
				YValues: []float64{3.14, 3.14},
			},
		},
	}

	xr, yr, yra := c.getRanges()
	testutil.AssertNil(t, c.checkRanges(xr, yr, yra))
}

func at(i image.Image, x, y int) drawing.Color {
	return drawing.ColorFromAlphaMixedRGBA(i.At(x, y).RGBA())
}

func TestChartE2ELine(t *testing.T) {
	// replaced new assertions helper

	c := Chart{
		Height:         50,
		Width:          50,
		TitleStyle:     Hidden(),
		XAxis:          HideXAxis(),
		YAxis:          HideYAxis(),
		YAxisSecondary: HideYAxis(),
		Canvas: Style{
			Padding: BoxZero,
		},
		Background: Style{
			Padding: BoxZero,
		},
		Series: []Series{
			ContinuousSeries{
				XValues: LinearRangeWithStep(0, 4, 1),
				YValues: LinearRangeWithStep(0, 4, 1),
			},
		},
	}

	var buffer = &bytes.Buffer{}
	err := c.Render(PNG, buffer)
	testutil.AssertNil(t, err)

	// do color tests ...

	i, err := png.Decode(buffer)
	testutil.AssertNil(t, err)

	// test the bottom and top of the line
	testutil.AssertEqual(t, drawing.ColorWhite, at(i, 0, 0))
	testutil.AssertEqual(t, drawing.ColorWhite, at(i, 49, 49))

	// test a line mid point
	defaultSeriesColor := GetDefaultColor(0)
	testutil.AssertEqual(t, defaultSeriesColor, at(i, 0, 49))
	testutil.AssertEqual(t, defaultSeriesColor, at(i, 49, 0))
	testutil.AssertEqual(t, drawing.ColorFromHex("bddbf6"), at(i, 24, 24))
}

func TestChartE2ELineWithFill(t *testing.T) {
	// replaced new assertions helper

	logBuffer := new(bytes.Buffer)

	c := Chart{
		Height: 50,
		Width:  50,
		Canvas: Style{
			Padding: BoxZero,
		},
		Background: Style{
			Padding: BoxZero,
		},
		TitleStyle:     Hidden(),
		XAxis:          HideXAxis(),
		YAxis:          HideYAxis(),
		YAxisSecondary: HideYAxis(),
		Series: []Series{
			ContinuousSeries{
				Style: Style{
					StrokeColor: drawing.ColorBlue,
					FillColor:   drawing.ColorRed,
				},
				XValues: LinearRangeWithStep(0, 4, 1),
				YValues: LinearRangeWithStep(0, 4, 1),
			},
		},
		Log: NewLogger(OptLoggerStdout(logBuffer), OptLoggerStderr(logBuffer)),
	}

	testutil.AssertEqual(t, 5, len(c.Series[0].(ContinuousSeries).XValues))
	testutil.AssertEqual(t, 5, len(c.Series[0].(ContinuousSeries).YValues))

	var buffer = &bytes.Buffer{}
	err := c.Render(PNG, buffer)
	testutil.AssertNil(t, err)

	i, err := png.Decode(buffer)
	testutil.AssertNil(t, err)

	// test the bottom and top of the line
	testutil.AssertEqual(t, drawing.ColorWhite, at(i, 0, 0))
	testutil.AssertEqual(t, drawing.ColorRed, at(i, 49, 49))

	// test a line mid point
	defaultSeriesColor := drawing.ColorBlue
	testutil.AssertEqual(t, defaultSeriesColor, at(i, 0, 49))
	testutil.AssertEqual(t, defaultSeriesColor, at(i, 49, 0))
}
