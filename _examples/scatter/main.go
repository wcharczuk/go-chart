package main

import (
	"log"
	"net/http"

	"github.com/wcharczuk/go-chart"
)

func drawChart(res http.ResponseWriter, req *http.Request) {
	graph := chart.Chart{
		Series: []chart.Series{
			chart.ContinuousSeries{
				Style: chart.Style{
					Show:        true,
					StrokeWidth: chart.Disabled,
					DotWidth:    3,
				},
				XValues: chart.Sequence.Random(32, 1024),
				YValues: chart.Sequence.Random(32, 1024),
			},
			chart.ContinuousSeries{
				Style: chart.Style{
					Show:        true,
					StrokeWidth: chart.Disabled,
					DotWidth:    5,
				},
				XValues: chart.Sequence.Random(16, 1024),
				YValues: chart.Sequence.Random(16, 1024),
			},
			chart.ContinuousSeries{
				Style: chart.Style{
					Show:        true,
					StrokeWidth: chart.Disabled,
					DotWidth:    7,
				},
				XValues: chart.Sequence.Random(8, 1024),
				YValues: chart.Sequence.Random(8, 1024),
			},
		},
	}

	res.Header().Set("Content-Type", "image/png")
	err := graph.Render(chart.PNG, res)
	if err != nil {
		log.Println(err.Error())
	}

}

func main() {
	http.HandleFunc("/", drawChart)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
