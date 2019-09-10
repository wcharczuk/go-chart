package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/wcharczuk/go-chart"
)

const style = "svg .background { fill: white; }" +
	"svg .canvas { fill: white; }" +
	"svg .blue.fill.stroke { fill: blue; stroke: lightblue; }" +
	"svg .green.fill.stroke { fill: green; stroke: lightgreen; }" +
	"svg .gray.fill.stroke { fill: gray; stroke: lightgray; }" +
	"svg .blue.text { fill: white; }" +
	"svg .green.text { fill: white; }" +
	"svg .gray.text { fill: white; }"

func svgWithCustomInlineCSS(res http.ResponseWriter, _ *http.Request) {
	res.Header().Set("Content-Type", chart.ContentTypeSVG)

	// Render the CSS with custom css
	err := pieChart().Render(chart.SVGWithCSS(style, ""), res)
	if err != nil {
		fmt.Printf("Error rendering pie chart: %v\n", err)
	}
}

func svgWithCustomInlineCSSNonce(res http.ResponseWriter, _ *http.Request) {
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Security-Policy/style-src
	// This should be randomly generated on every request!
	const nonce = "RAND0MBASE64"

	res.Header().Set("Content-Security-Policy", fmt.Sprintf("style-src 'nonce-%s'", nonce))
	res.Header().Set("Content-Type", chart.ContentTypeSVG)

	// Render the CSS with custom css and a nonce.
	// Try changing the nonce to a different string - your browser should block the CSS.
	err := pieChart().Render(chart.SVGWithCSS(style, nonce), res)
	if err != nil {
		fmt.Printf("Error rendering pie chart: %v\n", err)
	}
}

func svgWithCustomExternalCSS(res http.ResponseWriter, _ *http.Request) {
	// Add external CSS
	res.Write([]byte(
		`<?xml version="1.0" standalone="no"?>` +
			`<?xml-stylesheet href="/main.css" type="text/css"?>` +
			`<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">`))

	res.Header().Set("Content-Type", chart.ContentTypeSVG)
	err := pieChart().Render(chart.SVG, res)
	if err != nil {
		fmt.Printf("Error rendering pie chart: %v\n", err)
	}
}

func pieChart() chart.PieChart {
	return chart.PieChart{
		// Note that setting ClassName will cause all other inline styles to be dropped!
		Background: chart.Style{ClassName: "background"},
		Canvas: chart.Style{
			ClassName: "canvas",
		},
		Width:  512,
		Height: 512,
		Values: []chart.Value{
			{Value: 5, Label: "Blue", Style: chart.Style{ClassName: "blue"}},
			{Value: 5, Label: "Green", Style: chart.Style{ClassName: "green"}},
			{Value: 4, Label: "Gray", Style: chart.Style{ClassName: "gray"}},
		},
	}
}

func css(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/css")
	res.Write([]byte(style))
}

func main() {
	http.HandleFunc("/", svgWithCustomInlineCSS)
	http.HandleFunc("/nonce", svgWithCustomInlineCSSNonce)
	http.HandleFunc("/external", svgWithCustomExternalCSS)
	http.HandleFunc("/main.css", css)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
