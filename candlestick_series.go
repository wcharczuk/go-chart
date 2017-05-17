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

// String returns a string value for the candle value.
func (cv CandleValue) String() string {
	return fmt.Sprintf("candle %s high: %.2f low: %.2f open: %.2f close: %.2f", cv.Timestamp.Format("2006-01-02"), cv.High, cv.Low, cv.Open, cv.Close)
}

// IsZero returns if the value is zero or not.
func (cv CandleValue) IsZero() bool {
	return cv.Timestamp.IsZero()
}

// CandlestickSeries is a special type of series that takes a norma value provider
// and maps it to day value stats (high, low, open, close).
type CandlestickSeries struct {
	Name  string
	Style Style
	YAxis YAxisType

	// CandleValues will be used in place of creating them from the `InnerSeries`.
	CandleValues []CandleValue

	// InnerSeries is used if the `CandleValues` are not set.
	InnerSeries ValuesProvider
}

// GetName implements Series.GetName.
func (cs *CandlestickSeries) GetName() string {
	return cs.Name
}

// GetStyle implements Series.GetStyle.
func (cs *CandlestickSeries) GetStyle() Style {
	return cs.Style
}

// GetYAxis returns which yaxis the series is mapped to.
func (cs *CandlestickSeries) GetYAxis() YAxisType {
	return cs.YAxis
}

// Len returns the length of the series.
func (cs *CandlestickSeries) Len() int {
	return len(cs.GetCandleValues())
}

// GetBoundedValues returns the bounded values at a given index.
func (cs *CandlestickSeries) GetBoundedValues(index int) (x, y0, y1 float64) {
	value := cs.GetCandleValues()[index]
	return util.Time.ToFloat64(value.Timestamp), value.Low, value.High
}

// GetCandleValues returns the candle values.
func (cs CandlestickSeries) GetCandleValues() []CandleValue {
	if cs.CandleValues == nil {
		cs.CandleValues = cs.GenerateCandleValues()
	}
	return cs.CandleValues
}

// GenerateCandleValues returns the candlestick values for each day represented by the inner series.
func (cs CandlestickSeries) GenerateCandleValues() []CandleValue {
	if cs.InnerSeries == nil {
		return nil
	}

	totalValues := cs.InnerSeries.Len()
	if totalValues == 0 {
		return nil
	}

	var values []CandleValue
	var lastYear, lastMonth, lastDay int
	var year, month, day int

	var t time.Time
	var tv, lv, v float64

	tv, v = cs.InnerSeries.GetValues(0)
	t = util.Time.FromFloat64(tv)
	year, month, day = t.Year(), int(t.Month()), t.Day()

	lastYear, lastMonth, lastDay = year, month, day

	value := CandleValue{
		Timestamp: cs.newTimestamp(year, month, day),
		Open:      v,
		Low:       v,
		High:      v,
	}
	lv = v

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
				Open:      v,
				High:      v,
				Low:       v,
			}

			lastYear = year
			lastMonth = month
			lastDay = day
		} else {
			value.Low = math.Min(value.Low, v)
			value.High = math.Max(value.High, v)
		}
		lv = v
	}

	return values
}

func (cs CandlestickSeries) newTimestamp(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 12, 0, 0, 0, util.Date.Eastern())
}

// Render implements Series.Render.
func (cs CandlestickSeries) Render(r Renderer, canvasBox Box, xrange, yrange Range, defaults Style) {
	style := cs.Style.InheritFrom(defaults)
	Draw.CandlestickSeries(r, canvasBox, xrange, yrange, style, cs)
}

// Validate validates the series.
func (cs CandlestickSeries) Validate() error {
	if cs.CandleValues == nil && cs.InnerSeries == nil {
		return fmt.Errorf("candlestick series requires either `CandleValues` or `InnerSeries` to be set")
	}
	return nil
}
