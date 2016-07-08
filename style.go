package chart

import "image/color"

// Style is a simple style set.
type Style struct {
	Show        bool
	StrokeColor color.RGBA
	FillColor   color.RGBA
	StrokeWidth float64
	FontSize    float64
	FontColor   color.RGBA
	Padding     Box
}

// IsZero returns if the object is set or not.
func (s Style) IsZero() bool {
	return ColorIsZero(s.StrokeColor) && ColorIsZero(s.FillColor) && s.StrokeWidth == 0 && s.FontSize == 0
}

// GetStrokeColor returns the stroke color.
func (s Style) GetStrokeColor(defaults ...color.RGBA) color.RGBA {
	if ColorIsZero(s.StrokeColor) {
		if len(defaults) > 0 {
			return defaults[0]
		}
		return DefaultStrokeColor
	}
	return s.StrokeColor
}

// GetFillColor returns the fill color.
func (s Style) GetFillColor(defaults ...color.RGBA) color.RGBA {
	if ColorIsZero(s.FillColor) {
		if len(defaults) > 0 {
			return defaults[0]
		}
		return DefaultFillColor
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
func (s Style) GetFontColor(defaults ...color.RGBA) color.RGBA {
	if ColorIsZero(s.FontColor) {
		if len(defaults) > 0 {
			return defaults[0]
		}
		return DefaultTextColor
	}
	return s.FontColor
}
