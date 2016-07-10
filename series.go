package chart

import "math"

// Series is a entity data set. It constitutes an item to draw on the chart.
// The series interface is the bare minimum you need to implement to draw something on a chart.
type Series interface {
	GetName() string
	GetStyle() Style
	Render(c *Chart, r Renderer, canvasBox Box, xrange, yrange Range) error
}

// ValueProvider is a series that is a set of values.
type ValueProvider interface {
	Len() int
	GetValue(index int) (float64, float64)
}

// FormatterProvider is a series that has custom formatters.
type FormatterProvider interface {
	GetXFormatter() Formatter
	GetYFormatter() Formatter
}

// DrawLineSeries draws a line series with a renderer.
func DrawLineSeries(c *Chart, r Renderer, canvasBox Box, xrange, yrange Range, vs ValueProvider) error {
	if vs.Len() == 0 {
		return
	}

	cx := canvasBox.Left
	cy := canvasBox.Top
	cb := canvasBox.Bottom
	cw := canvasBox.Width

	v0x, v0y := vs.GetValue(0)
	x0 := cw - xrange.Translate(v0x)
	y0 := yrange.Translate(v0y)

	var vx, vy float64
	var x, y int

	fill := s.GetStyle().GetFillColor()
	if !fill.IsZero() {
		r.SetFillColor(fill)
		r.MoveTo(x0+cx, y0+cy)
		for i := 1; i < vs.Len(); i++ {
			vx, vy = vs.GetValue(i)
			x = cw - xrange.Translate(vx)
			y = yrange.Translate(vy)
			r.LineTo(x+cx, y+cy)
		}
		r.LineTo(x+cx, cb)
		r.LineTo(x0+cx, cb)
		r.Close()
		r.Fill()
	}

	stroke := s.GetStyle().GetStrokeColor(GetDefaultSeriesStrokeColor(index))
	r.SetStrokeColor(stroke)
	r.SetStrokeWidth(s.GetStyle().GetStrokeWidth(DefaultStrokeWidth))

	r.MoveTo(x0+cx, y0+cy)
	for i := 1; i < vs.Len(); i++ {
		vx, vy = vs.GetValue(i)
		x = cw - xrange.Translate(vx)
		y = yrange.Translate(vy)
		r.LineTo(x+cx, y+cy)
	}
	r.Stroke()
}

// DrawAnnotation draws an anotation with a renderer.
func DrawAnnotation(c *Chart, r Renderer, canvasBox Box, xrange, yrange, s Style, lx, ly int, lv string) {
	py := canvasBox.Top

	r.SetFontSize(s.GetFontSize(DefaultFinalLabelFontSize))
	textWidth, _ := r.MeasureText(ll)
	textHeight := int(math.Floor(DefaultFinalLabelFontSize))
	halfTextHeight := textHeight >> 1

	pt := s.Padding.GetTop(DefaultFinalLabelPadding.Top)
	pl := s.Padding.GetLeft(DefaultFinalLabelPadding.Left)
	pr := s.Padding.GetRight(DefaultFinalLabelPadding.Right)
	pb := s.Padding.GetBottom(DefaultFinalLabelPadding.Bottom)

	textX := lx + pl + DefaultFinalLabelDeltaWidth
	textY := ly + halfTextHeight

	ltlx := lx + pl + DefaultFinalLabelDeltaWidth
	ltly := ly - (pt + halfTextHeight)

	ltrx := lx + pl + pr + textWidth
	ltry := ly - (pt + halfTextHeight)

	lbrx := lx + pl + pr + textWidth
	lbry := ly + (pb + halfTextHeight)

	lblx := lx + DefaultFinalLabelDeltaWidth
	lbly := ly + (pb + halfTextHeight)

	//draw the shape...
	r.SetFillColor(s.GetFillColor(DefaultAnnotationFillColor))
	r.SetStrokeColor(s.GetStrokeColor())
	r.SetStrokeWidth(s.GetStrokeWidth())
	r.MoveTo(lx, ly)
	r.LineTo(ltlx, ltly)
	r.LineTo(ltrx, ltry)
	r.LineTo(lbrx, lbry)
	r.LineTo(lblx, lbly)
	r.LineTo(cx, ly)
	r.Close()
	r.FillStroke()

	r.SetFontColor(s.GetFontColor(DefaultTextColor))
	r.Text(ll, textX, textY)
}
