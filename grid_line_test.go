package chart

import (
	"testing"

	"github.com/blend/go-sdk/assert"
)

func TestGenerateGridLines(t *testing.T) {
	assert := assert.New(t)

	ticks := []Tick{
		{Value: 1.0, Label: "1.0"},
		{Value: 2.0, Label: "2.0"},
		{Value: 3.0, Label: "3.0"},
		{Value: 4.0, Label: "4.0"},
	}

	gl := GenerateGridLines(ticks, Style{}, Style{})
	assert.Len(2, gl)

	assert.Equal(2.0, gl[0].Value)
	assert.Equal(3.0, gl[1].Value)
}
