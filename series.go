package chart

// Series is an alias to Renderable.
type Series interface {
	GetYAxis() YAxisType
	Renderable
}
