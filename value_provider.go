package chart

// ValueProvider is a type that produces values.
type ValueProvider interface {
	Len() int
	GetValue(index int) (float64, float64)
}

// BoundedValueProvider allows series to return a range.
type BoundedValueProvider interface {
	Len() int
	GetBoundedValue(index int) (x, y1, y2 float64)
}
