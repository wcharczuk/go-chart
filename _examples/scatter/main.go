package main

import (
	"log"
	"net/http"

	_ "net/http/pprof"

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
				XValues: chart.Sequence.Random(8192, 1024),
				YValues: chart.Sequence.Random(8192, 1024),
			},
		},
	}

	res.Header().Set("Content-Type", "image/png")
	err := graph.Render(chart.PNG, res)
	if err != nil {
		log.Println(err.Error())
	}
}

func unit(res http.ResponseWriter, req *http.Request) {
	graph := chart.Chart{
		Height: 50,
		Width:  50,
		Canvas: chart.Style{
			Padding: chart.Box{IsSet: true},
		},
		Background: chart.Style{
			Padding: chart.Box{IsSet: true},
		},
		Series: []chart.Series{
			chart.ContinuousSeries{
				XValues: chart.Sequence.Float64(0, 4, 1),
				YValues: chart.Sequence.Float64(0, 4, 1),
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
	http.HandleFunc("/unit", unit)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
