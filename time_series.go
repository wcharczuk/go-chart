package chart

import (
	"fmt"
	"time"

	"github.com/blendlabs/go-util"
)

// TimeSeries is a line on a chart.
type TimeSeries struct {
	Name            string
	Style           Style
	FinalValueLabel Style

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

// GetXFormatter returns the x value formatter.
func (ts TimeSeries) GetXFormatter() Formatter {
	return func(v interface{}) string {
		if typed, isTyped := v.(time.Time); isTyped {
			return typed.Format(DefaultDateFormat)
		}
		if typed, isTyped := v.(int64); isTyped {
			return time.Unix(typed, 0).Format(DefaultDateFormat)
		}
		if typed, isTyped := v.(float64); isTyped {
			return time.Unix(int64(typed), 0).Format(DefaultDateFormat)
		}
		return util.StringEmpty
	}
}

// GetYFormatter returns the y value formatter.
func (ts TimeSeries) GetYFormatter() Formatter {
	return func(v interface{}) string {
		if typed, isTyped := v.(float64); isTyped {
			return fmt.Sprintf("%0.2f", typed)
		}
		return util.StringEmpty
	}
}
