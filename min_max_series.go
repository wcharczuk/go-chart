package chart

import "math"

// MinSeries draws a horizontal line at the minimum value of the inner series.
type MinSeries struct {
	Name        string
	Style       Style
	YAxis       YAxisType
	InnerSeries ValueProvider

	minValue *float64
}

// GetName returns the name of the time series.
func (ms MinSeries) GetName() string {
	return ms.Name
}

// GetStyle returns the line style.
func (ms MinSeries) GetStyle() Style {
	return ms.Style
}

// GetYAxis returns which YAxis the series draws on.
func (ms MinSeries) GetYAxis() YAxisType {
	return ms.YAxis
}

// Len returns the number of elements in the series.
func (ms MinSeries) Len() int {
	return ms.InnerSeries.Len()
}

// GetValue gets a value at a given index.
func (ms *MinSeries) GetValue(index int) (x, y float64) {
	ms.ensureMinValue()
	x, _ = ms.InnerSeries.GetValue(index)
	y = *ms.minValue
	return
}

// Render renders the series.
func (ms *MinSeries) Render(r Renderer, canvasBox Box, xrange, yrange Range, defaults Style) {
	style := ms.Style.InheritFrom(defaults)
	Draw.LineSeries(r, canvasBox, xrange, yrange, style, ms)
}

func (ms *MinSeries) ensureMinValue() {
	if ms.minValue == nil {
		minValue := math.MaxFloat64
		var y float64
		for x := 0; x < ms.InnerSeries.Len(); x++ {
			_, y = ms.InnerSeries.GetValue(x)
			if y < minValue {
				minValue = y
			}
		}
		ms.minValue = &minValue
	}
}

// MaxSeries draws a horizontal line at the maximum value of the inner series.
type MaxSeries struct {
	Name        string
	Style       Style
	YAxis       YAxisType
	InnerSeries ValueProvider

	maxValue *float64
}

// GetName returns the name of the time series.
func (ms MaxSeries) GetName() string {
	return ms.Name
}

// GetStyle returns the line style.
func (ms MaxSeries) GetStyle() Style {
	return ms.Style
}

// GetYAxis returns which YAxis the series draws on.
func (ms MaxSeries) GetYAxis() YAxisType {
	return ms.YAxis
}

// Len returns the number of elements in the series.
func (ms MaxSeries) Len() int {
	return ms.InnerSeries.Len()
}

// GetValue gets a value at a given index.
func (ms *MaxSeries) GetValue(index int) (x, y float64) {
	ms.ensureMaxValue()
	x, _ = ms.InnerSeries.GetValue(index)
	y = *ms.maxValue
	return
}

// Render renders the series.
func (ms *MaxSeries) Render(r Renderer, canvasBox Box, xrange, yrange Range, defaults Style) {
	style := ms.Style.InheritFrom(defaults)
	Draw.LineSeries(r, canvasBox, xrange, yrange, style, ms)
}

func (ms *MaxSeries) ensureMaxValue() {
	if ms.maxValue == nil {
		maxValue := -math.MaxFloat64
		var y float64
		for x := 0; x < ms.InnerSeries.Len(); x++ {
			_, y = ms.InnerSeries.GetValue(x)
			if y > maxValue {
				maxValue = y
			}
		}
		ms.maxValue = &maxValue
	}
}
