package chart

import (
	"testing"

	"github.com/wcharczuk/go-chart/v2/testutil"
)

func TestValuesValues(t *testing.T) {
	// replaced new assertions helper

	vs := []Value{
		{Value: 10, Label: "Blue"},
		{Value: 9, Label: "Green"},
		{Value: 8, Label: "Gray"},
		{Value: 7, Label: "Orange"},
		{Value: 6, Label: "HEANG"},
		{Value: 5, Label: "??"},
		{Value: 2, Label: "!!"},
	}

	values := Values(vs).Values()
	testutil.AssertLen(t, values, 7)
	testutil.AssertEqual(t, 10, values[0])
	testutil.AssertEqual(t, 9, values[1])
	testutil.AssertEqual(t, 8, values[2])
	testutil.AssertEqual(t, 7, values[3])
	testutil.AssertEqual(t, 6, values[4])
	testutil.AssertEqual(t, 5, values[5])
	testutil.AssertEqual(t, 2, values[6])
}

func TestValuesValuesNormalized(t *testing.T) {
	// replaced new assertions helper

	vs := []Value{
		{Value: 10, Label: "Blue"},
		{Value: 9, Label: "Green"},
		{Value: 8, Label: "Gray"},
		{Value: 7, Label: "Orange"},
		{Value: 6, Label: "HEANG"},
		{Value: 5, Label: "??"},
		{Value: 2, Label: "!!"},
	}

	values := Values(vs).ValuesNormalized()
	testutil.AssertLen(t, values, 7)
	testutil.AssertEqual(t, 0.2127, values[0])
	testutil.AssertEqual(t, 0.0425, values[6])
}

func TestValuesNormalize(t *testing.T) {
	// replaced new assertions helper

	vs := []Value{
		{Value: 10, Label: "Blue"},
		{Value: 9, Label: "Green"},
		{Value: 8, Label: "Gray"},
		{Value: 7, Label: "Orange"},
		{Value: 6, Label: "HEANG"},
		{Value: 5, Label: "??"},
		{Value: 2, Label: "!!"},
	}

	values := Values(vs).Normalize()
	testutil.AssertLen(t, values, 7)
	testutil.AssertEqual(t, 0.2127, values[0].Value)
	testutil.AssertEqual(t, 0.0425, values[6].Value)
}
