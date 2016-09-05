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
				{Value: 0.0, Label: "0.00"},
				{Value: 2.0, Label: "2.00"},
				{Value: 4.0, Label: "4.00"},
				{Value: 6.0, Label: "6.00"},
				{Value: 8.0, Label: "Eight"},
				{Value: 10.0, Label: "Ten"},
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
	graph.Elements = []chart.Renderable{
		func(r chart.Renderer, cb chart.Box, defaults chart.Style) {

			b := chart.Box{Top: 50, Left: 50, Right: 150, Bottom: 300}

			cx, cy := b.Center()

			chart.Draw.Box(r, chart.Box{Top: cy - 2, Left: cx - 2, Right: cx + 2, Bottom: cy + 2}, chart.Style{
				StrokeWidth: 2,
				StrokeColor: chart.ColorBlack,
			})

			chart.Draw.Box(r, b, chart.Style{
				StrokeWidth: 2,
				StrokeColor: chart.ColorBlue,
			})
			chart.Draw.Box(r, b.BoundedRotate(chart.Math.DegreesToRadians(45)), chart.Style{
				StrokeWidth: 2,
				StrokeColor: chart.ColorRed,
			})

			chart.Draw.BoxRotated(r, b, chart.Math.DegreesToRadians(45), chart.Style{
				StrokeWidth: 2,
				StrokeColor: chart.ColorOrange,
			})
		},
	}
	graph.Render(chart.PNG, res)
}

func main() {
	http.HandleFunc("/", drawChart)
	http.ListenAndServe(":8080", nil)
}
