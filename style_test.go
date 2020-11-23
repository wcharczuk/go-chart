package chart

import (
	"testing"

	"github.com/golang/freetype/truetype"
	"github.com/wcharczuk/go-chart/v2/drawing"
	"github.com/wcharczuk/go-chart/v2/testutil"
)

func TestStyleIsZero(t *testing.T) {
	// replaced new assertions helper
	zero := Style{}
	testutil.AssertTrue(t, zero.IsZero())

	strokeColor := Style{StrokeColor: drawing.ColorWhite}
	testutil.AssertFalse(t, strokeColor.IsZero())

	fillColor := Style{FillColor: drawing.ColorWhite}
	testutil.AssertFalse(t, fillColor.IsZero())

	strokeWidth := Style{StrokeWidth: 5.0}
	testutil.AssertFalse(t, strokeWidth.IsZero())

	fontSize := Style{FontSize: 12.0}
	testutil.AssertFalse(t, fontSize.IsZero())

	fontColor := Style{FontColor: drawing.ColorWhite}
	testutil.AssertFalse(t, fontColor.IsZero())

	font := Style{Font: &truetype.Font{}}
	testutil.AssertFalse(t, font.IsZero())
}

func TestStyleGetStrokeColor(t *testing.T) {
	// replaced new assertions helper

	unset := Style{}
	testutil.AssertEqual(t, drawing.ColorTransparent, unset.GetStrokeColor())
	testutil.AssertEqual(t, drawing.ColorWhite, unset.GetStrokeColor(drawing.ColorWhite))

	set := Style{StrokeColor: drawing.ColorWhite}
	testutil.AssertEqual(t, drawing.ColorWhite, set.GetStrokeColor())
	testutil.AssertEqual(t, drawing.ColorWhite, set.GetStrokeColor(drawing.ColorBlack))
}

func TestStyleGetFillColor(t *testing.T) {
	// replaced new assertions helper

	unset := Style{}
	testutil.AssertEqual(t, drawing.ColorTransparent, unset.GetFillColor())
	testutil.AssertEqual(t, drawing.ColorWhite, unset.GetFillColor(drawing.ColorWhite))

	set := Style{FillColor: drawing.ColorWhite}
	testutil.AssertEqual(t, drawing.ColorWhite, set.GetFillColor())
	testutil.AssertEqual(t, drawing.ColorWhite, set.GetFillColor(drawing.ColorBlack))
}

func TestStyleGetStrokeWidth(t *testing.T) {
	// replaced new assertions helper

	unset := Style{}
	testutil.AssertEqual(t, DefaultStrokeWidth, unset.GetStrokeWidth())
	testutil.AssertEqual(t, DefaultStrokeWidth+1, unset.GetStrokeWidth(DefaultStrokeWidth+1))

	set := Style{StrokeWidth: DefaultStrokeWidth + 2}
	testutil.AssertEqual(t, DefaultStrokeWidth+2, set.GetStrokeWidth())
	testutil.AssertEqual(t, DefaultStrokeWidth+2, set.GetStrokeWidth(DefaultStrokeWidth+1))
}

func TestStyleGetFontSize(t *testing.T) {
	// replaced new assertions helper

	unset := Style{}
	testutil.AssertEqual(t, DefaultFontSize, unset.GetFontSize())
	testutil.AssertEqual(t, DefaultFontSize+1, unset.GetFontSize(DefaultFontSize+1))

	set := Style{FontSize: DefaultFontSize + 2}
	testutil.AssertEqual(t, DefaultFontSize+2, set.GetFontSize())
	testutil.AssertEqual(t, DefaultFontSize+2, set.GetFontSize(DefaultFontSize+1))
}

func TestStyleGetFontColor(t *testing.T) {
	// replaced new assertions helper

	unset := Style{}
	testutil.AssertEqual(t, drawing.ColorTransparent, unset.GetFontColor())
	testutil.AssertEqual(t, drawing.ColorWhite, unset.GetFontColor(drawing.ColorWhite))

	set := Style{FontColor: drawing.ColorWhite}
	testutil.AssertEqual(t, drawing.ColorWhite, set.GetFontColor())
	testutil.AssertEqual(t, drawing.ColorWhite, set.GetFontColor(drawing.ColorBlack))
}

func TestStyleGetFont(t *testing.T) {
	// replaced new assertions helper

	f, err := GetDefaultFont()
	testutil.AssertNil(t, err)

	unset := Style{}
	testutil.AssertNil(t, unset.GetFont())
	testutil.AssertEqual(t, f, unset.GetFont(f))

	set := Style{Font: f}
	testutil.AssertNotNil(t, set.GetFont())
}

func TestStyleGetPadding(t *testing.T) {
	// replaced new assertions helper

	unset := Style{}
	testutil.AssertTrue(t, unset.GetPadding().IsZero())
	testutil.AssertFalse(t, unset.GetPadding(DefaultBackgroundPadding).IsZero())
	testutil.AssertEqual(t, DefaultBackgroundPadding, unset.GetPadding(DefaultBackgroundPadding))

	set := Style{Padding: DefaultBackgroundPadding}
	testutil.AssertFalse(t, set.GetPadding().IsZero())
	testutil.AssertEqual(t, DefaultBackgroundPadding, set.GetPadding())
	testutil.AssertEqual(t, DefaultBackgroundPadding, set.GetPadding(Box{
		Top:    DefaultBackgroundPadding.Top + 1,
		Left:   DefaultBackgroundPadding.Left + 1,
		Right:  DefaultBackgroundPadding.Right + 1,
		Bottom: DefaultBackgroundPadding.Bottom + 1,
	}))
}

func TestStyleWithDefaultsFrom(t *testing.T) {
	// replaced new assertions helper

	f, err := GetDefaultFont()
	testutil.AssertNil(t, err)

	unset := Style{}
	set := Style{
		StrokeColor: drawing.ColorWhite,
		StrokeWidth: 5.0,
		FillColor:   drawing.ColorWhite,
		FontColor:   drawing.ColorWhite,
		Font:        f,
		Padding:     DefaultBackgroundPadding,
	}

	coalesced := unset.InheritFrom(set)
	testutil.AssertEqual(t, set, coalesced)
}

func TestStyleGetStrokeOptions(t *testing.T) {
	// replaced new assertions helper

	set := Style{
		StrokeColor: drawing.ColorWhite,
		StrokeWidth: 5.0,
		FillColor:   drawing.ColorWhite,
		FontColor:   drawing.ColorWhite,
		Padding:     DefaultBackgroundPadding,
	}
	svgStroke := set.GetStrokeOptions()
	testutil.AssertFalse(t, svgStroke.StrokeColor.IsZero())
	testutil.AssertNotZero(t, svgStroke.StrokeWidth)
	testutil.AssertTrue(t, svgStroke.FillColor.IsZero())
	testutil.AssertTrue(t, svgStroke.FontColor.IsZero())
}

func TestStyleGetFillOptions(t *testing.T) {
	// replaced new assertions helper

	set := Style{
		StrokeColor: drawing.ColorWhite,
		StrokeWidth: 5.0,
		FillColor:   drawing.ColorWhite,
		FontColor:   drawing.ColorWhite,
		Padding:     DefaultBackgroundPadding,
	}
	svgFill := set.GetFillOptions()
	testutil.AssertFalse(t, svgFill.FillColor.IsZero())
	testutil.AssertZero(t, svgFill.StrokeWidth)
	testutil.AssertTrue(t, svgFill.StrokeColor.IsZero())
	testutil.AssertTrue(t, svgFill.FontColor.IsZero())
}

func TestStyleGetFillAndStrokeOptions(t *testing.T) {
	// replaced new assertions helper

	set := Style{
		StrokeColor: drawing.ColorWhite,
		StrokeWidth: 5.0,
		FillColor:   drawing.ColorWhite,
		FontColor:   drawing.ColorWhite,
		Padding:     DefaultBackgroundPadding,
	}
	svgFillAndStroke := set.GetFillAndStrokeOptions()
	testutil.AssertFalse(t, svgFillAndStroke.FillColor.IsZero())
	testutil.AssertNotZero(t, svgFillAndStroke.StrokeWidth)
	testutil.AssertFalse(t, svgFillAndStroke.StrokeColor.IsZero())
	testutil.AssertTrue(t, svgFillAndStroke.FontColor.IsZero())
}

func TestStyleGetTextOptions(t *testing.T) {
	// replaced new assertions helper

	set := Style{
		StrokeColor: drawing.ColorWhite,
		StrokeWidth: 5.0,
		FillColor:   drawing.ColorWhite,
		FontColor:   drawing.ColorWhite,
		Padding:     DefaultBackgroundPadding,
	}
	svgStroke := set.GetTextOptions()
	testutil.AssertTrue(t, svgStroke.StrokeColor.IsZero())
	testutil.AssertZero(t, svgStroke.StrokeWidth)
	testutil.AssertTrue(t, svgStroke.FillColor.IsZero())
	testutil.AssertFalse(t, svgStroke.FontColor.IsZero())
}
