package chart

import (
	"errors"
	"io"

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
func (c *Chart) Render(rp RendererProvider, w io.Writer) error {
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
	r.SetFont(font)
	r.SetDPI(c.GetDPI(DefaultDPI))

	xrange, yrange, yrangeAlt := c.getRanges()
	canvasBox := c.getCanvasBox(r, xrange, yrange, yrangeAlt)
	xrange, yrange, yrangeAlt = c.getRangeDomains(canvasBox, xrange, yrange, yrangeAlt)

	c.drawBackground(r)
	c.drawCanvas(r, canvasBox)
	c.drawAxes(r, canvasBox, xrange, yrange, yrangeAlt)
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
	return
}

func (c Chart) getCanvasBox(r Renderer, xrange, yrange, yrangeAlt Range) Box {
	dpl := c.Background.Padding.GetLeft(DefaultBackgroundPadding.Left)
	dpr := c.Background.Padding.GetRight(DefaultBackgroundPadding.Right)
	dpb := c.Background.Padding.GetBottom(DefaultBackgroundPadding.Bottom)

	if c.YAxisSecondary.Style.Show {
		dpl = c.getYAxisSecondaryWidth(r, yrangeAlt)
	}
	if c.YAxis.Style.Show {
		dpr = c.getYAxisWidth(r, yrange)
	}
	if c.XAxis.Style.Show {
		dpb = c.getXAxisHeight(r, xrange)
	}

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

func (c Chart) getYAxisSecondaryWidth(r Renderer, ra Range) int {
	var ll string
	ticks := c.YAxisSecondary.getTicks(ra)
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

func (c Chart) getYAxisWidth(r Renderer, ra Range) int {
	var ll string
	ticks := c.YAxis.getTicks(ra)
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

func (c Chart) getXAxisHeight(r Renderer, ra Range) int {
	r.SetFontSize(c.XAxis.Style.GetFontSize(DefaultFontSize))
	r.SetFont(c.XAxis.Style.GetFont(c.Font))
	var tl int
	ticks := c.YAxis.getTicks(ra)
	for _, t := range ticks {
		_, lh := r.MeasureText(t.Label)
		if lh > tl {
			tl = lh
		}
	}
	return tl + DefaultXAxisMargin
}

func (c Chart) getRangeDomains(canvasBox Box, xrange, yrange, yrangeAlt Range) (Range, Range, Range) {
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

func (c Chart) drawAxes(r Renderer, canvasBox Box, xrange, yrange, yrangeAlt Range) {
	if c.XAxis.Style.Show {
		c.XAxis.Render(r, canvasBox, xrange)
	}
	if c.YAxis.Style.Show {
		c.YAxis.Render(r, canvasBox, yrange, YAxisPrimary)
	}
	if c.YAxisSecondary.Style.Show {
		c.YAxisSecondary.Render(r, canvasBox, yrangeAlt, YAxisSecondary)
	}
}

func (c Chart) getSeriesDefaults(seriesIndex int) Style {
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
		s.Render(r, canvasBox, xrange, yrange, c.getSeriesDefaults(seriesIndex))
	} else if s.GetYAxis() == YAxisSecondary {
		s.Render(r, canvasBox, xrange, yrange, c.getSeriesDefaults(seriesIndex))
	}
}

func (c Chart) drawTitle(r Renderer) error {
	if len(c.Title) > 0 && c.TitleStyle.Show {
		r.SetFontColor(c.Canvas.GetFontColor(DefaultTextColor))
		titleFontSize := c.Canvas.GetFontSize(DefaultTitleFontSize)
		r.SetFontSize(titleFontSize)
		textWidth, _ := r.MeasureText(c.Title)
		titleX := (c.Width >> 1) - (textWidth >> 1)
		titleY := c.TitleStyle.Padding.GetTop(DefaultTitleTop) + int(titleFontSize)
		r.Text(c.Title, titleX, titleY)
	}
	return nil
}
