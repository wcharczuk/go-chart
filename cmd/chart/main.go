package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	chart "github.com/wcharczuk/go-chart"
)

var (
	outputPath        = flag.String("output", "", "The output file")
	inputFormat       = flag.String("format", "csv", "The input format, either 'csv' or 'tsv' (defaults to 'csv')")
	inputPath         = flag.String("f", "", "The input file")
	disableLinreg     = flag.Bool("disable-linreg", false, "If we should omit linear regressions")
	disableLastValues = flag.Bool("disable-last-values", false, "If we should omit last values")
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

	var series []chart.Series
	mainSeries := chart.ContinuousSeries{
		Name:    "Values",
		XValues: chart.LinearRange(1, float64(len(yvalues))),
		YValues: yvalues,
	}
	series = append(series, mainSeries)

	if !*disableLinreg {
		linRegSeries := &chart.LinearRegressionSeries{
			InnerSeries: mainSeries,
		}
		series = append(series, linRegSeries)
	}

	graph := chart.Chart{
		Series: series,
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
