package main

import (
	"net/http"

	"github.com/wcharczuk/go-chart"
)

func drawChart(res http.ResponseWriter, req *http.Request) {
	/*
	   In this example we set a rotation on the style for the custom ticks from the `custom_ticks` example.
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
			TickStyle: chart.Style{
				TextRotationDegrees: 45.0,
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
		XAxis: chart.XAxis{
			Style: chart.Style{
				Show: true,
			},
			TickStyle: chart.Style{
				TextRotationDegrees: 45.0,
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
	graph.Elements = []chart.Renderable{
		func(r chart.Renderer, cb chart.Box, defaults chart.Style) {

			b := chart.Box{50, 50, 90, 110}
			chart.Draw.Box(r, b, chart.Style{
				StrokeWidth: 2,
				StrokeColor: chart.ColorBlue,
			})
			chart.Draw.Box(r, b.Rotate(chart.Math.DegreesToRadians(45)), chart.Style{
				StrokeWidth: 2,
				StrokeColor: chart.ColorRed,
			})
		},
	}
	graph.Render(chart.PNG, res)
}

func main() {
	http.HandleFunc("/", drawChart)
	http.ListenAndServe(":8080", nil)
}
