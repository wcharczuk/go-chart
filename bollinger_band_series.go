package chart

import "math"

// BollingerBandsSeries draws bollinger bands for an inner series.
// Bollinger bands are defined by two lines, one at SMA+k*stddev, one at SMA-k*stdev.
type BollingerBandsSeries struct {
	Name  string
	Style Style
	YAxis YAxisType

	Period      int
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

// GetPeriod returns the window size.
func (bbs BollingerBandsSeries) GetPeriod() int {
	if bbs.Period == 0 {
		return DefaultSimpleMovingAveragePeriod
	}
	return bbs.Period
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
		bbs.valueBuffer = NewRingBufferWithCapacity(bbs.GetPeriod())
	}
	if bbs.valueBuffer.Len() >= bbs.GetPeriod() {
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

// GetBoundedLastValue returns the last bounded value for the series.
func (bbs *BollingerBandsSeries) GetBoundedLastValue() (x, y1, y2 float64) {
	if bbs.InnerSeries == nil {
		return
	}
	period := bbs.GetPeriod()
	seriesLength := bbs.InnerSeries.Len()
	startAt := seriesLength - period
	if startAt < 0 {
		startAt = 0
	}

	vb := NewRingBufferWithCapacity(period)
	for index := startAt; index < seriesLength; index++ {
		xn, yn := bbs.InnerSeries.GetValue(index)
		vb.Enqueue(yn)
		x = xn
	}

	ay := bbs.getAverage(vb)
	std := bbs.getStdDev(vb)

	y1 = ay + (bbs.GetK() * std)
	y2 = ay - (bbs.GetK() * std)

	return
}

// Render renders the series.
func (bbs *BollingerBandsSeries) Render(r Renderer, canvasBox Box, xrange, yrange Range, defaults Style) {
	s := bbs.Style.InheritFrom(defaults.InheritFrom(Style{
		StrokeWidth: 1.0,
		StrokeColor: DefaultAxisColor.WithAlpha(64),
		FillColor:   DefaultAxisColor.WithAlpha(32),
	}))

	Draw.BoundedSeries(r, canvasBox, xrange, yrange, s, bbs, bbs.GetPeriod())
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
