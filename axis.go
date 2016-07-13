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
	GetTicks(r Renderer, ra Range, vf ValueFormatter) []Tick
	GetGridLines(ticks []Tick) []GridLine
	Render(c *Chart, r Renderer, canvasBox Box, ra Range, ticks []Tick)
}
