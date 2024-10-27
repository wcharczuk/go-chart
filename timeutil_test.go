package chart

import (
	"testing"
	"time"
)

func TestTimeMinMax(t *testing.T) {
	// empty input
	min, max := TimeMinMax()
	if !min.IsZero() || !max.IsZero() {
		t.Errorf("Expected minimum and maximum to be zero time for empty input, but got min=%s, max=%s", min, max)
	}

	// non-empty input
	times := []time.Time{
		time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
		time.Date(2022, 2, 1, 0, 0, 0, 0, time.UTC),
		time.Date(2022, 3, 1, 0, 0, 0, 0, time.UTC),
	}
	expectedMin := time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)
	expectedMax := time.Date(2022, 3, 1, 0, 0, 0, 0, time.UTC)
	min, max = TimeMinMax(times...)
	if min != expectedMin || max != expectedMax {
		t.Errorf("Expected minimum=%s, maximum=%s for non-empty input, but got min=%s, max=%s", expectedMin, expectedMax, min, max)
	}
}

func TestTimeToFloat64(t *testing.T) {
	// zero time
	tf := TimeToFloat64(time.Time{})
	if tf != 0 {
		t.Errorf("Expected float64 representation of zero time to be 0, but got %f", tf)
	}

	// non-zero time
	tm := time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)
	expectedTF := float64(tm.UnixNano())
	tf = TimeToFloat64(tm)
	if tf != expectedTF {
		t.Errorf("Expected float64 representation of time %s to be %f, but got %f", tm, expectedTF, tf)
	}
}

func TestTimeFromFloat64(t *testing.T) {
	// zero float64
	expectedT := time.Time{}
	actualT := TimeFromFloat64(0)
	if actualT != expectedT {
		t.Errorf("Expected time from float64 representation of 0 to be zero time, but got %s", actualT)
	}

	// non-zero float64 represent nanoseconds
	expectedT = time.Date(2022, 1, 1, 0, 0, 0, 123456789, time.Local)
	nanosecondsFloat := float64(expectedT.UnixNano())
	actualT = TimeFromFloat64(nanosecondsFloat)
	if actualT.Equal(expectedT) {
		t.Errorf("Expected time from float64 representation %f to be %s, but got %s", nanosecondsFloat, expectedT, actualT)
	}
}
