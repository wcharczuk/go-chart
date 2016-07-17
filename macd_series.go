package chart

const (
	// DefaultMACDWindowPrimary is the long window.
	DefaultMACDWindowPrimary = 26
	// DefaultMACDWindowSecondary is the short window.
	DefaultMACDWindowSecondary = 12
)

// MACDSeries (or Moving Average Convergence Divergence) is a special type of series that
// computes the difference between two different EMA values for a given index, as denoted by WindowPrimary(26) and WindowSecondary(12).
type MACDSeries struct {
	Name        string
	Style       Style
	YAxis       YAxisType
	InnerSeries ValueProvider

	WindowPrimary   int
	WindowSecondary int

	Sigma float64
}

// GetWindows returns the primary and secondary window sizes.
func (macd MACDSeries) GetWindows() (w1, w2 int) {
	if macd.WindowPrimary == 0 {
		w1 = DefaultMACDWindowPrimary
	} else {
		w1 = macd.WindowPrimary
	}
	if macd.WindowSecondary == 0 {
		w2 = DefaultMACDWindowSecondary
	} else {
		w2 = macd.WindowSecondary
	}
	return
}

// GetSigma returns the smoothing factor for the serise.
func (macd MACDSeries) GetSigma(defaults ...float64) float64 {
	if macd.Sigma == 0 {
		if len(defaults) > 0 {
			return defaults[0]
		}
		return DefaultExponentialMovingAverageSigma
	}
	return macd.Sigma
}

// GetName returns the name of the time series.
func (macd MACDSeries) GetName() string {
	return macd.Name
}

// GetStyle returns the line style.
func (macd MACDSeries) GetStyle() Style {
	return macd.Style
}

// GetYAxis returns which YAxis the series draws on.
func (macd MACDSeries) GetYAxis() YAxisType {
	return macd.YAxis
}

// Len returns the number of elements in the series.
func (macd MACDSeries) Len() int {
	if macd.InnerSeries == nil {
		return 0
	}

	w1, _ := macd.GetWindows()
	innerLen := macd.InnerSeries.Len()
	if innerLen > w1 {
		return innerLen - w1
	}
	return 0
}

// GetValue gets a value at a given index.
func (macd MACDSeries) GetValue(index int) (x float64, y float64) {
	if macd.InnerSeries == nil {
		return
	}

	w1, w2 := macd.GetWindows()

	effectiveIndex := index + w1
	x, _ = macd.InnerSeries.GetValue(effectiveIndex)

	ema1 := macd.computeEMA(w1, effectiveIndex)
	ema2 := macd.computeEMA(w2, effectiveIndex)

	y = ema1 - ema2
	return
}

func (macd MACDSeries) computeEMA(windowSize int, index int) float64 {
	_, v := macd.InnerSeries.GetValue(index)
	if windowSize == 1 {
		return v
	}
	sig := macd.GetSigma()
	return sig*v + ((1.0 - sig) * macd.computeEMA(windowSize-1, index-1))
}

// Render renders the series.
func (macd MACDSeries) Render(r Renderer, canvasBox Box, xrange, yrange Range, defaults Style) {
	style := macd.Style.WithDefaultsFrom(defaults)
	DrawLineSeries(r, canvasBox, xrange, yrange, style, macd)
}
