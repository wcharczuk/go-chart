package chart

import (
	"math/rand"
	"testing"
	"time"

	assert "github.com/blendlabs/go-assert"
	"github.com/wcharczuk/go-chart/util"
)

func generateDummyStockData() (times []time.Time, prices []float64) {
	start := util.Date.On(util.NYSEOpen(), time.Date(2017, 05, 15, 0, 0, 0, 0, util.Date.Eastern()))
	cursor := start
	for day := 0; day < 60; day++ {

		if util.Date.IsWeekendDay(cursor.Weekday()) {
			cursor = start.AddDate(0, 0, day)
			continue
		}

		for hour := 0; hour < 7; hour++ {
			for minute := 0; minute < 60; minute++ {
				times = append(times, cursor)
				prices = append(prices, rand.Float64()*256)
				cursor = cursor.Add(time.Minute)
			}

			cursor = cursor.Add(time.Hour)
		}

		cursor = start.AddDate(0, 0, day)
	}

	return
}

func TestCandlestickSeriesCandleValues(t *testing.T) {
	assert := assert.New(t)

	xdata, ydata := generateDummyStockData()

	candleSeries := CandlestickSeries{
		XValues: xdata,
		YValues: ydata,
	}

	values := candleSeries.CandleValues()
	assert.Len(values, 43) // should be 60 days per the generator.
}
