package chart

import (
	"testing"

	"github.com/wcharczuk/go-chart/v2/testutil"
)

func TestSeqEach(t *testing.T) {
	// replaced new assertions helper

	values := Seq{NewArray(1, 2, 3, 4)}
	values.Each(func(i int, v float64) {
		testutil.AssertEqual(t, i, v-1)
	})
}

func TestSeqMap(t *testing.T) {
	// replaced new assertions helper

	values := Seq{NewArray(1, 2, 3, 4)}
	mapped := values.Map(func(i int, v float64) float64 {
		testutil.AssertEqual(t, i, v-1)
		return v * 2
	})
	testutil.AssertEqual(t, 4, mapped.Len())
}

func TestSeqFoldLeft(t *testing.T) {
	// replaced new assertions helper

	values := Seq{NewArray(1, 2, 3, 4)}
	ten := values.FoldLeft(func(_ int, vp, v float64) float64 {
		return vp + v
	})
	testutil.AssertEqual(t, 10, ten)

	orderTest := Seq{NewArray(10, 3, 2, 1)}
	four := orderTest.FoldLeft(func(_ int, vp, v float64) float64 {
		return vp - v
	})
	testutil.AssertEqual(t, 4, four)
}

func TestSeqFoldRight(t *testing.T) {
	// replaced new assertions helper

	values := Seq{NewArray(1, 2, 3, 4)}
	ten := values.FoldRight(func(_ int, vp, v float64) float64 {
		return vp + v
	})
	testutil.AssertEqual(t, 10, ten)

	orderTest := Seq{NewArray(10, 3, 2, 1)}
	notFour := orderTest.FoldRight(func(_ int, vp, v float64) float64 {
		return vp - v
	})
	testutil.AssertEqual(t, -14, notFour)
}

func TestSeqSum(t *testing.T) {
	// replaced new assertions helper

	values := Seq{NewArray(1, 2, 3, 4)}
	testutil.AssertEqual(t, 10, values.Sum())
}

func TestSeqAverage(t *testing.T) {
	// replaced new assertions helper

	values := Seq{NewArray(1, 2, 3, 4)}
	testutil.AssertEqual(t, 2.5, values.Average())

	valuesOdd := Seq{NewArray(1, 2, 3, 4, 5)}
	testutil.AssertEqual(t, 3, valuesOdd.Average())
}

func TestSequenceVariance(t *testing.T) {
	// replaced new assertions helper

	values := Seq{NewArray(1, 2, 3, 4, 5)}
	testutil.AssertEqual(t, 2, values.Variance())
}

func TestSequenceNormalize(t *testing.T) {
	// replaced new assertions helper

	normalized := ValueSequence(1, 2, 3, 4, 5).Normalize().Values()

	testutil.AssertNotEmpty(t, normalized)
	testutil.AssertLen(t, normalized, 5)
	testutil.AssertEqual(t, 0, normalized[0])
	testutil.AssertEqual(t, 0.25, normalized[1])
	testutil.AssertEqual(t, 1, normalized[4])
}

func TestLinearRange(t *testing.T) {
	// replaced new assertions helper

	values := LinearRange(1, 100)
	testutil.AssertLen(t, values, 100)
	testutil.AssertEqual(t, 1, values[0])
	testutil.AssertEqual(t, 100, values[99])
}

func TestLinearRangeWithStep(t *testing.T) {
	// replaced new assertions helper

	values := LinearRangeWithStep(0, 100, 5)
	testutil.AssertEqual(t, 100, values[20])
	testutil.AssertLen(t, values, 21)
}

func TestLinearRangeReversed(t *testing.T) {
	// replaced new assertions helper

	values := LinearRange(10.0, 1.0)
	testutil.AssertEqual(t, 10, len(values))
	testutil.AssertEqual(t, 10.0, values[0])
	testutil.AssertEqual(t, 1.0, values[9])
}

func TestLinearSequenceRegression(t *testing.T) {
	// replaced new assertions helper

	// note; this assumes a 1.0 step is implicitly set in the constructor.
	linearProvider := NewLinearSequence().WithStart(1.0).WithEnd(100.0)
	testutil.AssertEqual(t, 1, linearProvider.Start())
	testutil.AssertEqual(t, 100, linearProvider.End())
	testutil.AssertEqual(t, 100, linearProvider.Len())

	values := Seq{linearProvider}.Values()
	testutil.AssertLen(t, values, 100)
	testutil.AssertEqual(t, 1.0, values[0])
	testutil.AssertEqual(t, 100, values[99])
}
