package chart

const (
	// DefaultMACDPeriodPrimary is the long window.
	DefaultMACDPeriodPrimary = 26
	// DefaultMACDPeriodSecondary is the short window.
	DefaultMACDPeriodSecondary = 12
	// DefaultMACDSignalPeriod is the signal period to compute for the MACD.
	DefaultMACDSignalPeriod = 9
)

// MACDSeries computes the difference between the MACD line and the MACD Signal line.
// It is used in technical analysis and gives a lagging indicator of momentum.
type MACDSeries struct {
	Name        string
	Style       Style
	YAxis       YAxisType
	InnerSeries ValueProvider

	PrimaryPeriod   int
	SecondaryPeriod int
	SignalPeriod    int
}

// GetPeriods returns the primary and secondary periods.
func (macd MACDSeries) GetPeriods() (w1, w2, sig int) {
	if macd.PrimaryPeriod == 0 {
		w1 = DefaultMACDPeriodPrimary
	} else {
		w1 = macd.PrimaryPeriod
	}
	if macd.SecondaryPeriod == 0 {
		w2 = DefaultMACDPeriodSecondary
	} else {
		w2 = macd.SecondaryPeriod
	}
	if macd.SignalPeriod == 0 {
		sig = DefaultMACDSignalPeriod
	} else {
		sig = macd.SignalPeriod
	}
	return
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

	return macd.InnerSeries.Len()
}

// GetValue gets a value at a given index. For MACD it is the signal value.
func (macd MACDSeries) GetValue(index int) (x float64, y float64) {
	if macd.InnerSeries == nil {
		return
	}

	w1, w2, sig := macd.GetPeriods()

	x, _ = macd.InnerSeries.GetValue(index)

	signal := MACDSignalSeries{
		InnerSeries:     macd.InnerSeries,
		PrimaryPeriod:   w1,
		SecondaryPeriod: w2,
		SignalPeriod:    sig,
	}

	macdl := MACDLineSeries{
		InnerSeries:     macd.InnerSeries,
		PrimaryPeriod:   w1,
		SecondaryPeriod: w2,
	}

	_, lv := macdl.GetValue(index)
	_, sv := signal.GetValue(index)
	y = lv - sv

	return
}

// MACDSignalSeries computes the EMA of the MACDLineSeries.
type MACDSignalSeries struct {
	Name        string
	Style       Style
	YAxis       YAxisType
	InnerSeries ValueProvider

	PrimaryPeriod   int
	SecondaryPeriod int
	SignalPeriod    int
}

// GetPeriods returns the primary and secondary periods.
func (macds MACDSignalSeries) GetPeriods() (w1, w2, sig int) {
	if macds.PrimaryPeriod == 0 {
		w1 = DefaultMACDPeriodPrimary
	} else {
		w1 = macds.PrimaryPeriod
	}
	if macds.SecondaryPeriod == 0 {
		w2 = DefaultMACDPeriodSecondary
	} else {
		w2 = macds.SecondaryPeriod
	}
	if macds.SignalPeriod == 0 {
		sig = DefaultMACDSignalPeriod
	} else {
		sig = macds.SignalPeriod
	}
	return
}

// GetName returns the name of the time series.
func (macds MACDSignalSeries) GetName() string {
	return macds.Name
}

// GetStyle returns the line style.
func (macds MACDSignalSeries) GetStyle() Style {
	return macds.Style
}

// GetYAxis returns which YAxis the series draws on.
func (macds MACDSignalSeries) GetYAxis() YAxisType {
	return macds.YAxis
}

// Len returns the number of elements in the series.
func (macds MACDSignalSeries) Len() int {
	if macds.InnerSeries == nil {
		return 0
	}

	return macds.InnerSeries.Len()
}

// GetValue gets a value at a given index. For MACD it is the signal value.
func (macds MACDSignalSeries) GetValue(index int) (x float64, y float64) {
	if macds.InnerSeries == nil {
		return
	}

	w1, w2, sig := macds.GetPeriods()

	x, _ = macds.InnerSeries.GetValue(index)

	signal := EMASeries{
		InnerSeries: MACDLineSeries{
			InnerSeries:     macds.InnerSeries,
			PrimaryPeriod:   w1,
			SecondaryPeriod: w2,
		},
		Period: sig,
	}

	_, y = signal.GetValue(index)
	return
}

// Render renders the series.
func (macds MACDSignalSeries) Render(r Renderer, canvasBox Box, xrange, yrange Range, defaults Style) {
	style := macds.Style.WithDefaultsFrom(defaults)
	DrawLineSeries(r, canvasBox, xrange, yrange, style, macds)
}

// MACDLineSeries is a series that computes the inner ema1-ema2 value as a series.
type MACDLineSeries struct {
	Name        string
	Style       Style
	YAxis       YAxisType
	InnerSeries ValueProvider

	PrimaryPeriod   int
	SecondaryPeriod int

	Sigma float64
}

// GetName returns the name of the time series.
func (macdl MACDLineSeries) GetName() string {
	return macdl.Name
}

// GetStyle returns the line style.
func (macdl MACDLineSeries) GetStyle() Style {
	return macdl.Style
}

// GetYAxis returns which YAxis the series draws on.
func (macdl MACDLineSeries) GetYAxis() YAxisType {
	return macdl.YAxis
}

// GetPeriods returns the primary and secondary periods.
func (macdl MACDLineSeries) GetPeriods() (w1, w2 int) {
	if macdl.PrimaryPeriod == 0 {
		w1 = DefaultMACDPeriodPrimary
	} else {
		w1 = macdl.PrimaryPeriod
	}
	if macdl.SecondaryPeriod == 0 {
		w2 = DefaultMACDPeriodSecondary
	} else {
		w2 = macdl.SecondaryPeriod
	}
	return
}

// Len returns the number of elements in the series.
func (macdl MACDLineSeries) Len() int {
	if macdl.InnerSeries == nil {
		return 0
	}

	return macdl.InnerSeries.Len()
}

// GetValue gets a value at a given index. For MACD it is the signal value.
func (macdl MACDLineSeries) GetValue(index int) (x float64, y float64) {
	if macdl.InnerSeries == nil {
		return
	}

	w1, w2 := macdl.GetPeriods()

	x, _ = macdl.InnerSeries.GetValue(index)

	ema1 := EMASeries{
		InnerSeries: macdl.InnerSeries,
		Period:      w1,
	}

	ema2 := EMASeries{
		InnerSeries: macdl.InnerSeries,
		Period:      w2,
	}

	_, emav1 := ema1.GetValue(index)
	_, emav2 := ema2.GetValue(index)

	y = emav2 - emav1
	return
}

// Render renders the series.
func (macdl MACDLineSeries) Render(r Renderer, canvasBox Box, xrange, yrange Range, defaults Style) {
	style := macdl.Style.WithDefaultsFrom(defaults)
	DrawLineSeries(r, canvasBox, xrange, yrange, style, macdl)
}
