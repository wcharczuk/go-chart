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
		Values: []chart.PieChartValue{
			{Value: 0.3, Label: "Blue"},
			{Value: 0.2, Label: "Green"},
			{Value: 0.2, Label: "Gray"},
			{Value: 0.1, Label: "Orange"},
			{Value: 0.1, Label: "??"},
		},
	}

	res.Header().Set("Content-Type", "image/png")
	err := pie.Render(chart.PNG, res)
	if err != nil {
		fmt.Printf("Error rendering pie chart: %v\n", err)
	}
}

func main() {
	http.HandleFunc("/", drawChart)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
