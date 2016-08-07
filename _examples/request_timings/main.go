package main

import (
	"net/http"
	"strings"
	"time"

	util "github.com/blendlabs/go-util"
	"github.com/wcharczuk/go-chart"
)

func readData() ([]time.Time, []float64) {
	var xvalues []time.Time
	var yvalues []float64
	util.ReadFileByLines("requests.csv", func(line string) {
		parts := strings.Split(line, ",")
		year := util.ParseInt(parts[0])
		month := util.ParseInt(parts[1])
		day := util.ParseInt(parts[2])
		hour := util.ParseInt(parts[3])
		elapsedMillis := util.ParseFloat64(parts[4])
		xvalues = append(xvalues, time.Date(year, time.Month(month), day, hour, 0, 0, 0, time.UTC))
		yvalues = append(yvalues, elapsedMillis)
	})
	return xvalues, yvalues
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

	minSeries := &chart.MinSeries{
		Style: chart.Style{
			Show:            true,
			StrokeColor:     chart.ColorAlternateGray,
			StrokeDashArray: []float64{5.0, 5.0},
		},
		InnerSeries: mainSeries,
	}

	maxSeries := &chart.MaxSeries{
		Style: chart.Style{
			Show:            true,
			StrokeColor:     chart.ColorAlternateGray,
			StrokeDashArray: []float64{5.0, 5.0},
		},
		InnerSeries: mainSeries,
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
		YAxis: chart.YAxis{
			Name:      "Elapsed Millis",
			NameStyle: chart.StyleShow(),
			Style:     chart.StyleShow(),
		},
		XAxis: chart.XAxis{
			Style:          chart.StyleShow(),
			ValueFormatter: chart.TimeHourValueFormatter,
		},
		Series: []chart.Series{
			mainSeries,
			minSeries,
			maxSeries,
			chart.LastValueAnnotation(minSeries),
			chart.LastValueAnnotation(maxSeries),
			linreg,
			chart.LastValueAnnotation(linreg),
			sma,
			chart.LastValueAnnotation(sma),
		},
	}

	graph.Elements = []chart.Renderable{chart.Legend(&graph)}

	res.Header().Set("Content-Type", "image/png")
	graph.Render(chart.PNG, res)
}

func main() {
	http.HandleFunc("/", drawChart)
	http.ListenAndServe(":8080", nil)
}
