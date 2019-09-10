package main

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/wcharczuk/go-chart"
)

var lock sync.Mutex
var graph *chart.Chart
var ts *chart.TimeSeries

func addData(t time.Time, e time.Duration) {
	lock.Lock()
	ts.XValues = append(ts.XValues, t)
	ts.YValues = append(ts.YValues, chart.TimeMillis(e))
	lock.Unlock()
}

func drawChart(res http.ResponseWriter, req *http.Request) {
	start := time.Now()
	defer func() {
		addData(start, time.Since(start))
	}()
	if len(ts.XValues) == 0 {
		http.Error(res, "no data (yet)", http.StatusBadRequest)
		return
	}
	res.Header().Set("Content-Type", "image/png")
	if err := graph.Render(chart.PNG, res); err != nil {
		log.Printf("%v", err)
	}
}

func main() {
	ts = &chart.TimeSeries{
		XValues: []time.Time{},
		YValues: []float64{},
	}
	graph = &chart.Chart{
		Series: []chart.Series{ts},
	}
	http.HandleFunc("/", drawChart)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
