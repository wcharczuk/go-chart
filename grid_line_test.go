package chart

import (
	"testing"

	"github.com/blendlabs/go-assert"
)

func TestGenerateGridLines(t *testing.T) {
	assert := assert.New(t)

	ticks := []Tick{
		Tick{Value: 1.0, Label: "1.0"},
		Tick{Value: 2.0, Label: "2.0"},
		Tick{Value: 3.0, Label: "3.0"},
		Tick{Value: 4.0, Label: "4.0"},
	}

	gl := GenerateGridLines(ticks, true)
	assert.Len(gl, 4)
	assert.Equal(1.0, gl[0].Value)
	assert.Equal(2.0, gl[1].Value)
	assert.Equal(3.0, gl[2].Value)
	assert.Equal(4.0, gl[3].Value)

	assert.True(gl[0].IsVertical)
}
