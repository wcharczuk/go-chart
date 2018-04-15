package util

import (
	"testing"
	"time"

	"github.com/blend/go-sdk/assert"
)

func TestMinAndMax(t *testing.T) {
	assert := assert.New(t)
	values := []float64{1.0, 2.0, 3.0, 4.0}
	min, max := Math.MinAndMax(values...)
	assert.Equal(1.0, min)
	assert.Equal(4.0, max)
}

func TestMinAndMaxReversed(t *testing.T) {
	assert := assert.New(t)
	values := []float64{4.0, 2.0, 3.0, 1.0}
	min, max := Math.MinAndMax(values...)
	assert.Equal(1.0, min)
	assert.Equal(4.0, max)
}

func TestMinAndMaxEmpty(t *testing.T) {
	assert := assert.New(t)
	values := []float64{}
	min, max := Math.MinAndMax(values...)
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
	min, max := Math.MinAndMaxOfTime(values...)
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
	min, max := Math.MinAndMaxOfTime(values...)
	assert.Equal(values[0], min)
	assert.Equal(values[3], max)
}

func TestMinAndMaxOfTimeEmpty(t *testing.T) {
	assert := assert.New(t)
	values := []time.Time{}
	min, max := Math.MinAndMaxOfTime(values...)
	assert.Equal(time.Time{}, min)
	assert.Equal(time.Time{}, max)
}

func TestGetRoundToForDelta(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(100.0, Math.GetRoundToForDelta(1001.00))
	assert.Equal(10.0, Math.GetRoundToForDelta(101.00))
	assert.Equal(1.0, Math.GetRoundToForDelta(11.00))
}

func TestRoundUp(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(0.5, Math.RoundUp(0.49, 0.1))
	assert.Equal(1.0, Math.RoundUp(0.51, 1.0))
	assert.Equal(0.4999, Math.RoundUp(0.49988, 0.0001))
}

func TestRoundDown(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(0.5, Math.RoundDown(0.51, 0.1))
	assert.Equal(1.0, Math.RoundDown(1.01, 1.0))
	assert.Equal(0.5001, Math.RoundDown(0.50011, 0.0001))
}

func TestPercentDifference(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(0.5, Math.PercentDifference(1.0, 1.5))
	assert.Equal(-0.5, Math.PercentDifference(2.0, 1.0))
}

func TestNormalize(t *testing.T) {
	assert := assert.New(t)

	values := []float64{10, 9, 8, 7, 6}
	normalized := Math.Normalize(values...)
	assert.Len(5, normalized)
	assert.Equal(0.25, normalized[0])
	assert.Equal(0.1499, normalized[4])
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
		assert.Equal(r, Math.DegreesToRadians(d))
	}
}

func TestPercentToRadians(t *testing.T) {
	assert := assert.New(t)

	for d, r := range _degreesToRadians {
		assert.Equal(r, Math.PercentToRadians(d/360.0))
	}
}

func TestRadiansToDegrees(t *testing.T) {
	assert := assert.New(t)

	for d, r := range _degreesToRadians {
		assert.Equal(d, Math.RadiansToDegrees(r))
	}
}

func TestRadianAdd(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(_pi, Math.RadianAdd(_pi2, _pi2))
	assert.Equal(_3pi2, Math.RadianAdd(_pi2, _pi))
	assert.Equal(_pi, Math.RadianAdd(_pi, _2pi))
	assert.Equal(_pi, Math.RadianAdd(_pi, -_2pi))
}

func TestRotateCoordinate90(t *testing.T) {
	assert := assert.New(t)

	cx, cy := 10, 10
	x, y := 5, 10

	rx, ry := Math.RotateCoordinate(cx, cy, x, y, Math.DegreesToRadians(90))
	assert.Equal(10, rx)
	assert.Equal(5, ry)
}

func TestRotateCoordinate45(t *testing.T) {
	assert := assert.New(t)

	cx, cy := 10, 10
	x, y := 5, 10

	rx, ry := Math.RotateCoordinate(cx, cy, x, y, Math.DegreesToRadians(45))
	assert.Equal(7, rx)
	assert.Equal(7, ry)
}
