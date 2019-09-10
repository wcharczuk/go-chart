package chart

import (
	"bytes"
	"image"
	"image/png"
	"math"
	"testing"
	"time"

	"github.com/blend/go-sdk/assert"

	"github.com/wcharczuk/go-chart/drawing"
)

func TestChartGetDPI(t *testing.T) {
	assert := assert.New(t)

	unset := Chart{}
	assert.Equal(DefaultDPI, unset.GetDPI())
	assert.Equal(192, unset.GetDPI(192))

	set := Chart{DPI: 128}
	assert.Equal(128, set.GetDPI())
	assert.Equal(128, set.GetDPI(192))
}

func TestChartGetFont(t *testing.T) {
	assert := assert.New(t)

	f, err := GetDefaultFont()
	assert.Nil(err)

	unset := Chart{}
	assert.Nil(unset.GetFont())

	set := Chart{Font: f}
	assert.NotNil(set.GetFont())
}

func TestChartGetWidth(t *testing.T) {
	assert := assert.New(t)

	unset := Chart{}
	assert.Equal(DefaultChartWidth, unset.GetWidth())

	set := Chart{Width: DefaultChartWidth + 10}
	assert.Equal(DefaultChartWidth+10, set.GetWidth())
}

func TestChartGetHeight(t *testing.T) {
	assert := assert.New(t)

	unset := Chart{}
	assert.Equal(DefaultChartHeight, unset.GetHeight())

	set := Chart{Height: DefaultChartHeight + 10}
	assert.Equal(DefaultChartHeight+10, set.GetHeight())
}

func TestChartGetRanges(t *testing.T) {
	assert := assert.New(t)

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
	assert.Equal(-2.0, xrange.GetMin())
	assert.Equal(5.0, xrange.GetMax())

	assert.Equal(-2.1, yrange.GetMin())
	assert.Equal(4.5, yrange.GetMax())

	assert.Equal(10.0, yrangeAlt.GetMin())
	assert.Equal(14.0, yrangeAlt.GetMax())

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
	assert.Equal(9.8, xr2.GetMin())
	assert.Equal(19.8, xr2.GetMax())

	assert.Equal(9.9, yr2.GetMin())
	assert.Equal(19.9, yr2.GetMax())

	assert.Equal(9.7, yra2.GetMin())
	assert.Equal(19.7, yra2.GetMax())
}

func TestChartGetRangesUseTicks(t *testing.T) {
	assert := assert.New(t)

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
	assert.Equal(-2.0, xr.GetMin())
	assert.Equal(2.0, xr.GetMax())
	assert.Equal(0.0, yr.GetMin())
	assert.Equal(5.0, yr.GetMax())
	assert.True(yar.IsZero(), yar.String())
}

func TestChartGetRangesUseUserRanges(t *testing.T) {
	assert := assert.New(t)

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
	assert.Equal(-2.0, xr.GetMin())
	assert.Equal(2.0, xr.GetMax())
	assert.Equal(-5.0, yr.GetMin())
	assert.Equal(5.0, yr.GetMax())
	assert.True(yar.IsZero(), yar.String())
}

func TestChartGetBackgroundStyle(t *testing.T) {
	assert := assert.New(t)

	c := Chart{
		Background: Style{
			FillColor: drawing.ColorBlack,
		},
	}

	bs := c.getBackgroundStyle()
	assert.Equal(bs.FillColor.String(), drawing.ColorBlack.String())
}

func TestChartGetCanvasStyle(t *testing.T) {
	assert := assert.New(t)

	c := Chart{
		Canvas: Style{
			FillColor: drawing.ColorBlack,
		},
	}

	bs := c.getCanvasStyle()
	assert.Equal(bs.FillColor.String(), drawing.ColorBlack.String())
}

func TestChartGetDefaultCanvasBox(t *testing.T) {
	assert := assert.New(t)

	c := Chart{}
	canvasBoxDefault := c.getDefaultCanvasBox()
	assert.False(canvasBoxDefault.IsZero())
	assert.Equal(DefaultBackgroundPadding.Top, canvasBoxDefault.Top)
	assert.Equal(DefaultBackgroundPadding.Left, canvasBoxDefault.Left)
	assert.Equal(c.GetWidth()-DefaultBackgroundPadding.Right, canvasBoxDefault.Right)
	assert.Equal(c.GetHeight()-DefaultBackgroundPadding.Bottom, canvasBoxDefault.Bottom)

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
	assert.False(canvasBoxCustom.IsZero())
	assert.Equal(DefaultBackgroundPadding.Top+1, canvasBoxCustom.Top)
	assert.Equal(DefaultBackgroundPadding.Left+1, canvasBoxCustom.Left)
	assert.Equal(c.GetWidth()-(DefaultBackgroundPadding.Right+1), canvasBoxCustom.Right)
	assert.Equal(c.GetHeight()-(DefaultBackgroundPadding.Bottom+1), canvasBoxCustom.Bottom)
}

func TestChartGetValueFormatters(t *testing.T) {
	assert := assert.New(t)

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
	assert.NotNil(dxf)
	assert.NotNil(dyf)
	assert.NotNil(dyaf)
}

func TestChartHasAxes(t *testing.T) {
	assert := assert.New(t)

	assert.True(Chart{}.hasAxes())
	assert.False(Chart{XAxis: XAxis{Style: Hidden()}, YAxis: YAxis{Style: Hidden()}, YAxisSecondary: YAxis{Style: Hidden()}}.hasAxes())

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
	assert.True(x.hasAxes())

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
	assert.True(y.hasAxes())

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
	assert.True(ya.hasAxes())
}

func TestChartGetAxesTicks(t *testing.T) {
	assert := assert.New(t)

	r, err := PNG(1024, 1024)
	assert.Nil(err)

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
	assert.NotEmpty(xt)
	assert.NotEmpty(yt)
	assert.NotEmpty(yat)
}

func TestChartSingleSeries(t *testing.T) {
	assert := assert.New(t)
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
	assert.NotEmpty(buffer.Bytes())
}

func TestChartRegressionBadRanges(t *testing.T) {
	assert := assert.New(t)

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
	assert.True(true, "Render needs to finish.")
}

func TestChartRegressionBadRangesByUser(t *testing.T) {
	assert := assert.New(t)

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
	assert.True(true, "Render needs to finish.")
}

func TestChartValidatesSeries(t *testing.T) {
	assert := assert.New(t)

	c := Chart{
		Series: []Series{
			ContinuousSeries{
				XValues: LinearRange(1.0, 10.0),
				YValues: LinearRange(1.0, 10.0),
			},
		},
	}

	assert.Nil(c.validateSeries())

	c = Chart{
		Series: []Series{
			ContinuousSeries{
				XValues: LinearRange(1.0, 10.0),
			},
		},
	}

	assert.NotNil(c.validateSeries())
}

func TestChartCheckRanges(t *testing.T) {
	assert := assert.New(t)

	c := Chart{
		Series: []Series{
			ContinuousSeries{
				XValues: []float64{1.0, 2.0},
				YValues: []float64{3.10, 3.14},
			},
		},
	}

	xr, yr, yra := c.getRanges()
	assert.Nil(c.checkRanges(xr, yr, yra))
}

func TestChartCheckRangesWithRanges(t *testing.T) {
	assert := assert.New(t)

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
	assert.Nil(c.checkRanges(xr, yr, yra))
}

func at(i image.Image, x, y int) drawing.Color {
	return drawing.ColorFromAlphaMixedRGBA(i.At(x, y).RGBA())
}

func TestChartE2ELine(t *testing.T) {
	assert := assert.New(t)

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
	assert.Nil(err)

	// do color tests ...

	i, err := png.Decode(buffer)
	assert.Nil(err)

	// test the bottom and top of the line
	assert.Equal(drawing.ColorWhite, at(i, 0, 0))
	assert.Equal(drawing.ColorWhite, at(i, 49, 49))

	// test a line mid point
	defaultSeriesColor := GetDefaultColor(0)
	assert.Equal(defaultSeriesColor, at(i, 0, 49))
	assert.Equal(defaultSeriesColor, at(i, 49, 0))
	assert.Equal(drawing.ColorFromHex("bddbf6"), at(i, 24, 24))
}

func TestChartE2ELineWithFill(t *testing.T) {
	assert := assert.New(t)

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

	assert.Equal(5, len(c.Series[0].(ContinuousSeries).XValues))
	assert.Equal(5, len(c.Series[0].(ContinuousSeries).YValues))

	var buffer = &bytes.Buffer{}
	err := c.Render(PNG, buffer)
	assert.Nil(err)

	i, err := png.Decode(buffer)
	assert.Nil(err)

	// test the bottom and top of the line
	assert.Equal(drawing.ColorWhite, at(i, 0, 0))
	assert.Equal(drawing.ColorRed, at(i, 49, 49))

	// test a line mid point
	defaultSeriesColor := drawing.ColorBlue
	assert.Equal(defaultSeriesColor, at(i, 0, 49))
	assert.Equal(defaultSeriesColor, at(i, 49, 0))
}
