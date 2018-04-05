package seq

import (
	"testing"

	assert "github.com/blend/go-sdk/assert"
)

func TestRange(t *testing.T) {
	assert := assert.New(t)

	values := Range(1, 100)
	assert.Len(100, values)
	assert.Equal(1, values[0])
	assert.Equal(100, values[99])
}

func TestRangeWithStep(t *testing.T) {
	assert := assert.New(t)

	values := RangeWithStep(0, 100, 5)
	assert.Equal(100, values[20])
	assert.Len(21, values)
}

func TestRangeReversed(t *testing.T) {
	assert := assert.New(t)

	values := Range(10.0, 1.0)
	assert.Equal(10, len(values))
	assert.Equal(10.0, values[0])
	assert.Equal(1.0, values[9])
}

func TestValuesRegression(t *testing.T) {
	assert := assert.New(t)

	// note; this assumes a 1.0 step is implicitly set in the constructor.
	linearProvider := NewLinear().WithStart(1.0).WithEnd(100.0)
	assert.Equal(1, linearProvider.Start())
	assert.Equal(100, linearProvider.End())
	assert.Equal(100, linearProvider.Len())

	values := Seq{Provider: linearProvider}.Array()
	assert.Len(100, values)
	assert.Equal(1.0, values[0])
	assert.Equal(100, values[99])
}
