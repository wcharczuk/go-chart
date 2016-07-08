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
