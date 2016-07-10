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
	Render(c *Chart, r Renderer, canvasBox Box, ra Range)
}

// axis represents the basics of an axis implementation.
type axis struct {
	Name           string
	Style          Style
	ValueFormatter ValueFormatter
	Range          Range
	Ticks          []Tick
}

func (a axis) GetName() string {
	return a.Name
}

func (a axis) GetStyle() Style {
	return a.Style
}

func (a axis) getTicks(ra Range) []Tick {
	if len(a.Ticks) > 0 {
		return a.Ticks
	}
	return a.generateTicks(ra)
}

func (a axis) generateTicks(ra Range) []Tick {
	step := a.getTickStep(ra)
	return a.generateTicksWithStep(ra, step)
}

func (a axis) getTickCount(ra Range) int {
	return 0
}

func (a axis) getTickStep(ra Range) float64 {
	return 0.0
}

func (a axis) generateTicksWithStep(ra Range, step float64) []Tick {
	return []Tick{}
}
