package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/wcharczuk/go-chart"
)

func random(min, max float64) float64 {
	return rand.Float64()*(max-min) + min
}

func drawLargeChart(res http.ResponseWriter, r *http.Request) {
	/*
		Usage: http://localhost:8080?series=100&values=1000
	*/

	numSeriesString := r.FormValue("series")

	numSeriesInt64, err := strconv.ParseInt(numSeriesString, 10, 64)
	if err != nil {
		numSeriesInt64 = int64(1)
	}
	numSeries := int(numSeriesInt64)

	series := make([]chart.Series, numSeries)

	numValuesString := r.FormValue("values")

	numValuesInt64, err := strconv.ParseInt(numValuesString, 10, 64)
	if err != nil {
		numValuesInt64 = int64(100)
	}
	numValues := int(numValuesInt64)

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
			Name:      "Time",
			NameStyle: chart.StyleShow(),
			Style:     chart.StyleShow(),
		},
		YAxis: chart.YAxis{
			Name:      "Value",
			NameStyle: chart.StyleShow(),
			Style:     chart.StyleShow(),
		},
		Series: series,
	}

	res.Header().Set("Content-Type", "image/png")
	graph.Render(chart.PNG, res)
}

func main() {
	http.HandleFunc("/", drawLargeChart)
	http.HandleFunc("/favico.ico", func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte{})
	})
	http.ListenAndServe(":8080", nil)
}
