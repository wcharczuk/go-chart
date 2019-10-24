package chart

import (
	"github.com/wcharczuk/go-chart/drawing"
)

// DetailsBox adds a box with additional text
func DetailsBox(c *Chart, text []string, userDefaults ...Style) Renderable {
	return func(r Renderer, box Box, chartDefaults Style) {
		// default style
		defaults := Style{
			FillColor:   drawing.ColorWhite,
			FontColor:   DefaultTextColor,
			FontSize:    8.0,
			StrokeColor: DefaultAxisColor,
			StrokeWidth: DefaultAxisLineWidth,
		}

		var style Style
		if len(userDefaults) > 0 {
			style = userDefaults[0].InheritFrom(chartDefaults.InheritFrom(defaults))
		} else {
			style = chartDefaults.InheritFrom(defaults)
		}

		contentPadding := Box{
			Top:    box.Height(),
			Left:   5,
			Right:  5,
			Bottom: box.Height(),
		}

		contentBox := Box{
			Bottom: box.Height(),
			Left:   5,
		}

		content := Box{
			Top:    contentBox.Bottom - 5,
			Left:   contentBox.Left + contentPadding.Left,
			Right:  contentBox.Left + contentPadding.Left,
			Bottom: contentBox.Bottom - 5,
		}

		style.GetTextOptions().WriteToRenderer(r)

		// measure and add size of text to box height and width
		for _, t := range text {
			textbox := r.MeasureText(t)
			content.Top -= textbox.Height()
			right := content.Left + textbox.Width()
			content.Right = MaxInt(content.Right, right)
		}

		contentBox = contentBox.Grow(content)
		contentBox.Right = content.Right + contentPadding.Right
		contentBox.Top = content.Top - 5

		// draw the box
		Draw.Box(r, contentBox, style)

		style.GetTextOptions().WriteToRenderer(r)

		// add the text
		ycursor := content.Top
		x := content.Left
		for _, t := range text {
			textbox := r.MeasureText(t)
			y := ycursor + textbox.Height()
			r.Text(t, x, y)
			ycursor += textbox.Height()
		}
	}
}
