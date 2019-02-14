package chart

import (
	"testing"

	"github.com/blend/go-sdk/assert"
)

func TestLastValueAnnotationSeries(t *testing.T) {
	assert := assert.New(t)

	series := ContinuousSeries{
		XValues: []float64{1.0, 2.0, 3.0, 4.0, 5.0},
		YValues: []float64{5.0, 3.0, 3.0, 2.0, 1.0},
	}

	lva := LastValueAnnotationSeries(series)
	assert.NotEmpty(lva.Annotations)
	lvaa := lva.Annotations[0]
	assert.Equal(5, lvaa.XValue)
	assert.Equal(1, lvaa.YValue)
}
