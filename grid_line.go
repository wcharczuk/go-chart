package chart

// GridLine is a line on a graph canvas.
type GridLine struct {
	IsMinor    bool
	IsVertical bool
	Style      Style
	Value      float64
}

// Major returns if the gridline is a `major` line.
func (gl GridLine) Major() bool {
	return !gl.IsMinor
}

// Minor returns if the gridline is a `minor` line.
func (gl GridLine) Minor() bool {
	return gl.IsMinor
}

// Vertical returns if the line is vertical line or not.
func (gl GridLine) Vertical() bool {
	return gl.IsVertical
}

// Horizontal returns if the line is horizontal line or not.
func (gl GridLine) Horizontal() bool {
	return !gl.IsVertical
}

// Render renders the gridline
func (gl GridLine) Render(r Renderer, canvasBox Box, ra Range) {
	if gl.IsVertical {
		lineLeft := canvasBox.Left + ra.Translate(gl.Value)
		lineBottom := canvasBox.Bottom
		lineTop := canvasBox.Top

		r.SetStrokeColor(gl.Style.GetStrokeColor(DefaultAxisColor))
		r.SetStrokeWidth(gl.Style.GetStrokeWidth(DefaultAxisLineWidth))

		r.MoveTo(lineLeft, lineBottom)
		r.LineTo(lineLeft, lineTop)
		r.Stroke()
	} else {
		lineLeft := canvasBox.Left
		lineRight := canvasBox.Right
		lineHeight := canvasBox.Bottom - ra.Translate(gl.Value)

		r.SetStrokeColor(gl.Style.GetStrokeColor(DefaultAxisColor))
		r.SetStrokeWidth(gl.Style.GetStrokeWidth(DefaultAxisLineWidth))

		r.MoveTo(lineLeft, lineHeight)
		r.LineTo(lineRight, lineHeight)
		r.Stroke()
	}
}
