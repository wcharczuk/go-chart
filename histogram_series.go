package chart

// HistogramSeries is a special type of series that draws as a histogram.
// Some peculiarities; it will always be lower bounded at 0 (at the very least).
// This may alter ranges a bit and generally you want to put a histogram series on it's own y-axis.
type HistogramSeries struct {
	Name        string
	Style       Style
	YAxis       YAxisType
	InnerSeries ValueProvider
}

// GetName implements Series.GetName.
func (hs HistogramSeries) GetName() string {
	return hs.Name
}

// GetStyle implements Series.GetStyle.
func (hs HistogramSeries) GetStyle() Style {
	return hs.Style
}

// GetYAxis returns which yaxis the series is mapped to.
func (hs HistogramSeries) GetYAxis() YAxisType {
	return hs.YAxis
}

// Len implements BoundedValueProvider.Len.
func (hs HistogramSeries) Len() int {
	return hs.InnerSeries.Len()
}

// GetValue implements ValueProvider.GetValue.
func (hs HistogramSeries) GetValue(index int) (x, y float64) {
	return hs.InnerSeries.GetValue(index)
}

// GetBoundedValue implements BoundedValueProvider.GetBoundedValue
func (hs HistogramSeries) GetBoundedValue(index int) (x, y1, y2 float64) {
	vx, vy := hs.InnerSeries.GetValue(index)

	x = vx

	if vy > 0 {
		y1 = vy
		return
	}

	y2 = vy
	return
}

// Render implements Series.Render.
func (hs HistogramSeries) Render(r Renderer, canvasBox Box, xrange, yrange Range, defaults Style) {
	style := hs.Style.InheritFrom(defaults)
	Draw.HistogramSeries(r, canvasBox, xrange, yrange, style, hs)
}
