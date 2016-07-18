package chart

const (
	// DefaultEMASigma is the default exponential smoothing factor.
	DefaultEMASigma = 0.25
)

// EMASeries is a computed series.
type EMASeries struct {
	Name  string
	Style Style
	YAxis YAxisType

	Period      int
	InnerSeries ValueProvider
}

// GetName returns the name of the time series.
func (ema EMASeries) GetName() string {
	return ema.Name
}

// GetStyle returns the line style.
func (ema EMASeries) GetStyle() Style {
	return ema.Style
}

// GetYAxis returns which YAxis the series draws on.
func (ema EMASeries) GetYAxis() YAxisType {
	return ema.YAxis
}

// GetPeriod returns the window size.
func (ema EMASeries) GetPeriod(defaults ...int) int {
	if ema.Period == 0 {
		if len(defaults) > 0 {
			return defaults[0]
		}
		return ema.InnerSeries.Len()
	}
	return ema.Period
}

// Len returns the number of elements in the series.
func (ema EMASeries) Len() int {
	return ema.InnerSeries.Len()
}

// GetSigma returns the smoothing factor for the serise.
func (ema EMASeries) GetSigma() float64 {
	return 2.0 / (float64(ema.Period) + 1)
}

// GetValue gets a value at a given index.
func (ema EMASeries) GetValue(index int) (x float64, y float64) {
	if ema.InnerSeries == nil {
		return
	}
	vx, _ := ema.InnerSeries.GetValue(index)
	x = vx
	y = ema.compute(ema.GetPeriod(), index)
	return
}

// GetLastValue computes the last moving average value but walking back window size samples,
// and recomputing the last moving average chunk.
func (ema EMASeries) GetLastValue() (x float64, y float64) {
	if ema.InnerSeries == nil {
		return
	}
	lastIndex := ema.InnerSeries.Len() - 1
	x, _ = ema.InnerSeries.GetValue(lastIndex)
	y = ema.compute(ema.GetPeriod(), lastIndex)
	return
}

func (ema EMASeries) compute(period, index int) float64 {
	_, v := ema.InnerSeries.GetValue(index)
	if index == 0 {
		return v
	}
	previousEMA := ema.compute(period-1, index-1)
	return ((v - previousEMA) * ema.GetSigma()) + previousEMA
}

// Render renders the series.
func (ema EMASeries) Render(r Renderer, canvasBox Box, xrange, yrange Range, defaults Style) {
	style := ema.Style.WithDefaultsFrom(defaults)
	DrawLineSeries(r, canvasBox, xrange, yrange, style, ema)
}
