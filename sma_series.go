package chart

const (
	// DefaultSimpleMovingAveragePeriod is the default number of values to average.
	DefaultSimpleMovingAveragePeriod = 16
)

// SMASeries is a computed series.
type SMASeries struct {
	Name  string
	Style Style
	YAxis YAxisType

	Period      int
	InnerSeries ValueProvider
}

// GetName returns the name of the time series.
func (sma SMASeries) GetName() string {
	return sma.Name
}

// GetStyle returns the line style.
func (sma SMASeries) GetStyle() Style {
	return sma.Style
}

// GetYAxis returns which YAxis the series draws on.
func (sma SMASeries) GetYAxis() YAxisType {
	return sma.YAxis
}

// Len returns the number of elements in the series.
func (sma SMASeries) Len() int {
	return sma.InnerSeries.Len()
}

// GetPeriod returns the window size.
func (sma SMASeries) GetPeriod(defaults ...int) int {
	if sma.Period == 0 {
		if len(defaults) > 0 {
			return defaults[0]
		}
		return DefaultSimpleMovingAveragePeriod
	}
	return sma.Period
}

// GetValue gets a value at a given index.
func (sma SMASeries) GetValue(index int) (x, y float64) {
	if sma.InnerSeries == nil {
		return
	}
	px, _ := sma.InnerSeries.GetValue(index)
	x = px
	y = sma.getAverage(index)
	return
}

// GetLastValue computes the last moving average value but walking back window size samples,
// and recomputing the last moving average chunk.
func (sma SMASeries) GetLastValue() (x, y float64) {
	if sma.InnerSeries == nil {
		return
	}
	seriesLen := sma.InnerSeries.Len()
	px, _ := sma.InnerSeries.GetValue(seriesLen - 1)
	x = px
	y = sma.getAverage(seriesLen - 1)
	return
}

func (sma SMASeries) getAverage(index int) float64 {
	period := sma.GetPeriod()
	floor := Math.MaxInt(0, index-period)
	var accum float64
	var count float64
	for x := index; x >= floor; x-- {
		_, vy := sma.InnerSeries.GetValue(x)
		accum += vy
		count += 1.0
	}
	return accum / count
}

// Render renders the series.
func (sma SMASeries) Render(r Renderer, canvasBox Box, xrange, yrange Range, defaults Style) {
	style := sma.Style.InheritFrom(defaults)
	Draw.LineSeries(r, canvasBox, xrange, yrange, style, sma)
}
