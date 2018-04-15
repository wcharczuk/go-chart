package seq

import (
	"testing"

	assert "github.com/blend/go-sdk/assert"
)

func TestRandomRegression(t *testing.T) {
	assert := assert.New(t)

	randomProvider := NewRandom().WithLen(4096).WithMax(256)
	assert.Equal(4096, randomProvider.Len())
	assert.Equal(256, *randomProvider.Max())

	randomSequence := New(randomProvider)
	randomValues := randomSequence.Array()
	assert.Len(4096, randomValues)
	assert.InDelta(128, randomSequence.Average(), 10.0)
}
