package main

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/wcharczuk/go-chart"
)

func parseInt(str string) int {
	v, _ := strconv.Atoi(str)
	return v
}

func parseFloat64(str string) float64 {
	v, _ := strconv.ParseFloat(str, 64)
	return v
}

func readData() ([]time.Time, []float64) {
	var xvalues []time.Time
	var yvalues []float64
	chart.File.ReadByLines("requests.csv", func(line string) {
		parts := strings.Split(line, ",")
		year := parseInt(parts[0])
		month := parseInt(parts[1])
		day := parseInt(parts[2])
		hour := parseInt(parts[3])
		elapsedMillis := parseFloat64(parts[4])
		xvalues = append(xvalues, time.Date(year, time.Month(month), day, hour, 0, 0, 0, time.UTC))
		yvalues = append(yvalues, elapsedMillis)
	})
	return xvalues, yvalues
}

func releases() []chart.GridLine {
	return []chart.GridLine{
		{Value: chart.Time.ToFloat64(time.Date(2016, 8, 1, 9, 30, 0, 0, time.UTC))},
		{Value: chart.Time.ToFloat64(time.Date(2016, 8, 2, 9, 30, 0, 0, time.UTC))},
		{Value: chart.Time.ToFloat64(time.Date(2016, 8, 3, 9, 30, 0, 0, time.UTC))},
		{Value: chart.Time.ToFloat64(time.Date(2016, 8, 4, 9, 30, 0, 0, time.UTC))},
		{Value: chart.Time.ToFloat64(time.Date(2016, 8, 5, 9, 30, 0, 0, time.UTC))},
		{Value: chart.Time.ToFloat64(time.Date(2016, 8, 6, 9, 30, 0, 0, time.UTC))},
	}
}

func drawChart(res http.ResponseWriter, req *http.Request) {
	xvalues, yvalues := readData()
	mainSeries := chart.TimeSeries{
		Name: "Prod Request Timings",
		Style: chart.Style{
			Show:        true,
			StrokeColor: chart.ColorBlue,
			FillColor:   chart.ColorBlue.WithAlpha(100),
		},
		XValues: xvalues,
		YValues: yvalues,
	}

	linreg := &chart.LinearRegressionSeries{
		Name: "Linear Regression",
		Style: chart.Style{
			Show:            true,
			StrokeColor:     chart.ColorAlternateBlue,
			StrokeDashArray: []float64{5.0, 5.0},
		},
		InnerSeries: mainSeries,
	}

	sma := &chart.SMASeries{
		Name: "SMA",
		Style: chart.Style{
			Show:            true,
			StrokeColor:     chart.ColorRed,
			StrokeDashArray: []float64{5.0, 5.0},
		},
		InnerSeries: mainSeries,
	}

	graph := chart.Chart{
		Width:  1280,
		Height: 720,
		Background: chart.Style{
			Padding: chart.Box{
				Top: 50,
			},
		},
		YAxis: chart.YAxis{
			Name:      "Elapsed Millis",
			NameStyle: chart.StyleShow(),
			Style:     chart.StyleShow(),
			TickStyle: chart.Style{
				TextRotationDegrees: 45.0,
			},
		},
		XAxis: chart.XAxis{
			Style: chart.Style{
				Show: true,
			},
			ValueFormatter: chart.TimeHourValueFormatter,
			GridMajorStyle: chart.Style{
				Show:        true,
				StrokeColor: chart.ColorAlternateGray,
				StrokeWidth: 1.0,
			},
			GridLines: releases(),
		},
		Series: []chart.Series{
			mainSeries,
			linreg,
			chart.LastValueAnnotation(linreg),
			sma,
			chart.LastValueAnnotation(sma),
		},
	}

	graph.Elements = []chart.Renderable{chart.LegendThin(&graph)}

	res.Header().Set("Content-Type", "image/png")
	graph.Render(chart.PNG, res)
}

func main() {
	http.HandleFunc("/", drawChart)
	http.ListenAndServe(":8080", nil)
}
