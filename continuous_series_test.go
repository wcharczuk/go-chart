package chart

import (
	"fmt"
	"testing"

	"github.com/wcharczuk/go-chart/v2/testutil"
)

func TestContinuousSeries(t *testing.T) {
	// replaced new assertions helper

	cs := ContinuousSeries{
		Name:    "Test Series",
		XValues: LinearRange(1.0, 10.0),
		YValues: LinearRange(1.0, 10.0),
	}

	testutil.AssertEqual(t, "Test Series", cs.GetName())
	testutil.AssertEqual(t, 10, cs.Len())
	x0, y0 := cs.GetValues(0)
	testutil.AssertEqual(t, 1.0, x0)
	testutil.AssertEqual(t, 1.0, y0)

	xn, yn := cs.GetValues(9)
	testutil.AssertEqual(t, 10.0, xn)
	testutil.AssertEqual(t, 10.0, yn)

	xn, yn = cs.GetLastValues()
	testutil.AssertEqual(t, 10.0, xn)
	testutil.AssertEqual(t, 10.0, yn)
}

func TestContinuousSeriesValueFormatter(t *testing.T) {
	// replaced new assertions helper

	cs := ContinuousSeries{
		XValueFormatter: func(v interface{}) string {
			return fmt.Sprintf("%f foo", v)
		},
		YValueFormatter: func(v interface{}) string {
			return fmt.Sprintf("%f bar", v)
		},
	}

	xf, yf := cs.GetValueFormatters()
	testutil.AssertEqual(t, "0.100000 foo", xf(0.1))
	testutil.AssertEqual(t, "0.100000 bar", yf(0.1))
}

func TestContinuousSeriesValidate(t *testing.T) {
	// replaced new assertions helper

	cs := ContinuousSeries{
		Name:    "Test Series",
		XValues: LinearRange(1.0, 10.0),
		YValues: LinearRange(1.0, 10.0),
	}
	testutil.AssertNil(t, cs.Validate())

	cs = ContinuousSeries{
		Name:    "Test Series",
		XValues: LinearRange(1.0, 10.0),
	}
	testutil.AssertNotNil(t, cs.Validate())

	cs = ContinuousSeries{
		Name:    "Test Series",
		YValues: LinearRange(1.0, 10.0),
	}
	testutil.AssertNotNil(t, cs.Validate())
}
