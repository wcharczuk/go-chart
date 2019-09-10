package main

import (
	"os"

	"github.com/wcharczuk/go-chart"
)

func main() {
	sbc := chart.StackedBarChart{
		Title: "Test Stacked Bar Chart",
		Background: chart.Style{
			Padding: chart.Box{
				Top: 40,
			},
		},
		Height: 512,
		Bars: []chart.StackedBar{
			{
				Name: "This is a very long string to test word break wrapping.",
				Values: []chart.Value{
					{Value: 5, Label: "Blue"},
					{Value: 5, Label: "Green"},
					{Value: 4, Label: "Gray"},
					{Value: 3, Label: "Orange"},
					{Value: 3, Label: "Test"},
					{Value: 2, Label: "??"},
					{Value: 1, Label: "!!"},
				},
			},
			{
				Name: "Test",
				Values: []chart.Value{
					{Value: 10, Label: "Blue"},
					{Value: 5, Label: "Green"},
					{Value: 1, Label: "Gray"},
				},
			},
			{
				Name: "Test 2",
				Values: []chart.Value{
					{Value: 10, Label: "Blue"},
					{Value: 6, Label: "Green"},
					{Value: 4, Label: "Gray"},
				},
			},
		},
	}

	f, _ := os.Create("output.png")
	defer f.Close()
	sbc.Render(chart.PNG, f)
}
