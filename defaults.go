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
	DefaultChartWidth = 1024
	// DefaultStrokeWidth is the default chart stroke width.
	DefaultStrokeWidth = 0.0
	// DefaultDotWidth is the default chart dot width.
	DefaultDotWidth = 0.0
	// DefaultSeriesLineWidth is the default line width.
	DefaultSeriesLineWidth = 1.0
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
	// DefaultAnnotationDeltaWidth is the width of the left triangle out of annotations.
	DefaultAnnotationDeltaWidth = 10
	// DefaultAnnotationFontSize is the font size of annotations.
	DefaultAnnotationFontSize = 10.0
	// DefaultAxisFontSize is the font size of the axis labels.
	DefaultAxisFontSize = 10.0
	// DefaultTitleTop is the default distance from the top of the chart to put the title.
	DefaultTitleTop = 10

	// DefaultBackgroundStrokeWidth is the default stroke on the chart background.
	DefaultBackgroundStrokeWidth = 0.0
	// DefaultCanvasStrokeWidth is the default stroke on the chart canvas.
	DefaultCanvasStrokeWidth = 0.0

	// DefaultLineSpacing is the default vertical distance between lines of text.
	DefaultLineSpacing = 5

	// DefaultYAxisMargin is the default distance from the right of the canvas to the y axis labels.
	DefaultYAxisMargin = 10
	// DefaultXAxisMargin is the default distance from bottom of the canvas to the x axis labels.
	DefaultXAxisMargin = 10

	//DefaultVerticalTickHeight is half the margin.
	DefaultVerticalTickHeight = DefaultXAxisMargin >> 1
	//DefaultHorizontalTickWidth is half the margin.
	DefaultHorizontalTickWidth = DefaultYAxisMargin >> 1

	// DefaultTickCount is the default number of ticks to show
	DefaultTickCount = 10
	// DefaultTickCountSanityCheck is a hard limit on number of ticks to prevent infinite loops.
	DefaultTickCountSanityCheck = 1 << 10 //1024

	// DefaultMinimumTickHorizontalSpacing is the minimum distance between horizontal ticks.
	DefaultMinimumTickHorizontalSpacing = 20
	// DefaultMinimumTickVerticalSpacing is the minimum distance between vertical ticks.
	DefaultMinimumTickVerticalSpacing = 20

	// DefaultDateFormat is the default date format.
	DefaultDateFormat = "2006-01-02"
	// DefaultDateHourFormat is the date format for hour timestamp formats.
	DefaultDateHourFormat = "01-02 3PM"
	// DefaultDateMinuteFormat is the date format for minute range timestamp formats.
	DefaultDateMinuteFormat = "01-02 3:04PM"
	// DefaultFloatFormat is the default float format.
	DefaultFloatFormat = "%.2f"
	// DefaultPercentValueFormat is the default percent format.
	DefaultPercentValueFormat = "%0.2f%%"

	// DefaultBarSpacing is the default pixel spacing between bars.
	DefaultBarSpacing = 100
	// DefaultBarWidth is the default pixel width of bars in a bar chart.
	DefaultBarWidth = 50
)

var (
	// ColorWhite is white.
	ColorWhite = drawing.Color{R: 255, G: 255, B: 255, A: 255}
	// ColorBlue is the basic theme blue color.
	ColorBlue = drawing.Color{R: 0, G: 116, B: 217, A: 255}
	// ColorCyan is the basic theme cyan color.
	ColorCyan = drawing.Color{R: 0, G: 217, B: 210, A: 255}
	// ColorGreen is the basic theme green color.
	ColorGreen = drawing.Color{R: 0, G: 217, B: 101, A: 255}
	// ColorRed is the basic theme red color.
	ColorRed = drawing.Color{R: 217, G: 0, B: 116, A: 255}
	// ColorOrange is the basic theme orange color.
	ColorOrange = drawing.Color{R: 217, G: 101, B: 0, A: 255}
	// ColorYellow is the basic theme yellow color.
	ColorYellow = drawing.Color{R: 217, G: 210, B: 0, A: 255}
	// ColorBlack is the basic theme black color.
	ColorBlack = drawing.Color{R: 51, G: 51, B: 51, A: 255}
	// ColorLightGray is the basic theme light gray color.
	ColorLightGray = drawing.Color{R: 239, G: 239, B: 239, A: 255}

	// ColorAlternateBlue is a alternate theme color.
	ColorAlternateBlue = drawing.Color{R: 106, G: 195, B: 203, A: 255}
	// ColorAlternateGreen is a alternate theme color.
	ColorAlternateGreen = drawing.Color{R: 42, G: 190, B: 137, A: 255}
	// ColorAlternateGray is a alternate theme color.
	ColorAlternateGray = drawing.Color{R: 110, G: 128, B: 139, A: 255}
	// ColorAlternateYellow is a alternate theme color.
	ColorAlternateYellow = drawing.Color{R: 240, G: 174, B: 90, A: 255}
	// ColorAlternateLightGray is a alternate theme color.
	ColorAlternateLightGray = drawing.Color{R: 187, G: 190, B: 191, A: 255}

	// ColorTransparent is a transparent (alpha zero) color.
	ColorTransparent = drawing.Color{R: 1, G: 1, B: 1, A: 0}
)

var (
	// DefaultBackgroundColor is the default chart background color.
	// It is equivalent to css color:white.
	DefaultBackgroundColor = ColorWhite
	// DefaultBackgroundStrokeColor is the default chart border color.
	// It is equivalent to color:white.
	DefaultBackgroundStrokeColor = ColorWhite
	// DefaultCanvasColor is the default chart canvas color.
	// It is equivalent to css color:white.
	DefaultCanvasColor = ColorWhite
	// DefaultCanvasStrokeColor is the default chart canvas stroke color.
	// It is equivalent to css color:white.
	DefaultCanvasStrokeColor = ColorWhite
	// DefaultTextColor is the default chart text color.
	// It is equivalent to #333333.
	DefaultTextColor = ColorBlack
	// DefaultAxisColor is the default chart axis line color.
	// It is equivalent to #333333.
	DefaultAxisColor = ColorBlack
	// DefaultStrokeColor is the default chart border color.
	// It is equivalent to #efefef.
	DefaultStrokeColor = ColorLightGray
	// DefaultFillColor is the default fill color.
	// It is equivalent to #0074d9.
	DefaultFillColor = ColorBlue
	// DefaultAnnotationFillColor is the default annotation background color.
	DefaultAnnotationFillColor = ColorWhite
	// DefaultGridLineColor is the default grid line color.
	DefaultGridLineColor = ColorLightGray
)

var (
	// DefaultColors are a couple default series colors.
	DefaultColors = []drawing.Color{
		ColorBlue,
		ColorGreen,
		ColorRed,
		ColorCyan,
		ColorOrange,
	}

	// DefaultAlternateColors are a couple alternate colors.
	DefaultAlternateColors = []drawing.Color{
		ColorAlternateBlue,
		ColorAlternateGreen,
		ColorAlternateGray,
		ColorAlternateYellow,
		ColorBlue,
		ColorGreen,
		ColorRed,
		ColorCyan,
		ColorOrange,
	}
)

var (
	// DashArrayDots is a dash array that represents '....' style stroke dashes.
	DashArrayDots = []int{1, 1}
	// DashArrayDashesSmall is a dash array that represents '- - -' style stroke dashes.
	DashArrayDashesSmall = []int{3, 3}
	// DashArrayDashesMedium is a dash array that represents '-- -- --' style stroke dashes.
	DashArrayDashesMedium = []int{5, 5}
	// DashArrayDashesLarge is a dash array that represents '----- ----- -----' style stroke dashes.
	DashArrayDashesLarge = []int{10, 10}
)

// NewColor returns a new color.
func NewColor(r, g, b, a uint8) drawing.Color {
	return drawing.Color{R: r, G: g, B: b, A: a}
}

// GetDefaultColor returns a color from the default list by index.
// NOTE: the index will wrap around (using a modulo).
func GetDefaultColor(index int) drawing.Color {
	finalIndex := index % len(DefaultColors)
	return DefaultColors[finalIndex]
}

// GetAlternateColor returns a color from the default list by index.
// NOTE: the index will wrap around (using a modulo).
func GetAlternateColor(index int) drawing.Color {
	finalIndex := index % len(DefaultAlternateColors)
	return DefaultAlternateColors[finalIndex]
}

var (
	// DefaultAnnotationPadding is the padding around an annotation.
	DefaultAnnotationPadding = Box{Top: 5, Left: 5, Right: 5, Bottom: 5}
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
