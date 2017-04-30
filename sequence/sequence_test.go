package sequence

import (
	"testing"

	assert "github.com/blendlabs/go-assert"
)

func TestSequenceEach(t *testing.T) {
	assert := assert.New(t)

	values := Seq{Array([]float64{1, 2, 3, 4})}
	values.Each(func(i int, v float64) {
		assert.Equal(i, v)
	})
}

func TestSequenceMap(t *testing.T) {
	assert := assert.New(t)

	values := Seq{Array([]float64{1, 2, 3, 4})}
	mapped := values.Map(func(i int, v float64) float64 {
		assert.Equal(i, v)
		return v * 2
	})
	assert.Equal(4, mapped.Len())
}
