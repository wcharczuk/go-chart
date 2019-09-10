package main

//go:generate go run main.go

import (
	"os"

	"github.com/wcharczuk/go-chart"
	"github.com/wcharczuk/go-chart/drawing"
)

func main() {
	/*
	   In this example we set some custom colors for the series and the chart background and canvas.
	*/
	graph := chart.Chart{
		Background: chart.Style{
			FillColor: drawing.ColorBlue,
		},
		Canvas: chart.Style{
			FillColor: drawing.ColorFromHex("efefef"),
		},
		Series: []chart.Series{
			chart.ContinuousSeries{
				Style: chart.Style{
					StrokeColor: drawing.ColorRed,               // will supercede defaults
					FillColor:   drawing.ColorRed.WithAlpha(64), // will supercede defaults
				},
				XValues: []float64{1.0, 2.0, 3.0, 4.0, 5.0},
				YValues: []float64{1.0, 2.0, 3.0, 4.0, 5.0},
			},
		},
	}

	f, _ := os.Create("output.png")
	defer f.Close()
	graph.Render(chart.PNG, f)
}
