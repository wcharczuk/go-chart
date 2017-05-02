package sequence

import (
	"testing"

	assert "github.com/blendlabs/go-assert"
)

func TestValues(t *testing.T) {
	assert := assert.New(t)

	values := Values(1, 100)
	assert.Len(values, 100)
	assert.Equal(1, values[0])
	assert.Equal(100, values[99])
}

func TestValuesWithStep(t *testing.T) {
	assert := assert.New(t)

	values := ValuesWithStep(0, 100, 5)
	assert.Equal(100, values[20])
	assert.Len(values, 21)
}

func TestValuesReversed(t *testing.T) {
	assert := assert.New(t)

	values := Values(10.0, 1.0)
	assert.NotEmpty(values)
	assert.Len(values, 10)
}
