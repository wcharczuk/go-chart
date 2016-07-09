package drawing

import (
	"testing"

	"github.com/blendlabs/go-assert"
)

func TestColorFromHex(t *testing.T) {
	assert := assert.New(t)

	white := ColorFromHex("FFFFFF")
	assert.Equal(ColorWhite, white)

	shortWhite := ColorFromHex("FFF")
	assert.Equal(ColorWhite, shortWhite)

	black := ColorFromHex("000000")
	assert.Equal(ColorBlack, black)

	shortBlack := ColorFromHex("000")
	assert.Equal(ColorBlack, shortBlack)

	red := ColorFromHex("FF0000")
	assert.Equal(ColorRed, red)

	shortRed := ColorFromHex("F00")
	assert.Equal(ColorRed, shortRed)

	green := ColorFromHex("00FF00")
	assert.Equal(ColorGreen, green)

	shortGreen := ColorFromHex("0F0")
	assert.Equal(ColorGreen, shortGreen)

	blue := ColorFromHex("0000FF")
	assert.Equal(ColorBlue, blue)

	shortBlue := ColorFromHex("00F")
	assert.Equal(ColorBlue, shortBlue)
}
