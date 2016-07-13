package chart

import (
	"testing"

	"github.com/blendlabs/go-assert"
)

func TestYAxisGetTickCount(t *testing.T) {
	assert := assert.New(t)

	r, err := PNG(1024, 1024)
	assert.Nil(err)

	f, err := GetDefaultFont()
	assert.Nil(err)

	ya := YAxis{}
	yr := Range{Min: 10, Max: 100, Domain: 1024}
	styleDefaults := Style{
		Font:     f,
		FontSize: 10.0,
	}
	vf := FloatValueFormatter
	count := ya.getTickCount(r, yr, styleDefaults, vf)
	assert.Equal(34, count)
}

func TestYAxisGetTickStep(t *testing.T) {
	assert := assert.New(t)

	r, err := PNG(1024, 1024)
	assert.Nil(err)

	f, err := GetDefaultFont()
	assert.Nil(err)

	ya := YAxis{}
	yr := Range{Min: 10, Max: 100, Domain: 1024}
	styleDefaults := Style{
		Font:     f,
		FontSize: 10.0,
	}
	vf := FloatValueFormatter
	step := ya.getTickStep(r, yr, styleDefaults, vf)
	assert.Equal(yr.Delta()/34.0, step)
}

func TestYAxisGetTicks(t *testing.T) {
	assert := assert.New(t)

	r, err := PNG(1024, 1024)
	assert.Nil(err)

	f, err := GetDefaultFont()
	assert.Nil(err)

	ya := YAxis{}
	yr := Range{Min: 10, Max: 100, Domain: 1024}
	styleDefaults := Style{
		Font:     f,
		FontSize: 10.0,
	}
	vf := FloatValueFormatter
	ticks := ya.GetTicks(r, yr, styleDefaults, vf)
	assert.Len(ticks, 35)
}

func TestYAxisGetTicksWithUserDefaults(t *testing.T) {
	assert := assert.New(t)

	r, err := PNG(1024, 1024)
	assert.Nil(err)

	f, err := GetDefaultFont()
	assert.Nil(err)

	ya := YAxis{
		Ticks: []Tick{Tick{Value: 1.0, Label: "1.0"}},
	}
	yr := Range{Min: 10, Max: 100, Domain: 1024}
	styleDefaults := Style{
		Font:     f,
		FontSize: 10.0,
	}
	vf := FloatValueFormatter
	ticks := ya.GetTicks(r, yr, styleDefaults, vf)
	assert.Len(ticks, 1)
}
