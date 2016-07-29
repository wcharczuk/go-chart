package chart

// ContinuousSeries represents a line on a chart.
type ContinuousSeries struct {
	Name  string
	Style Style

	YAxis YAxisType

	XValues []float64
	YValues []float64
}

// GetName returns the name of the time series.
func (cs ContinuousSeries) GetName() string {
	return cs.Name
}

// GetStyle returns the line style.
func (cs ContinuousSeries) GetStyle() Style {
	return cs.Style
}

// Len returns the number of elements in the series.
func (cs ContinuousSeries) Len() int {
	return len(cs.XValues)
}

// GetValue gets a value at a given index.
func (cs ContinuousSeries) GetValue(index int) (float64, float64) {
	return cs.XValues[index], cs.YValues[index]
}

// GetLastValue gets the last value.
func (cs ContinuousSeries) GetLastValue() (float64, float64) {
	return cs.XValues[len(cs.XValues)-1], cs.YValues[len(cs.YValues)-1]
}

// GetValueFormatters returns value formatter defaults for the series.
func (cs ContinuousSeries) GetValueFormatters() (x, y ValueFormatter) {
	x = FloatValueFormatter
	y = FloatValueFormatter
	return
}

// GetYAxis returns which YAxis the series draws on.
func (cs ContinuousSeries) GetYAxis() YAxisType {
	return cs.YAxis
}

// Render renders the series.
func (cs ContinuousSeries) Render(r Renderer, canvasBox Box, xrange, yrange Range, defaults Style) {
	style := cs.Style.InheritFrom(defaults)
	Draw.LineSeries(r, canvasBox, xrange, yrange, style, cs)
}
