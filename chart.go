package chart

import (
	"errors"
	"io"
	"math"

	"github.com/golang/freetype/truetype"
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
	c.YAxisSecondary.AxisType = YAxisSecondary

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

	c.drawBackground(r)

	var xt, yt, yta []Tick
	xr, yr, yra := c.getRanges()
	canvasBox := c.getDefaultCanvasBox()
	xf, yf, yfa := c.getValueFormatters()
	xr, yr, yra = c.setRangeDomains(canvasBox, xr, yr, yra)

	if c.hasAxes() {
		xt, yt, yta = c.getAxesTicks(r, xr, yr, yra, xf, yf, yfa)
		canvasBox = c.getAxisAdjustedCanvasBox(r, canvasBox, xr, yr, yra, xt, yt, yta)
		xr, yr, yra = c.setRangeDomains(canvasBox, xr, yr, yra)
		xt, yt, yta = c.getAxesTicks(r, xr, yr, yra, xf, yf, yfa)
	}

	if c.hasAnnotationSeries() {
		canvasBox = c.getAnnotationAdjustedCanvasBox(r, canvasBox, xr, yr, yra, xf, yf, yfa)
		xr, yr, yra = c.setRangeDomains(canvasBox, xr, yr, yra)
		xt, yt, yta = c.getAxesTicks(r, xr, yr, yra, xf, yf, yfa)
	}

	c.drawCanvas(r, canvasBox)
	c.drawAxes(r, canvasBox, xr, yr, yra, xt, yt, yta)

	for index, series := range c.Series {
		c.drawSeries(r, canvasBox, xr, yr, yra, series, index)
	}
	c.drawTitle(r)

	return r.Save(w)
}

func (c Chart) getRanges() (xrange, yrange, yrangeAlt Range) {
	var minx, maxx float64 = math.MaxFloat64, 0
	var miny, maxy float64 = math.MaxFloat64, 0
	var minya, maxya float64 = math.MaxFloat64, 0

	for _, s := range c.Series {
		seriesAxis := s.GetYAxis()
		if vp, isValueProvider := s.(ValueProvider); isValueProvider {
			seriesLength := vp.Len()
			for index := 0; index < seriesLength; index++ {
				vx, vy := vp.GetValue(index)

				minx = math.Min(minx, vx)
				maxx = math.Max(maxx, vx)

				if seriesAxis == YAxisPrimary {
					miny = math.Min(miny, vy)
					maxy = math.Max(maxy, vy)
				} else if seriesAxis == YAxisSecondary {
					minya = math.Min(minya, vy)
					maxya = math.Max(maxya, vy)
				}
			}
		}
	}

	if !c.XAxis.Range.IsZero() {
		xrange.Min = c.XAxis.Range.Min
		xrange.Max = c.XAxis.Range.Max
	} else {
		xrange.Min = minx
		xrange.Max = maxx
	}

	if !c.YAxis.Range.IsZero() {
		yrange.Min = c.YAxis.Range.Min
		yrange.Max = c.YAxis.Range.Max
	} else {
		yrange.Min = miny
		yrange.Max = maxy
		yrange.Min, yrange.Max = yrange.GetRoundedRangeBounds()
	}

	if !c.YAxisSecondary.Range.IsZero() {
		yrangeAlt.Min = c.YAxisSecondary.Range.Min
		yrangeAlt.Max = c.YAxisSecondary.Range.Max
	} else {
		yrangeAlt.Min = minya
		yrangeAlt.Max = maxya
		yrangeAlt.Min, yrangeAlt.Max = yrangeAlt.GetRoundedRangeBounds()
	}

	return
}

func (c Chart) getDefaultCanvasBox() Box {
	dpt := c.Background.Padding.GetTop(DefaultBackgroundPadding.Top)
	dpl := c.Background.Padding.GetLeft(DefaultBackgroundPadding.Left)
	dpr := c.Background.Padding.GetRight(DefaultBackgroundPadding.Right)
	dpb := c.Background.Padding.GetBottom(DefaultBackgroundPadding.Bottom)

	return Box{
		Top:    dpt,
		Left:   dpl,
		Right:  c.Width - dpr,
		Bottom: c.Height - dpb,
	}
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

func (c Chart) hasAxes() bool {
	return c.XAxis.Style.Show || c.YAxis.Style.Show || c.YAxisSecondary.Style.Show
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

func (c Chart) getAxisAdjustedCanvasBox(r Renderer, canvasBox Box, xr, yr, yra Range, xticks, yticks, yticksAlt []Tick) Box {
	axesOuterBox := canvasBox.Clone()
	if c.XAxis.Style.Show {
		axesBounds := c.XAxis.Measure(r, canvasBox, xr, xticks)
		axesOuterBox = axesOuterBox.Grow(axesBounds)
	}
	if c.YAxis.Style.Show {
		axesBounds := c.YAxis.Measure(r, canvasBox, yr, yticks)
		axesOuterBox = axesOuterBox.Grow(axesBounds)
	}
	if c.YAxisSecondary.Style.Show {
		axesBounds := c.YAxisSecondary.Measure(r, canvasBox, yra, yticksAlt)
		axesOuterBox = axesOuterBox.Grow(axesBounds)
	}

	return canvasBox.OuterConstrain(c.asBox(), axesOuterBox)
}

func (c Chart) setRangeDomains(canvasBox Box, xr, yr, yra Range) (xr2, yr2, yra2 Range) {
	xr2.Min, xr2.Max = xr.Min, xr.Max
	xr2.Domain = canvasBox.Width()
	yr2.Min, yr2.Max = yr.Min, yr.Max
	yr2.Domain = canvasBox.Height()
	yra2.Min, yra2.Max = yra.Min, yra.Max
	yra2.Domain = canvasBox.Height()
	return
}

func (c Chart) hasAnnotationSeries() bool {
	for _, s := range c.Series {
		if as, isAnnotationSeries := s.(AnnotationSeries); isAnnotationSeries {
			if as.Style.Show {
				return true
			}
		}
	}
	return false
}

func (c Chart) getAnnotationAdjustedCanvasBox(r Renderer, canvasBox Box, xr, yr, yra Range, xf, yf, yfa ValueFormatter) Box {
	annotationSeriesBox := canvasBox.Clone()
	for seriesIndex, s := range c.Series {
		if as, isAnnotationSeries := s.(AnnotationSeries); isAnnotationSeries {
			if as.Style.Show {
				style := c.seriesStyleDefaults(seriesIndex)
				var annotationBounds Box
				if as.YAxis == YAxisPrimary {
					annotationBounds = as.Measure(r, canvasBox, xr, yr, style)
				} else if as.YAxis == YAxisSecondary {
					annotationBounds = as.Measure(r, canvasBox, xr, yra, style)
				}

				annotationSeriesBox = annotationSeriesBox.Grow(annotationBounds)
			}
		}
	}

	return canvasBox.OuterConstrain(c.asBox(), annotationSeriesBox)
}

func (c Chart) drawBackground(r Renderer) {
	DrawBox(r, Box{Right: c.Width, Bottom: c.Height}, c.Canvas.WithDefaultsFrom(Style{
		FillColor:   DefaultBackgroundColor,
		StrokeColor: DefaultBackgroundStrokeColor,
		StrokeWidth: DefaultStrokeWidth,
	}))
}

func (c Chart) drawCanvas(r Renderer, canvasBox Box) {
	DrawBox(r, canvasBox, c.Canvas.WithDefaultsFrom(Style{
		FillColor:   DefaultCanvasColor,
		StrokeColor: DefaultCanvasStrokeColor,
		StrokeWidth: DefaultStrokeWidth,
	}))
}

func (c Chart) drawAxes(r Renderer, canvasBox Box, xrange, yrange, yrangeAlt Range, xticks, yticks, yticksAlt []Tick) {
	if c.XAxis.Style.Show {
		c.XAxis.Render(r, canvasBox, xrange, xticks)
	}
	if c.YAxis.Style.Show {
		c.YAxis.Render(r, canvasBox, yrange, yticks)
	}
	if c.YAxisSecondary.Style.Show {
		c.YAxisSecondary.Render(r, canvasBox, yrangeAlt, yticksAlt)
	}
}

func (c Chart) drawSeries(r Renderer, canvasBox Box, xrange, yrange, yrangeAlt Range, s Series, seriesIndex int) {
	if s.GetYAxis() == YAxisPrimary {
		s.Render(r, canvasBox, xrange, yrange, c.seriesStyleDefaults(seriesIndex))
	} else if s.GetYAxis() == YAxisSecondary {
		s.Render(r, canvasBox, xrange, yrangeAlt, c.seriesStyleDefaults(seriesIndex))
	}
}

func (c Chart) drawTitle(r Renderer) {
	if len(c.Title) > 0 && c.TitleStyle.Show {
		r.SetFont(c.TitleStyle.GetFont(c.Font))
		r.SetFontColor(c.TitleStyle.GetFontColor(DefaultTextColor))
		titleFontSize := c.TitleStyle.GetFontSize(DefaultTitleFontSize)
		r.SetFontSize(titleFontSize)

		textBox := r.MeasureText(c.Title)

		textWidth := textBox.Width()
		textHeight := textBox.Height()

		titleX := (c.Width >> 1) - (textWidth >> 1)
		titleY := c.TitleStyle.Padding.GetTop(DefaultTitleTop) + textHeight

		r.Text(c.Title, titleX, titleY)
	}
}

func (c Chart) seriesStyleDefaults(seriesIndex int) Style {
	strokeColor := GetDefaultSeriesStrokeColor(seriesIndex)
	return Style{
		StrokeColor: strokeColor,
		StrokeWidth: DefaultStrokeWidth,
		Font:        c.Font,
		FontSize:    DefaultFontSize,
	}
}

func (c Chart) asBox() Box {
	return Box{Right: c.Width, Bottom: c.Height}
}
