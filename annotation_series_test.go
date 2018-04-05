package chart

import (
	"image/color"
	"testing"

	"github.com/blend/go-sdk/assert"
	"github.com/wcharczuk/go-chart/drawing"
)

func TestAnnotationSeriesMeasure(t *testing.T) {
	assert := assert.New(t)

	as := AnnotationSeries{
		Style: Style{
			Show: true,
		},
		Annotations: []Value2{
			{XValue: 1.0, YValue: 1.0, Label: "1.0"},
			{XValue: 2.0, YValue: 2.0, Label: "2.0"},
			{XValue: 3.0, YValue: 3.0, Label: "3.0"},
			{XValue: 4.0, YValue: 4.0, Label: "4.0"},
		},
	}

	r, err := PNG(110, 110)
	assert.Nil(err)

	f, err := GetDefaultFont()
	assert.Nil(err)

	xrange := &ContinuousRange{
		Min:    1.0,
		Max:    4.0,
		Domain: 100,
	}
	yrange := &ContinuousRange{
		Min:    1.0,
		Max:    4.0,
		Domain: 100,
	}

	cb := Box{
		Top:    5,
		Left:   5,
		Right:  105,
		Bottom: 105,
	}
	sd := Style{
		FontSize: 10.0,
		Font:     f,
	}

	box := as.Measure(r, cb, xrange, yrange, sd)
	assert.False(box.IsZero())
	assert.Equal(-5.0, box.Top)
	assert.Equal(5.0, box.Left)
	assert.Equal(146.0, box.Right) //the top,left annotation sticks up 5px and out ~44px.
	assert.Equal(115.0, box.Bottom)
}

func TestAnnotationSeriesRender(t *testing.T) {
	assert := assert.New(t)

	as := AnnotationSeries{
		Style: Style{
			Show:        true,
			FillColor:   drawing.ColorWhite,
			StrokeColor: drawing.ColorBlack,
		},
		Annotations: []Value2{
			{XValue: 1.0, YValue: 1.0, Label: "1.0"},
			{XValue: 2.0, YValue: 2.0, Label: "2.0"},
			{XValue: 3.0, YValue: 3.0, Label: "3.0"},
			{XValue: 4.0, YValue: 4.0, Label: "4.0"},
		},
	}

	r, err := PNG(110, 110)
	assert.Nil(err)

	f, err := GetDefaultFont()
	assert.Nil(err)

	xrange := &ContinuousRange{
		Min:    1.0,
		Max:    4.0,
		Domain: 100,
	}
	yrange := &ContinuousRange{
		Min:    1.0,
		Max:    4.0,
		Domain: 100,
	}

	cb := Box{
		Top:    5,
		Left:   5,
		Right:  105,
		Bottom: 105,
	}
	sd := Style{
		FontSize: 10.0,
		Font:     f,
	}

	as.Render(r, cb, xrange, yrange, sd)

	rr, isRaster := r.(*rasterRenderer)
	assert.True(isRaster)
	assert.NotNil(rr)

	c := rr.i.At(38, 70)
	converted, isRGBA := color.RGBAModel.Convert(c).(color.RGBA)
	assert.True(isRGBA)
	assert.Equal(0, converted.R)
	assert.Equal(0, converted.G)
	assert.Equal(0, converted.B)
}
