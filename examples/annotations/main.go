package main

import (
	"net/http"

	"github.com/wcharczuk/go-chart"
)

func drawChart(res http.ResponseWriter, req *http.Request) {

	/*
	   In this example we add an `Annotation` series, which is a special type of series that
	   draws annotation labels at given X and Y values (as translated by their respective ranges).

	   It is important to not that the chart automatically sizes the canvas box to fit the annotations,
	   As well as automatically assign a series color for the `Stroke` or border component of the series.

	   The annotation series is most often used by the original author to show the last value of another series, but
	   they can be used in other capacities as well.
	*/

	graph := chart.Chart{
		Series: []chart.Series{
			chart.ContinuousSeries{
				XValues: []float64{1.0, 2.0, 3.0, 4.0, 5.0},
				YValues: []float64{1.0, 2.0, 3.0, 4.0, 5.0},
			},
			chart.AnnotationSeries{
				Annotations: []chart.Annotation{
					{X: 1.0, Y: 1.0, Label: "One"},
					{X: 2.0, Y: 2.0, Label: "Two"},
					{X: 3.0, Y: 3.0, Label: "Three"},
					{X: 4.0, Y: 4.0, Label: "Four"},
					{X: 5.0, Y: 5.0, Label: "Five"},
				},
			},
		},
	}

	res.Header().Set("Content-Type", "image/png")
	graph.Render(chart.PNG, res)
}

func main() {
	http.HandleFunc("/", drawChart)
	http.ListenAndServe(":8080", nil)
}
