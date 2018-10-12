package main

import (
	"fmt"
	"github.com/wcharczuk/go-chart"
	"log"
	"net/http"
)

// Note: Additional examples on how to add Stylesheets are in the custom_stylesheets example

func inlineSVGWithClasses(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte(
		"<!DOCTYPE html><html><head>" +
			"<link rel=\"stylesheet\" type=\"text/css\" href=\"/main.css\">" +
			"</head>" +
			"<body>"))

	pie := chart.PieChart{
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

	err := pie.Render(chart.SVG, res)
	if err != nil {
		fmt.Printf("Error rendering pie chart: %v\n", err)
	}
	res.Write([]byte("</body>"))
}

func css(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/css")
	res.Write([]byte("svg .background { fill: white; }" +
		"svg .canvas { fill: white; }" +
		"svg path.blue { fill: blue; stroke: lightblue; }" +
		"svg path.green { fill: green; stroke: lightgreen; }" +
		"svg path.gray { fill: gray; stroke: lightgray; }" +
		"svg text.blue { fill: white; }" +
		"svg text.green { fill: white; }" +
		"svg text.gray { fill: white; }"))
}

func main() {
	http.HandleFunc("/", inlineSVGWithClasses)
	http.HandleFunc("/main.css", css)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
