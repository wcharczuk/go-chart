package chart

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"math"

	"github.com/golang/freetype/truetype"
	"github.com/wcharczuk/go-chart/drawing"
)

// PNG returns a new png/raster renderer.
func PNG(width, height int) (Renderer, error) {
	i := image.NewRGBA(image.Rect(0, 0, width, height))
	gc, err := drawing.NewRasterGraphicContext(i)
	if err == nil {
		return &rasterRenderer{
			i:  i,
			gc: gc,
		}, nil
	}
	return nil, err
}

// rasterRenderer renders chart commands to a bitmap.
type rasterRenderer struct {
	i  *image.RGBA
	gc *drawing.RasterGraphicContext

	fontSize  float64
	fontColor color.RGBA
	f         *truetype.Font
}

// SetDPI implements the interface method.
func (rr *rasterRenderer) SetDPI(dpi float64) {
	rr.gc.SetDPI(dpi)
}

// SetStrokeColor implements the interface method.
func (rr *rasterRenderer) SetStrokeColor(c color.RGBA) {
	rr.gc.SetStrokeColor(c)
}

// SetFillColor implements the interface method.
func (rr *rasterRenderer) SetFillColor(c color.RGBA) {
	rr.gc.SetFillColor(c)
}

// SetLineWidth implements the interface method.
func (rr *rasterRenderer) SetStrokeWidth(width float64) {
	rr.gc.SetLineWidth(width)
}

// MoveTo implements the interface method.
func (rr *rasterRenderer) MoveTo(x, y int) {
	rr.gc.MoveTo(float64(x), float64(y))
}

// LineTo implements the interface method.
func (rr *rasterRenderer) LineTo(x, y int) {
	rr.gc.LineTo(float64(x), float64(y))
}

// Close implements the interface method.
func (rr *rasterRenderer) Close() {
	rr.gc.Close()
}

// Stroke implements the interface method.
func (rr *rasterRenderer) Stroke() {
	rr.gc.Stroke()
}

// Fill implements the interface method.
func (rr *rasterRenderer) Fill() {
	rr.gc.Fill()
}

// FillStroke implements the interface method.
func (rr *rasterRenderer) FillStroke() {
	rr.gc.FillStroke()
}

// Circle implements the interface method.
func (rr *rasterRenderer) Circle(radius float64, x, y int) {
	xf := float64(x)
	yf := float64(y)
	rr.gc.MoveTo(xf-radius, yf)              //9
	rr.gc.QuadCurveTo(xf, yf, xf, yf-radius) //12
	rr.gc.QuadCurveTo(xf, yf, xf+radius, yf) //3
	rr.gc.QuadCurveTo(xf, yf, xf, yf+radius) //6
	rr.gc.QuadCurveTo(xf, yf, xf-radius, yf) //9
	rr.gc.Close()
	rr.gc.FillStroke()
}

// SetFont implements the interface method.
func (rr *rasterRenderer) SetFont(f *truetype.Font) {
	rr.f = f
	rr.gc.SetFont(f)
}

// SetFontSize implements the interface method.
func (rr *rasterRenderer) SetFontSize(size float64) {
	rr.fontSize = size
	rr.gc.SetFontSize(size)
}

// SetFontColor implements the interface method.
func (rr *rasterRenderer) SetFontColor(c color.RGBA) {
	rr.fontColor = c
	rr.gc.SetFillColor(c)
	rr.gc.SetStrokeColor(c)
}

// Text implements the interface method.
func (rr *rasterRenderer) Text(body string, x, y int) {
	rr.gc.CreateStringPath(body, float64(x), float64(y))
	rr.gc.Fill()
}

// MeasureText returns the height and width in pixels of a string.
func (rr *rasterRenderer) MeasureText(body string) (width int, height int) {
	l, t, r, b, err := rr.gc.GetStringBounds(body)
	if err != nil {
		return
	}
	dw := r - l
	dh := b - t
	width = int(math.Ceil(dw * (4.0 / 3.0)))
	height = int(math.Ceil(dh * (4.0 / 3.0)))
	return
}

// Save implements the interface method.
func (rr *rasterRenderer) Save(w io.Writer) error {
	return png.Encode(w, rr.i)
}
