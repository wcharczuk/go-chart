package chart

import (
	"io"

	"github.com/golang/freetype/truetype"
	"github.com/wcharczuk/go-chart/drawing"
)

// Renderer represents the basic methods required to draw a chart.
type Renderer interface {
	// GetDPI gets the DPI for the renderer.
	GetDPI() float64

	// SetDPI sets the DPI for the renderer.
	SetDPI(dpi float64)

	// SetStrokeColor sets the current stroke color.
	SetStrokeColor(drawing.Color)

	// SetFillColor sets the current fill color.
	SetFillColor(drawing.Color)

	// SetStrokeWidth sets the stroke width.
	SetStrokeWidth(width float64)

	// MoveTo moves the cursor to a given point.
	MoveTo(x, y int)

	// LineTo both starts a shape and draws a line to a given point
	// from the previous point.
	LineTo(x, y int)

	// Close finalizes a shape as drawn by LineTo.
	Close()

	// Stroke strokes the path.
	Stroke()

	// Fill fills the path, but does not stroke.
	Fill()

	// FillStroke fills and strokes a path.
	FillStroke()

	// Circle draws a circle at the given coords with a given radius.
	Circle(radius float64, x, y int)

	// SetFont sets a font for a text field.
	SetFont(*truetype.Font)

	// SetFontColor sets a font's color
	SetFontColor(drawing.Color)

	// SetFontSize sets the font size for a text field.
	SetFontSize(size float64)

	// Text draws a text blob.
	Text(body string, x, y int)

	// MeasureText measures text.
	MeasureText(body string) (width int, height int)

	// Save writes the image to the given writer.
	Save(w io.Writer) error
}
