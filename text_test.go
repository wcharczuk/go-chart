package chart

import (
	"testing"

	"github.com/wcharczuk/go-chart/v2/testutil"
)

func TestTextWrapWord(t *testing.T) {
	// replaced new assertions helper

	r, err := PNG(1024, 1024)
	testutil.AssertNil(t, err)
	f, err := GetDefaultFont()
	testutil.AssertNil(t, err)

	basicTextStyle := Style{Font: f, FontSize: 24}

	output := Text.WrapFitWord(r, "this is a test string", 100, basicTextStyle)
	testutil.AssertNotEmpty(t, output)
	testutil.AssertLen(t, output, 3)

	for _, line := range output {
		basicTextStyle.WriteToRenderer(r)
		lineBox := r.MeasureText(line)
		testutil.AssertTrue(t, lineBox.Width() < 100, line+": "+lineBox.String())
	}
	testutil.AssertEqual(t, "this is", output[0])
	testutil.AssertEqual(t, "a test", output[1])
	testutil.AssertEqual(t, "string", output[2])

	output = Text.WrapFitWord(r, "foo", 100, basicTextStyle)
	testutil.AssertLen(t, output, 1)
	testutil.AssertEqual(t, "foo", output[0])

	// test that it handles newlines.
	output = Text.WrapFitWord(r, "this\nis\na\ntest\nstring", 100, basicTextStyle)
	testutil.AssertLen(t, output, 5)

	// test that it handles newlines and long lines.
	output = Text.WrapFitWord(r, "this\nis\na\ntest\nstring that is very long", 100, basicTextStyle)
	testutil.AssertLen(t, output, 8)
}

func TestTextWrapRune(t *testing.T) {
	// replaced new assertions helper

	r, err := PNG(1024, 1024)
	testutil.AssertNil(t, err)
	f, err := GetDefaultFont()
	testutil.AssertNil(t, err)

	basicTextStyle := Style{Font: f, FontSize: 24}

	output := Text.WrapFitRune(r, "this is a test string", 150, basicTextStyle)
	testutil.AssertNotEmpty(t, output)
	testutil.AssertLen(t, output, 2)
	testutil.AssertEqual(t, "this is a t", output[0])
	testutil.AssertEqual(t, "est string", output[1])
}
