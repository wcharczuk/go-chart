package main

import (
	"math/rand"
	"os"
	"time"

	chart "github.com/wcharczuk/go-chart"
)

func random(min, max float64) float64 {
	return rand.Float64()*(max-min) + min
}

func main() {
	/*
		This is an example of using `DetailsBox` to add a box with additional
		information about the graph.
	*/
	xValues := make([]time.Time, 0)
	yValues := make([][]float64, 3)
	for i := 0; i < 50; i++ {
		xValues = append(xValues, time.Now().Add(time.Duration(i)*time.Minute))

		for j := 0; j < 3; j++ {
			yValues[j] = append(yValues[j], random(float64(0), float64(20)))
		}
	}

	   	seriesOne := chart.TimeSeries{
	   		Name:    "Series One",
	   		XValues: xValues,
	   		YValues: yValues[0],
	   	}

	   	seriesTwo := chart.TimeSeries{
	   		Name:    "Series Two",
	   		XValues: xValues,
	   		YValues: yValues[1],
		   }

		seriesThree := chart.TimeSeries{
			Name: "Series Three",
			XValues: xValues,
			YValues: yValues[2],
		}

	graph := chart.Chart{
		Series: []chart.Series{
			seriesOne,
			seriesTwo,
			seriesThree,
		},
		Background: chart.Style{
			Padding: chart.Box{
				Left: 120,
			},
		},
	}

	graph.Elements = []chart.Renderable{
		chart.LegendLeft(&graph),
		chart.DetailsBox(&graph, []string{
			"M-F 8:00 am - 6:00 pm",
			"Poll for data every 9 min.",
		}),
	}

	f, _ := os.Create("output.png")
	defer f.Close()
	graph.Render(chart.PNG, f)
}
