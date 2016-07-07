package chart

import "time"

// Series is a entity data set.
type Series interface {
	GetName() string
	GetStyle() Style
	Len() int
	GetValue(index int) Point

	GetXRange(domain int) Range
	GetYRange(domain int) Range
}

// TimeSeries is a line on a chart.
type TimeSeries struct {
	Name  string
	Style Style

	XValues []time.Time
	YValues []float64
}

// GetName returns the name of the time series.
func (ts TimeSeries) GetName() string {
	return ts.Name
}

// GetStyle returns the line style.
func (ts TimeSeries) GetStyle() Style {
	return ts.Style
}

// Len returns the number of elements in the series.
func (ts TimeSeries) Len() int {
	return len(ts.XValues)
}

// GetXRange returns the x range.
func (ts TimeSeries) GetXRange(domain int) Range {
	return NewRangeOfTime(domain, ts.XValues...)
}

// GetYRange returns the x range.
func (ts TimeSeries) GetYRange(domain int) Range {
	return NewRangeOfFloat64(domain, ts.YValues...)
}

// GetValue gets a value at a given index.
func (ts TimeSeries) GetValue(index int) Point {
	return Point{X: float64(ts.XValues[index].Unix()), Y: ts.YValues[index]}
}

// ContinousSeries represents a line on a chart.
type ContinousSeries struct {
	Name  string
	Style Style

	XValues []float64
	YValues []float64
}

// GetName returns the name of the time series.
func (cs ContinousSeries) GetName() string {
	return cs.Name
}

// GetStyle returns the line style.
func (cs ContinousSeries) GetStyle() Style {
	return cs.Style
}

// Len returns the number of elements in the series.
func (cs ContinousSeries) Len() int {
	return len(cs.XValues)
}

// GetValue gets a value at a given index.
func (cs ContinousSeries) GetValue(index int) Point {
	return Point{X: cs.XValues[index], Y: cs.YValues[index]}
}

// GetXRange returns the x range.
func (cs ContinousSeries) GetXRange(domain int) Range {
	return NewRangeOfFloat64(domain, cs.XValues...)
}

// GetYRange returns the x range.
func (cs ContinousSeries) GetYRange(domain int) Range {
	return NewRangeOfFloat64(domain, cs.YValues...)
}
