package chart

import (
	"testing"

	"github.com/wcharczuk/go-chart/v2/testutil"
)

func TestGenerateContinuousTicks(t *testing.T) {
	// replaced new assertions helper

	f, err := GetDefaultFont()
	testutil.AssertNil(t, err)

	r, err := PNG(1024, 1024)
	testutil.AssertNil(t, err)
	r.SetFont(f)

	ra := &ContinuousRange{
		Min:    0.0,
		Max:    10.0,
		Domain: 256,
	}

	vf := FloatValueFormatter

	ticks := GenerateContinuousTicks(r, ra, false, Style{}, vf)
	testutil.AssertNotEmpty(t, ticks)
	testutil.AssertLen(t, ticks, 11)
	testutil.AssertEqual(t, 0.0, ticks[0].Value)
	testutil.AssertEqual(t, 10, ticks[len(ticks)-1].Value)
}

func TestGenerateContinuousTicksDescending(t *testing.T) {
	// replaced new assertions helper

	f, err := GetDefaultFont()
	testutil.AssertNil(t, err)

	r, err := PNG(1024, 1024)
	testutil.AssertNil(t, err)
	r.SetFont(f)

	ra := &ContinuousRange{
		Min:        0.0,
		Max:        10.0,
		Domain:     256,
		Descending: true,
	}

	vf := FloatValueFormatter

	ticks := GenerateContinuousTicks(r, ra, false, Style{}, vf)
	testutil.AssertNotEmpty(t, ticks)
	testutil.AssertLen(t, ticks, 11)
	testutil.AssertEqual(t, 10.0, ticks[0].Value)
	testutil.AssertEqual(t, 9.0, ticks[1].Value)
	testutil.AssertEqual(t, 1.0, ticks[len(ticks)-2].Value)
	testutil.AssertEqual(t, 0.0, ticks[len(ticks)-1].Value)
}
