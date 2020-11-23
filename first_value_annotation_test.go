package chart

import (
	"testing"

	"github.com/wcharczuk/go-chart/v2/testutil"
)

func TestFirstValueAnnotation(t *testing.T) {
	// replaced new assertions helper

	series := ContinuousSeries{
		XValues: []float64{1.0, 2.0, 3.0, 4.0, 5.0},
		YValues: []float64{5.0, 3.0, 3.0, 2.0, 1.0},
	}

	fva := FirstValueAnnotation(series)
	testutil.AssertNotEmpty(t, fva.Annotations)
	fvaa := fva.Annotations[0]
	testutil.AssertEqual(t, 1, fvaa.XValue)
	testutil.AssertEqual(t, 5, fvaa.YValue)
}
