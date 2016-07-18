package chart

import (
	"math"

	"github.com/wcharczuk/go-chart/drawing"
)

// DrawLineSeries draws a line series with a renderer.
func DrawLineSeries(r Renderer, canvasBox Box, xrange, yrange Range, s Style, vs ValueProvider) {
	if vs.Len() == 0 {
		return
	}

	cb := canvasBox.Bottom
	cl := canvasBox.Left

	v0x, v0y := vs.GetValue(0)
	x0 := cl + xrange.Translate(v0x)
	y0 := cb - yrange.Translate(v0y)

	var vx, vy float64
	var x, y int

	fill := s.GetFillColor()
	if !fill.IsZero() {
		r.SetFillColor(fill)
		r.MoveTo(x0, y0)
		for i := 1; i < vs.Len(); i++ {
			vx, vy = vs.GetValue(i)
			x = cl + xrange.Translate(vx)
			y = cb - yrange.Translate(vy)
			r.LineTo(x, y)
		}
		r.LineTo(x, cb)
		r.LineTo(x0, cb)
		r.Close()
		r.Fill()
	}

	r.SetStrokeColor(s.GetStrokeColor())
	r.SetStrokeDashArray(s.GetStrokeDashArray())
	r.SetStrokeWidth(s.GetStrokeWidth(DefaultStrokeWidth))

	r.MoveTo(x0, y0)
	for i := 1; i < vs.Len(); i++ {
		vx, vy = vs.GetValue(i)
		x = cl + xrange.Translate(vx)
		y = cb - yrange.Translate(vy)
		r.LineTo(x, y)
	}
	r.Stroke()
}

// DrawBoundedSeries draws a series that implements BoundedValueProvider.
func DrawBoundedSeries(r Renderer, canvasBox Box, xrange, yrange Range, s Style, bbs BoundedValueProvider, drawOffsetIndexes ...int) {
	drawOffsetIndex := 0
	if len(drawOffsetIndexes) > 0 {
		drawOffsetIndex = drawOffsetIndexes[0]
	}

	r.SetStrokeColor(s.GetStrokeColor())
	r.SetStrokeDashArray(s.GetStrokeDashArray())
	r.SetStrokeWidth(s.GetStrokeWidth())
	r.SetFillColor(s.GetFillColor())

	cb := canvasBox.Bottom
	cl := canvasBox.Left

	v0x, v0y1, v0y2 := bbs.GetBoundedValue(0)
	x0 := cl + xrange.Translate(v0x)
	y0 := cb - yrange.Translate(v0y1)

	var vx, vy1, vy2 float64
	var x, y int

	xvalues := make([]float64, bbs.Len())
	xvalues[0] = v0x
	y2values := make([]float64, bbs.Len())
	y2values[0] = v0y2

	r.MoveTo(x0, y0)
	for i := 1; i < bbs.Len(); i++ {
		vx, vy1, vy2 = bbs.GetBoundedValue(i)

		xvalues[i] = vx
		y2values[i] = vy2

		x = cl + xrange.Translate(vx)
		y = cb - yrange.Translate(vy1)
		if i > drawOffsetIndex {
			r.LineTo(x, y)
		} else {
			r.MoveTo(x, y)
		}
	}
	y = cb - yrange.Translate(vy2)
	r.LineTo(x, y)
	for i := bbs.Len() - 1; i >= drawOffsetIndex; i-- {
		vx, vy2 = xvalues[i], y2values[i]
		x = cl + xrange.Translate(vx)
		y = cb - yrange.Translate(vy2)
		r.LineTo(x, y)
	}
	r.Close()
	r.FillStroke()
}

// DrawHistogramSeries draws a value provider as boxes from 0.
func DrawHistogramSeries(r Renderer, canvasBox Box, xrange, yrange Range, s Style, vs ValueProvider, barWidths ...int) {
	if vs.Len() == 0 {
		return
	}

	//calculate bar width?
	seriesLength := vs.Len()
	barWidth := int(math.Floor(float64(xrange.Domain) / float64(seriesLength)))
	if len(barWidths) > 0 {
		barWidth = barWidths[0]
	}

	cb := canvasBox.Bottom
	cl := canvasBox.Left

	//foreach datapoint, draw a box.
	for index := 0; index < seriesLength; index++ {
		vx, vy := vs.GetValue(index)
		y0 := yrange.Translate(0)
		x := cl + xrange.Translate(vx)
		y := yrange.Translate(vy)

		DrawBox(r, Box{
			Top:    cb - y0,
			Left:   x - (barWidth >> 1),
			Right:  x + (barWidth >> 1),
			Bottom: cb - y,
		}, s)
	}
}

// MeasureAnnotation measures how big an annotation would be.
func MeasureAnnotation(r Renderer, canvasBox Box, s Style, lx, ly int, label string) Box {
	r.SetFillColor(s.GetFillColor(DefaultAnnotationFillColor))
	r.SetStrokeColor(s.GetStrokeColor())
	r.SetStrokeWidth(s.GetStrokeWidth())
	r.SetFont(s.GetFont())
	r.SetFontColor(s.GetFontColor(DefaultTextColor))
	r.SetFontSize(s.GetFontSize(DefaultAnnotationFontSize))

	textBox := r.MeasureText(label)
	textWidth := textBox.Width()
	textHeight := textBox.Height()
	halfTextHeight := textHeight >> 1

	pt := s.Padding.GetTop(DefaultAnnotationPadding.Top)
	pl := s.Padding.GetLeft(DefaultAnnotationPadding.Left)
	pr := s.Padding.GetRight(DefaultAnnotationPadding.Right)
	pb := s.Padding.GetBottom(DefaultAnnotationPadding.Bottom)

	strokeWidth := s.GetStrokeWidth()

	top := ly - (pt + halfTextHeight)
	right := lx + pl + pr + textWidth + DefaultAnnotationDeltaWidth + int(strokeWidth)
	bottom := ly + (pb + halfTextHeight)

	return Box{
		Top:    top,
		Left:   lx,
		Right:  right,
		Bottom: bottom,
	}
}

// DrawAnnotation draws an anotation with a renderer.
func DrawAnnotation(r Renderer, canvasBox Box, style Style, lx, ly int, label string) {
	r.SetFillColor(style.GetFillColor())
	r.SetStrokeColor(style.GetStrokeColor())
	r.SetStrokeWidth(style.GetStrokeWidth())
	r.SetStrokeDashArray(style.GetStrokeDashArray())

	r.SetFont(style.GetFont())
	r.SetFontColor(style.GetFontColor(DefaultTextColor))
	r.SetFontSize(style.GetFontSize(DefaultAnnotationFontSize))

	textBox := r.MeasureText(label)
	textWidth := textBox.Width()
	halfTextHeight := textBox.Height() >> 1

	pt := style.Padding.GetTop(DefaultAnnotationPadding.Top)
	pl := style.Padding.GetLeft(DefaultAnnotationPadding.Left)
	pr := style.Padding.GetRight(DefaultAnnotationPadding.Right)
	pb := style.Padding.GetBottom(DefaultAnnotationPadding.Bottom)

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

	r.MoveTo(lx, ly)
	r.LineTo(ltx, lty)
	r.LineTo(rtx, rty)
	r.LineTo(rbx, rby)
	r.LineTo(lbx, lby)
	r.LineTo(lx, ly)
	r.Close()
	r.FillStroke()

	r.Text(label, textX, textY)
}

// DrawBox draws a box with a given style.
func DrawBox(r Renderer, b Box, s Style) {
	r.SetFillColor(s.GetFillColor())
	r.SetStrokeColor(s.GetStrokeColor())
	r.SetStrokeWidth(s.GetStrokeWidth(DefaultStrokeWidth))
	r.SetStrokeDashArray(s.GetStrokeDashArray())

	r.MoveTo(b.Left, b.Top)
	r.LineTo(b.Right, b.Top)
	r.LineTo(b.Right, b.Bottom)
	r.LineTo(b.Left, b.Bottom)
	r.LineTo(b.Left, b.Top)
	r.FillStroke()
}

// DrawText draws text with a given style.
func DrawText(r Renderer, text string, x, y int, s Style) {
	r.SetFontColor(s.GetFontColor(DefaultTextColor))
	r.SetStrokeColor(s.GetStrokeColor())
	r.SetStrokeWidth(s.GetStrokeWidth())
	r.SetFont(s.GetFont())
	r.SetFontSize(s.GetFontSize())

	r.Text(text, x, y)
}

// DrawTextCentered draws text with a given style centered.
func DrawTextCentered(r Renderer, text string, x, y int, s Style) {
	r.SetFontColor(s.GetFontColor(DefaultTextColor))
	r.SetStrokeColor(s.GetStrokeColor())
	r.SetStrokeWidth(s.GetStrokeWidth())
	r.SetFont(s.GetFont())
	r.SetFontSize(s.GetFontSize())

	tb := r.MeasureText(text)
	tx := x - (tb.Width() >> 1)
	ty := y - (tb.Height() >> 1)
	r.Text(text, tx, ty)
}

// CreateLegend returns a legend renderable function.
func CreateLegend(c *Chart, userDefaults ...Style) Renderable {
	return func(r Renderer, cb Box, chartDefaults Style) {
		legendDefaults := Style{
			FillColor:   drawing.ColorWhite,
			FontColor:   DefaultTextColor,
			FontSize:    8.0,
			StrokeColor: DefaultAxisColor,
			StrokeWidth: DefaultAxisLineWidth,
		}

		var legendStyle Style
		if len(userDefaults) > 0 {
			legendStyle = userDefaults[0].WithDefaultsFrom(chartDefaults.WithDefaultsFrom(legendDefaults))
		} else {
			legendStyle = chartDefaults.WithDefaultsFrom(legendDefaults)
		}

		// DEFAULTS
		legendPadding := Box{
			Top:    5,
			Left:   5,
			Right:  5,
			Bottom: 5,
		}
		lineTextGap := 5
		lineLengthMinimum := 25

		var labels []string
		var lines []Style
		for index, s := range c.Series {
			if s.GetStyle().IsZero() || s.GetStyle().Show {
				if _, isAnnotationSeries := s.(AnnotationSeries); !isAnnotationSeries {
					labels = append(labels, s.GetName())
					lines = append(lines, s.GetStyle().WithDefaultsFrom(c.styleDefaultsSeries(index)))
				}
			}
		}

		legend := Box{
			Top:  cb.Top,
			Left: cb.Left,
			// bottom and right will be sized by the legend content + relevant padding.
		}

		legendContent := Box{
			Top:    legend.Top + legendPadding.Top,
			Left:   legend.Left + legendPadding.Left,
			Right:  legend.Left + legendPadding.Left,
			Bottom: legend.Top + legendPadding.Top,
		}

		r.SetFont(legendStyle.GetFont())
		r.SetFontColor(legendStyle.GetFontColor())
		r.SetFontSize(legendStyle.GetFontSize())

		// measure
		labelCount := 0
		for x := 0; x < len(labels); x++ {
			if len(labels[x]) > 0 {
				tb := r.MeasureText(labels[x])
				if labelCount > 0 {
					legendContent.Bottom += DefaultMinimumTickVerticalSpacing
				}
				legendContent.Bottom += tb.Height()
				right := legendContent.Left + tb.Width() + lineTextGap + lineLengthMinimum
				legendContent.Right = MaxInt(legendContent.Right, right)
				labelCount++
			}
		}

		legend = legend.Grow(legendContent)
		legend.Right = legendContent.Right + legendPadding.Right
		legend.Bottom = legendContent.Bottom + legendPadding.Bottom

		DrawBox(r, legend, legendStyle)

		ycursor := legendContent.Top
		tx := legendContent.Left
		legendCount := 0
		for x := 0; x < len(labels); x++ {
			if len(labels[x]) > 0 {

				if legendCount > 0 {
					ycursor += DefaultMinimumTickVerticalSpacing
				}

				tb := r.MeasureText(labels[x])

				ty := ycursor + tb.Height()
				r.Text(labels[x], tx, ty)

				th2 := tb.Height() >> 1

				lx := tx + tb.Width() + lineTextGap
				ly := ty - th2
				lx2 := legendContent.Right - legendPadding.Right

				r.SetStrokeColor(lines[x].GetStrokeColor())
				r.SetStrokeWidth(lines[x].GetStrokeWidth())
				r.SetStrokeDashArray(lines[x].GetStrokeDashArray())

				r.MoveTo(lx, ly)
				r.LineTo(lx2, ly)
				r.Stroke()

				ycursor += tb.Height()
				legendCount++
			}
		}
	}
}
