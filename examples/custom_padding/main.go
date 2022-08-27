package main

//go:generate go run main.go

import (
	"github.com/d-Rickyy-b/go-chart-x/v2/pkg/chart"
	"github.com/d-Rickyy-b/go-chart-x/v2/pkg/drawing"
	"os"
)

func main() {
	graph := chart.Chart{
		Background: chart.Style{
			Padding: chart.Box{
				Top:    50,
				Left:   25,
				Right:  25,
				Bottom: 10,
			},
			FillColor: drawing.ColorFromHex("efefef"),
		},
		Series: []chart.Series{
			chart.ContinuousSeries{
				XValues: chart.Seq{Sequence: chart.NewLinearSequence().WithStart(1.0).WithEnd(100.0)}.Values(),
				YValues: chart.Seq{Sequence: chart.NewRandomSequence().WithLen(100).WithMin(100).WithMax(512)}.Values(),
			},
		},
	}

	f, _ := os.Create("output.png")
	defer f.Close()
	graph.Render(chart.PNG, f)
}
