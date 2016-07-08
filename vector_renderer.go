package chart

import (
	"bytes"
	"fmt"
	"image/color"
	"io"
	"strings"

	"golang.org/x/image/font"

	"github.com/ajstarks/svgo"
	"github.com/golang/freetype/truetype"
)

// SVG returns a new png/raster renderer.
func SVG(width, height int) Renderer {
	buffer := bytes.NewBuffer([]byte{})
	canvas := svg.New(buffer)
	canvas.Start(width, height)
	return &vectorRenderer{
		b: buffer,
		c: canvas,
		s: &Style{},
		p: []string{},
	}
}

// vectorRenderer renders chart commands to a bitmap.
type vectorRenderer struct {
	b  *bytes.Buffer
	c  *svg.SVG
	s  *Style
	f  *truetype.Font
	p  []string
	fc *font.Drawer
}

func (vr *vectorRenderer) SetStrokeColor(c color.RGBA) {
	vr.s.StrokeColor = c
}

// SetFillColor implements the interface method.
func (vr *vectorRenderer) SetFillColor(c color.RGBA) {
	vr.s.FillColor = c
}

// SetLineWidth implements the interface method.
func (vr *vectorRenderer) SetStrokeWidth(width float64) {
	vr.s.StrokeWidth = width
}

// MoveTo implements the interface method.
func (vr *vectorRenderer) MoveTo(x, y int) {
	vr.p = append(vr.p, fmt.Sprintf("M %d %d", x, y))
}

// LineTo implements the interface method.
func (vr *vectorRenderer) LineTo(x, y int) {
	vr.p = append(vr.p, fmt.Sprintf("L %d %d", x, y))
}

func (vr *vectorRenderer) Close() {
	vr.p = append(vr.p, fmt.Sprintf("Z"))
}

// Stroke draws the path with no fill.
func (vr *vectorRenderer) Stroke() {
	vr.s.FillColor = color.RGBA{}
	vr.s.FontColor = color.RGBA{}
	vr.drawPath()
}

// Fill draws the path with no stroke.
func (vr *vectorRenderer) Fill() {
	vr.s.StrokeColor = color.RGBA{}
	vr.s.StrokeWidth = 0
	vr.s.FontColor = color.RGBA{}
	vr.drawPath()
}

// FillStroke draws the path with both fill and stroke.
func (vr *vectorRenderer) FillStroke() {
	vr.s.FontColor = color.RGBA{}
	vr.drawPath()
}

func (vr *vectorRenderer) drawPath() {
	vr.c.Path(strings.Join(vr.p, "\n"), vr.s.SVG())
	vr.p = []string{}
}

// Circle implements the interface method.
func (vr *vectorRenderer) Circle(radius float64, x, y int) {
	vr.c.Circle(x, y, int(radius), vr.s.SVG())
}

// SetFont implements the interface method.
func (vr *vectorRenderer) SetFont(f *truetype.Font) {
	vr.f = f
}

// SetFontColor implements the interface method.
func (vr *vectorRenderer) SetFontColor(c color.RGBA) {
	vr.s.FontColor = c
}

// SetFontSize implements the interface method.
func (vr *vectorRenderer) SetFontSize(size float64) {
	vr.s.FontSize = size
}

func (vr *vectorRenderer) svgFontFace() string {
	family := "sans-serif"
	if vr.f != nil {
		name := vr.f.Name(truetype.NameIDFontFamily)
		if len(name) != 0 {
			family = fmt.Sprintf(`'%s',%s`, name, family)
		}
	}
	return fmt.Sprintf("font-family:%s", family)
}

// Text draws a text blob.
func (vr *vectorRenderer) Text(body string, x, y int) {
	vr.s.FillColor = color.RGBA{}
	vr.s.StrokeColor = color.RGBA{}
	vr.s.StrokeWidth = 0
	vr.c.Text(x, y, body, vr.s.SVG()+";"+vr.svgFontFace())
}

// MeasureText uses the truetype font drawer to measure the width of text.
func (vr *vectorRenderer) MeasureText(body string) int {
	if vr.fc == nil && vr.f != nil {
		vr.fc = &font.Drawer{
			Face: truetype.NewFace(vr.f, &truetype.Options{
				DPI:  DefaultDPI,
				Size: vr.s.FontSize,
			}),
		}
	}
	if vr.fc != nil {
		dimensions := vr.fc.MeasureString(body)
		return dimensions.Floor()
	}
	return 0
}

func (vr *vectorRenderer) Save(w io.Writer) error {
	vr.c.End()
	_, err := w.Write(vr.b.Bytes())
	return err
}
