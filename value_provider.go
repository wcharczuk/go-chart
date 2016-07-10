package chart

// ValueProvider is a type that produces values.
type ValueProvider interface {
	Len() int
	GetValue(index int) (float64, float64)
}
