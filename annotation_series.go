package chart

// Annotation is a label on the chart.
type Annotation struct {
	X, Y  float64
	Label string
}

// AnnotationSeries is a series of labels on the chart.
type AnnotationSeries struct {
	Name        string
	Style       Style
	YAxis       YAxisType
	Annotations []Annotation
}

// Render draws the series.
func (as AnnotationSeries) Render(r Renderer, canvasBox Box, xrange, yrange Range) {
	if as.Style.Show {
		for _, a := range as.Annotations {
			lx := xrange.Translate(a.X) + canvasBox.Left
			ly := yrange.Translate(a.Y) + canvasBox.Top
			DrawAnnotation(r, canvasBox, xrange, yrange, as.Style, lx, ly, a.Label)
		}
	}
}
