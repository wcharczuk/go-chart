package main

//go:generate go run main.go

import (
	"os"

	"github.com/wcharczuk/go-chart"
)

func main() {
	mainSeries := chart.ContinuousSeries{
		Name:    "A test series",
		XValues: chart.Seq{Sequence: chart.NewLinearSequence().WithStart(1.0).WithEnd(100.0)}.Values(),
		YValues: chart.Seq{Sequence: chart.NewRandomSequence().WithLen(100).WithMin(50).WithMax(150)}.Values(),
	}

	minSeries := &chart.MinSeries{
		Style: chart.Style{
			StrokeColor:     chart.ColorAlternateGray,
			StrokeDashArray: []float64{5.0, 5.0},
		},
		InnerSeries: mainSeries,
	}

	maxSeries := &chart.MaxSeries{
		Style: chart.Style{
			StrokeColor:     chart.ColorAlternateGray,
			StrokeDashArray: []float64{5.0, 5.0},
		},
		InnerSeries: mainSeries,
	}

	graph := chart.Chart{
		Width:  1920,
		Height: 1080,
		YAxis: chart.YAxis{
			Name: "Random Values",
			Range: &chart.ContinuousRange{
				Min: 25,
				Max: 175,
			},
		},
		XAxis: chart.XAxis{
			Name: "Random Other Values",
		},
		Series: []chart.Series{
			mainSeries,
			minSeries,
			maxSeries,
			chart.LastValueAnnotationSeries(minSeries),
			chart.LastValueAnnotationSeries(maxSeries),
		},
	}

	graph.Elements = []chart.Renderable{chart.Legend(&graph)}

	f, _ := os.Create("output.png")
	defer f.Close()
	graph.Render(chart.PNG, f)
}
