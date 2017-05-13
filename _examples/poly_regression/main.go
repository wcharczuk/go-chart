package main

import (
	"net/http"

	"github.com/wcharczuk/go-chart"
	"github.com/wcharczuk/go-chart/seq"
)

func drawChart(res http.ResponseWriter, req *http.Request) {

	/*
		In this example we add a new type of series, a `PolynomialRegressionSeries` that takes another series as a required argument.
		InnerSeries only needs to implement `ValuesProvider`, so really you could chain `PolynomialRegressionSeries` together if you wanted.
	*/

	mainSeries := chart.ContinuousSeries{
		Name:    "A test series",
		XValues: seq.Range(1.0, 100.0),                 //generates a []float64 from 1.0 to 100.0 in 1.0 step increments, or 100 elements.
		YValues: seq.RandomValuesWithAverage(100, 100), //generates a []float64 randomly from 0 to 100 with 100 elements.
	}

	polyRegSeries := &chart.PolynomialRegressionSeries{
		Degree:      3,
		InnerSeries: mainSeries,
	}

	graph := chart.Chart{
		Series: []chart.Series{
			mainSeries,
			polyRegSeries,
		},
	}

	res.Header().Set("Content-Type", "image/png")
	graph.Render(chart.PNG, res)
}

func main() {
	http.HandleFunc("/", drawChart)
	http.ListenAndServe(":8080", nil)
}
