package chart

import (
	"testing"

	assert "github.com/blendlabs/go-assert"
)

func TestSequenceFloat64(t *testing.T) {
	assert := assert.New(t)

	asc := Sequence.Float64(1.0, 10.0)
	assert.Len(asc, 10)

	desc := Sequence.Float64(10.0, 1.0)
	assert.Len(desc, 10)
}
