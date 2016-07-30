package chart

import (
	"errors"
	"fmt"
	"io"

	"github.com/golang/freetype/truetype"
)

// StackedBar is a bar within a StackedBarChart.
type StackedBar struct {
	Name   string
	Width  int
	Values []Value
}

// GetWidth returns the width of the bar.
func (sb StackedBar) GetWidth() int {
	if sb.Width == 0 {
		return 50
	}
	return sb.Width
}

// StackedBarChart is a chart that draws sections of a bar based on percentages.
type StackedBarChart struct {
	Title      string
	TitleStyle Style

	Width  int
	Height int
	DPI    float64

	Background Style
	Canvas     Style

	XAxis Style
	YAxis Style

	BarSpacing int

	Font        *truetype.Font
	defaultFont *truetype.Font

	Bars     []StackedBar
	Elements []Renderable
}

// GetDPI returns the dpi for the chart.
func (sbc StackedBarChart) GetDPI(defaults ...float64) float64 {
	if sbc.DPI == 0 {
		if len(defaults) > 0 {
			return defaults[0]
		}
		return DefaultDPI
	}
	return sbc.DPI
}

// GetFont returns the text font.
func (sbc StackedBarChart) GetFont() *truetype.Font {
	if sbc.Font == nil {
		return sbc.defaultFont
	}
	return sbc.Font
}

// GetWidth returns the chart width or the default value.
func (sbc StackedBarChart) GetWidth() int {
	if sbc.Width == 0 {
		return DefaultChartWidth
	}
	return sbc.Width
}

// GetHeight returns the chart height or the default value.
func (sbc StackedBarChart) GetHeight() int {
	if sbc.Height == 0 {
		return DefaultChartWidth
	}
	return sbc.Height
}

// GetBarSpacing returns the spacing between bars.
func (sbc StackedBarChart) GetBarSpacing() int {
	if sbc.BarSpacing == 0 {
		return 100
	}
	return sbc.BarSpacing
}

// Render renders the chart with the given renderer to the given io.Writer.
func (sbc StackedBarChart) Render(rp RendererProvider, w io.Writer) error {
	if len(sbc.Bars) == 0 {
		return errors.New("Please provide at least one bar.")
	}

	r, err := rp(sbc.GetWidth(), sbc.GetHeight())
	if err != nil {
		return err
	}

	if sbc.Font == nil {
		defaultFont, err := GetDefaultFont()
		if err != nil {
			return err
		}
		sbc.defaultFont = defaultFont
	}
	r.SetDPI(sbc.GetDPI(DefaultDPI))

	canvasBox := sbc.getAdjustedCanvasBox(sbc.getDefaultCanvasBox())
	sbc.drawBars(r, canvasBox)
	sbc.drawXAxis(r, canvasBox)
	sbc.drawYAxis(r, canvasBox)

	sbc.drawTitle(r)
	for _, a := range sbc.Elements {
		a(r, canvasBox, sbc.styleDefaultsElements())
	}

	return r.Save(w)
}

func (sbc StackedBarChart) drawBars(r Renderer, canvasBox Box) {
	xoffset := canvasBox.Left
	for _, bar := range sbc.Bars {
		sbc.drawBar(r, canvasBox, xoffset, bar)
		xoffset += sbc.GetBarSpacing()
	}
}

func (sbc StackedBarChart) drawBar(r Renderer, canvasBox Box, xoffset int, bar StackedBar) int {
	bxl := xoffset + Math.AbsInt(bar.GetWidth()>>1-sbc.GetBarSpacing()>>1)
	bxr := bxl + bar.GetWidth()

	normalizedBarComponents := Values(bar.Values).Normalize()
	yoffset := canvasBox.Top
	for index, bv := range normalizedBarComponents {
		barHeight := int(bv.Value * float64(canvasBox.Height()))
		barBox := Box{Top: yoffset, Left: bxl, Right: bxr, Bottom: yoffset + barHeight}
		Draw.Box(r, barBox, bv.Style.InheritFrom(sbc.styleDefaultsStackedBarValue(index)))
		yoffset += barHeight
	}

	return bxr
}

func (sbc StackedBarChart) drawXAxis(r Renderer, canvasBox Box) {
	if sbc.XAxis.Show {
		axisStyle := sbc.XAxis.InheritFrom(sbc.styleDefaultsAxes())
		axisStyle.WriteToRenderer(r)

		r.MoveTo(canvasBox.Left, canvasBox.Bottom)
		r.LineTo(canvasBox.Right, canvasBox.Bottom)
		r.Stroke()

		r.MoveTo(canvasBox.Left, canvasBox.Bottom)
		r.LineTo(canvasBox.Left, canvasBox.Bottom+DefaultVerticalTickHeight)
		r.Stroke()

		cursor := canvasBox.Left
		for _, bar := range sbc.Bars {

			spacing := (sbc.GetBarSpacing() >> 1)

			barLabelBox := Box{
				Top:    canvasBox.Bottom + DefaultXAxisMargin,
				Left:   cursor,
				Right:  cursor + bar.GetWidth() + spacing,
				Bottom: sbc.GetHeight(),
			}
			if len(bar.Name) > 0 {
				Draw.TextWithin(r, bar.Name, barLabelBox, axisStyle)
			}
			axisStyle.WriteToRenderer(r)
			r.MoveTo(barLabelBox.Right, canvasBox.Bottom)
			r.LineTo(barLabelBox.Right, canvasBox.Bottom+DefaultVerticalTickHeight)
			r.Stroke()
			cursor += bar.GetWidth() + spacing
		}
	}
}

func (sbc StackedBarChart) drawYAxis(r Renderer, canvasBox Box) {
	if sbc.YAxis.Show {
		axisStyle := sbc.YAxis.InheritFrom(sbc.styleDefaultsAxes())
		axisStyle.WriteToRenderer(r)
		r.MoveTo(canvasBox.Right, canvasBox.Top)
		r.LineTo(canvasBox.Right, canvasBox.Bottom)
		r.Stroke()

		r.MoveTo(canvasBox.Right, canvasBox.Bottom)
		r.LineTo(canvasBox.Right+DefaultHorizontalTickWidth, canvasBox.Bottom)
		r.Stroke()

		ticks := Sequence.Float64(1.0, 0.0, 0.2)
		for _, t := range ticks {
			ty := canvasBox.Bottom - int(t*float64(canvasBox.Height()))
			r.MoveTo(canvasBox.Right, ty)
			r.LineTo(canvasBox.Right+DefaultHorizontalTickWidth, ty)
			r.Stroke()

			text := fmt.Sprintf("%0.0f%%", t*100)
			Draw.Text(r, text, canvasBox.Right+DefaultYAxisMargin, ty, axisStyle)
		}

	}
}

func (sbc StackedBarChart) drawTitle(r Renderer) {
	if len(sbc.Title) > 0 && sbc.TitleStyle.Show {
		Draw.TextWithin(r, sbc.Title, sbc.Box(), sbc.styleDefaultsTitle())
	}
}

func (sbc StackedBarChart) getDefaultCanvasBox() Box {
	return sbc.Box()
}

func (sbc StackedBarChart) getAdjustedCanvasBox(canvasBox Box) Box {
	var totalWidth int
	for index, bar := range sbc.Bars {
		totalWidth += bar.GetWidth()
		if index < len(sbc.Bars)-1 {
			totalWidth += sbc.GetBarSpacing()
		}
	}

	return Box{
		Top:    canvasBox.Top,
		Left:   canvasBox.Left,
		Right:  canvasBox.Left + totalWidth,
		Bottom: canvasBox.Bottom,
	}
}

// Box returns the chart bounds as a box.
func (sbc StackedBarChart) Box() Box {
	dpr := sbc.Background.Padding.GetRight(10)
	dpb := sbc.Background.Padding.GetBottom(50)

	return Box{
		Top:    20,
		Left:   20,
		Right:  sbc.GetWidth() - dpr,
		Bottom: sbc.GetHeight() - dpb,
	}
}

func (sbc StackedBarChart) styleDefaultsStackedBarValue(index int) Style {
	return Style{
		StrokeColor: GetAlternateColor(index),
		StrokeWidth: 3.0,
		FillColor:   GetAlternateColor(index),
	}
}

func (sbc StackedBarChart) styleDefaultsTitle() Style {
	return sbc.TitleStyle.InheritFrom(Style{
		FontColor:           DefaultTextColor,
		Font:                sbc.GetFont(),
		FontSize:            sbc.getTitleFontSize(),
		TextHorizontalAlign: TextHorizontalAlignCenter,
		TextVerticalAlign:   TextVerticalAlignTop,
		TextWrap:            TextWrapWord,
	})
}

func (sbc StackedBarChart) getTitleFontSize() float64 {
	effectiveDimension := Math.MinInt(sbc.GetWidth(), sbc.GetHeight())
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

func (sbc StackedBarChart) styleDefaultsAxes() Style {
	return Style{
		StrokeColor:         DefaultAxisColor,
		Font:                sbc.GetFont(),
		FontSize:            DefaultAxisFontSize,
		FontColor:           DefaultAxisColor,
		TextHorizontalAlign: TextHorizontalAlignCenter,
		TextVerticalAlign:   TextVerticalAlignTop,
		TextWrap:            TextWrapWord,
	}
}
func (sbc StackedBarChart) styleDefaultsElements() Style {
	return Style{
		Font: sbc.GetFont(),
	}
}
