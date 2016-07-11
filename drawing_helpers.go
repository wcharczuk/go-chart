package chart

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

// MeasureAnnotation measures how big an annotation would be.
func MeasureAnnotation(r Renderer, canvasBox Box, xrange, yrange Range, s Style, lx, ly int, label string) Box {
	r.SetFont(s.GetFont())
	r.SetFontSize(s.GetFontSize(DefaultAnnotationFontSize))
	textWidth, textHeight := r.MeasureText(label)
	halfTextHeight := textHeight >> 1

	pt := s.Padding.GetTop(DefaultAnnotationPadding.Top)
	pl := s.Padding.GetLeft(DefaultAnnotationPadding.Left)
	pr := s.Padding.GetRight(DefaultAnnotationPadding.Right)
	pb := s.Padding.GetBottom(DefaultAnnotationPadding.Bottom)

	top := ly - (pt + halfTextHeight)
	right := lx + pl + pr + textWidth + DefaultAnnotationDeltaWidth
	bottom := ly + (pb + halfTextHeight)

	return Box{
		Top:    top,
		Left:   lx,
		Right:  right,
		Bottom: bottom,
		Width:  right - lx,
		Height: bottom - top,
	}
}

// DrawAnnotation draws an anotation with a renderer.
func DrawAnnotation(r Renderer, canvasBox Box, xrange, yrange Range, s Style, lx, ly int, label string) {
	r.SetFont(s.GetFont())
	r.SetFontSize(s.GetFontSize(DefaultAnnotationFontSize))
	textWidth, textHeight := r.MeasureText(label)
	halfTextHeight := textHeight >> 1

	pt := s.Padding.GetTop(DefaultAnnotationPadding.Top)
	pl := s.Padding.GetLeft(DefaultAnnotationPadding.Left)
	pr := s.Padding.GetRight(DefaultAnnotationPadding.Right)
	pb := s.Padding.GetBottom(DefaultAnnotationPadding.Bottom)

	textX := lx + pl + DefaultAnnotationDeltaWidth
	textY := ly + halfTextHeight

	ltx := lx + DefaultAnnotationDeltaWidth
	lty := ly - (pt + halfTextHeight)

	rtx := lx + pl + pr + textWidth + DefaultAnnotationDeltaWidth
	rty := ly - (pt + halfTextHeight)

	rbx := lx + pl + pr + textWidth + DefaultAnnotationDeltaWidth
	rby := ly + (pb + halfTextHeight)

	lbx := lx + DefaultAnnotationDeltaWidth
	lby := ly + (pb + halfTextHeight)

	//draw the shape...
	r.SetFillColor(s.GetFillColor(DefaultAnnotationFillColor))
	r.SetStrokeColor(s.GetStrokeColor())
	r.SetStrokeWidth(s.GetStrokeWidth())

	r.MoveTo(lx, ly)
	r.LineTo(ltx, lty)
	r.LineTo(rtx, rty)
	r.LineTo(rbx, rby)
	r.LineTo(lbx, lby)
	r.LineTo(lx, ly)
	r.Close()
	r.FillStroke()

	r.SetFontColor(s.GetFontColor(DefaultTextColor))
	r.Text(label, textX, textY)
}
