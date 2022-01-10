package chart

import (
	"testing"

	"github.com/wcharczuk/go-chart/v2/testutil"
)

func Test_LinearRange(t *testing.T) {
	// replaced new assertions helper

	values := LinearRange(1, 100)
	testutil.AssertLen(t, values, 100)
	testutil.AssertEqual(t, 1, values[0])
	testutil.AssertEqual(t, 100, values[99])
}

func Test_LinearRange_WithStep(t *testing.T) {
	// replaced new assertions helper

	values := LinearRangeWithStep(0, 100, 5)
	testutil.AssertEqual(t, 100, values[20])
	testutil.AssertLen(t, values, 21)
}

func Test_LinearRange_reversed(t *testing.T) {
	// replaced new assertions helper

	values := LinearRange(10.0, 1.0)
	testutil.AssertEqual(t, 10, len(values))
	testutil.AssertEqual(t, 10.0, values[0])
	testutil.AssertEqual(t, 1.0, values[9])
}

func Test_LinearSequence_Regression(t *testing.T) {
	// replaced new assertions helper

	// note; this assumes a 1.0 step is implicitly set in the constructor.
	linearProvider := NewLinearSequence().WithStart(1.0).WithEnd(100.0)
	testutil.AssertEqual(t, 1, linearProvider.Start())
	testutil.AssertEqual(t, 100, linearProvider.End())
	testutil.AssertEqual(t, 100, linearProvider.Len())

	values := Seq[float64]{linearProvider}.Values()
	testutil.AssertLen(t, values, 100)
	testutil.AssertEqual(t, 1.0, values[0])
	testutil.AssertEqual(t, 100, values[99])
}
