package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/wcharczuk/go-chart"
	"github.com/wcharczuk/go-chart/drawing"
)

func drawChart(res http.ResponseWriter, req *http.Request) {
	profitStyle := chart.Style{
		Show:        true,
		FillColor:   drawing.ColorFromHex("13c158"),
		StrokeColor: drawing.ColorFromHex("13c158"),
		StrokeWidth: 0,
	}

	lossStyle := chart.Style{
		Show:        true,
		FillColor:   drawing.ColorFromHex("c11313"),
		StrokeColor: drawing.ColorFromHex("c11313"),
		StrokeWidth: 0,
	}

	sbc := chart.BarChart{
		Title:      "Bar Chart Using BaseValue",
		TitleStyle: chart.StyleShow(),
		Background: chart.Style{
			Padding: chart.Box{
				Top: 40,
			},
		},
		Height:   512,
		BarWidth: 60,
		XAxis: chart.Style{
			Show: true,
		},
		YAxis: chart.YAxis{
			Style: chart.Style{
				Show: true,
			},
			Ticks: []chart.Tick{
				{-4.0, "-4"},
				{-2.0, "-2"},
				{0, "0"},
				{2.0, "2"},
				{4.0, "4"},
				{6.0, "6"},
				{8.0, "8"},
				{10.0, "10"},
				{12.0, "12"},
			},
		},
		UseBaseValue: true,
		BaseValue:    0.0,
		Bars: []chart.Value{
			{Value: 10.0, Style: profitStyle, Label: "Profit"},
			{Value: 12.0, Style: profitStyle, Label: "More Profit"},
			{Value: 8.0, Style: profitStyle, Label: "Still Profit"},
			{Value: -4.0, Style: lossStyle, Label: "Loss!"},
			{Value: 3.0, Style: profitStyle, Label: "Phew Ok"},
			{Value: -2.0, Style: lossStyle, Label: "Oh No!"},
		},
	}

	res.Header().Set("Content-Type", "image/png")
	err := sbc.Render(chart.PNG, res)
	if err != nil {
		fmt.Printf("Error rendering chart: %v\n", err)
	}
}

func port() string {
	if len(os.Getenv("PORT")) > 0 {
		return os.Getenv("PORT")
	}
	return "8080"
}

func main() {
	listenPort := fmt.Sprintf(":%s", port())
	fmt.Printf("Listening on %s\n", listenPort)
	http.HandleFunc("/", drawChart)
	log.Fatal(http.ListenAndServe(listenPort, nil))
}
