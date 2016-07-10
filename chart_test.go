package chart

import (
	"bytes"
	"testing"
	"time"

	"github.com/blendlabs/go-assert"
)

func TestChartSingleSeries(t *testing.T) {
	assert := assert.New(t)
	now := time.Now()
	c := Chart{
		Title:  "Hello!",
		Width:  1024,
		Height: 400,
		YAxis: YAxis{
			Range: Range{
				Min: 0.0,
				Max: 4.0,
			},
		},
		Series: []Series{
			TimeSeries{
				Name:    "goog",
				XValues: []time.Time{now.AddDate(0, 0, -3), now.AddDate(0, 0, -2), now.AddDate(0, 0, -1)},
				YValues: []float64{1.0, 2.0, 3.0},
			},
		},
	}

	buffer := bytes.NewBuffer([]byte{})
	c.Render(PNG, buffer)
	assert.NotEmpty(buffer.Bytes())
}
