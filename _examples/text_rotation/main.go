package main

import (
	"net/http"

	"github.com/wcharczuk/go-chart"
	"github.com/wcharczuk/go-chart/drawing"
)

func drawChart(res http.ResponseWriter, req *http.Request) {
	f, _ := chart.GetDefaultFont()
	r, _ := chart.PNG(1024, 1024)

	chart.Draw.Text(r, "Test", 64, 64, chart.Style{
		FontColor: drawing.ColorBlack,
		FontSize:  18,
		Font:      f,
	})

	chart.Draw.Text(r, "Test", 64, 64, chart.Style{
		FontColor:           drawing.ColorBlack,
		FontSize:            18,
		Font:                f,
		TextRotationDegrees: 45.0,
	})

	tb := chart.Draw.MeasureText(r, "Test", chart.Style{
		FontColor: drawing.ColorBlack,
		FontSize:  18,
		Font:      f,
	}).Shift(64, 64)

	tbc := tb.Corners().Rotate(45)

	chart.Draw.BoxCorners(r, tbc, chart.Style{
		StrokeColor: drawing.ColorRed,
		StrokeWidth: 2,
	})

	tbcb := tbc.Box()
	chart.Draw.Box(r, tbcb, chart.Style{
		StrokeColor: drawing.ColorBlue,
		StrokeWidth: 2,
	})

	res.Header().Set("Content-Type", "image/png")
	r.Save(res)
}

func main() {
	http.HandleFunc("/", drawChart)
	http.ListenAndServe(":8080", nil)
}
