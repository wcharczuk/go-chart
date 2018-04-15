package chart

import (
	"testing"

	assert "github.com/blend/go-sdk/assert"
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
	assert.Len(3, output)

	for _, line := range output {
		basicTextStyle.WriteToRenderer(r)
		lineBox := r.MeasureText(line)
		assert.True(lineBox.Width() < 100, line+": "+lineBox.String())
	}
	assert.Equal("this is", output[0])
	assert.Equal("a test", output[1])
	assert.Equal("string", output[2])

	output = Text.WrapFitWord(r, "foo", 100, basicTextStyle)
	assert.Len(1, output)
	assert.Equal("foo", output[0])

	// test that it handles newlines.
	output = Text.WrapFitWord(r, "this\nis\na\ntest\nstring", 100, basicTextStyle)
	assert.Len(5, output)

	// test that it handles newlines and long lines.
	output = Text.WrapFitWord(r, "this\nis\na\ntest\nstring that is very long", 100, basicTextStyle)
	assert.Len(8, output)
}

func TestTextWrapRune(t *testing.T) {
	assert := assert.New(t)

	r, err := PNG(1024, 1024)
	assert.Nil(err)
	f, err := GetDefaultFont()
	assert.Nil(err)

	basicTextStyle := Style{Font: f, FontSize: 24}

	output := Text.WrapFitRune(r, "this is a test string", 150, basicTextStyle)
	assert.NotEmpty(output)
	assert.Len(2, output)
	assert.Equal("this is a t", output[0])
	assert.Equal("est string", output[1])
}
