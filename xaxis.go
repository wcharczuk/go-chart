package chart

import (
	"math"
	"sort"
)

// XAxis represents the horizontal axis.
type XAxis struct {
	Name           string
	Style          Style
	ValueFormatter ValueFormatter
	Range          Range
	Ticks          []Tick

	GridLines      []GridLine
	GridMajorStyle Style
	GridMinorStyle Style
}

// GetName returns the name.
func (xa XAxis) GetName() string {
	return xa.Name
}

// GetStyle returns the style.
func (xa XAxis) GetStyle() Style {
	return xa.Style
}

// GetTicks returns the ticks for a series. It coalesces between user provided ticks and
// generated ticks.
func (xa XAxis) GetTicks(r Renderer, ra Range, defaults Style, vf ValueFormatter) []Tick {
	if len(xa.Ticks) > 0 {
		return xa.Ticks
	}
	if tp, isTickProvider := ra.(TicksProvider); isTickProvider {
		return tp.GetTicks(vf)
	}
	step := CalculateContinuousTickStep(r, ra, false, xa.Style.InheritFrom(defaults), vf)
	return GenerateContinuousTicksWithStep(ra, step, vf)
}

// GetGridLines returns the gridlines for the axis.
func (xa XAxis) GetGridLines(ticks []Tick) []GridLine {
	if len(xa.GridLines) > 0 {
		return xa.GridLines
	}
	return GenerateGridLines(ticks, xa.GridMajorStyle, xa.GridMinorStyle, true)
}

// Measure returns the bounds of the axis.
func (xa XAxis) Measure(r Renderer, canvasBox Box, ra Range, defaults Style, ticks []Tick) Box {
	xa.Style.InheritFrom(defaults).PersistToRenderer(r)
	sort.Sort(Ticks(ticks))

	var left, right, top, bottom = math.MaxInt32, 0, math.MaxInt32, 0
	for _, t := range ticks {
		v := t.Value
		lx := ra.Translate(v)
		tb := r.MeasureText(t.Label)

		tx := canvasBox.Left + lx
		ty := canvasBox.Bottom + DefaultXAxisMargin + tb.Height()

		top = MinInt(top, canvasBox.Bottom)
		left = MinInt(left, tx-(tb.Width()>>1))
		right = MaxInt(right, tx+(tb.Width()>>1))
		bottom = MaxInt(bottom, ty)
	}

	return Box{
		Top:    top,
		Left:   left,
		Right:  right,
		Bottom: bottom,
	}
}

// Render renders the axis
func (xa XAxis) Render(r Renderer, canvasBox Box, ra Range, defaults Style, ticks []Tick) {
	xa.Style.InheritFrom(defaults).PersistToRenderer(r)

	r.MoveTo(canvasBox.Left, canvasBox.Bottom)
	r.LineTo(canvasBox.Right, canvasBox.Bottom)
	r.Stroke()

	sort.Sort(Ticks(ticks))

	for _, t := range ticks {
		v := t.Value
		lx := ra.Translate(v)
		tb := r.MeasureText(t.Label)
		tx := canvasBox.Left + lx
		ty := canvasBox.Bottom + DefaultXAxisMargin + tb.Height()
		r.Text(t.Label, tx-tb.Width()>>1, ty)

		r.MoveTo(tx, canvasBox.Bottom)
		r.LineTo(tx, canvasBox.Bottom+DefaultVerticalTickHeight)
		r.Stroke()
	}

	if xa.GridMajorStyle.Show || xa.GridMinorStyle.Show {
		for _, gl := range xa.GetGridLines(ticks) {
			if (gl.IsMinor && xa.GridMinorStyle.Show) || (!gl.IsMinor && xa.GridMajorStyle.Show) {
				defaults := xa.GridMajorStyle
				if gl.IsMinor {
					defaults = xa.GridMinorStyle
				}
				gl.Render(r, canvasBox, ra, defaults)
			}
		}
	}
}
