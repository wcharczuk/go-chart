package main

//go:generate go run main.go

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/wcharczuk/go-chart"
)

func main() {
	log := chart.NewLogger()
	drawChart(log)
}

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
	err := chart.ReadLines("requests.csv", func(line string) error {
		parts := chart.SplitCSV(line)
		year := parseInt(parts[0])
		month := parseInt(parts[1])
		day := parseInt(parts[2])
		hour := parseInt(parts[3])
		elapsedMillis := parseFloat64(parts[4])
		xvalues = append(xvalues, time.Date(year, time.Month(month), day, hour, 0, 0, 0, time.UTC))
		yvalues = append(yvalues, elapsedMillis)
		return nil
	})
	if err != nil {
		fmt.Println(err.Error())
	}
	return xvalues, yvalues
}

func releases() []chart.GridLine {
	return []chart.GridLine{
		{Value: chart.TimeToFloat64(time.Date(2016, 8, 1, 9, 30, 0, 0, time.UTC))},
		{Value: chart.TimeToFloat64(time.Date(2016, 8, 2, 9, 30, 0, 0, time.UTC))},
		{Value: chart.TimeToFloat64(time.Date(2016, 8, 2, 15, 30, 0, 0, time.UTC))},
		{Value: chart.TimeToFloat64(time.Date(2016, 8, 4, 9, 30, 0, 0, time.UTC))},
		{Value: chart.TimeToFloat64(time.Date(2016, 8, 5, 9, 30, 0, 0, time.UTC))},
		{Value: chart.TimeToFloat64(time.Date(2016, 8, 6, 9, 30, 0, 0, time.UTC))},
	}
}

func drawChart(log chart.Logger) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		xvalues, yvalues := readData()
		mainSeries := chart.TimeSeries{
			Name: "Prod Request Timings",
			Style: chart.Style{
				StrokeColor: chart.ColorBlue,
				FillColor:   chart.ColorBlue.WithAlpha(100),
			},
			XValues: xvalues,
			YValues: yvalues,
		}

		linreg := &chart.LinearRegressionSeries{
			Name: "Linear Regression",
			Style: chart.Style{
				StrokeColor:     chart.ColorAlternateBlue,
				StrokeDashArray: []float64{5.0, 5.0},
			},
			InnerSeries: mainSeries,
		}

		sma := &chart.SMASeries{
			Name: "SMA",
			Style: chart.Style{
				StrokeColor:     chart.ColorRed,
				StrokeDashArray: []float64{5.0, 5.0},
			},
			InnerSeries: mainSeries,
		}

		graph := chart.Chart{
			Log:    log,
			Width:  1280,
			Height: 720,
			Background: chart.Style{
				Padding: chart.Box{
					Top: 50,
				},
			},
			YAxis: chart.YAxis{
				Name: "Elapsed Millis",
				TickStyle: chart.Style{
					TextRotationDegrees: 45.0,
				},
				ValueFormatter: func(v interface{}) string {
					return fmt.Sprintf("%d ms", int(v.(float64)))
				},
			},
			XAxis: chart.XAxis{
				ValueFormatter: chart.TimeHourValueFormatter,
				GridMajorStyle: chart.Style{
					StrokeColor: chart.ColorAlternateGray,
					StrokeWidth: 1.0,
				},
				GridLines: releases(),
			},
			Series: []chart.Series{
				mainSeries,
				linreg,
				chart.LastValueAnnotationSeries(linreg),
				sma,
				chart.LastValueAnnotationSeries(sma),
			},
		}

		graph.Elements = []chart.Renderable{chart.LegendThin(&graph)}

		f, _ := os.Create("output.png")
		defer f.Close()
		graph.Render(chart.PNG, f)
	}
}
