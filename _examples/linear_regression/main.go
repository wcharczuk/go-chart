package main

import (
	"net/http"

	"github.com/wcharczuk/go-chart"
	"github.com/wcharczuk/go-chart/seq"
)

func drawChart(res http.ResponseWriter, req *http.Request) {

	/*
		In this example we add a new type of series, a `SimpleMovingAverageSeries` that takes another series as a required argument.
		InnerSeries only needs to implement `ValuesProvider`, so really you could chain `SimpleMovingAverageSeries` together if you wanted.
	*/

	mainSeries := chart.ContinuousSeries{
		Name:    "A test series",
		XValues: seq.Range(1.0, 100.0),                 //generates a []float64 from 1.0 to 100.0 in 1.0 step increments, or 100 elements.
		YValues: seq.RandomValuesWithAverage(100, 100), //generates a []float64 randomly from 0 to 100 with 100 elements.
	}

	// note we create a LinearRegressionSeries series by assignin the inner series.
	// we need to use a reference because `.Render()` needs to modify state within the series.
	linRegSeries := &chart.LinearRegressionSeries{
		InnerSeries: mainSeries,
	} // we can optionally set the `WindowSize` property which alters how the moving average is calculated.

	graph := chart.Chart{
		Series: []chart.Series{
			mainSeries,
			linRegSeries,
		},
	}

	res.Header().Set("Content-Type", "image/png")
	graph.Render(chart.PNG, res)
}

func main() {
	http.HandleFunc("/", drawChart)
	http.ListenAndServe(":8080", nil)
}
