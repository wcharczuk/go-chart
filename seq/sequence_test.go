package seq

import (
	"testing"

	assert "github.com/blend/go-sdk/assert"
)

func TestSequenceEach(t *testing.T) {
	assert := assert.New(t)

	values := Seq{NewArray(1, 2, 3, 4)}
	values.Each(func(i int, v float64) {
		assert.Equal(i, v-1)
	})
}

func TestSequenceMap(t *testing.T) {
	assert := assert.New(t)

	values := Seq{NewArray(1, 2, 3, 4)}
	mapped := values.Map(func(i int, v float64) float64 {
		assert.Equal(i, v-1)
		return v * 2
	})
	assert.Equal(4, mapped.Len())
}

func TestSequenceFoldLeft(t *testing.T) {
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

func TestSequenceFoldRight(t *testing.T) {
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

func TestSequenceSum(t *testing.T) {
	assert := assert.New(t)

	values := Seq{NewArray(1, 2, 3, 4)}
	assert.Equal(10, values.Sum())
}

func TestSequenceAverage(t *testing.T) {
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

	normalized := Values(1, 2, 3, 4, 5).Normalize().Array()

	assert.NotEmpty(normalized)
	assert.Len(5, normalized)
	assert.Equal(0, normalized[0])
	assert.Equal(0.25, normalized[1])
	assert.Equal(1, normalized[4])
}
