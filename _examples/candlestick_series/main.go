package main

import (
	"math/rand"
	"net/http"
	"time"

	chart "github.com/wcharczuk/go-chart"
	util "github.com/wcharczuk/go-chart/util"
)

func stockData() (times []time.Time, prices []float64) {
	start := time.Date(2017, 05, 15, 9, 30, 0, 0, util.Date.Eastern())
	price := 256.0
	for day := 0; day < 60; day++ {
		cursor := start.AddDate(0, 0, day)

		if util.Date.IsNYSEHoliday(cursor) {
			continue
		}

		for minute := 0; minute < ((6 * 60) + 30); minute++ {
			cursor = cursor.Add(time.Minute)

			if rand.Float64() >= 0.5 {
				price = price + (rand.Float64() * (price * 0.01))
			} else {
				price = price - (rand.Float64() * (price * 0.01))
			}

			times = append(times, cursor)
			prices = append(prices, price)
		}
	}
	return
}

func drawChart(res http.ResponseWriter, req *http.Request) {
	xv, yv := stockData()

	priceSeries := chart.TimeSeries{
		Name: "SPY",
		Style: chart.Style{
			Show:        false,
			StrokeColor: chart.GetDefaultColor(0),
		},
		XValues: xv,
		YValues: yv,
	}

	candleSeries := chart.CandlestickSeries{
		Name:    "SPY",
		XValues: xv,
		YValues: yv,
	}

	graph := chart.Chart{
		XAxis: chart.XAxis{
			Style:        chart.Style{Show: true, FontSize: 8, TextRotationDegrees: 45},
			TickPosition: chart.TickPositionUnderTick,
			Range:        &chart.MarketHoursRange{},
		},
		YAxis: chart.YAxis{
			Style: chart.Style{Show: true},
		},
		Series: []chart.Series{
			candleSeries,
			priceSeries,
		},
	}

	res.Header().Set("Content-Type", "image/png")
	err := graph.Render(chart.PNG, res)
	if err != nil {
		panic(err)
	}
}

func main() {
	http.HandleFunc("/", drawChart)
	http.ListenAndServe(":8080", nil)
}
