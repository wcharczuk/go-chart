package main

import (
	"net/http"

	chart "github.com/wcharczuk/go-chart"
)

func drawChart(res http.ResponseWriter, req *http.Request) {
	/*
	   In this example we set a custom set of ticks to use for the y-axis. It can be (almost) whatever you want, including some custom labels for ticks.
	   Custom ticks will supercede a custom range, which will supercede automatic generation based on series values.
	*/

	graph := chart.Chart{
		YAxis: chart.YAxis{
			Style: chart.StyleShow(),
			Range: &chart.ContinuousRange{
				Min: 0.0,
				Max: 4.0,
			},
			Ticks: []chart.Tick{
				{Value: 0.0, Label: "0.00"},
				{Value: 2.0, Label: "2.00"},
				{Value: 4.0, Label: "4.00"},
				{Value: 6.0, Label: "6.00"},
				{Value: 8.0, Label: "Eight"},
				{Value: 10.0, Label: "Ten"},
			},
		},
		Series: []chart.Series{
			chart.ContinuousSeries{
				XValues: []float64{1.0, 2.0, 3.0, 4.0, 5.0},
				YValues: []float64{1.0, 2.0, 3.0, 4.0, 5.0},
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
