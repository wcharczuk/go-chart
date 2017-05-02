package sequence

// Values returns the array values of a linear sequence with a given start, end and optional step.
func Values(start, end float64) []float64 {
	return Seq{NewLinear().WithStart(start).WithEnd(end).WithStep(1.0)}.Array()
}

// ValuesWithStep returns the array values of a linear sequence with a given start, end and optional step.
func ValuesWithStep(start, end, step float64) []float64 {
	return Seq{NewLinear().WithStart(start).WithEnd(end).WithStep(step)}.Array()
}

// NewLinear returns a new linear generator.
func NewLinear() *Linear {
	return &Linear{}
}

// Linear is a stepwise generator.
type Linear struct {
	start float64
	end   float64
	step  float64
}

// Len returns the number of elements in the sequence.
func (lg Linear) Len() int {
	if lg.start < lg.end {
		return int((lg.end-lg.start)/lg.step) + 1
	}
	return int((lg.start-lg.end)/lg.step) + 1
}

// GetValue returns the value at a given index.
func (lg Linear) GetValue(index int) float64 {
	if lg.start < lg.end {
		return lg.start + (float64(index) * lg.step)
	}
	return lg.end + (float64(index) * lg.step)
}

// WithStart sets the start and returns the linear generator.
func (lg *Linear) WithStart(start float64) *Linear {
	lg.start = start
	return lg
}

// WithEnd sets the end and returns the linear generator.
func (lg *Linear) WithEnd(end float64) *Linear {
	lg.end = end
	return lg
}

// WithStep sets the step and returns the linear generator.
func (lg *Linear) WithStep(step float64) *Linear {
	lg.step = step
	return lg
}
