package chart

import (
	"testing"

	"github.com/wcharczuk/go-chart/v2/testutil"
)

func Test_ValueSequence_Normalize(t *testing.T) {
	// replaced new assertions helper

	normalized := ValueSequence(1, 2, 3, 4, 5).Normalize().Values()

	testutil.AssertNotEmpty(t, normalized)
	testutil.AssertLen(t, normalized, 5)
	testutil.AssertEqual(t, 0, normalized[0])
	testutil.AssertEqual(t, 0.25, normalized[1])
	testutil.AssertEqual(t, 1, normalized[4])
}
