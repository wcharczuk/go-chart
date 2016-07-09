package chart

import (
	"sync"

	"github.com/golang/freetype/truetype"
	"github.com/wcharczuk/go-chart/drawing"
)

const (
	// DefaultChartHeight is the default chart height.
	DefaultChartHeight = 400
	// DefaultChartWidth is the default chart width.
	DefaultChartWidth = 200
	// DefaultStrokeWidth is the default chart line/stroke width.
	DefaultStrokeWidth = 1.0
	// DefaultAxisLineWidth is the line width of the axis lines.
	DefaultAxisLineWidth = 1.0
	//DefaultDPI is the default dots per inch for the chart.
	DefaultDPI = 92.0
	// DefaultMinimumFontSize is the default minimum font size.
	DefaultMinimumFontSize = 8.0
	// DefaultFontSize is the default font size.
	DefaultFontSize = 10.0
	// DefaultTitleFontSize is the default title font size.
	DefaultTitleFontSize = 18.0
	// DefaultFinalLabelDeltaWidth is the width of the left triangle out of the final label.
	DefaultFinalLabelDeltaWidth = 10
	// DefaultFinalLabelFontSize is the font size of the final label.
	DefaultFinalLabelFontSize = 10.0
	// DefaultAxisFontSize is the font size of the axis labels.
	DefaultAxisFontSize = 10.0
	// DefaultTitleTop is the default distance from the top of the chart to put the title.
	DefaultTitleTop = 10
	// DefaultXAxisMargin is the default distance from bottom of the canvas to the x axis labels.
	DefaultXAxisMargin = 10
	// DefaultMinimumTickHorizontalSpacing is the minimum distance between horizontal ticks.
	DefaultMinimumTickHorizontalSpacing = 20
	// DefaultMinimumTickVerticalSpacing is the minimum distance between vertical ticks.
	DefaultMinimumTickVerticalSpacing = 20
	// DefaultDateFormat is the default date format.
	DefaultDateFormat = "2006-01-02"
	// DefaultMaxTickCount is the maximum number of ticks to draw
	DefaultMaxTickCount = 7
)

var (
	// DefaultBackgroundColor is the default chart background color.
	// It is equivalent to css color:white.
	DefaultBackgroundColor = drawing.Color{R: 255, G: 255, B: 255, A: 255}
	// DefaultBackgroundStrokeColor is the default chart border color.
	// It is equivalent to color:white.
	DefaultBackgroundStrokeColor = drawing.Color{R: 255, G: 255, B: 255, A: 255}
	// DefaultCanvasColor is the default chart canvas color.
	// It is equivalent to css color:white.
	DefaultCanvasColor = drawing.Color{R: 255, G: 255, B: 255, A: 255}
	// DefaultCanvasStrokColor is the default chart canvas stroke color.
	// It is equivalent to css color:white.
	DefaultCanvasStrokColor = drawing.Color{R: 255, G: 255, B: 255, A: 255}
	// DefaultTextColor is the default chart text color.
	// It is equivalent to #333333.
	DefaultTextColor = drawing.Color{R: 51, G: 51, B: 51, A: 255}
	// DefaultAxisColor is the default chart axis line color.
	// It is equivalent to #333333.
	DefaultAxisColor = drawing.Color{R: 51, G: 51, B: 51, A: 255}
	// DefaultStrokeColor is the default chart border color.
	// It is equivalent to #efefef.
	DefaultStrokeColor = drawing.Color{R: 239, G: 239, B: 239, A: 255}
	// DefaultFillColor is the default fill color.
	// It is equivalent to #0074d9.
	DefaultFillColor = drawing.Color{R: 0, G: 217, B: 116, A: 255}
	// DefaultFinalLabelBackgroundColor is the default final label background color.
	DefaultFinalLabelBackgroundColor = drawing.Color{R: 255, G: 255, B: 255, A: 255}
)

var (
	// DefaultSeriesStrokeColors are a couple default series colors.
	DefaultSeriesStrokeColors = []drawing.Color{
		drawing.Color{R: 0, G: 116, B: 217, A: 255},
		drawing.Color{R: 0, G: 217, B: 116, A: 255},
		drawing.Color{R: 217, G: 0, B: 116, A: 255},
	}
)

// GetDefaultSeriesStrokeColor returns a color from the default list by index.
// NOTE: the index will wrap around (using a modulo).g
func GetDefaultSeriesStrokeColor(index int) drawing.Color {
	finalIndex := index % len(DefaultSeriesStrokeColors)
	return DefaultSeriesStrokeColors[finalIndex]
}

var (
	// DefaultFinalLabelPadding is the padding around the final label.
	DefaultFinalLabelPadding = Box{Top: 5, Left: 0, Right: 7, Bottom: 5}
	// DefaultBackgroundPadding is the default canvas padding config.
	DefaultBackgroundPadding = Box{Top: 5, Left: 5, Right: 5, Bottom: 5}
)

var (
	_defaultFontLock sync.Mutex
	_defaultFont     *truetype.Font
)

// GetDefaultFont returns the default font (Roboto-Medium).
func GetDefaultFont() (*truetype.Font, error) {
	if _defaultFont == nil {
		_defaultFontLock.Lock()
		defer _defaultFontLock.Unlock()
		if _defaultFont == nil {
			font, err := truetype.Parse(roboto)
			if err != nil {
				return nil, err
			}
			_defaultFont = font
		}
	}
	return _defaultFont, nil
}
