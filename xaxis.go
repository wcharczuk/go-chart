package chart

import (
	"math"
	"sort"

	"github.com/wcharczuk/go-chart/drawing"
)

// XAxis represents the horizontal axis.
type XAxis struct {
	Name           string
	Style          Style
	ValueFormatter ValueFormatter
	Range          Range
	Ticks          []Tick
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
func (xa XAxis) GetTicks(r Renderer, ra Range, vf ValueFormatter) []Tick {
	if len(xa.Ticks) > 0 {
		return xa.Ticks
	}
	return xa.generateTicks(r, ra, vf)
}

func (xa XAxis) generateTicks(r Renderer, ra Range, vf ValueFormatter) []Tick {
	step := xa.getTickStep(r, ra, vf)
	return GenerateTicksWithStep(ra, step, vf)
}

func (xa XAxis) getTickStep(r Renderer, ra Range, vf ValueFormatter) float64 {
	tickCount := xa.getTickCount(r, ra, vf)
	step := ra.Delta() / float64(tickCount)
	return step
}

func (xa XAxis) getTickCount(r Renderer, ra Range, vf ValueFormatter) int {
	fontSize := xa.Style.GetFontSize(DefaultFontSize)
	r.SetFontSize(fontSize)

	// take a cut at determining the 'widest' value.
	l0 := vf(ra.Min)
	ln := vf(ra.Max)
	ll := l0
	if len(ln) > len(l0) {
		ll = ln
	}
	llw, _ := r.MeasureText(ll)
	textWidth := drawing.PointsToPixels(r.GetDPI(), float64(llw))
	width := textWidth + DefaultMinimumTickHorizontalSpacing
	count := int(math.Ceil(float64(ra.Domain) / float64(width)))
	return count
}

// Render renders the axis
func (xa XAxis) Render(r Renderer, canvasBox Box, ra Range, ticks []Tick) {
	tickFontSize := xa.Style.GetFontSize(DefaultFontSize)

	r.SetStrokeColor(xa.Style.GetStrokeColor(DefaultAxisColor))
	r.SetStrokeWidth(xa.Style.GetStrokeWidth(DefaultAxisLineWidth))

	r.MoveTo(canvasBox.Left, canvasBox.Bottom)
	r.LineTo(canvasBox.Right, canvasBox.Bottom)
	r.Stroke()

	r.SetFontColor(xa.Style.GetFontColor(DefaultAxisColor))
	r.SetFontSize(tickFontSize)

	sort.Sort(Ticks(ticks))

	var textHeight int
	for _, t := range ticks {
		_, th := r.MeasureText(t.Label)
		if th > textHeight {
			textHeight = th
		}
	}
	ty := canvasBox.Bottom + DefaultXAxisMargin + int(textHeight)
	for _, t := range ticks {
		v := t.Value
		x := ra.Translate(v)
		tx := canvasBox.Right - x
		r.Text(t.Label, tx, ty)
	}
}
