package main

import (
	"bytes"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/wcharczuk/go-chart"
	"github.com/wcharczuk/go-chart/drawing"
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

	s1 := chart.TimeSeries{
		Name:    "a",
		XValues: s1x,
		YValues: s1y,
		Style: chart.Style{
			Show: true,
			//FillColor: chart.GetDefaultSeriesStrokeColor(0).WithAlpha(64),
		},
	}

	s1lv := chart.AnnotationSeries{
		Name: fmt.Sprintf("Last Value"),
		Style: chart.Style{
			Show:        true,
			StrokeColor: chart.GetDefaultSeriesStrokeColor(0),
		},
		Annotations: []chart.Annotation{
			chart.Annotation{
				X:     float64(s1x[len(s1x)-1].Unix()),
				Y:     s1y[len(s1y)-1],
				Label: fmt.Sprintf("%s - %s", "test", chart.FloatValueFormatter(s1y[len(s1y)-1])),
			},
		},
	}

	s1ma := &chart.BollingerBandsSeries{
		Name: "BBS",
		Style: chart.Style{
			Show:        true,
			StrokeColor: chart.DefaultAxisColor,
			FillColor:   chart.DefaultAxisColor.WithAlpha(64),
		},
		K:           2.0,
		WindowSize:  10,
		InnerSeries: s1,
	}

	c := chart.Chart{
		Title: "A Test Chart",
		TitleStyle: chart.Style{
			Show: false,
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
			GridMajorStyle: chart.Style{
				Show: false,
			},
			GridMinorStyle: chart.Style{
				Show: true,
			},
		},
		Series: []chart.Series{
			s1,
			s1ma,
			s1lv,
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

func boxHandler(rc *web.RequestContext) web.ControllerResult {
	r, err := chart.PNG(1024, 1024)
	if err != nil {
		rc.API().InternalError(err)
	}

	f, err := chart.GetDefaultFont()
	if err != nil {
		return rc.API().InternalError(err)
	}

	//1:1 128wx128h @ 64,64
	a := chart.Box{Top: 64, Left: 64, Right: 192, Bottom: 192}

	// 3:2 256x170 @ 16, 16
	//b := chart.Box{Top: 16, Left: 16, Right: 256, Bottom: 170}

	// 2:3 170x256 @ 16, 16
	c := chart.Box{Top: 16, Left: 16, Right: 170, Bottom: 256}

	//fitb := a.Fit(b)
	fitc := a.Fit(c)
	//growb := a.Grow(b)
	//growc := a.Grow(c)
	//grow := a.Grow(b).Grow(c)

	conc := a.Constrain(c)

	boxStyle := chart.Style{
		StrokeColor: drawing.ColorBlack,
		StrokeWidth: 1.0,
		Font:        f,
		FontSize:    18.0,
	}

	computedBoxStyle := chart.Style{
		StrokeColor: drawing.ColorRed,
		StrokeWidth: 1.0,
		Font:        f,
		FontSize:    18.0,
	}

	chart.DrawBox(r, a, boxStyle)
	//chart.DrawBox(r, b, boxStyle)
	chart.DrawBox(r, c, boxStyle)
	//chart.DrawBox(r, fitb, computedBoxStyle)
	chart.DrawBox(r, fitc, computedBoxStyle)
	/*chart.DrawBox(r, growb, computedBoxStyle)
	chart.DrawBox(r, growc, computedBoxStyle)
	chart.DrawBox(r, grow, computedBoxStyle)*/
	chart.DrawBox(r, conc, computedBoxStyle)

	ax, ay := a.Center()
	chart.DrawTextCentered(r, "a", ax, ay, boxStyle.WithDefaultsFrom(chart.Style{
		FillColor: boxStyle.StrokeColor,
	}))

	/*bx, by := b.Center()
	chart.DrawTextCentered(r, "b", bx, by, boxStyle.WithDefaultsFrom(chart.Style{
		FillColor: boxStyle.StrokeColor,
	}))*/

	cx, cy := c.Center()
	chart.DrawTextCentered(r, "c", cx, cy, boxStyle.WithDefaultsFrom(chart.Style{
		FillColor: boxStyle.StrokeColor,
	}))

	/*fbx, fby := fitb.Center()
	chart.DrawTextCentered(r, "a fit b", fbx, fby, computedBoxStyle.WithDefaultsFrom(chart.Style{
		FillColor: computedBoxStyle.StrokeColor,
	}))*/

	fcx, fcy := fitc.Center()
	chart.DrawTextCentered(r, "a fit c", fcx, fcy, computedBoxStyle.WithDefaultsFrom(chart.Style{
		FillColor: computedBoxStyle.StrokeColor,
	}))

	/*gbx, gby := growb.Center()
	chart.DrawTextCentered(r, "a grow b", gbx, gby, computedBoxStyle.WithDefaultsFrom(chart.Style{
		FillColor: computedBoxStyle.StrokeColor,
	}))

	gcx, gcy := growc.Center()
	chart.DrawTextCentered(r, "a grow c", gcx, gcy, computedBoxStyle.WithDefaultsFrom(chart.Style{
		FillColor: computedBoxStyle.StrokeColor,
	}))*/

	ccx, ccy := conc.Center()
	chart.DrawTextCentered(r, "a const c", ccx, ccy, computedBoxStyle.WithDefaultsFrom(chart.Style{
		FillColor: computedBoxStyle.StrokeColor,
	}))

	rc.Response.Header().Set("Content-Type", "image/png")
	buffer := bytes.NewBuffer([]byte{})
	err = r.Save(buffer)
	return rc.Raw(buffer.Bytes())
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
	app.GET("/box", boxHandler)
	log.Fatal(app.Start())
}
