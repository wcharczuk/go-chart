package main

import (
	"fmt"
	"log"

	"github.com/blendlabs/go-util"
	"github.com/wcharczuk/go-chart"
	"github.com/wcharczuk/go-web"
)

func main() {
	app := web.New()
	app.SetName("Chart Test Server")
	app.SetLogger(web.NewStandardOutputLogger())
	app.GET("/", func(rc *web.RequestContext) web.ControllerResult {
		rc.Response.Header().Set("Content-Type", "image/png")
		c := chart.Chart{
			Title: "A Test Chart",
			TitleStyle: chart.Style{
				Show: true,
			},
			Width:  640,
			Height: 480,
			Axes: chart.Style{
				Show:        true,
				StrokeWidth: 1.0,
			},
			YRange: chart.Range{
				Min: 0.0,
				Max: 7.0,
				Formatter: func(v interface{}) string {
					if typed, isTyped := v.(float64); isTyped {
						return fmt.Sprintf("%.4f", typed)
					}
					return util.StringEmpty
				},
			},
			FinalValueLabel: chart.Style{
				Show: true,
			},
			Series: []chart.Series{
				chart.ContinuousSeries{
					Name:    "a",
					XValues: []float64{1.0, 2.0, 3.0, 4.0},
					YValues: []float64{2.5, 5.0, 2.0, 3.3},
				},
				chart.ContinuousSeries{
					Name:    "b",
					XValues: []float64{3.0, 4.0, 5.0, 6.0},
					YValues: []float64{6.0, 5.0, 4.0, 1.0},
				},
			},
		}

		err := c.Render(chart.PNG, rc.Response)
		if err != nil {
			return rc.API().InternalError(err)
		}
		return nil
	})
	log.Fatal(app.Start())
}
