package chart

// Value is a chart value.
type Value struct {
	Style Style
	Label string
	Value float64
}

// Values is an array of Value.
type Values []Value

// Values returns the values.
func (vs Values) Values() []float64 {
	values := make([]float64, len(vs))
	for index, v := range vs {
		values[index] = v.Value
	}
	return values
}

// ValuesNormalized returns normalized values.
func (vs Values) ValuesNormalized() []float64 {
	return Math.Normalize(vs.Values()...)
}

// Normalize returns the values normalized.
func (vs Values) Normalize() []Value {
	output := make([]Value, len(vs))
	total := Math.Sum(vs.Values()...)
	for index, v := range vs {
		output[index] = Value{
			Style: v.Style,
			Label: v.Label,
			Value: Math.RoundDown(v.Value/total, 0.0001),
		}
	}
	return output
}

// Value2 is a two axis value.
type Value2 struct {
	Style          Style
	Label          string
	XValue, YValue float64
}
