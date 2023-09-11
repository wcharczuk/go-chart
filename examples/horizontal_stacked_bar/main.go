package main

import (
	"os"

	"github.com/wcharczuk/go-chart/v2"
	"github.com/wcharczuk/go-chart/v2/drawing"
)

func main() {
	chart.DefaultBackgroundColor = chart.ColorTransparent
	chart.DefaultCanvasColor = chart.ColorTransparent

	barWidth := 80

	var (
		colorWhite          = drawing.Color{R: 241, G: 241, B: 241, A: 255}
		colorMariner        = drawing.Color{R: 60, G: 100, B: 148, A: 255}
		colorLightSteelBlue = drawing.Color{R: 182, G: 195, B: 220, A: 255}
		colorPoloBlue       = drawing.Color{R: 126, G: 155, B: 200, A: 255}
		colorSteelBlue      = drawing.Color{R: 73, G: 120, B: 177, A: 255}
	)

	stackedBarChart := chart.StackedBarChart{
		Title:      "Quarterly Sales",
		TitleStyle: chart.StyleShow(),
		Background: chart.Style{
			Padding: chart.Box{
				Top: 75,
			},
		},
		Width:        800,
		Height:       600,
		XAxis:        chart.StyleShow(),
		YAxis:        chart.StyleShow(),
		BarSpacing:   40,
		IsHorizontal: true,
		Bars: []chart.StackedBar{
			{
				Name:  "Q1",
				Width: barWidth,
				Values: []chart.Value{
					{
						Label: "32K",
						Value: 32,
						Style: chart.Style{
							StrokeWidth: .01,
							FillColor:   colorMariner,
							FontColor:   colorWhite,
						},
					},
					{
						Label: "46K",
						Value: 46,
						Style: chart.Style{
							StrokeWidth: .01,
							FillColor:   colorLightSteelBlue,
							FontColor:   colorWhite,
						},
					},
					{
						Label: "48K",
						Value: 48,
						Style: chart.Style{
							StrokeWidth: .01,
							FillColor:   colorPoloBlue,
							FontColor:   colorWhite,
						},
					},
					{
						Label: "42K",
						Value: 42,
						Style: chart.Style{
							StrokeWidth: .01,
							FillColor:   colorSteelBlue,
							FontColor:   colorWhite,
						},
					},
				},
			},
			{
				Name:  "Q2",
				Width: barWidth,
				Values: []chart.Value{
					{
						Label: "45K",
						Value: 45,
						Style: chart.Style{
							StrokeWidth: .01,
							FillColor:   colorMariner,
							FontColor:   colorWhite,
						},
					},
					{
						Label: "60K",
						Value: 60,
						Style: chart.Style{
							StrokeWidth: .01,
							FillColor:   colorLightSteelBlue,
							FontColor:   colorWhite,
						},
					},
					{
						Label: "62K",
						Value: 62,
						Style: chart.Style{
							StrokeWidth: .01,
							FillColor:   colorPoloBlue,
							FontColor:   colorWhite,
						},
					},
					{
						Label: "53K",
						Value: 53,
						Style: chart.Style{
							StrokeWidth: .01,
							FillColor:   colorSteelBlue,
							FontColor:   colorWhite,
						},
					},
				},
			},
			{
				Name:  "Q3",
				Width: barWidth,
				Values: []chart.Value{
					{
						Label: "54K",
						Value: 54,
						Style: chart.Style{
							StrokeWidth: .01,
							FillColor:   colorMariner,
							FontColor:   colorWhite,
						},
					},
					{
						Label: "58K",
						Value: 58,
						Style: chart.Style{
							StrokeWidth: .01,
							FillColor:   colorLightSteelBlue,
							FontColor:   colorWhite,
						},
					},
					{
						Label: "55K",
						Value: 55,
						Style: chart.Style{
							StrokeWidth: .01,
							FillColor:   colorPoloBlue,
							FontColor:   colorWhite,
						},
					},
					{
						Label: "47K",
						Value: 47,
						Style: chart.Style{
							StrokeWidth: .01,
							FillColor:   colorSteelBlue,
							FontColor:   colorWhite,
						},
					},
				},
			},
			{
				Name:  "Q4",
				Width: barWidth,
				Values: []chart.Value{
					{
						Label: "46K",
						Value: 46,
						Style: chart.Style{
							StrokeWidth: .01,
							FillColor:   colorMariner,
							FontColor:   colorWhite,
						},
					},
					{
						Label: "70K",
						Value: 70,
						Style: chart.Style{
							StrokeWidth: .01,
							FillColor:   colorLightSteelBlue,
							FontColor:   colorWhite,
						},
					},
					{
						Label: "74K",
						Value: 74,
						Style: chart.Style{
							StrokeWidth: .01,
							FillColor:   colorPoloBlue,
							FontColor:   colorWhite,
						},
					},
					{
						Label: "60K",
						Value: 60,
						Style: chart.Style{
							StrokeWidth: .01,
							FillColor:   colorSteelBlue,
							FontColor:   colorWhite,
						},
					},
				},
			},
		},
	}

	pngFile, err := os.Create("output.png")
	if err != nil {
		panic(err)
	}

	if err := stackedBarChart.Render(chart.PNG, pngFile); err != nil {
		panic(err)
	}

	if err := pngFile.Close(); err != nil {
		panic(err)
	}
}
