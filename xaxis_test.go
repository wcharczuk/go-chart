package chart

import (
	"testing"

	"github.com/blend/go-sdk/assert"
)

func TestXAxisGetTicks(t *testing.T) {
	assert := assert.New(t)

	r, err := PNG(1024, 1024)
	assert.Nil(err)

	f, err := GetDefaultFont()
	assert.Nil(err)

	xa := XAxis{}
	xr := &ContinuousRange{Min: 10, Max: 100, Domain: 1024}
	styleDefaults := Style{
		Font:     f,
		FontSize: 10.0,
	}
	vf := FloatValueFormatter
	ticks := xa.GetTicks(r, xr, styleDefaults, vf)
	assert.Len(16, ticks)
}

func TestXAxisGetTicksWithUserDefaults(t *testing.T) {
	assert := assert.New(t)

	r, err := PNG(1024, 1024)
	assert.Nil(err)

	f, err := GetDefaultFont()
	assert.Nil(err)

	xa := XAxis{
		Ticks: []Tick{{Value: 1.0, Label: "1.0"}},
	}
	xr := &ContinuousRange{Min: 10, Max: 100, Domain: 1024}
	styleDefaults := Style{
		Font:     f,
		FontSize: 10.0,
	}
	vf := FloatValueFormatter
	ticks := xa.GetTicks(r, xr, styleDefaults, vf)
	assert.Len(1, ticks)
}

func TestXAxisMeasure(t *testing.T) {
	assert := assert.New(t)

	f, err := GetDefaultFont()
	assert.Nil(err)
	style := Style{
		Font:     f,
		FontSize: 10.0,
	}
	r, err := PNG(100, 100)
	assert.Nil(err)
	ticks := []Tick{{Value: 1.0, Label: "1.0"}, {Value: 2.0, Label: "2.0"}, {Value: 3.0, Label: "3.0"}}
	xa := XAxis{}
	xab := xa.Measure(r, NewBox(0, 0, 100, 100), &ContinuousRange{Min: 1.0, Max: 3.0, Domain: 100}, style, ticks)
	assert.Equal(122, xab.Width())
	assert.Equal(21, xab.Height())
}
