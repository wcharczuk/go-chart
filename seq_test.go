package chart

import (
	"testing"

	assert "github.com/blend/go-sdk/assert"
)

func TestSeqEach(t *testing.T) {
	assert := assert.New(t)

	values := Seq{NewArray(1, 2, 3, 4)}
	values.Each(func(i int, v float64) {
		assert.Equal(i, v-1)
	})
}

func TestSeqMap(t *testing.T) {
	assert := assert.New(t)

	values := Seq{NewArray(1, 2, 3, 4)}
	mapped := values.Map(func(i int, v float64) float64 {
		assert.Equal(i, v-1)
		return v * 2
	})
	assert.Equal(4, mapped.Len())
}

func TestSeqFoldLeft(t *testing.T) {
	assert := assert.New(t)

	values := Seq{NewArray(1, 2, 3, 4)}
	ten := values.FoldLeft(func(_ int, vp, v float64) float64 {
		return vp + v
	})
	assert.Equal(10, ten)

	orderTest := Seq{NewArray(10, 3, 2, 1)}
	four := orderTest.FoldLeft(func(_ int, vp, v float64) float64 {
		return vp - v
	})
	assert.Equal(4, four)
}

func TestSeqFoldRight(t *testing.T) {
	assert := assert.New(t)

	values := Seq{NewArray(1, 2, 3, 4)}
	ten := values.FoldRight(func(_ int, vp, v float64) float64 {
		return vp + v
	})
	assert.Equal(10, ten)

	orderTest := Seq{NewArray(10, 3, 2, 1)}
	notFour := orderTest.FoldRight(func(_ int, vp, v float64) float64 {
		return vp - v
	})
	assert.Equal(-14, notFour)
}

func TestSeqSum(t *testing.T) {
	assert := assert.New(t)

	values := Seq{NewArray(1, 2, 3, 4)}
	assert.Equal(10, values.Sum())
}

func TestSeqAverage(t *testing.T) {
	assert := assert.New(t)

	values := Seq{NewArray(1, 2, 3, 4)}
	assert.Equal(2.5, values.Average())

	valuesOdd := Seq{NewArray(1, 2, 3, 4, 5)}
	assert.Equal(3, valuesOdd.Average())
}

func TestSequenceVariance(t *testing.T) {
	assert := assert.New(t)

	values := Seq{NewArray(1, 2, 3, 4, 5)}
	assert.Equal(2, values.Variance())
}

func TestSequenceNormalize(t *testing.T) {
	assert := assert.New(t)

	normalized := ValueSequence(1, 2, 3, 4, 5).Normalize().Values()

	assert.NotEmpty(normalized)
	assert.Len(normalized, 5)
	assert.Equal(0, normalized[0])
	assert.Equal(0.25, normalized[1])
	assert.Equal(1, normalized[4])
}

func TestLinearRange(t *testing.T) {
	assert := assert.New(t)

	values := LinearRange(1, 100)
	assert.Len(values, 100)
	assert.Equal(1, values[0])
	assert.Equal(100, values[99])
}

func TestLinearRangeWithStep(t *testing.T) {
	assert := assert.New(t)

	values := LinearRangeWithStep(0, 100, 5)
	assert.Equal(100, values[20])
	assert.Len(values, 21)
}

func TestLinearRangeReversed(t *testing.T) {
	assert := assert.New(t)

	values := LinearRange(10.0, 1.0)
	assert.Equal(10, len(values))
	assert.Equal(10.0, values[0])
	assert.Equal(1.0, values[9])
}

func TestLinearSequenceRegression(t *testing.T) {
	assert := assert.New(t)

	// note; this assumes a 1.0 step is implicitly set in the constructor.
	linearProvider := NewLinearSequence().WithStart(1.0).WithEnd(100.0)
	assert.Equal(1, linearProvider.Start())
	assert.Equal(100, linearProvider.End())
	assert.Equal(100, linearProvider.Len())

	values := Seq{linearProvider}.Values()
	assert.Len(values, 100)
	assert.Equal(1.0, values[0])
	assert.Equal(100, values[99])
}
