package matrix

import (
	"testing"

	assert "github.com/blend/go-sdk/assert"
)

func TestPoly(t *testing.T) {
	assert := assert.New(t)
	var xGiven = []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var yGiven = []float64{1, 6, 17, 34, 57, 86, 121, 162, 209, 262, 321}
	var degree = 2

	c, err := Poly(xGiven, yGiven, degree)
	assert.Nil(err)
	assert.Len(3, c)

	assert.InDelta(c[0], 0.999999999, DefaultEpsilon)
	assert.InDelta(c[1], 2, DefaultEpsilon)
	assert.InDelta(c[2], 3, DefaultEpsilon)
}
