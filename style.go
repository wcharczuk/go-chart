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

// String returns a text representation of the style.
func (s Style) String() string {
	if s.IsZero() {
		return "{}"
	}

	var output []string
	if s.Show {
		output = []string{"\"show\": true"}
	} else {
		output = []string{"\"show\": false"}
	}

	if !s.Padding.IsZero() {
		output = append(output, fmt.Sprintf("\"padding\": %s", s.Padding.String()))
	} else {
		output = append(output, "\"padding\": null")
	}

	if s.StrokeWidth >= 0 {
		output = append(output, fmt.Sprintf("\"stroke_width\": %0.2f", s.StrokeWidth))
	} else {
		output = append(output, "\"stroke_width\": null")
	}

	if !s.StrokeColor.IsZero() {
		output = append(output, fmt.Sprintf("\"stroke_color\": %s", s.StrokeColor.String()))
	} else {
		output = append(output, "\"stroke_color\": null")
	}

	if len(s.StrokeDashArray) > 0 {
		var elements []string
		for _, v := range s.StrokeDashArray {
			elements = append(elements, fmt.Sprintf("%.2f", v))
		}
		dashArray := strings.Join(elements, ", ")
		output = append(output, fmt.Sprintf("\"stroke_dash_array\": [%s]", dashArray))
	} else {
		output = append(output, "\"stroke_dash_array\": null")
	}

	if !s.FillColor.IsZero() {
		output = append(output, fmt.Sprintf("\"fill_color\": %s", s.FillColor.String()))
	} else {
		output = append(output, "\"fill_color\": null")
	}

	if s.FontSize != 0 {
		output = append(output, fmt.Sprintf("\"font_size\": \"%0.2fpt\"", s.FontSize))
	} else {
		output = append(output, "\"fill_color\": null")
	}

	if !s.FillColor.IsZero() {
		output = append(output, fmt.Sprintf("\"font_color\": %s", s.FillColor.String()))
	} else {
		output = append(output, "\"font_color\": null")
	}

	if s.Font != nil {
		output = append(output, fmt.Sprintf("\"font\": \"%s\"", s.Font.Name(truetype.NameIDFontFamily)))
	} else {
		output = append(output, "\"font_color\": null")
	}

	return "{" + strings.Join(output, ", ") + "}"
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

// PersistToRenderer passes the style onto a renderer.
func (s Style) PersistToRenderer(r Renderer) {
	r.SetStrokeColor(s.GetStrokeColor())
	r.SetStrokeWidth(s.GetStrokeWidth())
	r.SetStrokeDashArray(s.GetStrokeDashArray())
	r.SetFont(s.GetFont())
	r.SetFontColor(s.GetFontColor())
	r.SetFontSize(s.GetFontSize())
}

// InheritFrom coalesces two styles into a new style.
func (s Style) InheritFrom(defaults Style) (final Style) {
	final.StrokeColor = s.GetStrokeColor(defaults.StrokeColor)
	final.StrokeWidth = s.GetStrokeWidth(defaults.StrokeWidth)
	final.StrokeDashArray = s.GetStrokeDashArray(defaults.StrokeDashArray)
	final.FillColor = s.GetFillColor(defaults.FillColor)
	final.FontColor = s.GetFontColor(defaults.FontColor)
	final.FontSize = s.GetFontSize(defaults.FontSize)
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

	fontText := s.SVGFontFace()
	return strings.Join([]string{strokeWidthText, strokeText, fillText, fontSizeText, fontText}, ";")
}

// SVGStroke returns the stroke components.
func (s Style) SVGStroke() Style {
	return Style{
		StrokeDashArray: s.StrokeDashArray,
		StrokeColor:     s.StrokeColor,
		StrokeWidth:     s.StrokeWidth,
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
		StrokeDashArray: s.StrokeDashArray,
		FillColor:       s.FillColor,
		StrokeColor:     s.StrokeColor,
		StrokeWidth:     s.StrokeWidth,
	}
}

// SVGText returns just the text components of the style.
func (s Style) SVGText() Style {
	return Style{
		FontColor: s.FontColor,
		FontSize:  s.FontSize,
	}
}

// SVGFontFace returns the font face for the style.
func (s Style) SVGFontFace() string {
	family := "sans-serif"
	if s.GetFont() != nil {
		name := s.GetFont().Name(truetype.NameIDFontFamily)
		if len(name) != 0 {
			family = fmt.Sprintf(`'%s',%s`, name, family)
		}
	}
	return fmt.Sprintf("font-family:%s", family)
}

// SVGStrokeDashArray returns the stroke-dasharray property of a style.
func (s Style) SVGStrokeDashArray() string {
	if len(s.StrokeDashArray) > 0 {
		var values []string
		for _, v := range s.StrokeDashArray {
			values = append(values, fmt.Sprintf("%0.1f", v))
		}
		return "stroke-dasharray=\"" + strings.Join(values, ", ") + "\""
	}
	return ""
}
