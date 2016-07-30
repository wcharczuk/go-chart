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

	TickPosition tickPosition

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

// GetTickPosition returns the tick position option for the axis.
func (xa XAxis) GetTickPosition(defaults ...tickPosition) tickPosition {
	if xa.TickPosition == TickPositionUnset {
		if len(defaults) > 0 {
			return defaults[0]
		}
		return TickPositionUnderTick
	}
	return xa.TickPosition
}

// GetTicks returns the ticks for a series.
// The coalesce priority is:
// 	- User Supplied Ticks (i.e. Ticks array on the axis itself).
// 	- Range ticks (i.e. if the range provides ticks).
//	- Generating continuous ticks based on minimum spacing and canvas width.
func (xa XAxis) GetTicks(r Renderer, ra Range, defaults Style, vf ValueFormatter) []Tick {
	if len(xa.Ticks) > 0 {
		return xa.Ticks
	}
	if tp, isTickProvider := ra.(TicksProvider); isTickProvider {
		return tp.GetTicks(vf)
	}
	tickStyle := xa.Style.InheritFrom(defaults)
	return GenerateContinuousTicks(r, ra, false, tickStyle, vf)
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
	xa.Style.InheritFrom(defaults).WriteToRenderer(r)
	sort.Sort(Ticks(ticks))

	var left, right, top, bottom = math.MaxInt32, 0, math.MaxInt32, 0
	for _, t := range ticks {
		v := t.Value
		lx := ra.Translate(v)
		tb := r.MeasureText(t.Label)

		tx := canvasBox.Left + lx
		ty := canvasBox.Bottom + DefaultXAxisMargin + tb.Height()

		top = Math.MinInt(top, canvasBox.Bottom)
		left = Math.MinInt(left, tx-(tb.Width()>>1))
		right = Math.MaxInt(right, tx+(tb.Width()>>1))
		bottom = Math.MaxInt(bottom, ty)
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
	tickStyle := xa.Style.InheritFrom(defaults)

	tickStyle.GetStrokeOptions().WriteToRenderer(r)
	r.MoveTo(canvasBox.Left, canvasBox.Bottom)
	r.LineTo(canvasBox.Right, canvasBox.Bottom)
	r.Stroke()

	sort.Sort(Ticks(ticks))

	tp := xa.GetTickPosition()

	var tx, ty int
	for index, t := range ticks {
		v := t.Value
		lx := ra.Translate(v)

		tx = canvasBox.Left + lx

		tickStyle.GetStrokeOptions().WriteToRenderer(r)
		r.MoveTo(tx, canvasBox.Bottom)
		r.LineTo(tx, canvasBox.Bottom+DefaultVerticalTickHeight)
		r.Stroke()

		tickStyle.GetTextOptions().WriteToRenderer(r)
		tb := r.MeasureText(t.Label)

		switch tp {
		case TickPositionUnderTick, TickPositionUnset:
			ty = canvasBox.Bottom + DefaultXAxisMargin + tb.Height()
			r.Text(t.Label, tx-tb.Width()>>1, ty)
			break
		case TickPositionBetweenTicks:
			if index > 0 {
				llx := ra.Translate(ticks[index-1].Value)
				ltx := canvasBox.Left + llx
				Draw.TextWithin(r, t.Label, Box{
					Left:   ltx,
					Right:  tx,
					Top:    canvasBox.Bottom + DefaultXAxisMargin,
					Bottom: canvasBox.Bottom + DefaultXAxisMargin + tb.Height(),
				}, tickStyle.InheritFrom(Style{TextHorizontalAlign: TextHorizontalAlignCenter}))
			}
			break
		}

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
