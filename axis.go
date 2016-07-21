package chart

// YAxisType is a type of y-axis; it can either be primary or secondary.
type YAxisType int

const (
	// YAxisPrimary is the primary axis.
	YAxisPrimary YAxisType = 0
	// YAxisSecondary is the secondary axis.
	YAxisSecondary YAxisType = 1
)

// Axis is a chart feature detailing what values happen where.
type Axis interface {
	GetName() string
	GetStyle() Style

	GetTicks() []Tick
	GenerateTicks(r Renderer, ra Range, vf ValueFormatter) []Tick

	// GetGridLines returns the gridlines for the axis.
	GetGridLines(ticks []Tick) []GridLine

	// Measure should return an absolute box for the axis.
	// This is used when auto-fitting the canvas to the background.
	Measure(r Renderer, canvasBox Box, ra Range, style Style, ticks []Tick) Box

	// Render renders the axis.
	Render(r Renderer, canvasBox Box, ra Range, style Style, ticks []Tick)
}
