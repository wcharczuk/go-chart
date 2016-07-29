package chart

import (
	"testing"

	assert "github.com/blendlabs/go-assert"
)

func TestTextWrapWord(t *testing.T) {
	assert := assert.New(t)

	r, err := PNG(1024, 1024)
	assert.Nil(err)
	f, err := GetDefaultFont()
	assert.Nil(err)

	basicTextStyle := Style{Font: f, FontSize: 24}

	output := Text.WrapFitWord(r, "this is a test string", 100, basicTextStyle)
	assert.NotEmpty(output)
	assert.Len(output, 3)

	for _, line := range output {
		basicTextStyle.WriteToRenderer(r)
		lineBox := r.MeasureText(line)
		assert.True(lineBox.Width() < 100, line+": "+lineBox.String())
	}

	output = Text.WrapFitWord(r, "foo", 100, basicTextStyle)
	assert.Len(output, 1)
	assert.Equal("foo", output[0])
}
