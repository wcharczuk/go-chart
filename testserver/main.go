package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/wcharczuk/go-chart"
	"github.com/wcharczuk/go-web"
)

func chartHandler(rc *web.RequestContext) web.ControllerResult {
	rnd := rand.New(rand.NewSource(0))
	format, err := rc.RouteParameter("format")
	if err != nil {
		format = "png"
	}

	if format == "png" {
		rc.Response.Header().Set("Content-Type", "image/png")
	} else {
		rc.Response.Header().Set("Content-Type", "image/svg+xml")
	}

	var s1x []time.Time
	for x := 0; x < 100; x++ {
		s1x = append([]time.Time{time.Now().AddDate(0, 0, -1*x)}, s1x...)
	}
	var s1y []float64
	for x := 0; x < 100; x++ {
		s1y = append(s1y, rnd.Float64()*1024)
	}

	c := chart.Chart{
		Title: "A Test Chart",
		TitleStyle: chart.Style{
			Show: true,
		},
		Width:  1024,
		Height: 400,
		XAxis: chart.XAxis{
			Style: chart.Style{
				Show: true,
			},
		},
		YAxis: chart.YAxis{
			Style: chart.Style{
				Show: true,
			},
			Zero: chart.GridLine{
				Style: chart.Style{
					Show:        true,
					StrokeWidth: 1.0,
				},
			},
		},
		Series: []chart.Series{
			chart.TimeSeries{
				Name:    "a",
				XValues: s1x,
				YValues: s1y,
			},
		},
	}

	if format == "png" {
		err = c.Render(chart.PNG, rc.Response)
	} else {
		err = c.Render(chart.SVG, rc.Response)
	}
	if err != nil {
		return rc.API().InternalError(err)
	}
	rc.Response.WriteHeader(http.StatusOK)
	return nil
}

func main() {
	app := web.New()
	app.SetName("Chart Test Server")
	app.SetLogger(web.NewStandardOutputLogger())
	app.GET("/", chartHandler)
	app.GET("/format/:format", chartHandler)
	app.GET("/favico.ico", func(rc *web.RequestContext) web.ControllerResult {
		return rc.Raw([]byte{})
	})
	log.Fatal(app.Start())
}
