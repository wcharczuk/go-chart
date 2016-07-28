package chart

import (
	"testing"
	"time"

	"github.com/blendlabs/go-assert"
)

func TestMinAndMax(t *testing.T) {
	assert := assert.New(t)
	values := []float64{1.0, 2.0, 3.0, 4.0}
	min, max := MinAndMax(values...)
	assert.Equal(1.0, min)
	assert.Equal(4.0, max)
}

func TestMinAndMaxReversed(t *testing.T) {
	assert := assert.New(t)
	values := []float64{4.0, 2.0, 3.0, 1.0}
	min, max := MinAndMax(values...)
	assert.Equal(1.0, min)
	assert.Equal(4.0, max)
}

func TestMinAndMaxEmpty(t *testing.T) {
	assert := assert.New(t)
	values := []float64{}
	min, max := MinAndMax(values...)
	assert.Equal(0.0, min)
	assert.Equal(0.0, max)
}

func TestMinAndMaxOfTime(t *testing.T) {
	assert := assert.New(t)
	values := []time.Time{
		time.Now().AddDate(0, 0, -1),
		time.Now().AddDate(0, 0, -2),
		time.Now().AddDate(0, 0, -3),
		time.Now().AddDate(0, 0, -4),
	}
	min, max := MinAndMaxOfTime(values...)
	assert.Equal(values[3], min)
	assert.Equal(values[0], max)
}

func TestMinAndMaxOfTimeReversed(t *testing.T) {
	assert := assert.New(t)
	values := []time.Time{
		time.Now().AddDate(0, 0, -4),
		time.Now().AddDate(0, 0, -2),
		time.Now().AddDate(0, 0, -3),
		time.Now().AddDate(0, 0, -1),
	}
	min, max := MinAndMaxOfTime(values...)
	assert.Equal(values[0], min)
	assert.Equal(values[3], max)
}

func TestMinAndMaxOfTimeEmpty(t *testing.T) {
	assert := assert.New(t)
	values := []time.Time{}
	min, max := MinAndMaxOfTime(values...)
	assert.Equal(time.Time{}, min)
	assert.Equal(time.Time{}, max)
}

func TestSlices(t *testing.T) {
	assert := assert.New(t)

	s := Slices(10, 100)
	assert.Len(s, 10)
	assert.Equal(0, s[0])
	assert.Equal(10, s[1])
	assert.Equal(20, s[2])
	assert.Equal(90, s[9])
}

func TestGetRoundToForDelta(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(100.0, GetRoundToForDelta(1001.00))
	assert.Equal(10.0, GetRoundToForDelta(101.00))
	assert.Equal(1.0, GetRoundToForDelta(11.00))
}

func TestSeq(t *testing.T) {
	assert := assert.New(t)

	asc := Seq(1.0, 10.0)
	assert.Len(asc, 10)

	desc := Seq(10.0, 1.0)
	assert.Len(desc, 10)
}

func TestPercentDifference(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(0.5, PercentDifference(1.0, 1.5))
	assert.Equal(-0.5, PercentDifference(2.0, 1.0))
}

var (
	_degreesToRadians = map[float64]float64{
		0:   0, // !_2pi b/c no irrational nums in floats.
		45:  _pi4,
		90:  _pi2,
		135: _3pi4,
		180: _pi,
		225: _5pi4,
		270: _3pi2,
		315: _7pi4,
	}

	_compassToRadians = map[float64]float64{
		0:   _pi2,
		45:  _pi4,
		90:  0, // !_2pi b/c no irrational nums in floats.
		135: _7pi4,
		180: _3pi2,
		225: _5pi4,
		270: _pi,
		315: _3pi4,
	}
)

func TestDegreesToRadians(t *testing.T) {
	assert := assert.New(t)

	for d, r := range _degreesToRadians {
		assert.Equal(r, DegreesToRadians(d))
	}
}

func TestPercentToRadians(t *testing.T) {
	assert := assert.New(t)

	for d, r := range _degreesToRadians {
		assert.Equal(r, PercentToRadians(d/360.0))
	}
}

func TestRadiansToDegrees(t *testing.T) {
	assert := assert.New(t)

	for d, r := range _degreesToRadians {
		assert.Equal(d, RadiansToDegrees(r))
	}
}

func TestRadianAdd(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(_pi, RadianAdd(_pi2, _pi2))
	assert.Equal(_3pi2, RadianAdd(_pi2, _pi))
	assert.Equal(_pi, RadianAdd(_pi, _2pi))
	assert.Equal(_pi, RadianAdd(_pi, -_2pi))
}
