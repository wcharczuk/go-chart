package chart

// Series is an alias to Renderable.
type Series interface {
	GetYAxis() YAxisType
	GetStyle() Style
	Render(r Renderer, canvasBox Box, xrange, yrange Range, s Style)
}
