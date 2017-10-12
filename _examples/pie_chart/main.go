package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/wcharczuk/go-chart"
)

func drawChart(res http.ResponseWriter, req *http.Request) {
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

	res.Header().Set("Content-Type", "image/png")
	err := pie.Render(chart.PNG, res)
	if err != nil {
		fmt.Printf("Error rendering pie chart: %v\n", err)
	}
}

func drawChartRegression(res http.ResponseWriter, req *http.Request) {
	pie := chart.PieChart{
		Width:  512,
		Height: 512,
		Values: []chart.Value{
			{Value: 5, Label: "Blue"},
			{Value: 2, Label: "Two"},
			{Value: 1, Label: "One"},
		},
	}

	res.Header().Set("Content-Type", chart.ContentTypeSVG)
	err := pie.Render(chart.SVG, res)
	if err != nil {
		fmt.Printf("Error rendering pie chart: %v\n", err)
	}
}

func main() {
	http.HandleFunc("/", drawChart)
	http.HandleFunc("/reg", drawChartRegression)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
