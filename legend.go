package chart

import "github.com/wcharczuk/go-chart/drawing"

// Legend returns a legend renderable function.
func Legend(c *Chart, userDefaults ...Style) Renderable {
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
			legendStyle = userDefaults[0].InheritFrom(chartDefaults.InheritFrom(legendDefaults))
		} else {
			legendStyle = chartDefaults.InheritFrom(legendDefaults)
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
					lines = append(lines, s.GetStyle().InheritFrom(c.styleDefaultsSeries(index)))
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
				legendContent.Right = Math.MaxInt(legendContent.Right, right)
				labelCount++
			}
		}

		legend = legend.Grow(legendContent)
		legend.Right = legendContent.Right + legendPadding.Right
		legend.Bottom = legendContent.Bottom + legendPadding.Bottom

		Draw.Box(r, legend, legendStyle)

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
