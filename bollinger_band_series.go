package chart

import "math"

// BollingerBandsSeries draws bollinger bands for an inner series.
// Bollinger bands are defined by two lines, one at SMA+k*stddev, one at SMA-k*stdev.
type BollingerBandsSeries struct {
	Name  string
	Style Style
	YAxis YAxisType

	WindowSize  int
	K           float64
	InnerSeries ValueProvider

	valueBuffer *RingBuffer
}

// GetName returns the name of the time series.
func (bbs BollingerBandsSeries) GetName() string {
	return bbs.Name
}

// GetStyle returns the line style.
func (bbs BollingerBandsSeries) GetStyle() Style {
	return bbs.Style
}

// GetYAxis returns which YAxis the series draws on.
func (bbs BollingerBandsSeries) GetYAxis() YAxisType {
	return bbs.YAxis
}

// GetWindowSize returns the window size.
func (bbs BollingerBandsSeries) GetWindowSize(defaults ...int) int {
	if bbs.WindowSize == 0 {
		if len(defaults) > 0 {
			return defaults[0]
		}
		return DefaultMovingAverageWindowSize
	}
	return bbs.WindowSize
}

// GetK returns the K value.
func (bbs BollingerBandsSeries) GetK(defaults ...float64) float64 {
	if bbs.K == 0 {
		if len(defaults) > 0 {
			return defaults[0]
		}
		return 2.0
	}
	return bbs.K
}

// Len returns the number of elements in the series.
func (bbs *BollingerBandsSeries) Len() int {
	return bbs.InnerSeries.Len()
}

// GetBoundedValue gets the bounded value for the series.
func (bbs *BollingerBandsSeries) GetBoundedValue(index int) (x, y1, y2 float64) {
	if bbs.InnerSeries == nil {
		return
	}
	if bbs.valueBuffer == nil || index == 0 {
		bbs.valueBuffer = NewRingBufferWithCapacity(bbs.GetWindowSize())
	}
	if bbs.valueBuffer.Len() >= bbs.GetWindowSize() {
		bbs.valueBuffer.Dequeue()
	}
	px, py := bbs.InnerSeries.GetValue(index)
	bbs.valueBuffer.Enqueue(py)
	x = px

	ay := bbs.getAverage(bbs.valueBuffer)
	std := bbs.getStdDev(bbs.valueBuffer)

	y1 = ay + (bbs.GetK() * std)
	y2 = ay - (bbs.GetK() * std)
	return
}

// Render renders the series.
func (bbs BollingerBandsSeries) Render(r Renderer, canvasBox Box, xrange, yrange Range, defaults Style) {
	s := bbs.Style.WithDefaultsFrom(defaults)

	r.SetStrokeColor(s.GetStrokeColor())
	r.SetStrokeDashArray(s.GetStrokeDashArray())
	r.SetStrokeWidth(s.GetStrokeWidth())
	r.SetFillColor(s.GetFillColor())

	cb := canvasBox.Bottom
	cl := canvasBox.Left

	v0x, v0y1, v0y2 := bbs.GetBoundedValue(0)
	x0 := cl + xrange.Translate(v0x)
	y0 := cb - yrange.Translate(v0y1)

	var vx, vy1, vy2 float64
	var x, y int

	xvalues := make([]float64, bbs.Len())
	xvalues[0] = v0x
	y2values := make([]float64, bbs.Len())
	y2values[0] = v0y2

	r.MoveTo(x0, y0)
	for i := 1; i < bbs.Len(); i++ {
		vx, vy1, vy2 = bbs.GetBoundedValue(i)

		xvalues[i] = vx
		y2values[i] = vy2

		x = cl + xrange.Translate(vx)
		y = cb - yrange.Translate(vy1)
		if i > bbs.GetWindowSize() {
			r.LineTo(x, y)
		} else {
			r.MoveTo(x, y)
		}
	}
	y = cb - yrange.Translate(vy2)
	r.LineTo(x, y)
	for i := bbs.Len() - 1; i >= bbs.GetWindowSize(); i-- {
		vx, vy2 = xvalues[i], y2values[i]
		x = cl + xrange.Translate(vx)
		y = cb - yrange.Translate(vy2)
		r.LineTo(x, y)
	}
	r.Close()
	r.FillStroke()
}

func (bbs BollingerBandsSeries) getAverage(valueBuffer *RingBuffer) float64 {
	var accum float64
	valueBuffer.Each(func(v interface{}) {
		if typed, isTyped := v.(float64); isTyped {
			accum += typed
		}
	})
	return accum / float64(valueBuffer.Len())
}

func (bbs BollingerBandsSeries) getVariance(valueBuffer *RingBuffer) float64 {
	if valueBuffer.Len() == 0 {
		return 0
	}

	var variance float64
	m := bbs.getAverage(valueBuffer)

	valueBuffer.Each(func(v interface{}) {
		if n, isTyped := v.(float64); isTyped {
			variance += (float64(n) - m) * (float64(n) - m)
		}
	})

	return variance / float64(valueBuffer.Len())
}

func (bbs BollingerBandsSeries) getStdDev(valueBuffer *RingBuffer) float64 {
	return math.Pow(bbs.getVariance(valueBuffer), 0.5)
}
