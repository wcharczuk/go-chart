package chart

import (
	"bytes"
	"math"
	"testing"
	"time"

	"github.com/blendlabs/go-assert"
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
	assert.Equal(-2.0, xrange.Min)
	assert.Equal(5.0, xrange.Max)

	assert.Equal(-2.1, yrange.Min)
	assert.Equal(4.5, yrange.Max)

	assert.Equal(10.0, yrangeAlt.Min)
	assert.Equal(14.0, yrangeAlt.Max)

	cSet := Chart{
		XAxis: XAxis{
			Range: Range{Min: 9.8, Max: 19.8},
		},
		YAxis: YAxis{
			Range: Range{Min: 9.9, Max: 19.9},
		},
		YAxisSecondary: YAxis{
			Range: Range{Min: 9.7, Max: 19.7},
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
	assert.Equal(9.8, xr2.Min)
	assert.Equal(19.8, xr2.Max)

	assert.Equal(9.9, yr2.Min)
	assert.Equal(19.9, yr2.Max)

	assert.Equal(9.7, yra2.Min)
	assert.Equal(19.7, yra2.Max)
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

	assert.False(Chart{}.hasAxes())

	x := Chart{
		XAxis: XAxis{
			Style: Style{
				Show: true,
			},
		},
	}
	assert.True(x.hasAxes())

	y := Chart{
		YAxis: YAxis{
			Style: Style{
				Show: true,
			},
		},
	}
	assert.True(y.hasAxes())

	ya := Chart{
		YAxisSecondary: YAxis{
			Style: Style{
				Show: true,
			},
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
			Style: Style{Show: true},
			Range: Range{Min: 9.8, Max: 19.8},
		},
		YAxis: YAxis{
			Style: Style{Show: true},
			Range: Range{Min: 9.9, Max: 19.9},
		},
		YAxisSecondary: YAxis{
			Style: Style{Show: true},
			Range: Range{Min: 9.7, Max: 19.7},
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
			Range: Range{
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
			Range: Range{
				Min: math.Inf(-1),
				Max: math.Inf(1), // this could really happen? eh.
			},
		},
		Series: []Series{
			ContinuousSeries{
				XValues: Seq(1.0, 10.0),
				YValues: Seq(1.0, 10.0),
			},
		},
	}
	buffer := bytes.NewBuffer([]byte{})
	c.Render(PNG, buffer)
	assert.True(true, "Render needs to finish.")
}
