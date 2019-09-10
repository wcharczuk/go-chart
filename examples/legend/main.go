package main

//go:generate go run main.go

import (
	"os"

	chart "github.com/wcharczuk/go-chart"
)

func main() {

	/*
	   In this example we add a `Renderable` or a custom component to the `Elements` array.
	   In this specific case it is a pre-built renderable (`CreateLegend`) that draws a legend for the chart's series.
	   If you like, you can use `CreateLegend` as a template for writing your own renderable, or even your own legend.
	*/

	graph := chart.Chart{
		Background: chart.Style{
			Padding: chart.Box{
				Top:  20,
				Left: 20,
			},
		},
		Series: []chart.Series{
			chart.ContinuousSeries{
				Name:    "A test series",
				XValues: []float64{1.0, 2.0, 3.0, 4.0, 5.0},
				YValues: []float64{1.0, 2.0, 3.0, 4.0, 5.0},
			},
		},
	}

	//note we have to do this as a separate step because we need a reference to graph
	graph.Elements = []chart.Renderable{
		chart.Legend(&graph),
	}

	f, _ := os.Create("output.png")
	defer f.Close()
	graph.Render(chart.PNG, f)
}
