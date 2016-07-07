package chart

import (
	"testing"

	"github.com/blendlabs/go-assert"
)

func TestRangeTranslate(t *testing.T) {
	assert := assert.New(t)
	values := []float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0}
	r := Range{Domain: 1000}
	r.Min, r.Max = MinAndMax(values...)
	assert.Equal(428, r.Translate(5.0))
}
