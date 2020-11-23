package drawing

import (
	"testing"

	"image/color"

	"github.com/wcharczuk/go-chart/v2/testutil"
)

func TestColorFromHex(t *testing.T) {
	// replaced new assertions helper

	white := ColorFromHex("FFFFFF")
	testutil.AssertEqual(t, ColorWhite, white)

	shortWhite := ColorFromHex("FFF")
	testutil.AssertEqual(t, ColorWhite, shortWhite)

	black := ColorFromHex("000000")
	testutil.AssertEqual(t, ColorBlack, black)

	shortBlack := ColorFromHex("000")
	testutil.AssertEqual(t, ColorBlack, shortBlack)

	red := ColorFromHex("FF0000")
	testutil.AssertEqual(t, ColorRed, red)

	shortRed := ColorFromHex("F00")
	testutil.AssertEqual(t, ColorRed, shortRed)

	green := ColorFromHex("00FF00")
	testutil.AssertEqual(t, ColorGreen, green)

	shortGreen := ColorFromHex("0F0")
	testutil.AssertEqual(t, ColorGreen, shortGreen)

	blue := ColorFromHex("0000FF")
	testutil.AssertEqual(t, ColorBlue, blue)

	shortBlue := ColorFromHex("00F")
	testutil.AssertEqual(t, ColorBlue, shortBlue)
}

func TestColorFromAlphaMixedRGBA(t *testing.T) {
	// replaced new assertions helper

	black := ColorFromAlphaMixedRGBA(color.Black.RGBA())
	testutil.AssertTrue(t, black.Equals(ColorBlack), black.String())

	white := ColorFromAlphaMixedRGBA(color.White.RGBA())
	testutil.AssertTrue(t, white.Equals(ColorWhite), white.String())
}
