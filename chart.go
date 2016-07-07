package chart

import (
	"image/color"
	"io"

	"github.com/golang/freetype/truetype"
)

// Chart is what we're drawing.
/*
	The chart box model is as follows:
	0,0									width,0
			cl,ct					cr,ct
			cl,cb					cr,cb
	0, height							width,height
*/
type Chart struct {
	Title         string
	TitleFontSize float64

	Width   int
	Height  int
	Padding int

	BackgroundColor       color.RGBA
	CanvasBackgroundColor color.RGBA

	AxisShow     bool
	AxisStyle    Style
	AxisFontSize float64

	CanvasBorderShow  bool
	CanvasBorderStyle Style

	FinalValueLabelShow bool
	FinalValueStyle     Style

	FontColor color.RGBA
	Font      *truetype.Font

	Series []Series
}

// GetTitleFontSize calculates or returns the title font size.
func (c Chart) GetTitleFontSize() float64 {
	if c.TitleFontSize != 0 {
		if c.TitleFontSize > DefaultMinimumFontSize {
			return c.TitleFontSize
		}
	}
	fontSize := float64(c.Height >> 3)
	if fontSize > DefaultMinimumFontSize {
		return fontSize
	}
	return DefaultMinimumFontSize
}

// GetCanvasTop gets the top corner pixel.
func (c Chart) GetCanvasTop() int {
	return c.Padding
}

// GetCanvasLeft gets the left corner pixel.
func (c Chart) GetCanvasLeft() int {
	return c.Padding
}

// GetCanvasBottom gets the bottom corner pixel.
func (c Chart) GetCanvasBottom() int {
	return c.Height - c.Padding
}

// GetCanvasRight gets the right corner pixel.
func (c Chart) GetCanvasRight() int {
	return c.Width - c.Padding
}

// GetCanvasWidth returns the width of the canvas.
func (c Chart) GetCanvasWidth() int {
	if c.Padding > 0 {
		return c.Width - (c.Padding << 1)
	}
	return c.Width
}

// GetCanvasHeight returns the height of the canvas.
func (c Chart) GetCanvasHeight() int {
	if c.Padding > 0 {
		return c.Height - (c.Padding << 1)
	}
	return c.Height
}

// GetBackgroundColor returns the chart background color.
func (c Chart) GetBackgroundColor() color.RGBA {
	if ColorIsZero(c.BackgroundColor) {
		c.BackgroundColor = DefaultBackgroundColor
	}
	return c.BackgroundColor
}

// GetCanvasBackgroundColor returns the canvas background color.
func (c Chart) GetCanvasBackgroundColor() color.RGBA {
	if ColorIsZero(c.CanvasBackgroundColor) {
		c.CanvasBackgroundColor = DefaultCanvasColor
	}
	return c.CanvasBackgroundColor
}

// GetTextFont returns the text font.
func (c Chart) GetTextFont() (*truetype.Font, error) {
	if c.Font != nil {
		return c.Font, nil
	}
	return GetDefaultFont()
}

// GetTextFontColor returns the text font color.
func (c Chart) GetTextFontColor() color.RGBA {
	if ColorIsZero(c.FontColor) {
		c.FontColor = DefaultTextColor
	}
	return c.FontColor
}

// Render renders the chart with the given renderer to the given io.Writer.
func (c Chart) Render(provider RendererProvider, w io.Writer) error {
	r := provider(c.Width, c.Height)
	c.drawBackground(r)
	c.drawCanvas(r)
	c.drawAxes(r)

	for _, series := range c.Series {
		c.drawSeries(r, series)
	}
	err := c.drawTitle(r)
	if err != nil {
		return err
	}
	return r.Save(w)
}

func (c Chart) drawBackground(r Renderer) {
	r.SetStrokeColor(c.GetBackgroundColor())
	r.SetFillColor(c.GetBackgroundColor())
	r.SetLineWidth(0)
	r.MoveTo(0, 0)
	r.LineTo(c.Width, 0)
	r.LineTo(c.Width, c.Height)
	r.LineTo(0, c.Height)
	r.LineTo(0, 0)
	r.FillStroke()
	r.Close()
}

func (c Chart) drawCanvas(r Renderer) {
	if !c.CanvasBorderStyle.IsZero() {
		r.SetStrokeColor(c.CanvasBorderStyle.GetStrokeColor())
		r.SetLineWidth(c.CanvasBorderStyle.GetStrokeWidth())
	} else {
		r.SetStrokeColor(c.GetCanvasBackgroundColor())
		r.SetLineWidth(0)
	}
	r.SetFillColor(c.GetCanvasBackgroundColor())
	r.MoveTo(c.GetCanvasLeft(), c.GetCanvasTop())
	r.LineTo(c.GetCanvasRight(), c.GetCanvasTop())
	r.LineTo(c.GetCanvasRight(), c.GetCanvasBottom())
	r.LineTo(c.GetCanvasLeft(), c.GetCanvasBottom())
	r.LineTo(c.GetCanvasLeft(), c.GetCanvasTop())
	r.FillStroke()
	r.Close()
}

func (c Chart) drawAxes(r Renderer) {
	if c.AxisShow {
		if !c.AxisStyle.IsZero() {
			r.SetStrokeColor(c.AxisStyle.GetStrokeColor())
			r.SetLineWidth(c.AxisStyle.GetStrokeWidth())
		} else {
			r.SetStrokeColor(DefaultAxisColor)
			r.SetLineWidth(DefaultLineWidth)
		}
		r.MoveTo(c.GetCanvasLeft(), c.GetCanvasBottom())
		r.LineTo(c.GetCanvasRight(), c.GetCanvasBottom())
		r.LineTo(c.GetCanvasRight(), c.GetCanvasTop())
		r.Stroke()
	}
}

func (c Chart) drawSeries(r Renderer, s Series) {

}

func (c Chart) drawTitle(r Renderer) error {
	if len(c.Title) > 0 {
		font, err := c.GetTextFont()
		if err != nil {
			return err
		}
		r.SetFontColor(c.GetTextFontColor())
		r.SetFont(font)
		r.SetFontSize(c.GetTitleFontSize())
		textWidth := r.MeasureText(c.Title)
		titleX := (c.Width >> 1) - (textWidth >> 1)
		titleY := c.GetCanvasTop() + int(c.GetTitleFontSize()/2.0)
		r.Text(c.Title, titleX, titleY)
	}
	return nil
}
