package chart

import (
	"fmt"

	"github.com/blendlabs/go-util"
)

// ContinuousSeries represents a line on a chart.
type ContinuousSeries struct {
	Name            string
	Style           Style
	FinalValueLabel Style

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

// GetXFormatter returns the xs value formatter.
func (cs ContinuousSeries) GetXFormatter() Formatter {
	return func(v interface{}) string {
		if typed, isTyped := v.(float64); isTyped {
			return fmt.Sprintf("%0.2f", typed)
		}
		return util.StringEmpty
	}
}

// GetYFormatter returns the y value formatter.
func (cs ContinuousSeries) GetYFormatter() Formatter {
	return cs.GetXFormatter()
}

// Render renders the series.
func (cs ContinuousSeries) Render(c *Chart, r Renderer, canvasBox Box, xrange, yrange Range) error {
	DrawLineSeries(c, r, canvasBox, xrange, yrange, cs)

	if cs.FinalValueLabel.Show {
		asw := 0
		if c.Axes.Show {
			asw = int(c.Axes.GetStrokeWidth(DefaultAxisLineWidth))
		}

		_, lv := cs.GetValue(cs.Len() - 1)
		ll := yrange.Format(lv)
		lx := canvasBox.Right + asw
		ly := yrange.Translate(lv) + canvasBox.Top
		DrawAnnotation(c, r, canvasBox, xrange, yrange, cs.FinalValueLabel, lx, ly, ll)
	}
}
