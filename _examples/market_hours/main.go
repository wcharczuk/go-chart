package main

import (
	"net/http"

	"github.com/wcharczuk/go-chart"
	"github.com/wcharczuk/go-chart/seq"
	"github.com/wcharczuk/go-chart/util"
)

func drawChart(res http.ResponseWriter, req *http.Request) {
	start := util.Date.Date(2016, 7, 01, util.Date.Eastern())
	end := util.Date.Date(2016, 07, 21, util.Date.Eastern())
	xv := seq.Time.MarketHours(start, end, util.NYSEOpen(), util.NYSEClose(), util.Date.IsNYSEHoliday)
	yv := seq.New(seq.NewRandom().WithLen(len(xv)).WithAverage(200).WithScale(10)).Array()

	graph := chart.Chart{
		XAxis: chart.XAxis{
			Style:          chart.StyleShow(),
			TickPosition:   chart.TickPositionBetweenTicks,
			ValueFormatter: chart.TimeHourValueFormatter,
			Range: &chart.MarketHoursRange{
				MarketOpen:      util.NYSEOpen(),
				MarketClose:     util.NYSEClose(),
				HolidayProvider: util.Date.IsNYSEHoliday,
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
