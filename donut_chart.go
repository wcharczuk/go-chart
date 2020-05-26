package chart

import (
	"errors"
	"fmt"
	"io"

	"github.com/golang/freetype/truetype"
)

// DonutChart is a chart that draws sections of a circle based on percentages with an hole.
type DonutChart struct {
	Title      string
	TitleStyle Style

	ColorPalette ColorPalette

	Width  int
	Height int
	DPI    float64

	Background Style
	Canvas     Style
	SliceStyle Style

	Font        *truetype.Font
	defaultFont *truetype.Font

	Values   []Value
	Elements []Renderable
}

// GetDPI returns the dpi for the chart.
func (dc DonutChart) GetDPI(defaults ...float64) float64 {
	if dc.DPI == 0 {
		if len(defaults) > 0 {
			return defaults[0]
		}
		return DefaultDPI
	}
	return dc.DPI
}

// GetFont returns the text font.
func (dc DonutChart) GetFont() *truetype.Font {
	if dc.Font == nil {
		return dc.defaultFont
	}
	return dc.Font
}

// GetWidth returns the chart width or the default value.
func (dc DonutChart) GetWidth() int {
	if dc.Width == 0 {
		return DefaultChartWidth
	}
	return dc.Width
}

// GetHeight returns the chart height or the default value.
func (dc DonutChart) GetHeight() int {
	if dc.Height == 0 {
		return DefaultChartWidth
	}
	return dc.Height
}

// Render renders the chart with the given renderer to the given io.Writer.
func (dc DonutChart) Render(rp RendererProvider, w io.Writer) error {
	if len(dc.Values) == 0 {
		return errors.New("please provide at least one value")
	}

	r, err := rp(dc.GetWidth(), dc.GetHeight())
	if err != nil {
		return err
	}

	if dc.Font == nil {
		defaultFont, err := GetDefaultFont()
		if err != nil {
			return err
		}
		dc.defaultFont = defaultFont
	}
	r.SetDPI(dc.GetDPI(DefaultDPI))

	canvasBox := dc.getDefaultCanvasBox()
	canvasBox = dc.getCircleAdjustedCanvasBox(canvasBox)

	dc.drawBackground(r)
	dc.drawCanvas(r, canvasBox)

	finalValues, err := dc.finalizeValues(dc.Values)
	if err != nil {
		return err
	}
	dc.drawSlices(r, canvasBox, finalValues)
	dc.drawTitle(r)
	for _, a := range dc.Elements {
		a(r, canvasBox, dc.styleDefaultsElements())
	}

	return r.Save(w)
}

func (dc DonutChart) drawBackground(r Renderer) {
	Draw.Box(r, Box{
		Right:  dc.GetWidth(),
		Bottom: dc.GetHeight(),
	}, dc.getBackgroundStyle())
}

func (dc DonutChart) drawCanvas(r Renderer, canvasBox Box) {
	Draw.Box(r, canvasBox, dc.getCanvasStyle())
}

func (dc DonutChart) drawTitle(r Renderer) {
	if len(dc.Title) > 0 && !dc.TitleStyle.Hidden {
		Draw.TextWithin(r, dc.Title, dc.Box(), dc.styleDefaultsTitle())
	}
}

func (dc DonutChart) drawSlices(r Renderer, canvasBox Box, values []Value) {
	cx, cy := canvasBox.Center()
	diameter := MinInt(canvasBox.Width(), canvasBox.Height())
	radius := float64(diameter>>1) / 1.1
	labelRadius := (radius * 2.83) / 3.0

	// draw the donut slices
	var rads, delta, delta2, total float64
	var lx, ly int

	if len(values) == 1 {
		dc.styleDonutChartValueSingle(0).WriteToRenderer(r)
		r.MoveTo(cx, cy)
		r.ArcTo(cx, cy, (radius / 1.25), (radius / 1.25), DegreesToRadians(0), DegreesToRadians(359))
		r.LineTo(cx, cy)
		r.Close()
		r.FillStroke()
	} else {
		for index, v := range values {
			v.Style.InheritFrom(dc.styleDonutChartValue(index)).WriteToRenderer(r)
			r.MoveTo(cx, cy)
			rads = PercentToRadians(total)
			delta = PercentToRadians(v.Value)

			r.ArcTo(cx, cy, (radius / 1.25), (radius / 1.25), rads, delta)

			r.LineTo(cx, cy)
			r.Close()
			r.FillStroke()
			total = total + v.Value
		}
	}

	//making the donut hole
	v := Value{Value: 100, Label: "center"}
	styletemp := dc.SliceStyle.InheritFrom(Style{
		StrokeColor: ColorWhite, StrokeWidth: 4.0, FillColor: ColorWhite, FontColor: ColorWhite,
	})
	v.Style.InheritFrom(styletemp).WriteToRenderer(r)
	r.MoveTo(cx, cy)
	r.ArcTo(cx, cy, (radius / 3.5), (radius / 3.5), DegreesToRadians(0), DegreesToRadians(359))
	r.LineTo(cx, cy)
	r.Close()
	r.FillStroke()

	// draw the labels
	total = 0
	for index, v := range values {
		v.Style.InheritFrom(dc.styleDonutChartValue(index)).WriteToRenderer(r)
		if len(v.Label) > 0 {
			delta2 = PercentToRadians(total + (v.Value / 2.0))
			delta2 = RadianAdd(delta2, _pi2)
			lx, ly = CirclePoint(cx, cy, labelRadius, delta2)

			tb := r.MeasureText(v.Label)
			lx = lx - (tb.Width() >> 1)
			ly = ly + (tb.Height() >> 1)

			r.Text(v.Label, lx, ly)
		}
		total = total + v.Value
	}
}

func (dc DonutChart) finalizeValues(values []Value) ([]Value, error) {
	finalValues := Values(values).Normalize()
	if len(finalValues) == 0 {
		return nil, fmt.Errorf("donut chart must contain at least (1) non-zero value")
	}
	return finalValues, nil
}

func (dc DonutChart) getDefaultCanvasBox() Box {
	return dc.Box()
}

func (dc DonutChart) getCircleAdjustedCanvasBox(canvasBox Box) Box {
	circleDiameter := MinInt(canvasBox.Width(), canvasBox.Height())

	square := Box{
		Right:  circleDiameter,
		Bottom: circleDiameter,
	}

	return canvasBox.Fit(square)
}

func (dc DonutChart) getBackgroundStyle() Style {
	return dc.Background.InheritFrom(dc.styleDefaultsBackground())
}

func (dc DonutChart) getCanvasStyle() Style {
	return dc.Canvas.InheritFrom(dc.styleDefaultsCanvas())
}

func (dc DonutChart) styleDefaultsCanvas() Style {
	return Style{
		FillColor:   dc.GetColorPalette().CanvasColor(),
		StrokeColor: dc.GetColorPalette().CanvasStrokeColor(),
		StrokeWidth: DefaultStrokeWidth,
	}
}

func (dc DonutChart) styleDefaultsDonutChartValue() Style {
	return Style{
		StrokeColor: dc.GetColorPalette().TextColor(),
		StrokeWidth: 4.0,
		FillColor:   dc.GetColorPalette().TextColor(),
	}
}

func (dc DonutChart) styleDonutChartValue(index int) Style {
	return dc.SliceStyle.InheritFrom(Style{
		StrokeColor: ColorWhite,
		StrokeWidth: 4.0,
		FillColor:   dc.GetColorPalette().GetSeriesColor(index),
		FontSize:    dc.getScaledFontSize(),
		FontColor:   dc.GetColorPalette().TextColor(),
		Font:        dc.GetFont(),
	})
}

func (dc DonutChart) styleDonutChartValueSingle(index int) Style {
	return dc.SliceStyle.InheritFrom(Style{
		StrokeColor: dc.GetColorPalette().GetSeriesColor(index),
		StrokeWidth: 4.0,
		FillColor:   dc.GetColorPalette().GetSeriesColor(index),
		FontSize:    dc.getScaledFontSize(),
		FontColor:   dc.GetColorPalette().TextColor(),
		Font:        dc.GetFont(),
	})
}

func (dc DonutChart) getScaledFontSize() float64 {
	effectiveDimension := MinInt(dc.GetWidth(), dc.GetHeight())
	if effectiveDimension >= 2048 {
		return 48.0
	} else if effectiveDimension >= 1024 {
		return 24.0
	} else if effectiveDimension > 512 {
		return 18.0
	} else if effectiveDimension > 256 {
		return 12.0
	}
	return 10.0
}

func (dc DonutChart) styleDefaultsBackground() Style {
	return Style{
		FillColor:   dc.GetColorPalette().BackgroundColor(),
		StrokeColor: dc.GetColorPalette().BackgroundStrokeColor(),
		StrokeWidth: DefaultStrokeWidth,
	}
}

func (dc DonutChart) styleDefaultsElements() Style {
	return Style{
		Font: dc.GetFont(),
	}
}

func (dc DonutChart) styleDefaultsTitle() Style {
	return dc.TitleStyle.InheritFrom(Style{
		FontColor:           dc.GetColorPalette().TextColor(),
		Font:                dc.GetFont(),
		FontSize:            dc.getTitleFontSize(),
		TextHorizontalAlign: TextHorizontalAlignCenter,
		TextVerticalAlign:   TextVerticalAlignTop,
		TextWrap:            TextWrapWord,
	})
}

func (dc DonutChart) getTitleFontSize() float64 {
	effectiveDimension := MinInt(dc.GetWidth(), dc.GetHeight())
	if effectiveDimension >= 2048 {
		return 48
	} else if effectiveDimension >= 1024 {
		return 24
	} else if effectiveDimension >= 512 {
		return 18
	} else if effectiveDimension >= 256 {
		return 12
	}
	return 10
}

// GetColorPalette returns the color palette for the chart.
func (dc DonutChart) GetColorPalette() ColorPalette {
	if dc.ColorPalette != nil {
		return dc.ColorPalette
	}
	return AlternateColorPalette
}

// Box returns the chart bounds as a box.
func (dc DonutChart) Box() Box {
	dpr := dc.Background.Padding.GetRight(DefaultBackgroundPadding.Right)
	dpb := dc.Background.Padding.GetBottom(DefaultBackgroundPadding.Bottom)

	return Box{
		Top:    dc.Background.Padding.GetTop(DefaultBackgroundPadding.Top),
		Left:   dc.Background.Padding.GetLeft(DefaultBackgroundPadding.Left),
		Right:  dc.GetWidth() - dpr,
		Bottom: dc.GetHeight() - dpb,
	}
}
