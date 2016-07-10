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
	Axes       Style

	XRange Range
	YRange Range

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

	canvasBox := c.calculateCanvasBox(r)
	xrange, yrange := c.initRanges(canvasBox)

	c.drawBackground(r)
	c.drawCanvas(r, canvasBox)
	c.drawAxes(r, canvasBox, xrange, yrange)
	for index, series := range c.Series {
		c.drawSeries(r, canvasBox, series, xrange, yrange)
	}
	c.drawTitle(r)
	return r.Save(w)
}

func (c Chart) getAxisWidth() int {
	asw := 0
	if c.Axes.Show {
		asw = int(c.Axes.GetStrokeWidth(DefaultAxisLineWidth))
	}
	return asw
}

func (c Chart) calculateCanvasBox(r Renderer) Box {
	dpr := DefaultBackgroundPadding.Right
	finalLabelWidth := c.calculateFinalLabelWidth(r)
	if finalLabelWidth > dpr {
		dpr = finalLabelWidth
	}
	axisBottomHeight := c.calculateBottomLabelHeight()
	dpb := DefaultBackgroundPadding.Bottom
	if dpb < axisBottomHeight {
		dpb = axisBottomHeight
	}

	cb := Box{
		Top:    c.Background.Padding.GetTop(DefaultBackgroundPadding.Top),
		Left:   c.Background.Padding.GetLeft(DefaultBackgroundPadding.Left),
		Right:  c.Width - c.Background.Padding.GetRight(dpr),
		Bottom: c.Height - c.Background.Padding.GetBottom(dpb),
	}
	cb.Height = cb.Bottom - cb.Top
	cb.Width = cb.Right - cb.Left
	return cb
}

func (c Chart) calculateFinalLabelWidth(r Renderer) int {
	var finalLabelText string
	for _, s := range c.Series {
		if vs, isValueProvider := s.(ValueProvider); isValueProvider {
			_, lv := vs.GetValue(vs.Len() - 1)
			var ll string
			if c.YRange.Formatter != nil {
				ll = c.YRange.Formatter(lv)
			} else if fp, isFormatterProvider := s.(FormatterProvider); isFormatterProvider {
				ll = fp.GetYFormatter()(lv)
			}
			if len(finalLabelText) < len(ll) {
				finalLabelText = ll
			}
		}
	}

	r.SetFontSize(c.FinalValueLabel.GetFontSize(DefaultFinalLabelFontSize))
	textWidth, _ := r.MeasureText(finalLabelText)
	asw := c.getAxisWidth()

	pl := c.FinalValueLabel.Padding.GetLeft(DefaultFinalLabelPadding.Left)
	pr := c.FinalValueLabel.Padding.GetRight(DefaultFinalLabelPadding.Right)
	lsw := int(c.FinalValueLabel.GetStrokeWidth(DefaultAxisLineWidth))

	return DefaultYAxisMargin +
		pl + pr +
		textWidth + asw + 2*lsw
}

func (c Chart) calculateBottomLabelHeight() int {
	if c.Axes.Show {
		return c.getAxisWidth() + int(math.Ceil(c.Axes.GetFontSize(DefaultAxisFontSize))) + DefaultXAxisMargin
	}
	return 0
}

func (c Chart) initRanges(canvasBox Box) (xrange Range, yrange Range) {
	//iterate over each series, pull out the min/max for x,y
	var didSetFirstValues bool
	var globalMinY, globalMinX float64
	var globalMaxY, globalMaxX float64
	for _, s := range c.Series {
		if vp, isValueProvider := s.(ValueProvider); isValueProvider {
			seriesLength := vp.Len()
			for index := 0; index < seriesLength; index++ {
				vx, vy := vp.GetValue(index)
				if didSetFirstValues {
					if globalMinX > vx {
						globalMinX = vx
					}
					if globalMinY > vy {
						globalMinY = vy
					}
					if globalMaxX < vx {
						globalMaxX = vx
					}
					if globalMaxY < vy {
						globalMaxY = vy
					}
				} else {
					globalMinX, globalMaxX = vx, vx
					globalMinY, globalMaxY = vy, vy
					didSetFirstValues = true
				}
			}
		}
		if fp, isFormatterProvider := s.(FormatterProvider); isFormatterProvider {
			if xrange.Formatter == nil {
				xrange.Formatter = fp.GetXFormatter()
			}
			if yrange.Formatter == nil {
				yrange.Formatter = fp.GetYFormatter()
			}
		}
	}

	if c.XRange.IsZero() {
		xrange.Min = globalMinX
		xrange.Max = globalMaxX
	} else {
		xrange.Min = c.XRange.Min
		xrange.Max = c.XRange.Max
	}
	if c.XRange.Formatter != nil {
		xrange.Formatter = c.XRange.Formatter
	}
	if c.XRange.Ticks != nil {
		xrange.Ticks = c.XRange.Ticks
	}
	xrange.Domain = canvasBox.Width

	if c.YRange.IsZero() {
		yrange.Min = globalMinY
		yrange.Max = globalMaxY
	} else {
		yrange.Min = c.YRange.Min
		yrange.Max = c.YRange.Max
	}
	if c.YRange.Formatter != nil {
		yrange.Formatter = c.YRange.Formatter
	}
	if c.YRange.Ticks != nil {
		yrange.Ticks = c.YRange.Ticks
	}
	yrange.Domain = canvasBox.Height

	return
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

func (c Chart) drawAxes(r Renderer, canvasBox Box, xrange, yrange Range) {
	if c.Axes.Show {
		r.SetStrokeColor(c.Axes.GetStrokeColor(DefaultAxisColor))
		r.SetStrokeWidth(c.Axes.GetStrokeWidth(DefaultStrokeWidth))
		r.MoveTo(canvasBox.Left, canvasBox.Bottom)
		r.LineTo(canvasBox.Right, canvasBox.Bottom)
		r.LineTo(canvasBox.Right, canvasBox.Top)
		r.Stroke()

		c.drawXAxisLabels(r, canvasBox, xrange)
		c.drawYAxisLabels(r, canvasBox, yrange)
	}
}

func (c Chart) generateRangeTicks(r Range, tickCount int, offset float64) []Tick {
	var ticks []Tick
	rangeTicks := Slices(tickCount, r.Max-r.Min)
	for _, rv := range rangeTicks {
		ticks = append(ticks, Tick{
			RangeValue: rv + offset,
			Label:      r.Format(rv + offset),
		})
	}
	return ticks
}

func (c Chart) drawYAxisLabels(r Renderer, canvasBox Box, yrange Range) {
	tickFontSize := c.Axes.GetFontSize(DefaultAxisFontSize)
	asw := c.getAxisWidth()
	tx := canvasBox.Right + DefaultYAxisMargin + asw

	r.SetFontColor(c.Axes.GetFontColor(DefaultAxisColor))
	r.SetFontSize(tickFontSize)

	ticks := yrange.Ticks
	if ticks == nil {
		minimumTickHeight := tickFontSize + DefaultMinimumTickVerticalSpacing
		tickCount := int(math.Floor(float64(yrange.Domain) / float64(minimumTickHeight)))

		if tickCount > DefaultMaxTickCount {
			tickCount = DefaultMaxTickCount
		}
		ticks = c.generateRangeTicks(yrange, tickCount, yrange.Min)
	}

	for _, t := range ticks {
		v := t.RangeValue
		y := yrange.Translate(v)
		ty := int(y)
		r.Text(t.Label, tx, ty)
	}
}

func (c Chart) drawXAxisLabels(r Renderer, canvasBox Box, xrange Range) {
	tickFontSize := c.Axes.GetFontSize(DefaultAxisFontSize)
	ty := canvasBox.Bottom + DefaultXAxisMargin + int(tickFontSize)

	r.SetFontColor(c.Axes.GetFontColor(DefaultAxisColor))
	r.SetFontSize(tickFontSize)

	ticks := xrange.Ticks
	if ticks == nil {
		maxLabelWidth := 60
		minimumTickWidth := maxLabelWidth + DefaultMinimumTickHorizontalSpacing
		tickCount := int(math.Floor(float64(xrange.Domain) / float64(minimumTickWidth)))

		if tickCount > DefaultMaxTickCount {
			tickCount = DefaultMaxTickCount
		}
		ticks = c.generateRangeTicks(xrange, tickCount, xrange.Min)
	}

	for _, t := range ticks {
		v := t.RangeValue
		x := xrange.Translate(v)
		tx := canvasBox.Left + int(x)
		r.Text(t.Label, tx, ty)
	}
}

func (c Chart) drawSeries(r Renderer, canvasBox Box, s Series, xrange, yrange Range) {
	return s.Render(&c, r, canvasBox, xrange, yrange)
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
