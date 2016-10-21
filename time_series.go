package chart

import "time"

// TimeSeries is a line on a chart.
type TimeSeries struct {
	Name  string
	Style Style

	YAxis YAxisType

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
func (ts TimeSeries) GetValue(index int) (x, y float64) {
	x = Time.ToFloat64(ts.XValues[index])
	y = ts.YValues[index]
	return
}

// GetLastValue gets the last value.
func (ts TimeSeries) GetLastValue() (x, y float64) {
	x = Time.ToFloat64(ts.XValues[len(ts.XValues)-1])
	y = ts.YValues[len(ts.YValues)-1]
	return
}

// GetValueFormatters returns value formatter defaults for the series.
func (ts TimeSeries) GetValueFormatters() (x, y ValueFormatter) {
	x = TimeValueFormatter
	y = FloatValueFormatter
	return
}

// GetYAxis returns which YAxis the series draws on.
func (ts TimeSeries) GetYAxis() YAxisType {
	return ts.YAxis
}

// Render renders the series.
func (ts TimeSeries) Render(r Renderer, canvasBox Box, xrange, yrange Range, defaults Style) {
	style := ts.Style.InheritFrom(defaults)
	Draw.LineSeries(r, canvasBox, xrange, yrange, style, ts)
}
