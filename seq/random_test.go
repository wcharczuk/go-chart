package seq

import (
	"testing"

	assert "github.com/blendlabs/go-assert"
)

func TestRandomRegression(t *testing.T) {
	assert := assert.New(t)

	randomProvider := NewRandom().WithLen(100).WithAverage(256)
	assert.Equal(100, randomProvider.Len())
	assert.Equal(256, *randomProvider.Average())

	randomValues := New(randomProvider).Array()
	assert.Len(randomValues, 100)
}
