package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"time"

	chart "github.com/wcharczuk/go-chart"
)

var (
	outputPath        = flag.String("output", "", "The output file")
	disableLinreg     = flag.Bool("disable-linreg", false, "If we should omit linear regressions")
	disableLastValues = flag.Bool("disable-last-values", false, "If we should omit last values")
)

// NewLogger returns a new logger.
func NewLogger() *Logger {
	return &Logger{
		TimeFormat: time.RFC3339Nano,
		Stdout:     os.Stdout,
		Stderr:     os.Stderr,
	}
}

// Logger is a basic logger.
type Logger struct {
	TimeFormat string
	Stdout     io.Writer
	Stderr     io.Writer
}

// Info writes an info message.
func (l *Logger) Info(arguments ...interface{}) {
	l.Println(append([]interface{}{"[INFO]"}, arguments...)...)
}

// Infof writes an info message.
func (l *Logger) Infof(format string, arguments ...interface{}) {
	l.Println(append([]interface{}{"[INFO]"}, fmt.Sprintf(format, arguments...))...)
}

// Debug writes an debug message.
func (l *Logger) Debug(arguments ...interface{}) {
	l.Println(append([]interface{}{"[DEBUG]"}, arguments...)...)
}

// Debugf writes an debug message.
func (l *Logger) Debugf(format string, arguments ...interface{}) {
	l.Println(append([]interface{}{"[DEBUG]"}, fmt.Sprintf(format, arguments...))...)
}

// Error writes an error message.
func (l *Logger) Error(arguments ...interface{}) {
	l.Println(append([]interface{}{"[ERROR]"}, arguments...)...)
}

// Errorf writes an error message.
func (l *Logger) Errorf(format string, arguments ...interface{}) {
	l.Println(append([]interface{}{"[ERROR]"}, fmt.Sprintf(format, arguments...))...)
}

// Err writes an error message.
func (l *Logger) Err(err error) {
	if err != nil {
		l.Println(append([]interface{}{"[ERROR]"}, err.Error())...)
	}
}

// FatalErr writes an error message and exits.
func (l *Logger) FatalErr(err error) {
	if err != nil {
		l.Println(append([]interface{}{"[FATAL]"}, err.Error())...)
		os.Exit(1)
	}
}

// Println prints a new message.
func (l *Logger) Println(arguments ...interface{}) {
	fmt.Fprintln(l.Stdout, append([]interface{}{time.Now().UTC().Format(l.TimeFormat)}, arguments...)...)
}

// Errorln prints a new message.
func (l *Logger) Errorln(arguments ...interface{}) {
	fmt.Fprintln(l.Stderr, append([]interface{}{time.Now().UTC().Format(l.TimeFormat)}, arguments...)...)
}

func main() {
	log := NewLogger()

	rawData, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.FatalErr(err)
	}

	csvParts := chart.SplitCSV(string(rawData))

	yvalues, err := chart.ParseFloats(csvParts...)

	mainSeries := chart.ContinuousSeries{
		Name:    "A test series",
		XValues: chart.SeqRange(0, float64(len(csvParts))), //generates a []float64 from 1.0 to 100.0 in 1.0 step increments, or 100 elements.
		YValues: yvalues,
	}

	linRegSeries := &chart.LinearRegressionSeries{
		InnerSeries: mainSeries,
	}

	graph := chart.Chart{
		Series: []chart.Series{
			mainSeries,
			linRegSeries,
		},
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

	log.Info("rendering chart to", output.Name())
	if err := graph.Render(chart.PNG, output); err != nil {
		log.FatalErr(err)
	}
}
