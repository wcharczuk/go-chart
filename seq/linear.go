package seq

// Range returns the array values of a linear seq with a given start, end and optional step.
func Range(start, end float64) []float64 {
	return Seq{NewLinear().WithStart(start).WithEnd(end).WithStep(1.0)}.Array()
}

// RangeWithStep returns the array values of a linear seq with a given start, end and optional step.
func RangeWithStep(start, end, step float64) []float64 {
	return Seq{NewLinear().WithStart(start).WithEnd(end).WithStep(step)}.Array()
}

// NewLinear returns a new linear generator.
func NewLinear() *Linear {
	return &Linear{step: 1.0}
}

// Linear is a stepwise generator.
type Linear struct {
	start float64
	end   float64
	step  float64
}

// Start returns the start value.
func (lg Linear) Start() float64 {
	return lg.start
}

// End returns the end value.
func (lg Linear) End() float64 {
	return lg.end
}

// Step returns the step value.
func (lg Linear) Step() float64 {
	return lg.step
}

// Len returns the number of elements in the seq.
func (lg Linear) Len() int {
	if lg.start < lg.end {
		return int((lg.end-lg.start)/lg.step) + 1
	}
	return int((lg.start-lg.end)/lg.step) + 1
}

// GetValue returns the value at a given index.
func (lg Linear) GetValue(index int) float64 {
	fi := float64(index)
	if lg.start < lg.end {
		return lg.start + (fi * lg.step)
	}
	return lg.start - (fi * lg.step)
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
