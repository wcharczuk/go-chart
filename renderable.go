package chart

// Renderable is a type that can be rendered onto a chart.
type Renderable interface {
	Render(r Renderer, canvasBox Box, xrange, yrange Range, defaults Style)
}
