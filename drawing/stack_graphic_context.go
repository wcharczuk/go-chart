package drawing

import (
	"image"
	"image/color"

	"github.com/golang/freetype/truetype"
)

// StackGraphicContext is a context that does thngs.
type StackGraphicContext struct {
	current *ContextStack
}

// ContextStack is a graphic context implementation.
type ContextStack struct {
	Tr          Matrix
	Path        *Path
	LineWidth   float64
	Dash        []float64
	DashOffset  float64
	StrokeColor color.Color
	FillColor   color.Color
	FillRule    FillRule
	Cap         LineCap
	Join        LineJoin

	FontSize float64
	Font     *truetype.Font

	Scale float64

	Previous *ContextStack
}

// NewStackGraphicContext Create a new Graphic context from an image
func NewStackGraphicContext() *StackGraphicContext {
	gc := &StackGraphicContext{}
	gc.current = new(ContextStack)
	gc.current.Tr = NewIdentityMatrix()
	gc.current.Path = new(Path)
	gc.current.LineWidth = 1.0
	gc.current.StrokeColor = image.Black
	gc.current.FillColor = image.White
	gc.current.Cap = RoundCap
	gc.current.FillRule = FillRuleEvenOdd
	gc.current.Join = RoundJoin
	gc.current.FontSize = 10
	return gc
}

func (gc *StackGraphicContext) GetMatrixTransform() Matrix {
	return gc.current.Tr
}

func (gc *StackGraphicContext) SetMatrixTransform(Tr Matrix) {
	gc.current.Tr = Tr
}

func (gc *StackGraphicContext) ComposeMatrixTransform(Tr Matrix) {
	gc.current.Tr.Compose(Tr)
}

func (gc *StackGraphicContext) Rotate(angle float64) {
	gc.current.Tr.Rotate(angle)
}

func (gc *StackGraphicContext) Translate(tx, ty float64) {
	gc.current.Tr.Translate(tx, ty)
}

func (gc *StackGraphicContext) Scale(sx, sy float64) {
	gc.current.Tr.Scale(sx, sy)
}

func (gc *StackGraphicContext) SetStrokeColor(c color.Color) {
	gc.current.StrokeColor = c
}

func (gc *StackGraphicContext) SetFillColor(c color.Color) {
	gc.current.FillColor = c
}

func (gc *StackGraphicContext) SetFillRule(f FillRule) {
	gc.current.FillRule = f
}

func (gc *StackGraphicContext) SetLineWidth(lineWidth float64) {
	gc.current.LineWidth = lineWidth
}

func (gc *StackGraphicContext) SetLineCap(cap LineCap) {
	gc.current.Cap = cap
}

func (gc *StackGraphicContext) SetLineJoin(join LineJoin) {
	gc.current.Join = join
}

func (gc *StackGraphicContext) SetLineDash(dash []float64, dashOffset float64) {
	gc.current.Dash = dash
	gc.current.DashOffset = dashOffset
}

func (gc *StackGraphicContext) SetFontSize(fontSize float64) {
	gc.current.FontSize = fontSize
}

func (gc *StackGraphicContext) GetFontSize() float64 {
	return gc.current.FontSize
}

func (gc *StackGraphicContext) SetFont(f *truetype.Font) {
	gc.current.Font = f
}

func (gc *StackGraphicContext) GetFont() *truetype.Font {
	return gc.current.Font
}

func (gc *StackGraphicContext) BeginPath() {
	gc.current.Path.Clear()
}

func (gc *StackGraphicContext) IsEmpty() bool {
	return gc.current.Path.IsEmpty()
}

func (gc *StackGraphicContext) LastPoint() (float64, float64) {
	return gc.current.Path.LastPoint()
}

func (gc *StackGraphicContext) MoveTo(x, y float64) {
	gc.current.Path.MoveTo(x, y)
}

func (gc *StackGraphicContext) LineTo(x, y float64) {
	gc.current.Path.LineTo(x, y)
}

func (gc *StackGraphicContext) QuadCurveTo(cx, cy, x, y float64) {
	gc.current.Path.QuadCurveTo(cx, cy, x, y)
}

func (gc *StackGraphicContext) CubicCurveTo(cx1, cy1, cx2, cy2, x, y float64) {
	gc.current.Path.CubicCurveTo(cx1, cy1, cx2, cy2, x, y)
}

func (gc *StackGraphicContext) ArcTo(cx, cy, rx, ry, startAngle, angle float64) {
	gc.current.Path.ArcTo(cx, cy, rx, ry, startAngle, angle)
}

func (gc *StackGraphicContext) Close() {
	gc.current.Path.Close()
}

func (gc *StackGraphicContext) Save() {
	context := new(ContextStack)
	context.FontSize = gc.current.FontSize
	context.Font = gc.current.Font
	context.LineWidth = gc.current.LineWidth
	context.StrokeColor = gc.current.StrokeColor
	context.FillColor = gc.current.FillColor
	context.FillRule = gc.current.FillRule
	context.Dash = gc.current.Dash
	context.DashOffset = gc.current.DashOffset
	context.Cap = gc.current.Cap
	context.Join = gc.current.Join
	context.Path = gc.current.Path.Copy()
	context.Font = gc.current.Font
	context.Scale = gc.current.Scale
	copy(context.Tr[:], gc.current.Tr[:])
	context.Previous = gc.current
	gc.current = context
}

func (gc *StackGraphicContext) Restore() {
	if gc.current.Previous != nil {
		oldContext := gc.current
		gc.current = gc.current.Previous
		oldContext.Previous = nil
	}
}
