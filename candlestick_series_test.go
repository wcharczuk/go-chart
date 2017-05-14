package chart

import (
	"math/rand"
	"testing"
	"time"

	assert "github.com/blendlabs/go-assert"
	"github.com/wcharczuk/go-chart/util"
)

func generateDummyStockData() (times []time.Time, prices []float64) {
	start := util.Date.On(time.Date(2017, 05, 15, 6, 30, 0, 0, util.Date.Eastern()), util.NYSEOpen())
	var cursor time.Time
	for day := 0; day < 60; day++ {
		cursor = start.AddDate(0, 0, day)
		for hour := 0; hour < 7; hour++ {
			for minute := 0; minute < 60; minute++ {
				times = append(times, cursor)
				prices = append(prices, rand.Float64()*256)

				cursor = cursor.Add(time.Minute)
			}

			cursor = cursor.Add(time.Hour)
		}
	}
	return
}

func TestCandlestickSeriesCandleValues(t *testing.T) {
	assert := assert.New(t)

	xdata, ydata := generateDummyStockData()

	candleSeries := CandlestickSeries{
		InnerSeries: TimeSeries{
			XValues: xdata,
			YValues: ydata,
		},
	}

	values := candleSeries.CandleValues()
	assert.NotEmpty(values)
}
