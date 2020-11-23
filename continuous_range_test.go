package chart

import (
	"testing"

	"github.com/wcharczuk/go-chart/v2/testutil"
)

func TestRangeTranslate(t *testing.T) {
	// replaced new assertions helper
	values := []float64{1.0, 2.0, 2.5, 2.7, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0}
	r := ContinuousRange{Domain: 1000}
	r.Min, r.Max = MinMax(values...)

	// delta = ~7.0
	// value = ~5.0
	// domain = ~1000
	// 5/8 * 1000 ~=
	testutil.AssertEqual(t, 0, r.Translate(1.0))
	testutil.AssertEqual(t, 1000, r.Translate(8.0))
	testutil.AssertEqual(t, 572, r.Translate(5.0))
}
