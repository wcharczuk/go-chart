package main

//go:generate go run main.go

import (
	"github.com/d-Rickyy-b/go-chart-x/v2/pkg/chart"
	"os"
)

func main() {
	pie := chart.PieChart{
		Width:  512,
		Height: 512,
		Values: []chart.Value{
			{Value: 5, Label: "Blue"},
			{Value: 5, Label: "Green"},
			{Value: 4, Label: "Gray"},
			{Value: 4, Label: "Orange"},
			{Value: 3, Label: "Deep Blue"},
			{Value: 3, Label: "??"},
			{Value: 1, Label: "!!"},
		},
	}

	f, _ := os.Create("output.png")
	defer f.Close()
	pie.Render(chart.PNG, f)
}
