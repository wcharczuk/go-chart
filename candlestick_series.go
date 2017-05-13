package chart

import (
	"fmt"
	"time"
)

type CandleValue struct {
	Timestamp time.Time
	High      float64
	Low       float64
	Open      float64
	Close     float64
}

// CandlestickSeries is a special type of series that takes a norma value provider
// and maps it to day value stats (high, low, open, close).
type CandlestickSeries struct {
	Name        string
	Style       Style
	YAxis       YAxisType
	InnerSeries ValuesProvider
}

// GetName implements Series.GetName.
func (cs CandlestickSeries) GetName() string {
	return cs.Name
}

// GetStyle implements Series.GetStyle.
func (cs CandlestickSeries) GetStyle() Style {
	return cs.Style
}

// GetYAxis returns which yaxis the series is mapped to.
func (cs CandlestickSeries) GetYAxis() YAxisType {
	return cs.YAxis
}

// CandleValues returns the candlestick values for each day represented by the inner series.
func (cs CandlestickSeries) CandleValues() []CandleValue {
	// for each "day" represented by the inner series
	// compute the open (i.e. the first value at or near market open)
	// compute the close (i.e. the last value at or near market close)
	// compute the high, or the max
	// compute the low, or the min

	totalValues := cs.InnerSeries.Len()

	var values []CandleValue

	var day int
	for i := 0; i < totalValues; i++ {
		if day == 0 {
			// extract day value from time value
		}
		if 
	}

	return values
}

// Render implements Series.Render.
func (cs CandlestickSeries) Render(r Renderer, canvasBox Box, xrange, yrange Range, defaults Style) {
	//style := cs.Style.InheritFrom(defaults)
	//Draw.CandlestickSeries(r, canvasBox, xrange, yrange, style, cs)
}

// Validate validates the series.
func (cs CandlestickSeries) Validate() error {
	if cs.InnerSeries == nil {
		return fmt.Errorf("histogram series requires InnerSeries to be set")
	}
	return nil
}
