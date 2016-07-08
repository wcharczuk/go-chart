package main

import (
	"bytes"
	"log"
	"time"

	"github.com/wcharczuk/go-chart"
	"github.com/wcharczuk/go-web"
)

func main() {
	app := web.New()
	app.SetName("Chart Test Server")
	app.SetLogger(web.NewStandardOutputLogger())
	app.GET("/", func(rc *web.RequestContext) web.ControllerResult {
		rc.Response.Header().Set("Content-Type", "image/png")
		now := time.Now()
		c := chart.Chart{
			Title: "A Test Chart",
			TitleStyle: chart.Style{
				Show: true,
			},
			Width:  800,
			Height: 380,
			Axes: chart.Style{
				Show:        true,
				StrokeWidth: 1.0,
			},
			YRange: chart.Range{
				Min: 0.0,
				Max: 7.0,
			},
			FinalValueLabel: chart.Style{
				Show: true,
			},
			Series: []chart.Series{
				chart.TimeSeries{
					Name:    "a",
					XValues: []time.Time{now.AddDate(0, 0, -4), now.AddDate(0, 0, -3), now.AddDate(0, 0, -2), now.AddDate(0, 0, -1)},
					YValues: []float64{2.5, 5.0, 2.0, 3.0},
				},
				chart.TimeSeries{
					Name:    "b",
					XValues: []time.Time{now.AddDate(0, 0, -4), now.AddDate(0, 0, -3), now.AddDate(0, 0, -2), now.AddDate(0, 0, -1)},
					YValues: []float64{6.0, 5.0, 4.0, 1.0},
				},
			},
		}

		buffer := bytes.NewBuffer([]byte{})
		err := c.Render(chart.PNG, buffer)
		if err != nil {
			return rc.API().InternalError(err)
		}
		return rc.Raw(buffer.Bytes())
	})
	log.Fatal(app.Start())
}
