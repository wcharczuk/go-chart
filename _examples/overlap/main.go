package main

import (
	"net/http"

	chart "github.com/wcharczuk/go-chart"
	"github.com/wcharczuk/go-chart/drawing"
)

func conditionalColor(condition bool, trueColor drawing.Color, falseColor drawing.Color) drawing.Color {
	if condition {
		return trueColor
	}
	return falseColor
}

func drawChart(res http.ResponseWriter, req *http.Request) {
	r, _ := chart.PNG(1024, 1024)

	b0 := chart.Box{Left: 100, Top: 100, Right: 400, Bottom: 200}
	b1 := chart.Box{Left: 500, Top: 100, Right: 900, Bottom: 200}
	b0r := b0.Corners().Rotate(45).Shift(0, 200)

	chart.Draw.Box(r, b0, chart.Style{
		StrokeColor: drawing.ColorRed,
		StrokeWidth: 2,
		FillColor:   conditionalColor(b0.Corners().Overlaps(b1.Corners()), drawing.ColorRed, drawing.ColorTransparent),
	})

	chart.Draw.Box(r, b1, chart.Style{
		StrokeColor: drawing.ColorBlue,
		StrokeWidth: 2,
		FillColor:   conditionalColor(b1.Corners().Overlaps(b0.Corners()), drawing.ColorRed, drawing.ColorTransparent),
	})

	chart.Draw.Box2d(r, b0r, chart.Style{
		StrokeColor: drawing.ColorGreen,
		StrokeWidth: 2,
		FillColor:   conditionalColor(b0r.Overlaps(b0.Corners()), drawing.ColorRed, drawing.ColorTransparent),
	})

	res.Header().Set("Content-Type", "image/png")
	r.Save(res)
}

func main() {
	http.HandleFunc("/", drawChart)
	http.ListenAndServe(":8080", nil)
}
