package chart

import (
	"testing"
	"time"

	"github.com/blendlabs/go-assert"
)

func TestRangeTranslate(t *testing.T) {
	assert := assert.New(t)
	values := []float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0}
	r := NewRangeOfFloat64(1000, values...)
	assert.Equal(1.0, r.GetMin())
	assert.Equal(8.0, r.GetMax())
	assert.Equal(428, r.Translate(5.0))
}

func TestRangeOfTimeTranslate(t *testing.T) {
	assert := assert.New(t)
	values := []time.Time{
		time.Now().AddDate(0, 0, -1),
		time.Now().AddDate(0, 0, -2),
		time.Now().AddDate(0, 0, -3),
		time.Now().AddDate(0, 0, -4),
		time.Now().AddDate(0, 0, -5),
		time.Now().AddDate(0, 0, -6),
		time.Now().AddDate(0, 0, -7),
		time.Now().AddDate(0, 0, -8),
	}
	r := NewRangeOfTime(1000, values...)
	assert.Equal(values[7], r.GetMin())
	assert.Equal(values[0], r.GetMax())
	assert.Equal(571, r.Translate(time.Now().AddDate(0, 0, -5)))
}
