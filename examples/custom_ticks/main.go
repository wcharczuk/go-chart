package main

import (
	"net/http"

	"github.com/wcharczuk/go-chart"
)

func drawChart(res http.ResponseWriter, req *http.Request) {
	/*
	   In this example we set a custom set of ticks to use for the y-axis. It can be (almost) whatever you want, including some custom labels for ticks.
	   Custom ticks will supercede a custom range, which will supercede automatic generation based on series values.
	*/

	graph := chart.Chart{
		YAxis: chart.YAxis{
			Style: chart.Style{
				Show: true,
			},
			Range: &chart.ContinuousRange{
				Min: 0.0,
				Max: 4.0,
			},
			Ticks: []chart.Tick{
				{0.0, "0.00"},
				{2.0, "2.00"},
				{4.0, "4.00"},
				{6.0, "6.00"},
				{8.0, "Eight"},
				{10.0, "Ten"},
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
