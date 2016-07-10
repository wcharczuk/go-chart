package chart

import (
	"errors"
	"io"

	"github.com/golang/freetype/truetype"
	"github.com/wcharczuk/go-chart/drawing"
)

// Chart is what we're drawing.
type Chart struct {
	Title      string
	TitleStyle Style

	Width  int
	Height int
	DPI    float64

	Background Style
	Canvas     Style

	XAxis          XAxis
	YAxis          YAxis
	YAxisSecondary YAxis

	Font   *truetype.Font
	Series []Series
}

// GetDPI returns the dpi for the chart.
func (c Chart) GetDPI(defaults ...float64) float64 {
	if c.DPI == 0 {
		if len(defaults) > 0 {
			return defaults[0]
		}
		return DefaultDPI
	}
	return c.DPI
}

// GetFont returns the text font.
func (c Chart) GetFont() (*truetype.Font, error) {
	if c.Font == nil {
		f, err := GetDefaultFont()
		if err != nil {
			return nil, err
		}
		return f, nil
	}
	return c.Font, nil
}

// Render renders the chart with the given renderer to the given io.Writer.
func (c Chart) Render(rp RendererProvider, w io.Writer) error {
	if len(c.Series) == 0 {
		return errors.New("Please provide at least one series")
	}
	r, err := rp(c.Width, c.Height)
	if err != nil {
		return err
	}

	font, err := c.GetFont()
	if err != nil {
		return err
	}
	c.Font = font
	r.SetFont(font)
	r.SetDPI(c.GetDPI(DefaultDPI))

	xrange, yrange, yrangeAlt := c.getRanges()
	canvasBox := c.getDefaultCanvasBox()
	xf, yf, yfa := c.getValueFormatters()

	xrange, yrange, yrangeAlt = c.setRangeDomains(canvasBox, xrange, yrange, yrangeAlt)
	xticks, yticks, yticksAlt := c.getAxesTicks(r, xrange, yrange, yrangeAlt, xf, yf, yfa)
	canvasBox = c.getAdjustedCanvasBox(r, canvasBox, xticks, yticks, yticksAlt)

	// we do a second pass to take the updated domains into account
	xrange, yrange, yrangeAlt = c.setRangeDomains(canvasBox, xrange, yrange, yrangeAlt)
	xticks, yticks, yticksAlt = c.getAxesTicks(r, xrange, yrange, yrangeAlt, xf, yf, yfa)
	canvasBox = c.getAdjustedCanvasBox(r, canvasBox, xticks, yticks, yticksAlt)

	c.drawBackground(r)
	c.drawCanvas(r, canvasBox)
	c.drawAxes(r, canvasBox, xrange, yrange, yrangeAlt, xticks, yticks, yticksAlt)
	for index, series := range c.Series {
		c.drawSeries(r, canvasBox, xrange, yrange, yrangeAlt, series, index)
	}
	c.drawTitle(r)
	return r.Save(w)
}

func (c Chart) getRanges() (xrange, yrange, yrangeAlt Range) {
	//iterate over each series, pull out the min/max for x,y
	var didSetFirstValues bool

	var globalMinX, globalMaxX float64
	var globalMinY, globalMaxY float64
	var globalMinYA, globalMaxYA float64

	for _, s := range c.Series {
		if vp, isValueProvider := s.(ValueProvider); isValueProvider {
			seriesLength := vp.Len()
			for index := 0; index < seriesLength; index++ {
				vx, vy := vp.GetValue(index)
				if didSetFirstValues {
					if globalMinX > vx {
						globalMinX = vx
					}
					if globalMaxX < vx {
						globalMaxX = vx
					}
					if s.GetYAxis() == YAxisPrimary {
						if globalMinY > vy {
							globalMinY = vy
						}
						if globalMaxY < vy {
							globalMaxY = vy
						}
					} else if s.GetYAxis() == YAxisSecondary {
						if globalMinYA > vy {
							globalMinYA = vy
						}
						if globalMaxYA < vy {
							globalMaxYA = vy
						}
					}
				} else {
					globalMinX, globalMaxX = vx, vx
					if s.GetYAxis() == YAxisPrimary {
						globalMinY, globalMaxY = vy, vy
					} else if s.GetYAxis() == YAxisSecondary {
						globalMinYA, globalMaxYA = vy, vy
					}
					didSetFirstValues = true
				}
			}
		}
	}

	if !c.XAxis.Range.IsZero() {
		xrange.Min = c.XAxis.Range.Min
		xrange.Max = c.XAxis.Range.Max
	} else {
		xrange.Min = globalMinX
		xrange.Max = globalMaxX
	}

	if !c.YAxis.Range.IsZero() {
		yrange.Min = c.YAxis.Range.Min
		yrange.Max = c.YAxis.Range.Max
	} else {
		yrange.Min = globalMinY
		yrange.Max = globalMaxY
	}

	if !c.YAxisSecondary.Range.IsZero() {
		yrangeAlt.Min = c.YAxisSecondary.Range.Min
		yrangeAlt.Max = c.YAxisSecondary.Range.Max
	} else {
		yrangeAlt.Min = globalMinYA
		yrangeAlt.Max = globalMaxYA
	}
	return
}

func (c Chart) getDefaultCanvasBox() Box {
	dpl := c.Background.Padding.GetLeft(DefaultBackgroundPadding.Left)
	dpr := c.Background.Padding.GetRight(DefaultBackgroundPadding.Right)
	dpb := c.Background.Padding.GetBottom(DefaultBackgroundPadding.Bottom)

	cb := Box{
		Top:    c.Background.Padding.GetTop(DefaultBackgroundPadding.Top),
		Left:   dpl,
		Right:  c.Width - dpr,
		Bottom: c.Height - dpb,
	}
	cb.Height = cb.Bottom - cb.Top
	cb.Width = cb.Right - cb.Left
	return cb
}

func (c Chart) getValueFormatters() (x, y, ya ValueFormatter) {
	for _, s := range c.Series {
		if vfp, isVfp := s.(ValueFormatterProvider); isVfp {
			sx, sy := vfp.GetValueFormatters()
			if s.GetYAxis() == YAxisPrimary {
				x = sx
				y = sy
			} else if s.GetYAxis() == YAxisSecondary {
				x = sx
				ya = sy
			}
		}
	}
	if c.XAxis.ValueFormatter != nil {
		x = c.XAxis.ValueFormatter
	}
	if c.YAxis.ValueFormatter != nil {
		y = c.YAxis.ValueFormatter
	}
	if c.YAxisSecondary.ValueFormatter != nil {
		ya = c.YAxisSecondary.ValueFormatter
	}
	return
}

func (c Chart) getAxesTicks(r Renderer, xr, yr, yar Range, xf, yf, yfa ValueFormatter) (xticks, yticks, yticksAlt []Tick) {
	if c.XAxis.Style.Show {
		xticks = c.XAxis.GetTicks(r, xr, xf)
	}
	if c.YAxis.Style.Show {
		yticks = c.YAxis.GetTicks(r, yr, yf)
	}
	if c.YAxisSecondary.Style.Show {
		yticksAlt = c.YAxisSecondary.GetTicks(r, yar, yf)
	}
	return
}

func (c Chart) getAdjustedCanvasBox(r Renderer, defaults Box, xticks, yticks, yticksAlt []Tick) Box {
	canvasBox := Box{}

	var dpl, dpr, dpb int
	if c.XAxis.Style.Show {
		dpb = c.getXAxisHeight(r, xticks)
	}
	if c.YAxis.Style.Show {
		dpr = c.getYAxisWidth(r, yticks)
	}
	if c.YAxisSecondary.Style.Show {
		dpl = c.getYAxisSecondaryWidth(r, yticksAlt)
	}

	canvasBox.Top = defaults.Top
	if dpl != 0 {
		canvasBox.Left = c.Canvas.Padding.GetLeft(dpl)
	} else {
		canvasBox.Left = defaults.Left
	}
	if dpr != 0 {
		canvasBox.Right = c.Width - c.Canvas.Padding.GetRight(dpr)
	} else {
		canvasBox.Right = defaults.Right
	}
	if dpb != 0 {
		canvasBox.Bottom = c.Height - c.Canvas.Padding.GetBottom(dpb)
	} else {
		canvasBox.Bottom = defaults.Bottom
	}

	canvasBox.Width = canvasBox.Right - canvasBox.Left
	canvasBox.Height = canvasBox.Bottom - canvasBox.Top
	return canvasBox
}

func (c Chart) getXAxisHeight(r Renderer, ticks []Tick) int {
	r.SetFontSize(c.XAxis.Style.GetFontSize(DefaultFontSize))
	r.SetFont(c.XAxis.Style.GetFont(c.Font))
	var tl int
	for _, t := range ticks {
		_, lh := r.MeasureText(t.Label)
		if lh > tl {
			tl = lh
		}
	}
	return tl + DefaultXAxisMargin
}

func (c Chart) getYAxisWidth(r Renderer, ticks []Tick) int {
	var ll string
	for _, t := range ticks {
		if len(t.Label) > len(ll) {
			ll = t.Label
		}
	}
	r.SetFontSize(c.YAxis.Style.GetFontSize(DefaultFontSize))
	r.SetFont(c.YAxis.Style.GetFont(c.Font))
	tw, _ := r.MeasureText(ll)
	return tw + DefaultYAxisMargin
}

func (c Chart) getYAxisSecondaryWidth(r Renderer, ticks []Tick) int {
	var ll string
	for _, t := range ticks {
		if len(t.Label) > len(ll) {
			ll = t.Label
		}
	}

	r.SetFontSize(c.YAxisSecondary.Style.GetFontSize(DefaultFontSize))
	r.SetFont(c.YAxisSecondary.Style.GetFont(c.Font))
	tw, _ := r.MeasureText(ll)
	return tw + DefaultYAxisMargin
}

func (c Chart) setRangeDomains(canvasBox Box, xrange, yrange, yrangeAlt Range) (Range, Range, Range) {
	xrange.Domain = canvasBox.Width
	yrange.Domain = canvasBox.Height
	yrangeAlt.Domain = canvasBox.Height
	return xrange, yrange, yrangeAlt
}

func (c Chart) drawBackground(r Renderer) {
	r.SetFillColor(c.Background.GetFillColor(DefaultBackgroundColor))
	r.SetStrokeColor(c.Background.GetStrokeColor(DefaultBackgroundStrokeColor))
	r.SetStrokeWidth(c.Background.GetStrokeWidth(DefaultStrokeWidth))
	r.MoveTo(0, 0)
	r.LineTo(c.Width, 0)
	r.LineTo(c.Width, c.Height)
	r.LineTo(0, c.Height)
	r.LineTo(0, 0)
	r.Close()
	r.FillStroke()
}

func (c Chart) drawCanvas(r Renderer, canvasBox Box) {
	r.SetFillColor(c.Canvas.GetFillColor(DefaultCanvasColor))
	r.SetStrokeColor(c.Canvas.GetStrokeColor(DefaultCanvasStrokColor))
	r.SetStrokeWidth(c.Canvas.GetStrokeWidth(DefaultStrokeWidth))
	r.MoveTo(canvasBox.Left, canvasBox.Top)
	r.LineTo(canvasBox.Right, canvasBox.Top)
	r.LineTo(canvasBox.Right, canvasBox.Bottom)
	r.LineTo(canvasBox.Left, canvasBox.Bottom)
	r.LineTo(canvasBox.Left, canvasBox.Top)
	r.Close()
	r.FillStroke()
}

func (c Chart) drawAxes(r Renderer, canvasBox Box, xrange, yrange, yrangeAlt Range, xticks, yticks, yticksAlt []Tick) {
	if c.XAxis.Style.Show {
		c.XAxis.Render(r, canvasBox, xrange, xticks)
	}
	if c.YAxis.Style.Show {
		c.YAxis.Render(r, canvasBox, yrange, YAxisPrimary, yticks)
	}
	if c.YAxisSecondary.Style.Show {
		c.YAxisSecondary.Render(r, canvasBox, yrangeAlt, YAxisSecondary, yticksAlt)
	}
}

func (c Chart) getSeriesStyleDefaults(seriesIndex int) Style {
	strokeColor := GetDefaultSeriesStrokeColor(seriesIndex)
	return Style{
		StrokeColor: strokeColor,
		StrokeWidth: DefaultStrokeWidth,
		FillColor:   strokeColor.WithAlpha(100),
		Font:        c.Font,
		FontSize:    DefaultFontSize,
	}
}

func (c Chart) drawSeries(r Renderer, canvasBox Box, xrange, yrange, yrangeAlt Range, s Series, seriesIndex int) {
	if s.GetYAxis() == YAxisPrimary {
		s.Render(r, canvasBox, xrange, yrange, c.getSeriesStyleDefaults(seriesIndex))
	} else if s.GetYAxis() == YAxisSecondary {
		s.Render(r, canvasBox, xrange, yrangeAlt, c.getSeriesStyleDefaults(seriesIndex))
	}
}

func (c Chart) drawTitle(r Renderer) {
	if len(c.Title) > 0 && c.TitleStyle.Show {
		r.SetFont(c.TitleStyle.GetFont(c.Font))
		r.SetFontColor(c.TitleStyle.GetFontColor(DefaultTextColor))
		titleFontSize := c.TitleStyle.GetFontSize(DefaultTitleFontSize)
		r.SetFontSize(titleFontSize)

		textWidthPoints, textHeightPoints := r.MeasureText(c.Title)

		textWidth := int(drawing.PointsToPixels(r.GetDPI(), float64(textWidthPoints)))
		textHeight := int(drawing.PointsToPixels(r.GetDPI(), float64(textHeightPoints)))

		titleX := (c.Width >> 1) - (textWidth >> 1)
		titleY := c.TitleStyle.Padding.GetTop(DefaultTitleTop) + textHeight

		r.Text(c.Title, titleX, titleY)
	}
}
