package chart

import "math"

// DrawLineSeries draws a line series with a renderer.
func DrawLineSeries(r Renderer, canvasBox Box, xrange, yrange Range, s Style, vs ValueProvider) {
	if vs.Len() == 0 {
		return
	}

	ct := canvasBox.Top
	cb := canvasBox.Bottom
	cr := canvasBox.Right

	v0x, v0y := vs.GetValue(0)
	x0 := cr - xrange.Translate(v0x)
	y0 := yrange.Translate(v0y) + ct

	var vx, vy float64
	var x, y int

	fill := s.GetFillColor()
	if !fill.IsZero() {
		r.SetFillColor(fill)
		r.MoveTo(x0, y0)
		for i := 1; i < vs.Len(); i++ {
			vx, vy = vs.GetValue(i)
			x = cr - xrange.Translate(vx)
			y = yrange.Translate(vy) + ct
			r.LineTo(x, y)
		}
		r.LineTo(x, cb)
		r.LineTo(x0, cb)
		r.Close()
		r.Fill()
	}

	stroke := s.GetStrokeColor()
	r.SetStrokeColor(stroke)
	r.SetStrokeWidth(s.GetStrokeWidth(DefaultStrokeWidth))

	r.MoveTo(x0, y0)
	for i := 1; i < vs.Len(); i++ {
		vx, vy = vs.GetValue(i)
		x = cr - xrange.Translate(vx)
		y = yrange.Translate(vy) + ct
		r.LineTo(x, y)
	}
	r.Stroke()
}

// DrawAnnotation draws an anotation with a renderer.
func DrawAnnotation(r Renderer, canvasBox Box, xrange, yrange Range, s Style, lx, ly int, label string) {
	r.SetFontSize(s.GetFontSize(DefaultAnnotationFontSize))
	textWidth, _ := r.MeasureText(label)
	textHeight := int(math.Floor(DefaultAnnotationFontSize))
	halfTextHeight := textHeight >> 1

	pt := s.Padding.GetTop(DefaultAnnotationPadding.Top)
	pl := s.Padding.GetLeft(DefaultAnnotationPadding.Left)
	pr := s.Padding.GetRight(DefaultAnnotationPadding.Right)
	pb := s.Padding.GetBottom(DefaultAnnotationPadding.Bottom)

	textX := lx + pl + DefaultAnnotationDeltaWidth
	textY := ly + halfTextHeight

	ltlx := lx + pl + DefaultAnnotationDeltaWidth
	ltly := ly - (pt + halfTextHeight)

	ltrx := lx + pl + pr + textWidth
	ltry := ly - (pt + halfTextHeight)

	lbrx := lx + pl + pr + textWidth
	lbry := ly + (pb + halfTextHeight)

	lblx := lx + DefaultAnnotationDeltaWidth
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
	r.LineTo(lx, ly)
	r.Close()
	r.FillStroke()

	r.SetFontColor(s.GetFontColor(DefaultTextColor))
	r.Text(label, textX, textY)
}
