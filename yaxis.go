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

	AxisType YAxisType

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
	var ticks []Tick
	if len(ya.Ticks) > 0 {
		ticks = ya.Ticks
	} else {
		ticks = ya.generateTicks(r, ra, vf)
	}

	return ticks
}

func (ya YAxis) generateTicks(r Renderer, ra Range, vf ValueFormatter) []Tick {
	step := ya.getTickStep(r, ra, vf)
	ticks := GenerateTicksWithStep(ra, step, vf)
	return ticks
}

func (ya YAxis) getTickStep(r Renderer, ra Range, vf ValueFormatter) float64 {
	tickCount := ya.getTickCount(r, ra, vf)
	step := ra.Delta() / float64(tickCount)
	return step
}

func (ya YAxis) getTickCount(r Renderer, ra Range, vf ValueFormatter) int {
	//given the domain, figure out how many ticks we can draw ...
	label := vf(ra.Min)
	tb := r.MeasureText(label)
	return int(math.Ceil(float64(ra.Domain) / float64(tb.Height()+DefaultMinimumTickVerticalSpacing)))
}

// Measure returns the bounds of the axis.
func (ya YAxis) Measure(r Renderer, canvasBox Box, ra Range, ticks []Tick) Box {
	defaultFont, _ := GetDefaultFont()
	r.SetFont(ya.Style.GetFont(defaultFont))
	r.SetFontSize(ya.Style.GetFontSize(DefaultFontSize))

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
			maxx = MaxInt(maxx, tx+tb.Width())
		} else if ya.AxisType == YAxisSecondary {
			minx = MinInt(minx, finalTextX)
			maxx = MaxInt(maxx, tx)
		}
		miny = MinInt(miny, ly-tb.Height()>>1)
		maxy = MaxInt(maxy, ly+tb.Height()>>1)
	}

	return Box{
		Top:    miny,
		Left:   minx,
		Right:  maxx,
		Bottom: maxy,
	}
}

// Render renders the axis.
func (ya YAxis) Render(r Renderer, canvasBox Box, ra Range, ticks []Tick) {
	r.SetStrokeColor(ya.Style.GetStrokeColor(DefaultAxisColor))
	r.SetStrokeWidth(ya.Style.GetStrokeWidth(DefaultAxisLineWidth))

	fontColor := ya.Style.GetFontColor(DefaultAxisColor)
	r.SetFontColor(fontColor)
	fontSize := ya.Style.GetFontSize(DefaultFontSize)
	r.SetFontSize(fontSize)

	sort.Sort(Ticks(ticks))

	var lx int
	var tx int
	if ya.AxisType == YAxisPrimary {
		lx = canvasBox.Right
		tx = lx + DefaultYAxisMargin
	} else if ya.AxisType == YAxisSecondary {
		lx = canvasBox.Left
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
		ya.Zero.Render(r, canvasBox, ra)
	}
}
