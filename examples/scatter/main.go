package main

import (
	"log"
	"net/http"

	_ "net/http/pprof"

	"github.com/wcharczuk/go-chart"
	"github.com/wcharczuk/go-chart/drawing"
)

func drawChart(res http.ResponseWriter, req *http.Request) {

	viridisByY := func(xr, yr chart.Range, index int, x, y float64) drawing.Color {
		return chart.Viridis(y, yr.GetMin(), yr.GetMax())
	}

	graph := chart.Chart{
		Series: []chart.Series{
			chart.ContinuousSeries{
				Style: chart.Style{
					StrokeWidth:      chart.Disabled,
					DotWidth:         5,
					DotColorProvider: viridisByY,
				},
				XValues: chart.Seq{Sequence: chart.NewLinearSequence().WithStart(0).WithEnd(127)}.Values(),
				YValues: chart.Seq{Sequence: chart.NewRandomSequence().WithLen(128).WithMin(0).WithMax(1024)}.Values(),
			},
		},
	}

	res.Header().Set("Content-Type", chart.ContentTypePNG)
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
			Padding: chart.BoxZero,
		},
		Background: chart.Style{
			Padding: chart.BoxZero,
		},
		Series: []chart.Series{
			chart.ContinuousSeries{
				XValues: chart.Seq{Sequence: chart.NewLinearSequence().WithStart(0).WithEnd(4)}.Values(),
				YValues: chart.Seq{Sequence: chart.NewLinearSequence().WithStart(0).WithEnd(4)}.Values(),
			},
		},
	}

	res.Header().Set("Content-Type", chart.ContentTypePNG)
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
