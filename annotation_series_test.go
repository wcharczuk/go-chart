package chart

import (
	"image/color"
	"testing"

	"github.com/wcharczuk/go-chart/v2/drawing"
	"github.com/wcharczuk/go-chart/v2/testutil"
)

func TestAnnotationSeriesMeasure(t *testing.T) {
	// replaced new assertions helper

	as := AnnotationSeries{
		Annotations: []Value2{
			{XValue: 1.0, YValue: 1.0, Label: "1.0"},
			{XValue: 2.0, YValue: 2.0, Label: "2.0"},
			{XValue: 3.0, YValue: 3.0, Label: "3.0"},
			{XValue: 4.0, YValue: 4.0, Label: "4.0"},
		},
	}

	r, err := PNG(110, 110)
	testutil.AssertNil(t, err)

	f, err := GetDefaultFont()
	testutil.AssertNil(t, err)

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
	testutil.AssertFalse(t, box.IsZero())
	testutil.AssertEqual(t, -5.0, box.Top)
	testutil.AssertEqual(t, 5.0, box.Left)
	testutil.AssertEqual(t, 146.0, box.Right) //the top,left annotation sticks up 5px and out ~44px.
	testutil.AssertEqual(t, 115.0, box.Bottom)
}

func TestAnnotationSeriesRender(t *testing.T) {
	// replaced new assertions helper

	as := AnnotationSeries{
		Style: Style{
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
	testutil.AssertNil(t, err)

	f, err := GetDefaultFont()
	testutil.AssertNil(t, err)

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
	testutil.AssertTrue(t, isRaster)
	testutil.AssertNotNil(t, rr)

	c := rr.i.At(38, 70)
	converted, isRGBA := color.RGBAModel.Convert(c).(color.RGBA)
	testutil.AssertTrue(t, isRGBA)
	testutil.AssertEqual(t, 0, converted.R)
	testutil.AssertEqual(t, 0, converted.G)
	testutil.AssertEqual(t, 0, converted.B)
}
