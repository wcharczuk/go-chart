package chart

import (
	"testing"

	"github.com/wcharczuk/go-chart/v2/matrix"
	"github.com/wcharczuk/go-chart/v2/testutil"
)

func TestPolynomialRegression(t *testing.T) {
	// replaced new assertions helper

	var xv []float64
	var yv []float64

	for i := 0; i < 100; i++ {
		xv = append(xv, float64(i))
		yv = append(yv, float64(i*i))
	}

	values := ContinuousSeries{
		XValues: xv,
		YValues: yv,
	}

	poly := &PolynomialRegressionSeries{
		InnerSeries: values,
		Degree:      2,
	}

	for i := 0; i < 100; i++ {
		_, y := poly.GetValues(i)
		testutil.AssertInDelta(t, float64(i*i), y, matrix.DefaultEpsilon)
	}
}
