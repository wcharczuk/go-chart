package chart

import (
	"testing"

	"github.com/wcharczuk/go-chart/v2/testutil"
)

func TestGenerateGridLines(t *testing.T) {
	// replaced new assertions helper

	ticks := []Tick{
		{Value: 1.0, Label: "1.0"},
		{Value: 2.0, Label: "2.0"},
		{Value: 3.0, Label: "3.0"},
		{Value: 4.0, Label: "4.0"},
	}

	gl := GenerateGridLines(ticks, Style{}, Style{})
	testutil.AssertLen(t, gl, 2)

	testutil.AssertEqual(t, 2.0, gl[0].Value)
	testutil.AssertEqual(t, 3.0, gl[1].Value)
}
