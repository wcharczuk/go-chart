package chart

import (
	"testing"

	"github.com/blend/go-sdk/assert"
	"github.com/golang/freetype/truetype"
	"github.com/wcharczuk/go-chart/drawing"
)

func TestStyleIsZero(t *testing.T) {
	assert := assert.New(t)
	zero := Style{}
	assert.True(zero.IsZero())

	strokeColor := Style{StrokeColor: drawing.ColorWhite}
	assert.False(strokeColor.IsZero())

	fillColor := Style{FillColor: drawing.ColorWhite}
	assert.False(fillColor.IsZero())

	strokeWidth := Style{StrokeWidth: 5.0}
	assert.False(strokeWidth.IsZero())

	fontSize := Style{FontSize: 12.0}
	assert.False(fontSize.IsZero())

	fontColor := Style{FontColor: drawing.ColorWhite}
	assert.False(fontColor.IsZero())

	font := Style{Font: &truetype.Font{}}
	assert.False(font.IsZero())
}

func TestStyleGetStrokeColor(t *testing.T) {
	assert := assert.New(t)

	unset := Style{}
	assert.Equal(drawing.ColorTransparent, unset.GetStrokeColor())
	assert.Equal(drawing.ColorWhite, unset.GetStrokeColor(drawing.ColorWhite))

	set := Style{StrokeColor: drawing.ColorWhite}
	assert.Equal(drawing.ColorWhite, set.GetStrokeColor())
	assert.Equal(drawing.ColorWhite, set.GetStrokeColor(drawing.ColorBlack))
}

func TestStyleGetFillColor(t *testing.T) {
	assert := assert.New(t)

	unset := Style{}
	assert.Equal(drawing.ColorTransparent, unset.GetFillColor())
	assert.Equal(drawing.ColorWhite, unset.GetFillColor(drawing.ColorWhite))

	set := Style{FillColor: drawing.ColorWhite}
	assert.Equal(drawing.ColorWhite, set.GetFillColor())
	assert.Equal(drawing.ColorWhite, set.GetFillColor(drawing.ColorBlack))
}

func TestStyleGetStrokeWidth(t *testing.T) {
	assert := assert.New(t)

	unset := Style{}
	assert.Equal(DefaultStrokeWidth, unset.GetStrokeWidth())
	assert.Equal(DefaultStrokeWidth+1, unset.GetStrokeWidth(DefaultStrokeWidth+1))

	set := Style{StrokeWidth: DefaultStrokeWidth + 2}
	assert.Equal(DefaultStrokeWidth+2, set.GetStrokeWidth())
	assert.Equal(DefaultStrokeWidth+2, set.GetStrokeWidth(DefaultStrokeWidth+1))
}

func TestStyleGetFontSize(t *testing.T) {
	assert := assert.New(t)

	unset := Style{}
	assert.Equal(DefaultFontSize, unset.GetFontSize())
	assert.Equal(DefaultFontSize+1, unset.GetFontSize(DefaultFontSize+1))

	set := Style{FontSize: DefaultFontSize + 2}
	assert.Equal(DefaultFontSize+2, set.GetFontSize())
	assert.Equal(DefaultFontSize+2, set.GetFontSize(DefaultFontSize+1))
}

func TestStyleGetFontColor(t *testing.T) {
	assert := assert.New(t)

	unset := Style{}
	assert.Equal(drawing.ColorTransparent, unset.GetFontColor())
	assert.Equal(drawing.ColorWhite, unset.GetFontColor(drawing.ColorWhite))

	set := Style{FontColor: drawing.ColorWhite}
	assert.Equal(drawing.ColorWhite, set.GetFontColor())
	assert.Equal(drawing.ColorWhite, set.GetFontColor(drawing.ColorBlack))
}

func TestStyleGetFont(t *testing.T) {
	assert := assert.New(t)

	f, err := GetDefaultFont()
	assert.Nil(err)

	unset := Style{}
	assert.Nil(unset.GetFont())
	assert.Equal(f, unset.GetFont(f))

	set := Style{Font: f}
	assert.NotNil(set.GetFont())
}

func TestStyleGetPadding(t *testing.T) {
	assert := assert.New(t)

	unset := Style{}
	assert.True(unset.GetPadding().IsZero())
	assert.False(unset.GetPadding(DefaultBackgroundPadding).IsZero())
	assert.Equal(DefaultBackgroundPadding, unset.GetPadding(DefaultBackgroundPadding))

	set := Style{Padding: DefaultBackgroundPadding}
	assert.False(set.GetPadding().IsZero())
	assert.Equal(DefaultBackgroundPadding, set.GetPadding())
	assert.Equal(DefaultBackgroundPadding, set.GetPadding(Box{
		Top:    DefaultBackgroundPadding.Top + 1,
		Left:   DefaultBackgroundPadding.Left + 1,
		Right:  DefaultBackgroundPadding.Right + 1,
		Bottom: DefaultBackgroundPadding.Bottom + 1,
	}))
}

func TestStyleWithDefaultsFrom(t *testing.T) {
	assert := assert.New(t)

	f, err := GetDefaultFont()
	assert.Nil(err)

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
	assert.Equal(set, coalesced)
}

func TestStyleGetStrokeOptions(t *testing.T) {
	assert := assert.New(t)

	set := Style{
		StrokeColor: drawing.ColorWhite,
		StrokeWidth: 5.0,
		FillColor:   drawing.ColorWhite,
		FontColor:   drawing.ColorWhite,
		Padding:     DefaultBackgroundPadding,
	}
	svgStroke := set.GetStrokeOptions()
	assert.False(svgStroke.StrokeColor.IsZero())
	assert.NotZero(svgStroke.StrokeWidth)
	assert.True(svgStroke.FillColor.IsZero())
	assert.True(svgStroke.FontColor.IsZero())
}

func TestStyleGetFillOptions(t *testing.T) {
	assert := assert.New(t)

	set := Style{
		StrokeColor: drawing.ColorWhite,
		StrokeWidth: 5.0,
		FillColor:   drawing.ColorWhite,
		FontColor:   drawing.ColorWhite,
		Padding:     DefaultBackgroundPadding,
	}
	svgFill := set.GetFillOptions()
	assert.False(svgFill.FillColor.IsZero())
	assert.Zero(svgFill.StrokeWidth)
	assert.True(svgFill.StrokeColor.IsZero())
	assert.True(svgFill.FontColor.IsZero())
}

func TestStyleGetFillAndStrokeOptions(t *testing.T) {
	assert := assert.New(t)

	set := Style{
		StrokeColor: drawing.ColorWhite,
		StrokeWidth: 5.0,
		FillColor:   drawing.ColorWhite,
		FontColor:   drawing.ColorWhite,
		Padding:     DefaultBackgroundPadding,
	}
	svgFillAndStroke := set.GetFillAndStrokeOptions()
	assert.False(svgFillAndStroke.FillColor.IsZero())
	assert.NotZero(svgFillAndStroke.StrokeWidth)
	assert.False(svgFillAndStroke.StrokeColor.IsZero())
	assert.True(svgFillAndStroke.FontColor.IsZero())
}

func TestStyleGetTextOptions(t *testing.T) {
	assert := assert.New(t)

	set := Style{
		StrokeColor: drawing.ColorWhite,
		StrokeWidth: 5.0,
		FillColor:   drawing.ColorWhite,
		FontColor:   drawing.ColorWhite,
		Padding:     DefaultBackgroundPadding,
	}
	svgStroke := set.GetTextOptions()
	assert.True(svgStroke.StrokeColor.IsZero())
	assert.Zero(svgStroke.StrokeWidth)
	assert.True(svgStroke.FillColor.IsZero())
	assert.False(svgStroke.FontColor.IsZero())
}
