package chart

import "image/color"

// Style is a simple style set.
type Style struct {
	StrokeColor color.RGBA
	FillColor   color.RGBA
	StrokeWidth int
}

// IsZero returns if the object is set or not.
func (s Style) IsZero() bool {
	return ColorIsZero(s.StrokeColor) && ColorIsZero(s.FillColor) && s.StrokeWidth == 0
}

// GetStrokeColor returns the stroke color.
func (s Style) GetStrokeColor() color.RGBA {
	if ColorIsZero(s.StrokeColor) {
		return DefaultLineColor
	}
	return s.StrokeColor
}

// GetFillColor returns the fill color.
func (s Style) GetFillColor() color.RGBA {
	if ColorIsZero(s.FillColor) {
		return DefaultFillColor
	}
	return s.FillColor
}

// GetStrokeWidth returns the stroke width.
func (s Style) GetStrokeWidth() int {
	if s.StrokeWidth == 0 {
		return DefaultLineWidth
	}
	return s.StrokeWidth
}
