package chart

import (
	"bytes"
	"testing"

	"github.com/wcharczuk/go-chart/v2/testutil"
)

func TestDonutChart(t *testing.T) {
	// replaced new assertions helper

	pie := DonutChart{
		Canvas: Style{
			FillColor: ColorLightGray,
		},
		Values: []Value{
			{Value: 10, Label: "Blue"},
			{Value: 9, Label: "Green"},
			{Value: 8, Label: "Gray"},
			{Value: 7, Label: "Orange"},
			{Value: 6, Label: "HEANG"},
			{Value: 5, Label: "??"},
			{Value: 2, Label: "!!"},
		},
	}

	b := bytes.NewBuffer([]byte{})
	pie.Render(PNG, b)
	testutil.AssertNotZero(t, b.Len())
}

func TestDonutChartDropsZeroValues(t *testing.T) {
	// replaced new assertions helper

	pie := DonutChart{
		Canvas: Style{
			FillColor: ColorLightGray,
		},
		Values: []Value{
			{Value: 5, Label: "Blue"},
			{Value: 5, Label: "Green"},
			{Value: 0, Label: "Gray"},
		},
	}

	b := bytes.NewBuffer([]byte{})
	err := pie.Render(PNG, b)
	testutil.AssertNil(t, err)
}

func TestDonutChartAllZeroValues(t *testing.T) {
	// replaced new assertions helper

	pie := DonutChart{
		Canvas: Style{
			FillColor: ColorLightGray,
		},
		Values: []Value{
			{Value: 0, Label: "Blue"},
			{Value: 0, Label: "Green"},
			{Value: 0, Label: "Gray"},
		},
	}

	b := bytes.NewBuffer([]byte{})
	err := pie.Render(PNG, b)
	testutil.AssertNotNil(t, err)
}
