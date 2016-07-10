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

// GetName returns the name of the time series.
func (as AnnotationSeries) GetName() string {
	return as.Name
}

// GetStyle returns the line style.
func (as AnnotationSeries) GetStyle() Style {
	return as.Style
}

// GetYAxis returns which YAxis the series draws on.
func (as AnnotationSeries) GetYAxis() YAxisType {
	return as.YAxis
}

// Render draws the series.
func (as AnnotationSeries) Render(r Renderer, canvasBox Box, xrange, yrange Range, defaults Style) {
	if as.Style.Show {
		style := as.Style.WithDefaultsFrom(Style{
			FillColor:   DefaultAnnotationFillColor,
			FontSize:    DefaultAnnotationFontSize,
			StrokeColor: defaults.StrokeColor,
			StrokeWidth: defaults.StrokeWidth,
			Padding:     DefaultAnnotationPadding,
		})
		for _, a := range as.Annotations {
			lx := canvasBox.Right - xrange.Translate(a.X)
			ly := yrange.Translate(a.Y) + canvasBox.Top
			DrawAnnotation(r, canvasBox, xrange, yrange, style, lx, ly, a.Label)
		}
	}
}
