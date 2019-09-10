package main

//go:generate go run main.go

import (
	"os"

	"github.com/wcharczuk/go-chart"
)

func main() {
	pie := chart.DonutChart{
		Width:  512,
		Height: 512,
		Values: []chart.Value{
			{Value: 5, Label: "Blue"},
			{Value: 5, Label: "Green"},
			{Value: 4, Label: "Gray"},
			{Value: 4, Label: "Orange"},
			{Value: 3, Label: "Deep Blue"},
			{Value: 3, Label: "test"},
		},
	}

	f, _ := os.Create("output.png")
	defer f.Close()
	pie.Render(chart.PNG, f)
}
