package chart

import (
	"fmt"
	"math"
	"testing"

	"github.com/wcharczuk/go-chart/v2/testutil"
)

func TestBollingerBandSeries(t *testing.T) {
	// replaced new assertions helper

	s1 := mockValuesProvider{
		X: LinearRange(1.0, 100.0),
		Y: RandomValuesWithMax(100, 1024),
	}

	bbs := &BollingerBandsSeries{
		InnerSeries: s1,
	}

	xvalues := make([]float64, 100)
	y1values := make([]float64, 100)
	y2values := make([]float64, 100)

	for x := 0; x < 100; x++ {
		xvalues[x], y1values[x], y2values[x] = bbs.GetBoundedValues(x)
	}

	for x := bbs.GetPeriod(); x < 100; x++ {
		testutil.AssertTrue(t, y1values[x] > y2values[x], fmt.Sprintf("%v vs. %v", y1values[x], y2values[x]))
	}
}

func TestBollingerBandLastValue(t *testing.T) {
	// replaced new assertions helper

	s1 := mockValuesProvider{
		X: LinearRange(1.0, 100.0),
		Y: LinearRange(1.0, 100.0),
	}

	bbs := &BollingerBandsSeries{
		InnerSeries: s1,
	}

	x, y1, y2 := bbs.GetBoundedLastValues()
	testutil.AssertEqual(t, 100.0, x)
	testutil.AssertEqual(t, 101, math.Floor(y1))
	testutil.AssertEqual(t, 83, math.Floor(y2))
}
