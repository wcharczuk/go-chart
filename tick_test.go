package chart

import (
	"testing"

	"github.com/blendlabs/go-assert"
)

func TestGenerateTicksWithStep(t *testing.T) {
	assert := assert.New(t)

	ticks := GenerateTicksWithStep(Range{Min: 1.0, Max: 10.0, Domain: 100}, 1.0, FloatValueFormatter)
	assert.Len(ticks, 10)
}
