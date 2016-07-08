package chart

import (
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

	Background      Style
	Canvas          Style
	Axes            Style
	FinalValueLabel Style

	XRange Range
	YRange Range

	Font   *truetype.Font
	Series []Series
}

// GetFont returns the text font.
func (c Chart) GetFont() (*truetype.Font, error) {
	if c.Font != nil {
		return c.Font, nil
	}
	return GetDefaultFont()
}

// Render renders the chart with the given renderer to the given io.Writer.
func (c *Chart) Render(provider RendererProvider, w io.Writer) error {
	r := provider(c.Width, c.Height)
	if c.hasText() {
		font, err := c.GetFont()
		if err != nil {
			return err
		}
		r.SetFont(font)
	}

	canvasBox := c.calculateCanvasBox(r)
	xrange, yrange := c.initRanges(canvasBox)

	c.drawBackground(r)
	c.drawCanvas(r, canvasBox)
	c.drawAxes(r, canvasBox, xrange, yrange)
	for index, series := range c.Series {
		c.drawSeries(r, canvasBox, index, series, xrange, yrange)
	}
	c.drawTitle(r)
	return r.Save(w)
}

func (c Chart) hasText() bool {
	return c.TitleStyle.Show || c.Axes.Show
}

func (c Chart) calculateCanvasBox(r Renderer) Box {
	dpr := DefaultBackgroundPadding.Right
	finalLabelWidth := c.calculateFinalLabelWidth(r)
	if finalLabelWidth > dpr {
		dpr = finalLabelWidth
	}
	dpb := DefaultBackgroundPadding.Bottom

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
	if !c.FinalValueLabel.Show {
		return 0
	}
	var finalLabelText string
	for _, s := range c.Series {
		_, ll := s.GetLabel(s.Len() - 1)
		if len(finalLabelText) < len(ll) {
			finalLabelText = ll
		}
	}

	r.SetFontSize(c.FinalValueLabel.GetFontSize(DefaultFinalLabelFontSize))
	textWidth := r.MeasureText(finalLabelText)

	pl := c.FinalValueLabel.Padding.GetLeft(DefaultFinalLabelPadding.Left)
	pr := c.FinalValueLabel.Padding.GetRight(DefaultFinalLabelPadding.Right)

	asw := int(c.Axes.GetStrokeWidth(DefaultAxisLineWidth))
	lsw := int(c.FinalValueLabel.GetStrokeWidth(DefaultAxisLineWidth))

	return DefaultFinalLabelDeltaWidth +
		pl + pr +
		textWidth + asw + 2*lsw
}

func (c Chart) initRanges(canvasBox Box) (xrange Range, yrange Range) {
	//iterate over each series, pull out the min/max for x,y
	var didSetFirstValues bool
	var globalMinY, globalMinX float64
	var globalMaxY, globalMaxX float64
	for _, s := range c.Series {
		seriesLength := s.Len()
		for index := 0; index < seriesLength; index++ {
			vx, vy := s.GetValue(index)
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

	if c.XRange.IsZero() {
		xrange.Min = globalMinX
		xrange.Max = globalMaxX
	} else {
		xrange.Min = c.XRange.Min
		xrange.Max = c.XRange.Max
	}
	xrange.Domain = canvasBox.Width

	if c.YRange.IsZero() {
		yrange.Min = globalMinY
		yrange.Max = globalMaxY
	} else {
		yrange.Min = c.YRange.Min
		yrange.Max = c.YRange.Max
	}
	yrange.Domain = canvasBox.Height

	return
}

func (c Chart) drawBackground(r Renderer) {
	r.SetFillColor(c.Background.GetFillColor(DefaultBackgroundColor))
	r.SetStrokeColor(c.Background.GetStrokeColor(DefaultBackgroundStrokeColor))
	r.SetLineWidth(c.Background.GetStrokeWidth(DefaultStrokeWidth))
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
	r.SetLineWidth(c.Canvas.GetStrokeWidth(DefaultStrokeWidth))
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
		r.SetLineWidth(c.Axes.GetStrokeWidth(DefaultStrokeWidth))
		r.MoveTo(canvasBox.Left, canvasBox.Bottom)
		r.LineTo(canvasBox.Right, canvasBox.Bottom)
		r.LineTo(canvasBox.Right, canvasBox.Top)
		r.Stroke()

		c.drawAxesLabels(r, xrange, yrange)
	}
}

func (c Chart) drawAxesLabels(r Renderer, xrange, yrange Range) {
	// do x axis
	// do y axis
}

func (c Chart) drawSeries(r Renderer, canvasBox Box, index int, s Series, xrange, yrange Range) {
	r.SetStrokeColor(s.GetStyle().GetStrokeColor(GetDefaultSeriesStrokeColor(index)))
	r.SetLineWidth(s.GetStyle().GetStrokeWidth(DefaultStrokeWidth))

	if s.Len() == 0 {
		return
	}

	cx := canvasBox.Left
	cy := canvasBox.Top
	cw := canvasBox.Width

	v0x, v0y := s.GetValue(0)
	x0 := cw - xrange.Translate(v0x)
	y0 := yrange.Translate(v0y)
	r.MoveTo(x0+cx, y0+cy)

	var vx, vy float64
	var x, y int
	for i := 1; i < s.Len(); i++ {
		vx, vy = s.GetValue(i)
		x = cw - xrange.Translate(vx)
		y = yrange.Translate(vy)
		r.LineTo(x+cx, y+cy)
	}
	r.Stroke()

	c.drawFinalValueLabel(r, canvasBox, index, s, yrange)
}

func (c Chart) drawFinalValueLabel(r Renderer, canvasBox Box, index int, s Series, yrange Range) {
	if c.FinalValueLabel.Show {
		_, lv := s.GetValue(s.Len() - 1)
		_, ll := s.GetLabel(s.Len() - 1)

		py := canvasBox.Top
		ly := yrange.Translate(lv) + py

		r.SetFontSize(c.FinalValueLabel.GetFontSize(DefaultFinalLabelFontSize))
		textWidth := r.MeasureText(ll)
		textHeight := int(math.Floor(DefaultFinalLabelFontSize))
		halfTextHeight := textHeight >> 1

		cx := canvasBox.Right + int(c.Axes.GetStrokeWidth(DefaultAxisLineWidth))

		pt := c.FinalValueLabel.Padding.GetTop(DefaultFinalLabelPadding.Top)
		pl := c.FinalValueLabel.Padding.GetLeft(DefaultFinalLabelPadding.Left)
		pr := c.FinalValueLabel.Padding.GetRight(DefaultFinalLabelPadding.Right)
		pb := c.FinalValueLabel.Padding.GetBottom(DefaultFinalLabelPadding.Bottom)

		textX := cx + pl + DefaultFinalLabelDeltaWidth
		textY := ly + halfTextHeight

		ltlx := cx + pl + DefaultFinalLabelDeltaWidth
		ltly := ly - (pt + halfTextHeight)

		ltrx := cx + pl + pr + textWidth
		ltry := ly - (pt + halfTextHeight)

		lbrx := cx + pl + pr + textWidth
		lbry := ly + (pb + halfTextHeight)

		lblx := cx + DefaultFinalLabelDeltaWidth
		lbly := ly + (pb + halfTextHeight)

		//draw the shape...
		r.SetFillColor(c.FinalValueLabel.GetFillColor(DefaultFinalLabelBackgroundColor))
		r.SetStrokeColor(c.FinalValueLabel.GetStrokeColor(s.GetStyle().GetStrokeColor(GetDefaultSeriesStrokeColor(index))))
		r.SetLineWidth(c.FinalValueLabel.GetStrokeWidth(DefaultAxisLineWidth))
		r.MoveTo(cx, ly)
		r.LineTo(ltlx, ltly)
		r.LineTo(ltrx, ltry)
		r.LineTo(lbrx, lbry)
		r.LineTo(lblx, lbly)
		r.LineTo(cx, ly)
		r.Close()
		r.FillStroke()

		r.SetFontColor(c.FinalValueLabel.GetFontColor(DefaultTextColor))
		r.Text(ll, textX, textY)
	}
}

func (c Chart) drawTitle(r Renderer) error {
	if len(c.Title) > 0 && c.TitleStyle.Show {
		r.SetFontColor(c.Canvas.GetFontColor(DefaultTextColor))
		titleFontSize := c.Canvas.GetFontSize(DefaultTitleFontSize)
		r.SetFontSize(titleFontSize)
		textWidth := r.MeasureText(c.Title)
		titleX := (c.Width >> 1) - (textWidth >> 1)
		titleY := c.TitleStyle.Padding.GetTop(DefaultTitleTop) + int(titleFontSize)
		r.Text(c.Title, titleX, titleY)
	}
	return nil
}
