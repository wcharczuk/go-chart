package main

//go:generate go run main.go

import (
	"fmt"
	"os"

	"github.com/wcharczuk/go-chart"
)

func main() {
	/*
	   In this example we use a custom `ValueFormatter` for the y axis, letting us specify how to format text of the y-axis ticks.
	   You can also do this for the x-axis, or the secondary y-axis.
	   This example also shows what the chart looks like with the x-axis left off or not shown.
	*/

	graph := chart.Chart{
		YAxis: chart.YAxis{
			ValueFormatter: func(v interface{}) string {
				if vf, isFloat := v.(float64); isFloat {
					return fmt.Sprintf("%0.6f", vf)
				}
				return ""
			},
		},
		Series: []chart.Series{
			chart.ContinuousSeries{
				XValues: []float64{1.0, 2.0, 3.0, 4.0, 5.0},
				YValues: []float64{1.0, 2.0, 3.0, 4.0, 5.0},
			},
		},
	}
	f, _ := os.Create("output.png")
	defer f.Close()
	graph.Render(chart.PNG, f)
}
