package chart

import (
	"testing"
	"time"

	"github.com/wcharczuk/go-chart/v2/testutil"
)

func TestTimeSeriesGetValue(t *testing.T) {
	// replaced new assertions helper

	ts := TimeSeries{
		Name: "Test",
		XValues: []time.Time{
			time.Now().AddDate(0, 0, -5),
			time.Now().AddDate(0, 0, -4),
			time.Now().AddDate(0, 0, -3),
			time.Now().AddDate(0, 0, -2),
			time.Now().AddDate(0, 0, -1),
		},
		YValues: []float64{
			1.0, 2.0, 3.0, 4.0, 5.0,
		},
	}

	x0, y0 := ts.GetValues(0)
	testutil.AssertNotZero(t, x0)
	testutil.AssertEqual(t, 1.0, y0)
}

func TestTimeSeriesValidate(t *testing.T) {
	// replaced new assertions helper

	cs := TimeSeries{
		Name: "Test Series",
		XValues: []time.Time{
			time.Now().AddDate(0, 0, -5),
			time.Now().AddDate(0, 0, -4),
			time.Now().AddDate(0, 0, -3),
			time.Now().AddDate(0, 0, -2),
			time.Now().AddDate(0, 0, -1),
		},
		YValues: []float64{
			1.0, 2.0, 3.0, 4.0, 5.0,
		},
	}
	testutil.AssertNil(t, cs.Validate())

	cs = TimeSeries{
		Name: "Test Series",
		XValues: []time.Time{
			time.Now().AddDate(0, 0, -5),
			time.Now().AddDate(0, 0, -4),
			time.Now().AddDate(0, 0, -3),
			time.Now().AddDate(0, 0, -2),
			time.Now().AddDate(0, 0, -1),
		},
	}
	testutil.AssertNotNil(t, cs.Validate())

	cs = TimeSeries{
		Name: "Test Series",
		YValues: []float64{
			1.0, 2.0, 3.0, 4.0, 5.0,
		},
	}
	testutil.AssertNotNil(t, cs.Validate())
}
