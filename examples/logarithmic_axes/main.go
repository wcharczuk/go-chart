package main

//go:generate go run main.go

import (
	"os"

	"github.com/wcharczuk/go-chart"
)

func main() {

	/*
	   In this example we set the primary YAxis to have logarithmic range.
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
				YValues: []float64{1, 10, 100, 1000, 10000},
			},
		},
		YAxis: chart.YAxis{
			Style:     chart.Shown(),
			NameStyle: chart.Shown(),
			Range: &chart.LogarithmicRange{},
		},
	}

	f, _ := os.Create("output.png")
	defer f.Close()
	graph.Render(chart.PNG, f)
}
