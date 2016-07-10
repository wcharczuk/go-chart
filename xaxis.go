package chart

import "github.com/wcharczuk/go-chart/drawing"

// XAxis represents the horizontal axis.
type XAxis struct {
	axis
}

// Render renders the axis
func (xa XAxis) Render(r Renderer, canvasBox Box, ra Range) {
	tickFontSize := xa.Style.GetFontSize(DefaultFontSize)
	tickHeight := drawing.PointsToPixels(r.GetDPI(), tickFontSize)
	ty := canvasBox.Bottom + DefaultXAxisMargin + int(tickHeight)

	r.SetFontColor(xa.Style.GetFontColor(DefaultAxisColor))
	r.SetFontSize(tickFontSize)

	ticks := xa.getTicks(ra)

	for _, t := range ticks {
		v := t.Value
		x := ra.Translate(v)
		tx := canvasBox.Left + int(x)
		r.Text(t.Label, tx, ty)
	}
}
