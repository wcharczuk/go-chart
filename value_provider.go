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

// LastValueProvider is a special type of value provider that can return it's (potentially computed) last value.
type LastValueProvider interface {
	GetLastValue() (x, y float64)
}

// BoundedLastValueProvider is a special type of value provider that can return it's (potentially computed) bounded last value.
type BoundedLastValueProvider interface {
	GetBoundedLastValue() (x, y1, y2 float64)
}

// FullValueProvider is an interface that combines `ValueProvider` and `LastValueProvider`
type FullValueProvider interface {
	ValueProvider
	LastValueProvider
}

// FullBoundedValueProvider is an interface that combines `BoundedValueProvider` and `BoundedLastValueProvider`
type FullBoundedValueProvider interface {
	BoundedValueProvider
	BoundedLastValueProvider
}
