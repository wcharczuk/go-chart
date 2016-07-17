package main

import (
	"net/http"

	"github.com/wcharczuk/go-chart"
)

func drawChart(res http.ResponseWriter, req *http.Request) {
	xvalues, yvalues := getMockChartData()

	priceSeries := chart.TimeSeries{
		XValues: xvalues,
		YValues: yvalues,
	}

	graph := chart.Chart{
		Series: []chart.Series{
			priceSeries,
		},
	}

	res.Header().Set("Content-Type", "image/svg+xml")
	graph.Render(chart.SVG, res)
}

func main() {
	http.HandleFunc("/", drawChart)
	http.ListenAndServe(":8080", nil)
}
