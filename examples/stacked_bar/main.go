package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/wcharczuk/go-chart"
)

func drawChart(res http.ResponseWriter, req *http.Request) {
	sbc := chart.StackedBarChart{
		Background: chart.Style{
			Padding: chart.Box{Top: 50, Left: 50, Right: 50, Bottom: 50},
		},
		XAxis: chart.Style{
			Show: true,
		},
		YAxis: chart.Style{
			Show: true,
		},
		Bars: []chart.StackedBar{
			{
				Name: "Funnel",
				Values: []chart.Value{
					{Value: 5, Label: "Blue"},
					{Value: 5, Label: "Green"},
					{Value: 4, Label: "Gray"},
					{Value: 4, Label: "Orange"},
					{Value: 3, Label: "Test"},
					{Value: 3, Label: "??"},
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
		},
	}

	res.Header().Set("Content-Type", "image/svg+xml")
	err := sbc.Render(chart.SVG, res)
	if err != nil {
		fmt.Printf("Error rendering chart: %v\n", err)
	}
}

func main() {
	http.HandleFunc("/", drawChart)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
