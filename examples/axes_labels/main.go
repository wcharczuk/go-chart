package main

//go:generate go run main.go

import (
	"os"

	chart "github.com/wcharczuk/go-chart"
)

func main() {

	/*
	   The below will draw the same chart as the `basic` example, except with both the x and y axes turned on.
	   In this case, both the x and y axis ticks are generated automatically, the x and y ranges are established automatically, the canvas "box" is adjusted to fit the space the axes occupy so as not to clip.
	*/

	graph := chart.Chart{
		XAxis: chart.XAxis{
			Name: "The XAxis",
		},
		YAxis: chart.YAxis{
			Name: "The YAxis",
		},
		Series: []chart.Series{
			chart.ContinuousSeries{
				Style: chart.Style{
					StrokeColor: chart.GetDefaultColor(0).WithAlpha(64),
					FillColor:   chart.GetDefaultColor(0).WithAlpha(64),
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
