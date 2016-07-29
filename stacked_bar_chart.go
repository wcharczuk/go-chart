package chart

import (
	"errors"
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
		return 20
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
	bxl := xoffset
	bxr := xoffset + bar.GetWidth()

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

func (sbc StackedBarChart) drawTitle(r Renderer) {
	if len(sbc.Title) > 0 && sbc.TitleStyle.Show {
		r.SetFont(sbc.TitleStyle.GetFont(sbc.GetFont()))
		r.SetFontColor(sbc.TitleStyle.GetFontColor(DefaultTextColor))
		titleFontSize := sbc.TitleStyle.GetFontSize(DefaultTitleFontSize)
		r.SetFontSize(titleFontSize)

		textBox := r.MeasureText(sbc.Title)

		textWidth := textBox.Width()
		textHeight := textBox.Height()

		titleX := (sbc.GetWidth() >> 1) - (textWidth >> 1)
		titleY := sbc.TitleStyle.Padding.GetTop(DefaultTitleTop) + textHeight

		r.Text(sbc.Title, titleX, titleY)
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

	return canvasBox.OuterConstrain(sbc.Box(), Box{
		Top:    canvasBox.Top,
		Left:   canvasBox.Left,
		Right:  canvasBox.Left + totalWidth,
		Bottom: canvasBox.Bottom,
	})
}

// Box returns the chart bounds as a box.
func (sbc StackedBarChart) Box() Box {
	dpr := sbc.Background.Padding.GetRight(DefaultBackgroundPadding.Right)
	dpb := sbc.Background.Padding.GetBottom(DefaultBackgroundPadding.Bottom)

	return Box{
		Top:    sbc.Background.Padding.GetTop(DefaultBackgroundPadding.Top),
		Left:   sbc.Background.Padding.GetLeft(DefaultBackgroundPadding.Left),
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

func (sbc StackedBarChart) styleDefaultsElements() Style {
	return Style{
		Font: sbc.GetFont(),
	}
}
