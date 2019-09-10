package main

//go:generate go run main.go

import (
	"os"

	"github.com/wcharczuk/go-chart"
)

func main() {
	mainSeries := chart.ContinuousSeries{
		Name:    "A test series",
		XValues: chart.Seq{Sequence: chart.NewLinearSequence().WithStart(1.0).WithEnd(100.0)}.Values(),        //generates a []float64 from 1.0 to 100.0 in 1.0 step increments, or 100 elements.
		YValues: chart.Seq{Sequence: chart.NewRandomSequence().WithLen(100).WithMin(0).WithMax(100)}.Values(), //generates a []float64 randomly from 0 to 100 with 100 elements.
	}

	// note we create a SimpleMovingAverage series by assignin the inner series.
	// we need to use a reference because `.Render()` needs to modify state within the series.
	smaSeries := &chart.SMASeries{
		InnerSeries: mainSeries,
	} // we can optionally set the `WindowSize` property which alters how the moving average is calculated.

	graph := chart.Chart{
		Series: []chart.Series{
			mainSeries,
			smaSeries,
		},
	}

	f, _ := os.Create("output.png")
	defer f.Close()
	graph.Render(chart.PNG, f)
}
