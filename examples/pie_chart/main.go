package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/wcharczuk/go-chart"
)

func drawChart(res http.ResponseWriter, req *http.Request) {
	pie := chart.PieChart{
		Canvas: chart.Style{
			FillColor: chart.ColorLightGray,
		},
		Values: []chart.Value{
			{Value: 5, Label: "Blue"},
			{Value: 5, Label: "Green"},
			{Value: 4, Label: "Gray"},
			{Value: 4, Label: "Orange"},
			{Value: 3, Label: "Test"},
			{Value: 3, Label: "??"},
			{Value: 1, Label: "!!"},
		},
	}

	res.Header().Set("Content-Type", "image/svg+xml")
	err := pie.Render(chart.SVG, res)
	if err != nil {
		fmt.Printf("Error rendering pie chart: %v\n", err)
	}
}

func main() {
	http.HandleFunc("/", drawChart)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
