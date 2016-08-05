package main

import (
	"net/http"

	"github.com/wcharczuk/go-chart"
)

func drawChart(res http.ResponseWriter, req *http.Request) {
	start := chart.Date.Date(2016, 7, 01, chart.Date.Eastern())
	end := chart.Date.Date(2016, 07, 21, chart.Date.Eastern())
	xv := chart.Sequence.MarketHours(start, end, chart.NYSEOpen, chart.NYSEClose, chart.Date.IsNYSEHoliday)
	yv := chart.Sequence.RandomWithAverage(len(xv), 200, 10)

	graph := chart.Chart{
		XAxis: chart.XAxis{
			Style:          chart.StyleShow(),
			TickPosition:   chart.TickPositionBetweenTicks,
			ValueFormatter: chart.TimeHourValueFormatter,
			Range: &chart.MarketHoursRange{
				MarketOpen:      chart.NYSEOpen,
				MarketClose:     chart.NYSEClose,
				HolidayProvider: chart.Date.IsNYSEHoliday,
			},
		},
		YAxis: chart.YAxis{
			Style: chart.StyleShow(),
		},
		Series: []chart.Series{
			chart.TimeSeries{
				XValues: xv,
				YValues: yv,
			},
		},
	}

	res.Header().Set("Content-Type", "image/png")
	graph.Render(chart.PNG, res)
}

func main() {
	http.HandleFunc("/", drawChart)
	http.ListenAndServe(":8080", nil)
}
