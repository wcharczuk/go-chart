package chart

import (
	"fmt"
	"testing"

	assert "github.com/blend/go-sdk/assert"
	"github.com/wcharczuk/go-chart/seq"
)

func TestContinuousSeries(t *testing.T) {
	assert := assert.New(t)

	cs := ContinuousSeries{
		Name:    "Test Series",
		XValues: seq.Range(1.0, 10.0),
		YValues: seq.Range(1.0, 10.0),
	}

	assert.Equal("Test Series", cs.GetName())
	assert.Equal(10, cs.Len())
	x0, y0 := cs.GetValues(0)
	assert.Equal(1.0, x0)
	assert.Equal(1.0, y0)

	xn, yn := cs.GetValues(9)
	assert.Equal(10.0, xn)
	assert.Equal(10.0, yn)

	xn, yn = cs.GetLastValues()
	assert.Equal(10.0, xn)
	assert.Equal(10.0, yn)
}

func TestContinuousSeriesValueFormatter(t *testing.T) {
	assert := assert.New(t)

	cs := ContinuousSeries{
		XValueFormatter: func(v interface{}) string {
			return fmt.Sprintf("%f foo", v)
		},
		YValueFormatter: func(v interface{}) string {
			return fmt.Sprintf("%f bar", v)
		},
	}

	xf, yf := cs.GetValueFormatters()
	assert.Equal("0.100000 foo", xf(0.1))
	assert.Equal("0.100000 bar", yf(0.1))
}

func TestContinuousSeriesValidate(t *testing.T) {
	assert := assert.New(t)

	cs := ContinuousSeries{
		Name:    "Test Series",
		XValues: seq.Range(1.0, 10.0),
		YValues: seq.Range(1.0, 10.0),
	}
	assert.Nil(cs.Validate())

	cs = ContinuousSeries{
		Name:    "Test Series",
		XValues: seq.Range(1.0, 10.0),
	}
	assert.NotNil(cs.Validate())

	cs = ContinuousSeries{
		Name:    "Test Series",
		YValues: seq.Range(1.0, 10.0),
	}
	assert.NotNil(cs.Validate())
}
