package chart

// YAxis is a veritcal rule of the range.
// There can be (2) y-axes; a primary and secondary.
type YAxis struct {
	axis
}

// Render renders the axis.
func (ya YAxis) Render(r Renderer, canvasBox Box, ra Range, axisType YAxisType) {
	var tx int
	if axisType == YAxisPrimary {
		tx = canvasBox.Right + DefaultYAxisMargin
	} else if axisType == YAxisSecondary {
		tx = canvasBox.Left - DefaultYAxisMargin
	}

	r.SetFontColor(ya.Style.GetFontColor(DefaultAxisColor))
	r.SetFontSize(ya.Style.GetFontSize(DefaultFontSize))

	ticks := ya.getTicks(ra)
	for _, t := range ticks {
		v := t.Value
		y := ra.Translate(v)
		ty := int(y)
		r.Text(t.Label, tx, ty)
	}
}
