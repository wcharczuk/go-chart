package chart

import (
	"fmt"
	"time"

	"math"

	"github.com/wcharczuk/go-chart/util"
)

// CandleValue is a day's data for a candlestick plot.
type CandleValue struct {
	Timestamp time.Time
	High      float64
	Low       float64
	Open      float64
	Close     float64
}

// IsZero returns if the value is zero or not.
func (cv CandleValue) IsZero() bool {
	return cv.Timestamp.IsZero()
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
	if totalValues == 0 {
		return nil
	}

	var value CandleValue
	var values []CandleValue
	var lastYear, lastMonth, lastDay int
	var year, month, day int

	var tv float64
	var t time.Time
	var lv, v float64

	tv, v = cs.InnerSeries.GetValues(0)
	t = util.Time.FromFloat64(tv)
	year, month, day = t.Year(), int(t.Month()), t.Day()
	value.Timestamp = cs.newTimestamp(year, month, day)
	value.Open, value.Low, value.High = v, v, v

	for i := 1; i < totalValues; i++ {
		tv, v = cs.InnerSeries.GetValues(i)
		t = util.Time.FromFloat64(tv)
		year, month, day = t.Year(), int(t.Month()), t.Day()

		// if we've transitioned to a new day or we're on the last value
		if lastYear != year || lastMonth != month || lastDay != day || i == (totalValues-1) {
			value.Close = lv
			values = append(values, value)

			value = CandleValue{
				Timestamp: cs.newTimestamp(year, month, day),
			}
		}

		value.Low = math.Min(value.Low, v)
		value.High = math.Max(value.Low, v)
		lv = v
	}

	return values
}

func (cs CandlestickSeries) newTimestamp(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 12, 0, 0, 0, util.Date.Eastern())
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
