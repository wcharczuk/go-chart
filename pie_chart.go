package chart

import (
	"errors"
	"io"

	"github.com/golang/freetype/truetype"
)

// PieChart is a chart that draws sections of a circle based on percentages.
type PieChart struct {
	Title      string
	TitleStyle Style

	Width  int
	Height int
	DPI    float64

	Background Style
	Canvas     Style

	Font        *truetype.Font
	defaultFont *truetype.Font

	Values   []Value
	Elements []Renderable
}

// GetDPI returns the dpi for the chart.
func (pc PieChart) GetDPI(defaults ...float64) float64 {
	if pc.DPI == 0 {
		if len(defaults) > 0 {
			return defaults[0]
		}
		return DefaultDPI
	}
	return pc.DPI
}

// GetFont returns the text font.
func (pc PieChart) GetFont() *truetype.Font {
	if pc.Font == nil {
		return pc.defaultFont
	}
	return pc.Font
}

// GetWidth returns the chart width or the default value.
func (pc PieChart) GetWidth() int {
	if pc.Width == 0 {
		return DefaultChartWidth
	}
	return pc.Width
}

// GetHeight returns the chart height or the default value.
func (pc PieChart) GetHeight() int {
	if pc.Height == 0 {
		return DefaultChartWidth
	}
	return pc.Height
}

// Render renders the chart with the given renderer to the given io.Writer.
func (pc PieChart) Render(rp RendererProvider, w io.Writer) error {
	if len(pc.Values) == 0 {
		return errors.New("Please provide at least one value.")
	}

	r, err := rp(pc.GetWidth(), pc.GetHeight())
	if err != nil {
		return err
	}

	if pc.Font == nil {
		defaultFont, err := GetDefaultFont()
		if err != nil {
			return err
		}
		pc.defaultFont = defaultFont
	}
	r.SetDPI(pc.GetDPI(DefaultDPI))

	canvasBox := pc.getDefaultCanvasBox()
	canvasBox = pc.getCircleAdjustedCanvasBox(canvasBox)

	pc.drawBackground(r)
	pc.drawCanvas(r, canvasBox)

	finalValues := pc.finalizeValues(pc.Values)
	pc.drawSlices(r, canvasBox, finalValues)
	pc.drawTitle(r)
	for _, a := range pc.Elements {
		a(r, canvasBox, pc.styleDefaultsElements())
	}

	return r.Save(w)
}

func (pc PieChart) drawBackground(r Renderer) {
	DrawBox(r, Box{
		Right:  pc.GetWidth(),
		Bottom: pc.GetHeight(),
	}, pc.getBackgroundStyle())
}

func (pc PieChart) drawCanvas(r Renderer, canvasBox Box) {
	DrawBox(r, canvasBox, pc.getCanvasStyle())
}

func (pc PieChart) drawTitle(r Renderer) {
	if len(pc.Title) > 0 && pc.TitleStyle.Show {
		r.SetFont(pc.TitleStyle.GetFont(pc.GetFont()))
		r.SetFontColor(pc.TitleStyle.GetFontColor(DefaultTextColor))
		titleFontSize := pc.TitleStyle.GetFontSize(DefaultTitleFontSize)
		r.SetFontSize(titleFontSize)

		textBox := r.MeasureText(pc.Title)

		textWidth := textBox.Width()
		textHeight := textBox.Height()

		titleX := (pc.GetWidth() >> 1) - (textWidth >> 1)
		titleY := pc.TitleStyle.Padding.GetTop(DefaultTitleTop) + textHeight

		r.Text(pc.Title, titleX, titleY)
	}
}

func (pc PieChart) drawSlices(r Renderer, canvasBox Box, values []Value) {
	cx, cy := canvasBox.Center()
	diameter := MinInt(canvasBox.Width(), canvasBox.Height())
	radius := float64(diameter >> 1)
	labelRadius := (radius * 2.0) / 3.0

	// draw the pie slices
	var rads, delta, delta2, total float64
	var lx, ly int
	for index, v := range values {
		v.Style.InheritFrom(pc.stylePieChartValue(index)).PersistToRenderer(r)

		r.MoveTo(cx, cy)
		rads = PercentToRadians(total)
		delta = PercentToRadians(v.Value)

		r.ArcTo(cx, cy, radius, radius, rads, delta)

		r.LineTo(cx, cy)
		r.Close()
		r.FillStroke()
		total = total + v.Value
	}

	// draw the labels
	total = 0
	for index, v := range values {
		v.Style.InheritFrom(pc.stylePieChartValue(index)).PersistToRenderer(r)
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

func (pc PieChart) finalizeValues(values []Value) []Value {
	return Values(values).Normalize()
}

func (pc PieChart) getDefaultCanvasBox() Box {
	return pc.Box()
}

func (pc PieChart) getCircleAdjustedCanvasBox(canvasBox Box) Box {
	circleDiameter := MinInt(canvasBox.Width(), canvasBox.Height())

	square := Box{
		Right:  circleDiameter,
		Bottom: circleDiameter,
	}

	return canvasBox.Fit(square)
}

func (pc PieChart) getBackgroundStyle() Style {
	return pc.Background.InheritFrom(pc.styleDefaultsBackground())
}

func (pc PieChart) getCanvasStyle() Style {
	return pc.Canvas.InheritFrom(pc.styleDefaultsCanvas())
}

func (pc PieChart) styleDefaultsCanvas() Style {
	return Style{
		FillColor:   DefaultCanvasColor,
		StrokeColor: DefaultCanvasStrokeColor,
		StrokeWidth: DefaultStrokeWidth,
	}
}

func (pc PieChart) styleDefaultsPieChartValue() Style {
	return Style{
		StrokeColor: ColorWhite,
		StrokeWidth: 5.0,
		FillColor:   ColorWhite,
	}
}

func (pc PieChart) stylePieChartValue(index int) Style {
	return Style{
		StrokeColor: ColorWhite,
		StrokeWidth: 5.0,
		FillColor:   GetAlternateColor(index),
		FontSize:    24.0,
		FontColor:   ColorWhite,
		Font:        pc.GetFont(),
	}
}

func (pc PieChart) styleDefaultsBackground() Style {
	return Style{
		FillColor:   DefaultBackgroundColor,
		StrokeColor: DefaultBackgroundStrokeColor,
		StrokeWidth: DefaultStrokeWidth,
	}
}

func (pc PieChart) styleDefaultsElements() Style {
	return Style{
		Font: pc.GetFont(),
	}
}

// Box returns the chart bounds as a box.
func (pc PieChart) Box() Box {
	dpr := pc.Background.Padding.GetRight(DefaultBackgroundPadding.Right)
	dpb := pc.Background.Padding.GetBottom(DefaultBackgroundPadding.Bottom)

	return Box{
		Top:    pc.Background.Padding.GetTop(DefaultBackgroundPadding.Top),
		Left:   pc.Background.Padding.GetLeft(DefaultBackgroundPadding.Left),
		Right:  pc.GetWidth() - dpr,
		Bottom: pc.GetHeight() - dpb,
	}
}
