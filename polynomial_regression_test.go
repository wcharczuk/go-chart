package chart

import (
	"testing"

	assert "github.com/blend/go-sdk/assert"
	"github.com/wcharczuk/go-chart/matrix"
)

func TestPolynomialRegression(t *testing.T) {
	assert := assert.New(t)

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
		assert.InDelta(float64(i*i), y, matrix.DefaultEpsilon)
	}
}
