package main

//go:generate go run main.go

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/wcharczuk/go-chart"
)

func random(min, max float64) float64 {
	return rand.Float64()*(max-min) + min
}

func main() {
	numValues := 1024
	numSeries := 100
	series := make([]chart.Series, numSeries)

	for i := 0; i < numSeries; i++ {
		xValues := make([]time.Time, numValues)
		yValues := make([]float64, numValues)

		for j := 0; j < numValues; j++ {
			xValues[j] = time.Now().AddDate(0, 0, (numValues-j)*-1)
			yValues[j] = random(float64(-500), float64(500))
		}

		series[i] = chart.TimeSeries{
			Name:    fmt.Sprintf("aaa.bbb.hostname-%v.ccc.ddd.eee.fff.ggg.hhh.iii.jjj.kkk.lll.mmm.nnn.value", i),
			XValues: xValues,
			YValues: yValues,
		}
	}

	graph := chart.Chart{
		XAxis: chart.XAxis{
			Name: "Time",
		},
		YAxis: chart.YAxis{
			Name: "Value",
		},
		Series: series,
	}

	f, _ := os.Create("output.png")
	defer f.Close()
	graph.Render(chart.PNG, f)
}
