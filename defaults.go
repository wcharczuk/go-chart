package chart

import (
	"image/color"
	"sync"

	"github.com/golang/freetype/truetype"
)

const (
	// DefaultChartHeight is the default chart height.
	DefaultChartHeight = 400
	// DefaultChartWidth is the default chart width.
	DefaultChartWidth = 200
	// DefaultPadding is the default gap between the image border and
	// chart content (referred to as the "canvas").
	DefaultPadding = 10
	// DefaultLineWidth is the default chart line width.
	DefaultLineWidth = 2.0
	//DefaultDPI is the default dots per inch for the chart.
	DefaultDPI = 120.0
	// DefaultMinimumFontSize is the default minimum font size.
	DefaultMinimumFontSize = 8.0
)

var (
	// DefaultBackgroundColor is the default chart background color.
	// It is equivalent to css color:white.
	DefaultBackgroundColor = color.RGBA{R: 255, G: 255, B: 255, A: 255}
	// DefaultCanvasColor is the default chart canvas color.
	// It is equivalent to css color:white.
	DefaultCanvasColor = color.RGBA{R: 255, G: 255, B: 255, A: 255}
	// DefaultTextColor is the default chart text color.
	// It is equivalent to #333333.
	DefaultTextColor = color.RGBA{R: 51, G: 51, B: 51, A: 255}
	// DefaultAxisColor is the default chart axis line color.
	// It is equivalent to #333333.
	DefaultAxisColor = color.RGBA{R: 51, G: 51, B: 51, A: 255}
	// DefaultBorderColor is the default chart border color.
	// It is equivalent to #efefef.
	DefaultBorderColor = color.RGBA{R: 239, G: 239, B: 239, A: 255}
	// DefaultLineColor is the default (1st) series line color.
	// It is equivalent to #0074d9.
	DefaultLineColor = color.RGBA{R: 0, G: 217, B: 116, A: 255}
	// DefaultFillColor is the default fill color.
	// It is equivalent to #0074d9.
	DefaultFillColor = color.RGBA{R: 0, G: 217, B: 116, A: 255}
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
