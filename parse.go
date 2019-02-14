package chart

import (
	"strconv"
	"time"

	"github.com/blend/go-sdk/exception"
)

// ParseFloats parses a list of floats.
func ParseFloats(values ...string) ([]float64, error) {
	var output []float64
	var parsedValue float64
	var err error
	for _, value := range values {
		if parsedValue, err = strconv.ParseFloat(value, 64); err != nil {
			return nil, exception.New(err)
		}
		output = append(output, parsedValue)
	}
	return output, nil
}

// ParseTimes parses a list of times with a given format.
func ParseTimes(layout string, values ...string) ([]time.Time, error) {
	var output []time.Time
	var parsedValue time.Time
	var err error
	for _, value := range values {
		if parsedValue, err = time.Parse(layout, value); err != nil {
			return nil, exception.New(err)
		}
		output = append(output, parsedValue)
	}
	return output, nil
}
