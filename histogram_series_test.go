package chart

import (
	"testing"

	"github.com/wcharczuk/go-chart/v2/testutil"
)

func TestHistogramSeries(t *testing.T) {
	// replaced new assertions helper

	cs := ContinuousSeries{
		Name:    "Test Series",
		XValues: LinearRange(1.0, 20.0),
		YValues: LinearRange(10.0, -10.0),
	}

	hs := HistogramSeries{
		InnerSeries: cs,
	}

	for x := 0; x < hs.Len(); x++ {
		csx, csy := cs.GetValues(0)
		hsx, hsy1, hsy2 := hs.GetBoundedValues(0)
		testutil.AssertEqual(t, csx, hsx)
		testutil.AssertTrue(t, hsy1 > 0)
		testutil.AssertTrue(t, hsy2 <= 0)
		testutil.AssertTrue(t, csy < 0 || (csy > 0 && csy == hsy1))
		testutil.AssertTrue(t, csy > 0 || (csy < 0 && csy == hsy2))
	}
}
