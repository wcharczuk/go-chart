package chart

import (
	"testing"

	"github.com/blendlabs/go-assert"
)

func TestXAxisGetTickCount(t *testing.T) {
	assert := assert.New(t)

	r, err := PNG(1024, 1024)
	assert.Nil(err)

	f, err := GetDefaultFont()
	assert.Nil(err)

	xa := XAxis{}
	xr := Range{Min: 10, Max: 100, Domain: 1024}
	styleDefaults := Style{
		Font:     f,
		FontSize: 10.0,
	}
	vf := FloatValueFormatter
	count := xa.getTickCount(r, xr, styleDefaults, vf)
	assert.Equal(16, count)
}

func TestXAxisGetTickStep(t *testing.T) {
	assert := assert.New(t)

	r, err := PNG(1024, 1024)
	assert.Nil(err)

	f, err := GetDefaultFont()
	assert.Nil(err)

	xa := XAxis{}
	xr := Range{Min: 10, Max: 100, Domain: 1024}
	styleDefaults := Style{
		Font:     f,
		FontSize: 10.0,
	}
	vf := FloatValueFormatter
	step := xa.getTickStep(r, xr, styleDefaults, vf)
	assert.Equal(xr.Delta()/16.0, step)
}

func TestXAxisGetTicks(t *testing.T) {
	assert := assert.New(t)

	r, err := PNG(1024, 1024)
	assert.Nil(err)

	f, err := GetDefaultFont()
	assert.Nil(err)

	xa := XAxis{}
	xr := Range{Min: 10, Max: 100, Domain: 1024}
	styleDefaults := Style{
		Font:     f,
		FontSize: 10.0,
	}
	vf := FloatValueFormatter
	ticks := xa.GetTicks(r, xr, styleDefaults, vf)
	assert.Len(ticks, 17)
}

func TestXAxisGetTicksWithUserDefaults(t *testing.T) {
	assert := assert.New(t)

	r, err := PNG(1024, 1024)
	assert.Nil(err)

	f, err := GetDefaultFont()
	assert.Nil(err)

	xa := XAxis{
		Ticks: []Tick{Tick{Value: 1.0, Label: "1.0"}},
	}
	xr := Range{Min: 10, Max: 100, Domain: 1024}
	styleDefaults := Style{
		Font:     f,
		FontSize: 10.0,
	}
	vf := FloatValueFormatter
	ticks := xa.GetTicks(r, xr, styleDefaults, vf)
	assert.Len(ticks, 1)
}
