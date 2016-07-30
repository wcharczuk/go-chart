package chart

// LinearRegressionSeries is a series that plots the n-nearest neighbors
// linear regression for the values.
type LinearRegressionSeries struct {
	Name  string
	Style Style
	YAxis YAxisType

	Window      int
	Offset      int
	InnerSeries ValueProvider

	m       float64
	b       float64
	avgx    float64
	stddevx float64
}

// GetName returns the name of the time series.
func (lrs LinearRegressionSeries) GetName() string {
	return lrs.Name
}

// GetStyle returns the line style.
func (lrs LinearRegressionSeries) GetStyle() Style {
	return lrs.Style
}

// GetYAxis returns which YAxis the series draws on.
func (lrs LinearRegressionSeries) GetYAxis() YAxisType {
	return lrs.YAxis
}

// Len returns the number of elements in the series.
func (lrs LinearRegressionSeries) Len() int {
	return Math.MinInt(lrs.GetWindow(), lrs.InnerSeries.Len()-lrs.GetOffset())
}

// GetWindow returns the window size.
func (lrs LinearRegressionSeries) GetWindow() int {
	if lrs.Window == 0 {
		return lrs.InnerSeries.Len()
	}
	return lrs.Window
}

// GetEndIndex returns the effective window end.
func (lrs LinearRegressionSeries) GetEndIndex() int {
	return Math.MinInt(lrs.GetOffset()+(lrs.Len()), (lrs.InnerSeries.Len() - 1))
}

// GetOffset returns the data offset.
func (lrs LinearRegressionSeries) GetOffset() int {
	if lrs.Offset == 0 {
		return 0
	}
	return lrs.Offset
}

// GetValue gets a value at a given index.
func (lrs *LinearRegressionSeries) GetValue(index int) (x, y float64) {
	if lrs.InnerSeries == nil {
		return
	}
	if lrs.m == 0 && lrs.b == 0 {
		lrs.computeCoefficients()
	}
	offset := lrs.GetOffset()
	effectiveIndex := Math.MinInt(index+offset, lrs.InnerSeries.Len())
	x, y = lrs.InnerSeries.GetValue(effectiveIndex)
	y = (lrs.m * lrs.normalize(x)) + lrs.b
	return
}

// GetLastValue computes the last moving average value but walking back window size samples,
// and recomputing the last moving average chunk.
func (lrs *LinearRegressionSeries) GetLastValue() (x, y float64) {
	if lrs.InnerSeries == nil {
		return
	}
	if lrs.m == 0 && lrs.b == 0 {
		lrs.computeCoefficients()
	}
	endIndex := lrs.GetEndIndex()
	x, y = lrs.InnerSeries.GetValue(endIndex)
	y = (lrs.m * lrs.normalize(x)) + lrs.b
	return
}

func (lrs *LinearRegressionSeries) normalize(xvalue float64) float64 {
	return (xvalue - lrs.avgx) / lrs.stddevx
}

// computeCoefficients computes the `m` and `b` terms in the linear formula given by `y = mx+b`.
func (lrs *LinearRegressionSeries) computeCoefficients() {
	startIndex := lrs.GetOffset()
	endIndex := lrs.GetEndIndex()

	p := float64(endIndex - startIndex)

	xvalues := NewRingBufferWithCapacity(lrs.Len())
	for index := startIndex; index < endIndex; index++ {

		x, _ := lrs.InnerSeries.GetValue(index)
		xvalues.Enqueue(x)
	}

	lrs.avgx = xvalues.Average()
	lrs.stddevx = xvalues.StdDev()

	var sumx, sumy, sumxx, sumxy float64
	for index := startIndex; index < endIndex; index++ {
		x, y := lrs.InnerSeries.GetValue(index)

		x = lrs.normalize(x)

		sumx += x
		sumy += y
		sumxx += x * x
		sumxy += x * y
	}

	lrs.m = (p*sumxy - sumx*sumy) / (p*sumxx - sumx*sumx)
	lrs.b = (sumy / p) - (lrs.m * sumx / p)
}

// Render renders the series.
func (lrs *LinearRegressionSeries) Render(r Renderer, canvasBox Box, xrange, yrange Range, defaults Style) {
	style := lrs.Style.InheritFrom(defaults)
	Draw.LineSeries(r, canvasBox, xrange, yrange, style, lrs)
}
