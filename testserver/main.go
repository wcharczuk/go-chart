package main

import (
	"log"
	"net/http"

	"github.com/wcharczuk/go-chart"
	"github.com/wcharczuk/go-chart/drawing"
	"github.com/wcharczuk/go-web"
)

func chartHandler(rc *web.RequestContext) web.ControllerResult {
	format, err := rc.RouteParameter("format")
	if err != nil {
		format = "png"
	}

	if format == "png" {
		rc.Response.Header().Set("Content-Type", "image/png")
	} else {
		rc.Response.Header().Set("Content-Type", "image/svg+xml")
	}

	s1x := []float64{2.0, 3.0, 4.0, 5.0}
	s1y := []float64{2.5, 5.0, 2.0, 3.3}

	s2x := []float64{0.0, 0.5, 1.0, 1.5}
	s2y := []float64{1.1, 1.2, 1.0, 1.3}

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
			Range: chart.Range{
				Min: 0.0,
				Max: 7.0,
			},
		},
		YAxisSecondary: chart.YAxis{
			Style: chart.Style{
				Show: true,
			},
			Range: chart.Range{
				Min: 0.8,
				Max: 1.5,
			},
		},
		Series: []chart.Series{
			chart.ContinuousSeries{
				Name:    "a",
				XValues: s1x,
				YValues: s1y,
			},
			chart.ContinuousSeries{
				Name:    "b",
				YAxis:   chart.YAxisSecondary,
				XValues: s2x,
				YValues: s2y,
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

func measureTestHandler(rc *web.RequestContext) web.ControllerResult {
	format, err := rc.RouteParameter("format")
	if err != nil {
		format = "png"
	}

	if format == "png" {
		rc.Response.Header().Set("Content-Type", "image/png")
	} else {
		rc.Response.Header().Set("Content-Type", "image/svg+xml")
	}

	var r chart.Renderer
	if format == "png" {
		r, err = chart.PNG(512, 512)
	} else {
		r, err = chart.SVG(512, 512)
	}
	if err != nil {
		return rc.API().InternalError(err)
	}

	f, err := chart.GetDefaultFont()
	if err != nil {
		return rc.API().InternalError(err)
	}
	r.SetDPI(96)
	r.SetFont(f)
	r.SetFontSize(24.0)
	r.SetFontColor(drawing.ColorBlack)
	r.SetStrokeColor(drawing.ColorBlack)

	tx, ty := 64, 64

	tw, th := r.MeasureText("test")
	r.MoveTo(tx, ty)
	r.LineTo(tx+tw, ty)
	r.LineTo(tx+tw, ty-th)
	r.LineTo(tx, ty-th)
	r.LineTo(tx, ty)
	r.Stroke()

	r.Text("test", tx, ty)

	r.Save(rc.Response)
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
	app.GET("/measure", measureTestHandler)
	log.Fatal(app.Start())
}
