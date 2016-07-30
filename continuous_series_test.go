package chart

import (
	"testing"

	assert "github.com/blendlabs/go-assert"
)

func TestContinuousSeries(t *testing.T) {
	assert := assert.New(t)

	cs := ContinuousSeries{
		Name:    "Test Series",
		XValues: Sequence.Float64(1.0, 10.0),
		YValues: Sequence.Float64(1.0, 10.0),
	}

	assert.Equal("Test Series", cs.GetName())
	assert.Equal(10, cs.Len())
	x0, y0 := cs.GetValue(0)
	assert.Equal(1.0, x0)
	assert.Equal(1.0, y0)

	xn, yn := cs.GetValue(9)
	assert.Equal(10.0, xn)
	assert.Equal(10.0, yn)

	xn, yn = cs.GetLastValue()
	assert.Equal(10.0, xn)
	assert.Equal(10.0, yn)
}
