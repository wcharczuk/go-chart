package chart

import (
	"fmt"
	"time"
)

// Series is a entity data set.
type Series interface {
	GetName() string
	GetStyle() Style
	Len() int
	GetValue(index int) (float64, float64)
	GetLabel(index int) (string, string)
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

// GetValue gets a value at a given index.
func (ts TimeSeries) GetValue(index int) (x float64, y float64) {
	x = float64(ts.XValues[index].Unix())
	y = ts.YValues[index]
	return
}

// GetLabel gets a label for the values at a given index.
func (ts TimeSeries) GetLabel(index int) (xLabel string, yLabel string) {
	xLabel = ts.XValues[index].Format(DefaultDateFormat)
	yLabel = fmt.Sprintf("%0.2f", ts.YValues[index])
	return
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
func (cs ContinousSeries) GetValue(index int) (interface{}, float64) {
	return cs.XValues[index], cs.YValues[index]
}

// GetLabel gets a label for the values at a given index.
func (cs ContinousSeries) GetLabel(index int) (xLabel string, yLabel string) {
	xLabel = fmt.Sprintf("%0.2f", cs.XValues[index])
	yLabel = fmt.Sprintf("%0.2f", cs.YValues[index])
	return
}
