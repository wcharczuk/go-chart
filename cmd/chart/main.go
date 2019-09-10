package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/wcharczuk/go-chart"
)

var (
	outputPath = flag.String("output", "", "The output file")

	inputFormat = flag.String("format", "csv", "The input format, either 'csv' or 'tsv' (defaults to 'csv')")
	inputPath   = flag.String("f", "", "The input file")
	reverse     = flag.Bool("reverse", false, "If we should reverse the inputs")

	hideLegend     = flag.Bool("hide-legend", false, "If we should omit the chart legend")
	hideSMA        = flag.Bool("hide-sma", false, "If we should omit simple moving average")
	hideLinreg     = flag.Bool("hide-linreg", false, "If we should omit linear regressions")
	hideLastValues = flag.Bool("hide-last-values", false, "If we should omit last values")
)

func main() {
	flag.Parse()
	log := chart.NewLogger()

	var rawData []byte
	var err error
	if *inputPath != "" {
		if *inputPath == "-" {
			rawData, err = ioutil.ReadAll(os.Stdin)
			if err != nil {
				log.FatalErr(err)
			}
		} else {
			rawData, err = ioutil.ReadFile(*inputPath)
			if err != nil {
				log.FatalErr(err)
			}
		}
	} else if len(flag.Args()) > 0 {
		rawData = []byte(flag.Args()[0])
	} else {
		flag.Usage()
		os.Exit(1)
	}

	var parts []string
	switch *inputFormat {
	case "csv":
		parts = chart.SplitCSV(string(rawData))
	case "tsv":
		parts = strings.Split(string(rawData), "\t")
	default:
		log.FatalErr(fmt.Errorf("invalid format; must be 'csv' or 'tsv'"))
	}

	yvalues, err := chart.ParseFloats(parts...)
	if err != nil {
		log.FatalErr(err)
	}

	if *reverse {
		yvalues = chart.ValueSequence(yvalues...).Reverse().Values()
	}

	var series []chart.Series
	mainSeries := chart.ContinuousSeries{
		Name:    "Values",
		XValues: chart.LinearRange(1, float64(len(yvalues))),
		YValues: yvalues,
	}
	series = append(series, mainSeries)

	smaSeries := &chart.SMASeries{
		Name: "SMA",
		Style: chart.Style{
			Hidden:          *hideSMA,
			StrokeColor:     chart.ColorRed,
			StrokeDashArray: []float64{5.0, 5.0},
		},
		InnerSeries: mainSeries,
	}
	series = append(series, smaSeries)

	linRegSeries := &chart.LinearRegressionSeries{
		Name: "Values - Lin. Reg.",
		Style: chart.Style{
			Hidden: *hideLinreg,
		},
		InnerSeries: mainSeries,
	}
	series = append(series, linRegSeries)

	mainLastValue := chart.LastValueAnnotationSeries(mainSeries)
	mainLastValue.Style = chart.Style{
		Hidden: *hideLastValues,
	}
	series = append(series, mainLastValue)

	linregLastValue := chart.LastValueAnnotationSeries(linRegSeries)
	linregLastValue.Style = chart.Style{
		Hidden: (*hideLastValues || *hideLinreg),
	}
	series = append(series, linregLastValue)

	smaLastValue := chart.LastValueAnnotationSeries(smaSeries)
	smaLastValue.Style = chart.Style{
		Hidden: (*hideLastValues || *hideSMA),
	}
	series = append(series, smaLastValue)

	graph := chart.Chart{
		Background: chart.Style{
			Padding: chart.Box{
				Top: 50,
			},
		},
		Series: series,
	}

	if !*hideLegend {
		graph.Elements = []chart.Renderable{chart.LegendThin(&graph)}
	}

	var output *os.File
	if *outputPath != "" {
		output, err = os.Create(*outputPath)
		if err != nil {
			log.FatalErr(err)
		}
	} else {
		output, err = ioutil.TempFile("", "*.png")
		if err != nil {
			log.FatalErr(err)
		}
	}

	if err := graph.Render(chart.PNG, output); err != nil {
		log.FatalErr(err)
	}

	fmt.Fprintln(os.Stdout, output.Name())
	os.Exit(0)
}
