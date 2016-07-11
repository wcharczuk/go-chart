package chart

import (
	"math"
	"sort"
)

// YAxis is a veritcal rule of the range.
// There can be (2) y-axes; a primary and secondary.
type YAxis struct {
	Name           string
	Style          Style
	ValueFormatter ValueFormatter
	Range          Range
	Ticks          []Tick
}

// GetName returns the name.
func (ya YAxis) GetName() string {
	return ya.Name
}

// GetStyle returns the style.
func (ya YAxis) GetStyle() Style {
	return ya.Style
}

// GetTicks returns the ticks for a series. It coalesces between user provided ticks and
// generated ticks.
func (ya YAxis) GetTicks(r Renderer, ra Range, vf ValueFormatter) []Tick {
	if len(ya.Ticks) > 0 {
		return ya.Ticks
	}
	return ya.generateTicks(r, ra, vf)
}

func (ya YAxis) generateTicks(r Renderer, ra Range, vf ValueFormatter) []Tick {
	step := ya.getTickStep(r, ra, vf)
	return ya.generateTicksWithStep(ra, step, vf)
}

func (ya YAxis) getTickCount(r Renderer, ra Range, vf ValueFormatter) int {
	textHeight := int(ya.Style.GetFontSize(DefaultFontSize))
	height := textHeight + DefaultMinimumTickVerticalSpacing
	count := int(math.Ceil(float64(ra.Domain) / float64(height)))
	return count
}

func (ya YAxis) getTickStep(r Renderer, ra Range, vf ValueFormatter) float64 {
	tickCount := ya.getTickCount(r, ra, vf)
	return ra.Delta() / float64(tickCount)
}

func (ya YAxis) generateTicksWithStep(ra Range, step float64, vf ValueFormatter) []Tick {
	var ticks []Tick
	for cursor := ra.Min; cursor < ra.Max; cursor += step {
		ticks = append(ticks, Tick{
			Value: cursor,
			Label: vf(cursor),
		})
		if len(ticks) == 20 {
			return ticks
		}
	}
	return ticks
}

// Render renders the axis.
func (ya YAxis) Render(r Renderer, canvasBox Box, ra Range, axisType YAxisType, ticks []Tick) {
	r.SetStrokeColor(ya.Style.GetStrokeColor(DefaultAxisColor))
	r.SetStrokeWidth(ya.Style.GetStrokeWidth(DefaultAxisLineWidth))
	r.SetFontColor(ya.Style.GetFontColor(DefaultAxisColor))
	fontSize := ya.Style.GetFontSize(DefaultFontSize)
	r.SetFontSize(fontSize)

	sort.Sort(Ticks(ticks))

	var lx int
	var tx int
	if axisType == YAxisPrimary {
		lx = canvasBox.Right
		tx = canvasBox.Right + DefaultYAxisMargin

		r.MoveTo(lx, canvasBox.Bottom)
		r.LineTo(lx, canvasBox.Top)
		r.Stroke()

		for _, t := range ticks {
			v := t.Value
			ly := ra.Translate(v) + canvasBox.Top

			th := int(fontSize) >> 1
			ty := ly + th

			r.Text(t.Label, tx, ty)

			r.MoveTo(lx, ly)
			r.LineTo(lx+DefaultVerticalTickWidth, ly)
			r.Stroke()
		}
	} else if axisType == YAxisSecondary {
		lx = canvasBox.Left

		r.MoveTo(lx, canvasBox.Bottom)
		r.LineTo(lx, canvasBox.Top)
		r.Stroke()

		for _, t := range ticks {
			v := t.Value
			ly := ra.Translate(v) + canvasBox.Top

			ptw, _ := r.MeasureText(t.Label)

			tw := ptw
			th := int(fontSize)

			tx = lx - (int(tw) + (DefaultYAxisMargin >> 1))
			ty := ly + th>>1

			r.Text(t.Label, tx, ty)

			r.MoveTo(lx, ly)
			r.LineTo(lx-DefaultVerticalTickWidth, ly)
			r.Stroke()
		}
	}

}
