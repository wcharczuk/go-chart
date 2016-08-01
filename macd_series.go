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

	signal *MACDSignalSeries
	macdl  *MACDLineSeries
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
func (macd *MACDSeries) GetValue(index int) (x float64, y float64) {
	if macd.InnerSeries == nil {
		return
	}

	if macd.signal == nil || macd.macdl == nil {
		macd.ensureChildSeries()
	}

	_, lv := macd.macdl.GetValue(index)
	_, sv := macd.signal.GetValue(index)

	x, _ = macd.InnerSeries.GetValue(index)
	y = lv - sv

	return
}

func (macd *MACDSeries) ensureChildSeries() {
	w1, w2, sig := macd.GetPeriods()

	macd.signal = &MACDSignalSeries{
		InnerSeries:     macd.InnerSeries,
		PrimaryPeriod:   w1,
		SecondaryPeriod: w2,
		SignalPeriod:    sig,
	}

	macd.macdl = &MACDLineSeries{
		InnerSeries:     macd.InnerSeries,
		PrimaryPeriod:   w1,
		SecondaryPeriod: w2,
	}
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

	signal *EMASeries
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
func (macds *MACDSignalSeries) Len() int {
	if macds.InnerSeries == nil {
		return 0
	}

	return macds.InnerSeries.Len()
}

// GetValue gets a value at a given index. For MACD it is the signal value.
func (macds *MACDSignalSeries) GetValue(index int) (x float64, y float64) {
	if macds.InnerSeries == nil {
		return
	}

	if macds.signal == nil {
		macds.ensureSignal()
	}
	x, _ = macds.InnerSeries.GetValue(index)
	_, y = macds.signal.GetValue(index)
	return
}

func (macds *MACDSignalSeries) ensureSignal() {
	w1, w2, sig := macds.GetPeriods()

	macds.signal = &EMASeries{
		InnerSeries: &MACDLineSeries{
			InnerSeries:     macds.InnerSeries,
			PrimaryPeriod:   w1,
			SecondaryPeriod: w2,
		},
		Period: sig,
	}
}

// Render renders the series.
func (macds *MACDSignalSeries) Render(r Renderer, canvasBox Box, xrange, yrange Range, defaults Style) {
	style := macds.Style.InheritFrom(defaults)
	Draw.LineSeries(r, canvasBox, xrange, yrange, style, macds)
}

// MACDLineSeries is a series that computes the inner ema1-ema2 value as a series.
type MACDLineSeries struct {
	Name        string
	Style       Style
	YAxis       YAxisType
	InnerSeries ValueProvider

	PrimaryPeriod   int
	SecondaryPeriod int

	ema1 *EMASeries
	ema2 *EMASeries

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
func (macdl *MACDLineSeries) Len() int {
	if macdl.InnerSeries == nil {
		return 0
	}

	return macdl.InnerSeries.Len()
}

// GetValue gets a value at a given index. For MACD it is the signal value.
func (macdl *MACDLineSeries) GetValue(index int) (x float64, y float64) {
	if macdl.InnerSeries == nil {
		return
	}
	if macdl.ema1 == nil && macdl.ema2 == nil {
		macdl.ensureEMASeries()
	}

	x, _ = macdl.InnerSeries.GetValue(index)

	_, emav1 := macdl.ema1.GetValue(index)
	_, emav2 := macdl.ema2.GetValue(index)

	y = emav2 - emav1
	return
}

func (macdl *MACDLineSeries) ensureEMASeries() {
	w1, w2 := macdl.GetPeriods()

	macdl.ema1 = &EMASeries{
		InnerSeries: macdl.InnerSeries,
		Period:      w1,
	}
	macdl.ema2 = &EMASeries{
		InnerSeries: macdl.InnerSeries,
		Period:      w2,
	}
}

// Render renders the series.
func (macdl *MACDLineSeries) Render(r Renderer, canvasBox Box, xrange, yrange Range, defaults Style) {
	style := macdl.Style.InheritFrom(defaults)
	Draw.LineSeries(r, canvasBox, xrange, yrange, style, macdl)
}
