package chart

// Series is an alias to Renderable.
type Series interface {
	GetYAxis() YAxisType
	Render(r Renderer, canvasBox Box, xrange, yrange Range, s Style)
}
