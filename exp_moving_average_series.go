package chart

const (
	// DefaultExponentialMovingAverageSigma is the default exponential smoothing factor.
	DefaultExponentialMovingAverageSigma = 0.25
)

// ExponentialMovingAverageSeries is a computed series.
type ExponentialMovingAverageSeries struct {
	Name  string
	Style Style
	YAxis YAxisType

	// Sigma is the 'smoothing factor' parameter.
	Sigma       float64
	InnerSeries ValueProvider

	valueBuffer []float64
}

// GetName returns the name of the time series.
func (mas ExponentialMovingAverageSeries) GetName() string {
	return mas.Name
}

// GetStyle returns the line style.
func (mas ExponentialMovingAverageSeries) GetStyle() Style {
	return mas.Style
}

// GetYAxis returns which YAxis the series draws on.
func (mas ExponentialMovingAverageSeries) GetYAxis() YAxisType {
	return mas.YAxis
}

// Len returns the number of elements in the series.
func (mas *ExponentialMovingAverageSeries) Len() int {
	return mas.InnerSeries.Len()
}

// GetSigma returns the smoothing factor for the serise.
func (mas ExponentialMovingAverageSeries) GetSigma(defaults ...float64) float64 {
	if mas.Sigma == 0 {
		if len(defaults) > 0 {
			return defaults[0]
		}
		return DefaultExponentialMovingAverageSigma
	}
	return mas.Sigma
}

// GetValue gets a value at a given index.
func (mas *ExponentialMovingAverageSeries) GetValue(index int) (x float64, y float64) {
	if mas.InnerSeries == nil {
		return
	}
	if mas.valueBuffer == nil || index == 0 {
		mas.valueBuffer = make([]float64, mas.InnerSeries.Len())
	}
	vx, vy := mas.InnerSeries.GetValue(index)
	x = vx
	if index == 0 {
		mas.valueBuffer[0] = vy
		y = vy
		return
	}

	sig := mas.GetSigma()
	mas.valueBuffer[index] = sig*vy + ((1.0 - sig) * mas.valueBuffer[index-1])
	y = mas.valueBuffer[index]
	return
}

// GetLastValue computes the last moving average value but walking back window size samples,
// and recomputing the last moving average chunk.
func (mas ExponentialMovingAverageSeries) GetLastValue() (x float64, y float64) {
	if mas.InnerSeries == nil {
		return
	}

	seriesLength := mas.InnerSeries.Len()
	for index := 0; index < seriesLength; index++ {
		x, _ = mas.GetValue(index)
	}

	y = mas.valueBuffer[seriesLength-1]
	return
}

// Render renders the series.
func (mas *ExponentialMovingAverageSeries) Render(r Renderer, canvasBox Box, xrange, yrange Range, defaults Style) {
	style := mas.Style.WithDefaultsFrom(defaults)
	DrawLineSeries(r, canvasBox, xrange, yrange, style, mas)
}
