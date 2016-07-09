package chart

import (
	"bytes"
	"fmt"
	"image/color"
	"io"
	"math"
	"strings"

	"golang.org/x/image/font"

	"github.com/golang/freetype/truetype"
)

// SVG returns a new png/raster renderer.
func SVG(width, height int) (Renderer, error) {
	buffer := bytes.NewBuffer([]byte{})
	canvas := newCanvas(buffer)
	canvas.Start(width, height)
	return &vectorRenderer{
		b: buffer,
		c: canvas,
		s: &Style{},
		p: []string{},
	}, nil
}

// vectorRenderer renders chart commands to a bitmap.
type vectorRenderer struct {
	dpi float64
	b   *bytes.Buffer
	c   *canvas
	s   *Style
	f   *truetype.Font
	p   []string
	fc  *font.Drawer
}

// SetDPI implements the interface method.
func (vr *vectorRenderer) SetDPI(dpi float64) {
	vr.dpi = dpi
}

// SetStrokeColor implements the interface method.
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
func (vr *vectorRenderer) MeasureText(body string) (width, height int) {
	if vr.f != nil {
		vr.fc = &font.Drawer{
			Face: truetype.NewFace(vr.f, &truetype.Options{
				DPI:  vr.dpi,
				Size: vr.s.FontSize,
			}),
		}
		width = vr.fc.MeasureString(body).Ceil()
		height = int(math.Ceil(vr.s.FontSize))
	}
	return
}

func (vr *vectorRenderer) Save(w io.Writer) error {
	vr.c.End()
	_, err := w.Write(vr.b.Bytes())
	return err
}

func newCanvas(w io.Writer) *canvas {
	return &canvas{
		w: w,
	}
}

type canvas struct {
	w      io.Writer
	width  int
	height int
}

func (c *canvas) Start(width, height int) {
	c.width = width
	c.height = height
	c.w.Write([]byte(fmt.Sprintf(`<svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" width="%d" height="%d">\n`, c.width, c.height)))
}

func (c *canvas) Path(d string, style ...string) {
	if len(style) > 0 {
		c.w.Write([]byte(fmt.Sprintf(`<path d="%s" style="%s"/>\n`, d, style[0])))
	} else {
		c.w.Write([]byte(fmt.Sprintf(`<path d="%s"/>\n`, d)))
	}
}

func (c *canvas) Text(x, y int, body string, style ...string) {
	if len(style) > 0 {
		c.w.Write([]byte(fmt.Sprintf(`<text x="%d" y="%d" style="%s">%s</text>`, x, y, style[0], body)))
	} else {
		c.w.Write([]byte(fmt.Sprintf(`<text x="%d" y="%d">%s</text>`, x, y, body)))
	}
}

func (c *canvas) Circle(x, y, r int, style ...string) {
	if len(style) > 0 {
		c.w.Write([]byte(fmt.Sprintf(`<circle cx="%d" cy="%d" r="%d" style="%s">`, x, y, r, style[0])))
	} else {
		c.w.Write([]byte(fmt.Sprintf(`<circle cx="%d" cy="%d" r="%d">`, x, y, r)))
	}
}

func (c *canvas) End() {
	c.w.Write([]byte("</svg>"))
}
