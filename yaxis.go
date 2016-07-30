package chart

import (
	"math"
	"sort"
)

// YAxis is a veritcal rule of the range.
// There can be (2) y-axes; a primary and secondary.
type YAxis struct {
	Name  string
	Style Style

	Zero GridLine

	AxisType yAxisType

	ValueFormatter ValueFormatter
	Range          Range

	Ticks     []Tick
	GridLines []GridLine

	GridMajorStyle Style
	GridMinorStyle Style
}

// GetName returns the name.
func (ya YAxis) GetName() string {
	return ya.Name
}

// GetStyle returns the style.
func (ya YAxis) GetStyle() Style {
	return ya.Style
}

// GetTicks returns the ticks for a series.
// The coalesce priority is:
// 	- User Supplied Ticks (i.e. Ticks array on the axis itself).
// 	- Range ticks (i.e. if the range provides ticks).
//	- Generating continuous ticks based on minimum spacing and canvas width.
func (ya YAxis) GetTicks(r Renderer, ra Range, defaults Style, vf ValueFormatter) []Tick {
	if len(ya.Ticks) > 0 {
		return ya.Ticks
	}
	if tp, isTickProvider := ra.(TicksProvider); isTickProvider {
		return tp.GetTicks(vf)
	}
	step := CalculateContinuousTickStep(r, ra, true, ya.Style.InheritFrom(defaults), vf)
	return GenerateContinuousTicksWithStep(ra, step, vf, true)
}

// GetGridLines returns the gridlines for the axis.
func (ya YAxis) GetGridLines(ticks []Tick) []GridLine {
	if len(ya.GridLines) > 0 {
		return ya.GridLines
	}
	return GenerateGridLines(ticks, ya.GridMajorStyle, ya.GridMinorStyle, false)
}

// Measure returns the bounds of the axis.
func (ya YAxis) Measure(r Renderer, canvasBox Box, ra Range, defaults Style, ticks []Tick) Box {
	ya.Style.InheritFrom(defaults).WriteToRenderer(r)

	sort.Sort(Ticks(ticks))

	var tx int
	if ya.AxisType == YAxisPrimary {
		tx = canvasBox.Right + DefaultYAxisMargin
	} else if ya.AxisType == YAxisSecondary {
		tx = canvasBox.Left - DefaultYAxisMargin
	}

	var minx, maxx, miny, maxy = math.MaxInt32, 0, math.MaxInt32, 0
	for _, t := range ticks {
		v := t.Value
		ly := canvasBox.Bottom - ra.Translate(v)

		tb := r.MeasureText(t.Label)
		finalTextX := tx
		if ya.AxisType == YAxisSecondary {
			finalTextX = tx - tb.Width()
		}

		if ya.AxisType == YAxisPrimary {
			minx = canvasBox.Right
			maxx = Math.MaxInt(maxx, tx+tb.Width())
		} else if ya.AxisType == YAxisSecondary {
			minx = Math.MinInt(minx, finalTextX)
			maxx = Math.MaxInt(maxx, tx)
		}
		miny = Math.MinInt(miny, ly-tb.Height()>>1)
		maxy = Math.MaxInt(maxy, ly+tb.Height()>>1)
	}

	return Box{
		Top:    miny,
		Left:   minx,
		Right:  maxx,
		Bottom: maxy,
	}
}

// Render renders the axis.
func (ya YAxis) Render(r Renderer, canvasBox Box, ra Range, defaults Style, ticks []Tick) {
	ya.Style.InheritFrom(defaults).WriteToRenderer(r)

	sort.Sort(Ticks(ticks))

	sw := ya.Style.GetStrokeWidth(defaults.StrokeWidth)

	var lx int
	var tx int
	if ya.AxisType == YAxisPrimary {
		lx = canvasBox.Right + int(sw)
		tx = lx + DefaultYAxisMargin
	} else if ya.AxisType == YAxisSecondary {
		lx = canvasBox.Left - int(sw)
		tx = lx - DefaultYAxisMargin
	}

	r.MoveTo(lx, canvasBox.Bottom)
	r.LineTo(lx, canvasBox.Top)
	r.Stroke()

	for _, t := range ticks {
		v := t.Value
		ly := canvasBox.Bottom - ra.Translate(v)

		tb := r.MeasureText(t.Label)

		finalTextX := tx
		finalTextY := ly + tb.Height()>>1
		if ya.AxisType == YAxisSecondary {
			finalTextX = tx - tb.Width()
		}

		r.Text(t.Label, finalTextX, finalTextY)

		r.MoveTo(lx, ly)
		if ya.AxisType == YAxisPrimary {
			r.LineTo(lx+DefaultHorizontalTickWidth, ly)
		} else if ya.AxisType == YAxisSecondary {
			r.LineTo(lx-DefaultHorizontalTickWidth, ly)
		}
		r.Stroke()
	}

	if ya.Zero.Style.Show {
		ya.Zero.Render(r, canvasBox, ra, Style{})
	}

	if ya.GridMajorStyle.Show || ya.GridMinorStyle.Show {
		for _, gl := range ya.GetGridLines(ticks) {
			if (gl.IsMinor && ya.GridMinorStyle.Show) || (!gl.IsMinor && ya.GridMajorStyle.Show) {
				defaults := ya.GridMajorStyle
				if gl.IsMinor {
					defaults = ya.GridMinorStyle
				}
				gl.Render(r, canvasBox, ra, defaults)
			}
		}
	}
}
