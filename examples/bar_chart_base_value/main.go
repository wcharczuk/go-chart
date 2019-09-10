package main

//go:generate go run main.go

import (
	"os"

	"github.com/wcharczuk/go-chart"
	"github.com/wcharczuk/go-chart/drawing"
)

func main() {
	profitStyle := chart.Style{
		FillColor:   drawing.ColorFromHex("13c158"),
		StrokeColor: drawing.ColorFromHex("13c158"),
		StrokeWidth: 0,
	}

	lossStyle := chart.Style{
		FillColor:   drawing.ColorFromHex("c11313"),
		StrokeColor: drawing.ColorFromHex("c11313"),
		StrokeWidth: 0,
	}

	sbc := chart.BarChart{
		Title: "Bar Chart Using BaseValue",
		Background: chart.Style{
			Padding: chart.Box{
				Top: 40,
			},
		},
		Height:   512,
		BarWidth: 60,
		YAxis: chart.YAxis{
			Ticks: []chart.Tick{
				{Value: -4.0, Label: "-4"},
				{Value: -2.0, Label: "-2"},
				{Value: 0, Label: "0"},
				{Value: 2.0, Label: "2"},
				{Value: 4.0, Label: "4"},
				{Value: 6.0, Label: "6"},
				{Value: 8.0, Label: "8"},
				{Value: 10.0, Label: "10"},
				{Value: 12.0, Label: "12"},
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

	f, _ := os.Create("output.png")
	defer f.Close()
	sbc.Render(chart.PNG, f)
}
