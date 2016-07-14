package chart

import (
	"fmt"
	"strings"

	"github.com/golang/freetype/truetype"
	"github.com/wcharczuk/go-chart/drawing"
)

// Style is a simple style set.
type Style struct {
	Show    bool
	Padding Box

	StrokeWidth     float64
	StrokeColor     drawing.Color
	StrokeDashArray []float64

	FillColor drawing.Color
	FontSize  float64
	FontColor drawing.Color
	Font      *truetype.Font
}

// IsZero returns if the object is set or not.
func (s Style) IsZero() bool {
	return s.StrokeColor.IsZero() && s.FillColor.IsZero() && s.StrokeWidth == 0 && s.FontColor.IsZero() && s.FontSize == 0 && s.Font == nil
}

// GetStrokeColor returns the stroke color.
func (s Style) GetStrokeColor(defaults ...drawing.Color) drawing.Color {
	if s.StrokeColor.IsZero() {
		if len(defaults) > 0 {
			return defaults[0]
		}
		return drawing.ColorTransparent
	}
	return s.StrokeColor
}

// GetFillColor returns the fill color.
func (s Style) GetFillColor(defaults ...drawing.Color) drawing.Color {
	if s.FillColor.IsZero() {
		if len(defaults) > 0 {
			return defaults[0]
		}
		return drawing.ColorTransparent
	}
	return s.FillColor
}

// GetStrokeWidth returns the stroke width.
func (s Style) GetStrokeWidth(defaults ...float64) float64 {
	if s.StrokeWidth == 0 {
		if len(defaults) > 0 {
			return defaults[0]
		}
		return DefaultStrokeWidth
	}
	return s.StrokeWidth
}

// GetStrokeDashArray returns the stroke dash array.
func (s Style) GetStrokeDashArray(defaults ...[]float64) []float64 {
	if len(s.StrokeDashArray) == 0 {
		if len(defaults) > 0 {
			return defaults[0]
		}
		return nil
	}
	return s.StrokeDashArray
}

// GetFontSize gets the font size.
func (s Style) GetFontSize(defaults ...float64) float64 {
	if s.FontSize == 0 {
		if len(defaults) > 0 {
			return defaults[0]
		}
		return DefaultFontSize
	}
	return s.FontSize
}

// GetFontColor gets the font size.
func (s Style) GetFontColor(defaults ...drawing.Color) drawing.Color {
	if s.FontColor.IsZero() {
		if len(defaults) > 0 {
			return defaults[0]
		}
		return drawing.ColorTransparent
	}
	return s.FontColor
}

// GetFont returns the font face.
func (s Style) GetFont(defaults ...*truetype.Font) *truetype.Font {
	if s.Font == nil {
		if len(defaults) > 0 {
			return defaults[0]
		}
		return nil
	}
	return s.Font
}

// GetPadding returns the padding.
func (s Style) GetPadding(defaults ...Box) Box {
	if s.Padding.IsZero() {
		if len(defaults) > 0 {
			return defaults[0]
		}
		return Box{}
	}
	return s.Padding
}

// WithDefaultsFrom coalesces two styles into a new style.
func (s Style) WithDefaultsFrom(defaults Style) (final Style) {
	final.StrokeColor = s.GetStrokeColor(defaults.StrokeColor)
	final.StrokeWidth = s.GetStrokeWidth(defaults.StrokeWidth)
	final.StrokeDashArray = s.GetStrokeDashArray(defaults.StrokeDashArray)
	final.FillColor = s.GetFillColor(defaults.FillColor)
	final.FontColor = s.GetFontColor(defaults.FontColor)
	final.Font = s.GetFont(defaults.Font)
	final.Padding = s.GetPadding(defaults.Padding)
	return
}

// SVG returns the style as a svg style string.
func (s Style) SVG(dpi float64) string {
	sw := s.StrokeWidth
	sc := s.StrokeColor
	fc := s.FillColor
	fs := s.FontSize
	fnc := s.FontColor

	strokeWidthText := "stroke-width:0"
	if sw != 0 {
		strokeWidthText = "stroke-width:" + fmt.Sprintf("%d", int(sw))
	}

	strokeText := "stroke:none"
	if !sc.IsZero() {
		strokeText = "stroke:" + sc.String()
	}

	fillText := "fill:none"
	if !fc.IsZero() {
		fillText = "fill:" + fc.String()
	}

	fontSizeText := ""
	if fs != 0 {
		fontSizeText = "font-size:" + fmt.Sprintf("%.1fpx", drawing.PointsToPixels(dpi, fs))
	}

	if !fnc.IsZero() {
		fillText = "fill:" + fnc.String()
	}
	return strings.Join([]string{strokeWidthText, strokeText, fillText, fontSizeText}, ";")
}

// SVGStroke returns the stroke components.
func (s Style) SVGStroke() Style {
	return Style{
		StrokeColor: s.StrokeColor,
		StrokeWidth: s.StrokeWidth,
	}
}

// SVGFill returns the fill components.
func (s Style) SVGFill() Style {
	return Style{
		FillColor: s.FillColor,
	}
}

// SVGFillAndStroke returns the fill and stroke components.
func (s Style) SVGFillAndStroke() Style {
	return Style{
		FillColor:   s.FillColor,
		StrokeColor: s.StrokeColor,
		StrokeWidth: s.StrokeWidth,
	}
}

// SVGText returns just the text components of the style.
func (s Style) SVGText() Style {
	return Style{
		FontColor: s.FontColor,
		FontSize:  s.FontSize,
	}
}
