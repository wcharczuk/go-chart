package main

//go:generate go run main.go

import (
	"os"

	"github.com/wcharczuk/go-chart/v2"
	"github.com/wcharczuk/go-chart/v2/drawing"
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
                                        StrokeMaxSpanGap: 2, // do not span gaps more than 2.0 wide
				},
				XValues: []float64{1.0, 2.0, 3.0, 4.0, 5.0, 8.0, 10.0},
				YValues: []float64{1.0, 2.0, 3.0, 4.0, 5.0, 8.0, 10.0},
			},
		},
	}

	f, _ := os.Create("output.png")
	defer f.Close()
	graph.Render(chart.PNG, f)
}
