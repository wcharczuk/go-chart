package matrix

import (
	"testing"

	"github.com/wcharczuk/go-chart/v2/testutil"
)

func TestPoly(t *testing.T) {
	// replaced new assertions helper
	var xGiven = []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var yGiven = []float64{1, 6, 17, 34, 57, 86, 121, 162, 209, 262, 321}
	var degree = 2

	c, err := Poly(xGiven, yGiven, degree)
	testutil.AssertNil(t, err)
	testutil.AssertLen(t, c, 3)

	testutil.AssertInDelta(t, c[0], 0.999999999, DefaultEpsilon)
	testutil.AssertInDelta(t, c[1], 2, DefaultEpsilon)
	testutil.AssertInDelta(t, c[2], 3, DefaultEpsilon)
}
