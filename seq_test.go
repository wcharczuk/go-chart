package chart

import (
	"testing"

	"github.com/wcharczuk/go-chart/v2/testutil"
)

func Test_Seq_Each(t *testing.T) {
	values := Seq[int]{Array(1, 2, 3, 4)}
	values.Each(func(i int, v int) {
		testutil.AssertEqual(t, i, v-1)
	})
}

func Test_Seq_Map(t *testing.T) {
	values := Seq[int]{Array(1, 2, 3, 4)}
	mapped := values.Map(func(i int, v int) int {
		testutil.AssertEqual(t, i, v-1)
		return v * 2
	})
	testutil.AssertEqual(t, 4, mapped.Len())
}

func Test_Seq_FoldLeft(t *testing.T) {
	values := Seq[int]{Array(1, 2, 3, 4)}
	ten := values.FoldLeft(func(_ int, vp, v int) int {
		return vp + v
	})
	testutil.AssertEqual(t, 10, ten)

	orderTest := Seq[int]{Array(10, 3, 2, 1)}
	four := orderTest.FoldLeft(func(_ int, vp, v int) int {
		return vp - v
	})
	testutil.AssertEqual(t, 4, four)
}

func Test_Seq_FoldRight(t *testing.T) {
	values := Seq[int]{Array(1, 2, 3, 4)}
	ten := values.FoldRight(func(_ int, vp, v int) int {
		return vp + v
	})
	testutil.AssertEqual(t, 10, ten)

	orderTest := Seq[int]{Array(10, 3, 2, 1)}
	notFour := orderTest.FoldRight(func(_ int, vp, v int) int {
		return vp - v
	})
	testutil.AssertEqual(t, -14, notFour)
}

func Test_Seq_Sum(t *testing.T) {
	values := Seq[int]{Array(1, 2, 3, 4)}
	testutil.AssertEqual(t, 10, values.Sum())
}

func Test_Seq_Average(t *testing.T) {
	values := Seq[int]{Array(1, 2, 3, 4)}
	testutil.AssertEqual(t, 2.5, values.Average())

	valuesOdd := Seq[int]{Array(1, 2, 3, 4, 5)}
	testutil.AssertEqual(t, 3, valuesOdd.Average())
}

func Test_Seq_Variance(t *testing.T) {
	values := Seq[int]{Array(1, 2, 3, 4, 5)}
	testutil.AssertEqual(t, 2, values.Variance())
}

func Test_LinearRange(t *testing.T) {
	// replaced new assertions helper

	values := LinearRange(1, 100)
	testutil.AssertLen(t, values, 100)
	testutil.AssertEqual(t, 1, values[0])
	testutil.AssertEqual(t, 100, values[99])
}

func Test_LinearRange_WithStep(t *testing.T) {
	// replaced new assertions helper

	values := LinearRangeWithStep(0, 100, 5)
	testutil.AssertEqual(t, 100, values[20])
	testutil.AssertLen(t, values, 21)
}

func Test_LinearRange_reversed(t *testing.T) {
	// replaced new assertions helper

	values := LinearRange(10.0, 1.0)
	testutil.AssertEqual(t, 10, len(values))
	testutil.AssertEqual(t, 10.0, values[0])
	testutil.AssertEqual(t, 1.0, values[9])
}

func Test_LinearSequence_Regression(t *testing.T) {
	// replaced new assertions helper

	// note; this assumes a 1.0 step is implicitly set in the constructor.
	linearProvider := NewLinearSequence().WithStart(1.0).WithEnd(100.0)
	testutil.AssertEqual(t, 1, linearProvider.Start())
	testutil.AssertEqual(t, 100, linearProvider.End())
	testutil.AssertEqual(t, 100, linearProvider.Len())

	values := Seq[float64]{linearProvider}.Values()
	testutil.AssertLen(t, values, 100)
	testutil.AssertEqual(t, 1.0, values[0])
	testutil.AssertEqual(t, 100, values[99])
}
