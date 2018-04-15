package chart

import (
	"testing"

	"github.com/blend/go-sdk/assert"
	"github.com/wcharczuk/go-chart/seq"
)

var (
	emaXValues = seq.Range(1.0, 50.0)
	emaYValues = []float64{
		1, 2, 3, 4, 5, 4, 3, 2,
		1, 2, 3, 4, 5, 4, 3, 2,
		1, 2, 3, 4, 5, 4, 3, 2,
		1, 2, 3, 4, 5, 4, 3, 2,
		1, 2, 3, 4, 5, 4, 3, 2,
		1, 2, 3, 4, 5, 4, 3, 2,
		1, 2,
	}
	emaExpected = []float64{
		1,
		1.074074074,
		1.216735254,
		1.422903013,
		1.68787316,
		1.859141815,
		1.943649828,
		1.947823915,
		1.877614736,
		1.886680311,
		1.969148437,
		2.119581886,
		2.33294619,
		2.456431658,
		2.496695979,
		2.459903685,
		2.351762671,
		2.325706177,
		2.375653867,
		2.495975803,
		2.681459077,
		2.779128775,
		2.795489607,
		2.73656445,
		2.607930047,
		2.562898191,
		2.595276103,
		2.699329725,
		2.869749746,
		2.953471987,
		2.956918506,
		2.886035654,
		2.746329309,
		2.691045657,
		2.713931163,
		2.809195522,
		2.971477335,
		3.047664199,
		3.044133518,
		2.966790294,
		2.821102124,
		2.760279745,
		2.778036801,
		2.868552593,
		3.026437586,
		3.098553321,
		3.091253075,
		3.010419514,
		2.86149955,
		2.797684768,
	}
	emaDelta = 0.0001
)

func TestEMASeries(t *testing.T) {
	assert := assert.New(t)

	mockSeries := mockValuesProvider{
		emaXValues,
		emaYValues,
	}
	assert.Equal(50, mockSeries.Len())

	ema := &EMASeries{
		InnerSeries: mockSeries,
		Period:      26,
	}

	sig := ema.GetSigma()
	assert.Equal(2.0/(26.0+1), sig)

	var yvalues []float64
	for x := 0; x < ema.Len(); x++ {
		_, y := ema.GetValues(x)
		yvalues = append(yvalues, y)
	}

	for index, yv := range yvalues {
		assert.InDelta(yv, emaExpected[index], emaDelta)
	}

	lvx, lvy := ema.GetLastValues()
	assert.Equal(50.0, lvx)
	assert.InDelta(lvy, emaExpected[49], emaDelta)
}
