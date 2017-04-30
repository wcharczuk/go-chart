package sequence

// NewLinear returns a new linear generator.
func NewLinear() *Linear {
	return &Linear{}
}

// Linear is a stepwise generator.
type Linear struct {
	offset float64
	limit  float64
	step   float64
}

// Len returns the number of elements in the sequence.
func (lg Linear) Len() int {
	return int((lg.limit - lg.offset) / lg.step)
}

// GetValue returns the value at a given index.
func (lg Linear) GetValue(index int) float64 {
	return lg.offset + (float64(index) * lg.step)
}

// WithOffset sets the offset and returns the linear generator.
func (lg *Linear) WithOffset(offset float64) *Linear {
	lg.offset = offset
	return lg
}

// WithLimit sets the step and returns the linear generator.
func (lg *Linear) WithLimit(limit float64) *Linear {
	lg.limit = limit
	return lg
}

// WithStep sets the step and returns the linear generator.
func (lg *Linear) WithStep(step float64) *Linear {
	lg.step = step
	return lg
}
