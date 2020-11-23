package chart

import (
	"testing"

	"github.com/wcharczuk/go-chart/v2/testutil"
)

func TestPercentageDifferenceSeries(t *testing.T) {
	// replaced new assertions helper

	cs := ContinuousSeries{
		XValues: LinearRange(1.0, 10.0),
		YValues: LinearRange(1.0, 10.0),
	}

	pcs := PercentChangeSeries{
		Name:        "Test Series",
		InnerSeries: cs,
	}

	testutil.AssertEqual(t, "Test Series", pcs.GetName())
	testutil.AssertEqual(t, 10, pcs.Len())
	x0, y0 := pcs.GetValues(0)
	testutil.AssertEqual(t, 1.0, x0)
	testutil.AssertEqual(t, 0, y0)

	xn, yn := pcs.GetValues(9)
	testutil.AssertEqual(t, 10.0, xn)
	testutil.AssertEqual(t, 9.0, yn)

	xn, yn = pcs.GetLastValues()
	testutil.AssertEqual(t, 10.0, xn)
	testutil.AssertEqual(t, 9.0, yn)
}
